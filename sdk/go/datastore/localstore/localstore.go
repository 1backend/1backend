/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package localstore

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/indexplanner"
	"github.com/1backend/1backend/sdk/go/reflector"
	"github.com/flusflas/dipper"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

const localStoreTable = "rows"

type LocalStore struct {
	instance      any
	db            *sql.DB
	tx            *sql.Tx
	mu            *sync.RWMutex
	inTransaction bool
	filePath      string
	debug         bool
	autoIndexes   *indexplanner.Tracker
}

type storedRow struct {
	id   string
	data string
}

func NewLocalStore(instance any, filePath string) (*LocalStore, error) {
	if filePath == "" {
		tempFile, err := os.CreateTemp("", uuid.NewString())
		if err != nil {
			return nil, err
		}
		filePath = tempFile.Name()
		if err := tempFile.Close(); err != nil {
			return nil, err
		}
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", filePath)
	if err != nil {
		return nil, err
	}

	ls := &LocalStore{
		instance: instance,
		db:       db,
		mu:       &sync.RWMutex{},
		filePath: filePath,
		autoIndexes: indexplanner.NewTracker(indexplanner.TrackerOptions{
			Backend:   "localstore",
			Supported: false,
		}),
	}

	if err := ls.initDB(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return ls, nil
}

func (s *LocalStore) initDB() error {
	statements := []string{
		"PRAGMA busy_timeout = 5000;",
		"PRAGMA journal_mode = WAL;",
		"PRAGMA synchronous = NORMAL;",
		fmt.Sprintf(
			"CREATE TABLE IF NOT EXISTS %s (id TEXT PRIMARY KEY, data TEXT NOT NULL CHECK (json_valid(data)));",
			localStoreTable,
		),
	}

	for _, statement := range statements {
		if _, err := s.db.Exec(statement); err != nil {
			return err
		}
	}

	return nil
}

func (s *LocalStore) SetDebug(debug bool) {
	s.debug = debug
}

func (s *LocalStore) AutoIndexStats() datastore.AutoIndexStats {
	if s.autoIndexes == nil {
		return datastore.AutoIndexStats{
			Supported: false,
			Backend:   "localstore",
		}
	}
	return s.autoIndexes.Stats()
}

func (s *LocalStore) Create(obj datastore.Row) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.create(obj)
}

func (s *LocalStore) create(obj datastore.Row) error {
	data, err := marshalRow(obj)
	if err != nil {
		return err
	}

	_, err = s.exec(
		fmt.Sprintf("INSERT INTO %s (id, data) VALUES (?, ?);", localStoreTable),
		obj.GetId(),
		data,
	)
	if isConstraintError(err) {
		return datastore.ErrEntryAlreadyExists
	}
	return err
}

func (s *LocalStore) CreateMany(objs []datastore.Row) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.withWriteTx(func(tx *sql.Tx) error {
		for _, obj := range objs {
			data, err := marshalRow(obj)
			if err != nil {
				return err
			}

			_, err = tx.Exec(
				fmt.Sprintf("INSERT INTO %s (id, data) VALUES (?, ?);", localStoreTable),
				obj.GetId(),
				data,
			)
			if isConstraintError(err) {
				return datastore.ErrEntryAlreadyExists
			}
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *LocalStore) Upsert(obj datastore.Row, opts ...datastore.UpsertOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	options := datastore.ParseUpsertOptions(opts...)
	if len(options.Fields) == 0 {
		return s.upsertFull(obj)
	}
	return s.upsertPartial(obj, options.Fields...)
}

func (s *LocalStore) UpsertMany(objs []datastore.Row, opts ...datastore.UpsertOption) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	options := datastore.ParseUpsertOptions(opts...)
	return s.withWriteTx(func(tx *sql.Tx) error {
		for _, obj := range objs {
			if len(options.Fields) == 0 {
				if err := upsertFullTx(tx, obj); err != nil {
					return err
				}
				continue
			}

			if err := s.upsertPartialTx(tx, obj, options.Fields...); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *LocalStore) upsertFull(obj datastore.Row) error {
	return s.withWriteTx(func(tx *sql.Tx) error {
		return upsertFullTx(tx, obj)
	})
}

func upsertFullTx(tx *sql.Tx, obj datastore.Row) error {
	data, err := marshalRow(obj)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		fmt.Sprintf(
			"INSERT INTO %s (id, data) VALUES (?, ?) ON CONFLICT(id) DO UPDATE SET data = excluded.data;",
			localStoreTable,
		),
		obj.GetId(),
		data,
	)
	return err
}

func (s *LocalStore) upsertPartial(obj datastore.Row, fields ...string) error {
	return s.withWriteTx(func(tx *sql.Tx) error {
		return s.upsertPartialTx(tx, obj, fields...)
	})
}

func (s *LocalStore) upsertPartialTx(tx *sql.Tx, obj datastore.Row, fields ...string) error {
	existing, found, err := rowByID(tx, obj.GetId())
	if err != nil {
		return err
	}
	if !found {
		return upsertFullTx(tx, obj)
	}

	partialObj, err := reflector.DeepCopyIntoMap(obj)
	if err != nil {
		return err
	}

	updated, err := patchJSON(existing.data, func(target any) error {
		for _, field := range fields {
			value := getField(partialObj, field)
			if err := setField(&target, field, value); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		fmt.Sprintf("UPDATE %s SET data = ? WHERE id = ?;", localStoreTable),
		updated,
		obj.GetId(),
	)
	return err
}

func (s *LocalStore) Patch(id string, fields map[string]any) error {
	return s.PatchMany([]datastore.Patch{{ID: id, Fields: fields}})
}

func (s *LocalStore) PatchMany(updates []datastore.Patch) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.withWriteTx(func(tx *sql.Tx) error {
		for _, update := range updates {
			existing, found, err := rowByID(tx, update.ID)
			if err != nil {
				return err
			}

			baseData := existing.data
			if !found {
				base := map[string]any{}
				if err := setField(&base, s.idFieldName(), update.ID); err != nil {
					return err
				}
				bs, err := json.Marshal(base)
				if err != nil {
					return err
				}
				baseData = string(bs)
			}

			updated, err := patchJSON(baseData, func(target any) error {
				for field, value := range update.Fields {
					if err := setField(&target, field, value); err != nil {
						return err
					}
				}
				return nil
			})
			if err != nil {
				return err
			}

			_, err = tx.Exec(
				fmt.Sprintf(
					"INSERT INTO %s (id, data) VALUES (?, ?) ON CONFLICT(id) DO UPDATE SET data = excluded.data;",
					localStoreTable,
				),
				update.ID,
				updated,
			)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *LocalStore) idFieldName() string {
	t := reflect.TypeOf(s.instance)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return jsonFieldName(t.Field(0))
}

func (s *LocalStore) Query(filters ...datastore.Filter) datastore.QueryBuilder {
	q := &QueryBuilder{store: s}
	q.filters = append(q.filters, filters...)
	return q
}

func (s *LocalStore) BeginTransaction() (datastore.DataStore, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.inTransaction {
		return nil, errors.New("already in a transaction")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	return &LocalStore{
		instance:      s.instance,
		db:            s.db,
		tx:            tx,
		mu:            s.mu,
		inTransaction: true,
		filePath:      s.filePath,
		debug:         s.debug,
		autoIndexes:   s.autoIndexes,
	}, nil
}

func (s *LocalStore) Commit() error {
	if !s.inTransaction || s.tx == nil {
		return errors.New("not in a transaction")
	}

	err := s.tx.Commit()
	s.tx = nil
	s.inTransaction = false
	return err
}

func (s *LocalStore) Rollback() error {
	if !s.inTransaction || s.tx == nil {
		return errors.New("not in a transaction")
	}

	err := s.tx.Rollback()
	s.tx = nil
	s.inTransaction = false
	return err
}

func (s *LocalStore) IsInTransaction() bool {
	return s.inTransaction
}

func (s *LocalStore) Close() error {
	if s.inTransaction {
		if s.tx != nil {
			err := s.tx.Rollback()
			if err != nil && !errors.Is(err, sql.ErrTxDone) {
				return err
			}
			s.tx = nil
		}
		s.inTransaction = false
		return nil
	}

	return s.db.Close()
}

func (s *LocalStore) Refresh() error {
	return nil
}

func (s *LocalStore) exec(query string, args ...any) (sql.Result, error) {
	if s.debug {
		fmt.Printf("[localstore] %s args=%v\n", query, args)
	}
	if s.tx != nil {
		return s.tx.Exec(query, args...)
	}
	return s.db.Exec(query, args...)
}

func (s *LocalStore) query(query string, args ...any) (*sql.Rows, error) {
	if s.debug {
		fmt.Printf("[localstore] %s args=%v\n", query, args)
	}
	if s.tx != nil {
		return s.tx.Query(query, args...)
	}
	return s.db.Query(query, args...)
}

func (s *LocalStore) queryRow(query string, args ...any) *sql.Row {
	if s.debug {
		fmt.Printf("[localstore] %s args=%v\n", query, args)
	}
	if s.tx != nil {
		return s.tx.QueryRow(query, args...)
	}
	return s.db.QueryRow(query, args...)
}

func (s *LocalStore) withWriteTx(fn func(tx *sql.Tx) error) error {
	if s.tx != nil {
		return fn(s.tx)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

type QueryBuilder struct {
	store             *LocalStore
	filters           []datastore.Filter
	orderFields       []string
	orderDescs        []bool
	orderSortingTypes []datastore.SortingType
	limit             int64
	after             []any
	selectFields      []string
}

func (q *QueryBuilder) OrderBy(options ...datastore.OrderBy) datastore.QueryBuilder {
	if len(options) == 0 {
		return q
	}

	q.orderFields = make([]string, len(options))
	q.orderDescs = make([]bool, len(options))
	q.orderSortingTypes = make([]datastore.SortingType, len(options))

	for i, option := range options {
		q.orderFields[i] = option.Field
		q.orderDescs[i] = option.Desc
		q.orderSortingTypes[i] = option.SortingType
	}

	return q
}

func (q *QueryBuilder) Limit(limit int64) datastore.QueryBuilder {
	q.limit = limit
	return q
}

func (q *QueryBuilder) After(value ...any) datastore.QueryBuilder {
	if len(value) == 0 {
		return q
	}

	q.after = make([]any, len(value))
	for i := range value {
		q.after[i] = normalizeParam(value[i])
	}

	return q
}

func (q *QueryBuilder) Select(fields ...string) datastore.QueryBuilder {
	q.selectFields = fields
	return q
}

func (q *QueryBuilder) Find() ([]datastore.Row, error) {
	query, params, err := q.buildSelectQuery(false, 0)
	if err != nil {
		return nil, err
	}

	rows, err := q.store.query(query, params...)
	q.observe()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []datastore.Row
	for rows.Next() {
		var raw string
		if err := rows.Scan(&raw); err != nil {
			return nil, err
		}

		obj, err := q.store.unmarshalRow(raw)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (q *QueryBuilder) FindOne() (datastore.Row, bool, error) {
	var empty datastore.Row

	query, params, err := q.buildSelectQuery(false, 1)
	if err != nil {
		return empty, false, err
	}

	var raw string
	err = q.store.queryRow(query, params...).Scan(&raw)
	q.observe()
	if errors.Is(err, sql.ErrNoRows) {
		return empty, false, nil
	}
	if err != nil {
		return empty, false, err
	}

	obj, err := q.store.unmarshalRow(raw)
	if err != nil {
		return empty, false, err
	}

	return obj, true, nil
}

func (q *QueryBuilder) Count() (int64, error) {
	query, params, err := q.buildSelectQuery(true, 0)
	if err != nil {
		return 0, err
	}

	var count int64
	err = q.store.queryRow(query, params...).Scan(&count)
	q.observe()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *QueryBuilder) Update(obj datastore.Row) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	ids, err := q.matchingIDs(q.store.tx)
	if err != nil {
		return err
	}
	if len(ids) == 0 {
		q.observe()
		return errors.New("no records to update")
	}

	data, err := marshalRow(obj)
	if err != nil {
		return err
	}

	err = q.store.withWriteTx(func(tx *sql.Tx) error {
		for _, id := range ids {
			if _, err := tx.Exec(fmt.Sprintf("UPDATE %s SET data = ? WHERE id = ?;", localStoreTable), data, id); err != nil {
				return err
			}
		}
		return nil
	})
	q.observe()
	return err
}

func (q *QueryBuilder) Upsert(obj datastore.Row) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	err := q.store.withWriteTx(func(tx *sql.Tx) error {
		ids, err := q.matchingIDs(tx)
		if err != nil {
			return err
		}
		if len(ids) == 0 {
			return upsertFullTx(tx, obj)
		}

		data, err := marshalRow(obj)
		if err != nil {
			return err
		}
		for _, id := range ids {
			if _, err := tx.Exec(fmt.Sprintf("UPDATE %s SET data = ? WHERE id = ?;", localStoreTable), data, id); err != nil {
				return err
			}
		}
		return nil
	})
	q.observe()
	return err
}

func (q *QueryBuilder) UpdateFields(fields map[string]interface{}) error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	err := q.store.withWriteTx(func(tx *sql.Tx) error {
		rows, err := q.matchingRows(tx)
		if err != nil {
			return err
		}

		for _, row := range rows {
			updated, err := patchJSON(row.data, func(target any) error {
				for field, value := range fields {
					if err := setField(&target, field, value); err != nil {
						return err
					}
				}
				return nil
			})
			if err != nil {
				return err
			}

			if _, err := tx.Exec(fmt.Sprintf("UPDATE %s SET data = ? WHERE id = ?;", localStoreTable), updated, row.id); err != nil {
				return err
			}
		}

		return nil
	})
	q.observe()
	return err
}

func (q *QueryBuilder) Delete() error {
	q.store.mu.Lock()
	defer q.store.mu.Unlock()

	where, params, err := q.buildWhere(false)
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s", localStoreTable)
	if where != "" {
		query += " WHERE " + where
	}

	_, err = q.store.exec(query, params...)
	q.observe()
	return err
}

func (q *QueryBuilder) buildSelectQuery(count bool, forceLimit int64) (string, []any, error) {
	where, params, err := q.buildWhere(true)
	if err != nil {
		return "", nil, err
	}

	selected := "data"
	if count {
		selected = "COUNT(*)"
	}

	query := fmt.Sprintf("SELECT %s FROM %s", selected, localStoreTable)
	if where != "" {
		query += " WHERE " + where
	}

	if !count && len(q.orderFields) > 0 {
		orderParts := make([]string, 0, len(q.orderFields))
		for i, field := range q.orderFields {
			if i < len(q.orderSortingTypes) && q.orderSortingTypes[i] == datastore.SortingTypeRandom {
				orderParts = append(orderParts, "RANDOM()")
				continue
			}

			expr := fieldExpr(field, q.sortingTypeAt(i))
			if i < len(q.orderDescs) && q.orderDescs[i] {
				expr += " DESC"
			}
			orderParts = append(orderParts, expr)
		}
		query += " ORDER BY " + strings.Join(orderParts, ", ")
	}

	limit := q.limit
	if forceLimit > 0 {
		limit = forceLimit
	}
	if !count && limit > 0 {
		query += " LIMIT ?"
		params = append(params, limit)
	}

	return query, params, nil
}

func (q *QueryBuilder) buildWhere(includeAfter bool) (string, []any, error) {
	var parts []string
	var params []any

	for _, filter := range q.filters {
		sqlPart, sqlParams, err := q.compileFilter(filter)
		if err != nil {
			return "", nil, err
		}
		if sqlPart == "" {
			continue
		}
		parts = append(parts, sqlPart)
		params = append(params, sqlParams...)
	}

	if includeAfter && len(q.after) > 0 && len(q.orderFields) > 0 {
		afterSQL, afterParams, err := q.compileAfter()
		if err != nil {
			return "", nil, err
		}
		if afterSQL != "" {
			parts = append(parts, afterSQL)
			params = append(params, afterParams...)
		}
	}

	return strings.Join(parts, " AND "), params, nil
}

func (q *QueryBuilder) compileFilter(filter datastore.Filter) (string, []any, error) {
	switch filter.Op {
	case datastore.OpOr:
		var parts []string
		var params []any
		for _, subFilter := range filter.SubFilters {
			part, subParams, err := q.compileFilter(subFilter)
			if err != nil {
				return "", nil, err
			}
			if part == "" {
				continue
			}
			parts = append(parts, part)
			params = append(params, subParams...)
		}
		if len(parts) == 0 {
			return "0", nil, nil
		}
		return "(" + strings.Join(parts, " OR ") + ")", params, nil
	case datastore.OpEquals:
		return compileEquals(filter)
	case datastore.OpIsInList:
		return compileIsInList(filter)
	case datastore.OpIntersects:
		return compileIntersects(filter)
	case datastore.OpContainsSubstring:
		return compileStringMatch(filter, false)
	case datastore.OpStartsWith:
		return compileStringMatch(filter, true)
	case datastore.OpLessThan:
		return compileComparison(filter, "<")
	case datastore.OpLessThanOrEqual:
		return compileComparison(filter, "<=")
	case datastore.OpGreaterThan:
		return compileComparison(filter, ">")
	case datastore.OpGreaterThanOrEqual:
		return compileComparison(filter, ">=")
	default:
		return "", nil, fmt.Errorf("unknown filter %v", filter)
	}
}

func compileEquals(filter datastore.Filter) (string, []any, error) {
	values, err := filterValues(filter)
	if err != nil {
		return "", nil, err
	}
	if len(values) == 0 {
		return "0", nil, nil
	}

	value := normalizeParam(values[0])

	var parts []string
	var params []any
	for _, field := range filter.Fields {
		path := jsonPath(field)
		if value == nil {
			parts = append(parts, fmt.Sprintf(
				"(json_type(data, %s) IS NULL OR json_type(data, %s) = 'null')",
				sqlStringLiteral(path),
				sqlStringLiteral(path),
			))
			continue
		}

		parts = append(parts, fmt.Sprintf(
			"((json_type(data, %s) = 'array' AND EXISTS (SELECT 1 FROM json_each(data, %s) WHERE value = ?)) OR (json_type(data, %s) != 'array' AND %s = ?))",
			sqlStringLiteral(path),
			sqlStringLiteral(path),
			sqlStringLiteral(path),
			fieldExpr(field, datastore.SortingTypeDefault),
		))
		params = append(params, value, value)
	}

	if len(parts) == 0 {
		return "0", nil, nil
	}
	return "(" + strings.Join(parts, " OR ") + ")", params, nil
}

func compileIsInList(filter datastore.Filter) (string, []any, error) {
	values, err := filterValues(filter)
	if err != nil {
		return "", nil, err
	}
	if len(values) == 0 {
		return "0", nil, nil
	}

	placeholders := strings.TrimRight(strings.Repeat("?,", len(values)), ",")
	paramsForField := make([]any, 0, len(values))
	for _, value := range values {
		paramsForField = append(paramsForField, normalizeParam(value))
	}

	var parts []string
	var params []any
	for _, field := range filter.Fields {
		parts = append(parts, fmt.Sprintf(
			"(json_type(data, %s) != 'array' AND %s IN (%s))",
			sqlStringLiteral(jsonPath(field)),
			fieldExpr(field, datastore.SortingTypeDefault),
			placeholders,
		))
		params = append(params, paramsForField...)
	}

	if len(parts) == 0 {
		return "0", nil, nil
	}
	return "(" + strings.Join(parts, " OR ") + ")", params, nil
}

func compileIntersects(filter datastore.Filter) (string, []any, error) {
	values, err := filterValues(filter)
	if err != nil {
		return "", nil, err
	}
	if len(values) == 0 {
		return "0", nil, nil
	}

	valuesJSON, err := json.Marshal(values)
	if err != nil {
		return "", nil, err
	}

	var parts []string
	var params []any
	for _, field := range filter.Fields {
		parts = append(parts, fmt.Sprintf(
			"(json_type(data, %s) = 'array' AND EXISTS (SELECT 1 FROM json_each(data, %s) AS stored JOIN json_each(?) AS wanted ON stored.value = wanted.value))",
			sqlStringLiteral(jsonPath(field)),
			sqlStringLiteral(jsonPath(field)),
		))
		params = append(params, string(valuesJSON))
	}

	if len(parts) == 0 {
		return "0", nil, nil
	}
	return "(" + strings.Join(parts, " OR ") + ")", params, nil
}

func compileStringMatch(filter datastore.Filter, startsWith bool) (string, []any, error) {
	values, err := filterValues(filter)
	if err != nil {
		return "", nil, err
	}
	if len(values) == 0 {
		return "0", nil, nil
	}

	value, ok := normalizeParam(values[0]).(string)
	if !ok {
		return "0", nil, nil
	}

	var parts []string
	var params []any
	for _, field := range filter.Fields {
		path := jsonPath(field)
		var scalarExpr string
		var arrayExpr string
		if startsWith {
			scalarExpr = fmt.Sprintf("substr(CAST(%s AS TEXT), 1, length(?)) = ?", fieldExpr(field, datastore.SortingTypeDefault))
			arrayExpr = fmt.Sprintf("substr(CAST(value AS TEXT), 1, length(?)) = ?")
		} else {
			scalarExpr = fmt.Sprintf("instr(CAST(%s AS TEXT), ?) > 0", fieldExpr(field, datastore.SortingTypeDefault))
			arrayExpr = "instr(CAST(value AS TEXT), ?) > 0"
		}

		parts = append(parts, fmt.Sprintf(
			"((json_type(data, %s) = 'array' AND EXISTS (SELECT 1 FROM json_each(data, %s) WHERE %s)) OR (json_type(data, %s) != 'array' AND %s))",
			sqlStringLiteral(path),
			sqlStringLiteral(path),
			arrayExpr,
			sqlStringLiteral(path),
			scalarExpr,
		))
		if startsWith {
			params = append(params, value, value, value, value)
		} else {
			params = append(params, value, value)
		}
	}

	if len(parts) == 0 {
		return "0", nil, nil
	}
	return "(" + strings.Join(parts, " OR ") + ")", params, nil
}

func compileComparison(filter datastore.Filter, op string) (string, []any, error) {
	values, err := filterValues(filter)
	if err != nil {
		return "", nil, err
	}
	if len(values) != 1 {
		return "", nil, errors.New("comparison operators require exactly one value")
	}

	value := normalizeParam(values[0])
	sortingType := datastore.SortingTypeDefault
	if isNumeric(value) {
		sortingType = datastore.SortingTypeNumeric
	}

	var parts []string
	var params []any
	for _, field := range filter.Fields {
		expr := fieldExpr(field, sortingType)
		if sortingType == datastore.SortingTypeNumeric {
			expr = numericFieldExpr(field)
		}
		parts = append(parts, fmt.Sprintf(
			"(json_type(data, %s) != 'array' AND %s %s ?)",
			sqlStringLiteral(jsonPath(field)),
			expr,
			op,
		))
		params = append(params, value)
	}

	if len(parts) == 0 {
		return "0", nil, nil
	}
	return "(" + strings.Join(parts, " OR ") + ")", params, nil
}

func (q *QueryBuilder) compileAfter() (string, []any, error) {
	var orParts []string
	var params []any

	max := len(q.orderFields)
	if len(q.after) < max {
		max = len(q.after)
	}

	for i := 0; i < max; i++ {
		var andParts []string
		for j := 0; j < i; j++ {
			andParts = append(andParts, fmt.Sprintf("%s = ?", fieldExpr(q.orderFields[j], q.sortingTypeAt(j))))
			params = append(params, normalizeParam(q.after[j]))
		}

		op := ">"
		if i < len(q.orderDescs) && q.orderDescs[i] {
			op = "<"
		}
		andParts = append(andParts, fmt.Sprintf("%s %s ?", fieldExpr(q.orderFields[i], q.sortingTypeAt(i)), op))
		params = append(params, normalizeParam(q.after[i]))

		orParts = append(orParts, "("+strings.Join(andParts, " AND ")+")")
	}

	if len(orParts) == 0 {
		return "", nil, nil
	}
	return "(" + strings.Join(orParts, " OR ") + ")", params, nil
}

func (q *QueryBuilder) sortingTypeAt(index int) datastore.SortingType {
	if index < len(q.orderSortingTypes) && q.orderSortingTypes[index] != "" {
		return q.orderSortingTypes[index]
	}
	if index < len(q.after) && isNumeric(q.after[index]) {
		return datastore.SortingTypeNumeric
	}
	return datastore.SortingTypeDefault
}

func (q *QueryBuilder) matchingRows(tx *sql.Tx) ([]storedRow, error) {
	where, params, err := q.buildWhere(false)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT id, data FROM %s", localStoreTable)
	if where != "" {
		query += " WHERE " + where
	}

	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.Query(query, params...)
	} else {
		rows, err = q.store.query(query, params...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []storedRow
	for rows.Next() {
		var row storedRow
		if err := rows.Scan(&row.id, &row.data); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (q *QueryBuilder) matchingIDs(tx *sql.Tx) ([]string, error) {
	rows, err := q.matchingRows(tx)
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(rows))
	for _, row := range rows {
		ids = append(ids, row.id)
	}
	return ids, nil
}

func (q *QueryBuilder) observe() {
	if q.store.autoIndexes == nil {
		return
	}
	plan := indexplanner.PlanQuery(q.store.instance, q.queryDefinition(), indexplanner.PlanOptions{
		Dialect: indexplanner.DialectGeneric,
	})
	q.store.autoIndexes.Observe(plan)
}

func (q *QueryBuilder) queryDefinition() datastore.Query {
	afterJSON := ""
	if len(q.after) > 0 {
		if bs, err := json.Marshal(q.after); err == nil {
			afterJSON = string(bs)
		}
	}

	orderBys := make([]datastore.OrderBy, 0, len(q.orderFields))
	for i, field := range q.orderFields {
		orderBy := datastore.OrderBy{
			Field: field,
		}
		if i < len(q.orderDescs) {
			orderBy.Desc = q.orderDescs[i]
		}
		if i < len(q.orderSortingTypes) {
			orderBy.SortingType = q.orderSortingTypes[i]
		}
		orderBys = append(orderBys, orderBy)
	}

	return datastore.Query{
		Filters:   append([]datastore.Filter(nil), q.filters...),
		AfterJson: afterJSON,
		Limit:     q.limit,
		OrderBys:  orderBys,
	}
}

func (s *LocalStore) unmarshalRow(raw string) (datastore.Row, error) {
	instanceType := reflect.TypeOf(s.instance)
	if instanceType.Kind() == reflect.Ptr {
		obj := reflect.New(instanceType.Elem())
		if err := json.Unmarshal([]byte(raw), obj.Interface()); err != nil {
			return nil, err
		}
		row, ok := obj.Interface().(datastore.Row)
		if !ok {
			return nil, fmt.Errorf("%T does not implement datastore.Row", obj.Interface())
		}
		return row, nil
	}

	obj := reflect.New(instanceType)
	if err := json.Unmarshal([]byte(raw), obj.Interface()); err != nil {
		return nil, err
	}
	row, ok := obj.Elem().Interface().(datastore.Row)
	if !ok {
		return nil, fmt.Errorf("%T does not implement datastore.Row", obj.Elem().Interface())
	}
	return row, nil
}

func marshalRow(obj datastore.Row) (string, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func rowByID(tx *sql.Tx, id string) (storedRow, bool, error) {
	var row storedRow
	err := tx.QueryRow(fmt.Sprintf("SELECT id, data FROM %s WHERE id = ?;", localStoreTable), id).Scan(&row.id, &row.data)
	if errors.Is(err, sql.ErrNoRows) {
		return storedRow{}, false, nil
	}
	if err != nil {
		return storedRow{}, false, err
	}
	return row, true, nil
}

func patchJSON(raw string, patch func(target any) error) (string, error) {
	var target any
	if strings.TrimSpace(raw) == "" {
		target = map[string]any{}
	} else if err := json.Unmarshal([]byte(raw), &target); err != nil {
		return "", err
	}
	if target == nil {
		target = map[string]any{}
	}

	if err := patch(target); err != nil {
		return "", err
	}

	bs, err := json.Marshal(target)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func filterValues(filter datastore.Filter) ([]any, error) {
	var values []any
	if filter.ValuesJson == "" {
		return values, nil
	}
	if err := json.Unmarshal([]byte(filter.ValuesJson), &values); err != nil {
		return nil, err
	}
	return values, nil
}

func fieldExpr(field string, sortingType datastore.SortingType) string {
	expr := fmt.Sprintf("json_extract(data, %s)", sqlStringLiteral(jsonPath(field)))
	switch sortingType {
	case datastore.SortingTypeNumeric:
		return fmt.Sprintf("COALESCE(CAST(%s AS REAL), 0)", expr)
	default:
		return expr
	}
}

func numericFieldExpr(field string) string {
	return fmt.Sprintf("CAST(json_extract(data, %s) AS REAL)", sqlStringLiteral(jsonPath(field)))
}

func jsonPath(field string) string {
	parts := strings.Split(field, ".")
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		parts[i] = strings.ToLower(parts[i][:1]) + parts[i][1:]
	}
	return "$." + strings.Join(parts, ".")
}

func sqlStringLiteral(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "''") + "'"
}

func normalizeParam(value any) any {
	if value == nil {
		return nil
	}

	if t, ok := value.(time.Time); ok {
		return t.Format(time.RFC3339Nano)
	}

	rv := reflect.ValueOf(value)
	for rv.IsValid() && rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
		value = rv.Interface()
	}
	if !rv.IsValid() {
		return nil
	}

	switch rv.Kind() {
	case reflect.String:
		return rv.String()
	case reflect.Bool:
		return rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint()
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
		bs, err := json.Marshal(value)
		if err != nil {
			return fmt.Sprintf("%v", value)
		}
		return string(bs)
	default:
		return value
	}
}

func isNumeric(value any) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}

func isConstraintError(err error) bool {
	if err == nil {
		return false
	}
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "constraint") || strings.Contains(message, "unique")
}

func jsonFieldName(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "-" {
		return strings.ToLower(field.Name[:1]) + field.Name[1:]
	}
	if tag != "" {
		name := strings.Split(tag, ",")[0]
		if name != "" {
			return name
		}
	}
	return strings.ToLower(field.Name[:1]) + field.Name[1:]
}

func fixFieldName(s string) string {
	parts := strings.Split(s, ".")
	for i := range parts {
		if parts[i] == "" {
			continue
		}
		parts[i] = strings.ToLower(parts[i][:1]) + parts[i][1:]
	}

	return strings.Join(parts, ".")
}

func getField(obj any, field string) interface{} {
	field = fixFieldName(field)

	v := dipper.Get(obj, field)
	if err := dipper.Error(v); err != nil {
		return nil
	}
	return v
}

func setField(obj any, field string, value interface{}) error {
	field = fixFieldName(field)
	return dipper.Set(obj, field, value)
}
