/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package telemetry

import (
	"context"
	"encoding/json"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/indexplanner"
	"go.opentelemetry.io/otel/attribute"
)

// InstrumentDataStore wraps a datastore so high-level datastore operation
// latency, error, and query-shape metrics are emitted for callers using the
// datastore.DataStore interface.
func InstrumentDataStore(backend, table string, instance any, store datastore.DataStore) datastore.DataStore {
	if store == nil {
		return nil
	}

	ensureInstruments()
	RegisterAutoIndexStatsProvider(backend, table, store)

	return &instrumentedDataStore{
		store:    store,
		backend:  normalizeBackend(backend),
		table:    table,
		instance: instance,
	}
}

type instrumentedDataStore struct {
	store    datastore.DataStore
	backend  string
	table    string
	instance any
}

func (s *instrumentedDataStore) Create(obj datastore.Row) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "create", started, err)
	}()
	return s.store.Create(obj)
}

func (s *instrumentedDataStore) CreateMany(objs []datastore.Row) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "create_many", started, err, attribute.Int("datastore.rows", len(objs)))
	}()
	return s.store.CreateMany(objs)
}

func (s *instrumentedDataStore) Upsert(obj datastore.Row, opts ...datastore.UpsertOption) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "upsert", started, err)
	}()
	return s.store.Upsert(obj, opts...)
}

func (s *instrumentedDataStore) UpsertMany(objs []datastore.Row, opts ...datastore.UpsertOption) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "upsert_many", started, err, attribute.Int("datastore.rows", len(objs)))
	}()
	return s.store.UpsertMany(objs, opts...)
}

func (s *instrumentedDataStore) Patch(id string, fields map[string]any) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "patch", started, err)
	}()
	return s.store.Patch(id, fields)
}

func (s *instrumentedDataStore) PatchMany(updates []datastore.Patch) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "patch_many", started, err, attribute.Int("datastore.rows", len(updates)))
	}()
	return s.store.PatchMany(updates)
}

func (s *instrumentedDataStore) Query(filters ...datastore.Filter) datastore.QueryBuilder {
	return &instrumentedQueryBuilder{
		query:    s.store.Query(filters...),
		backend:  s.backend,
		table:    s.table,
		instance: s.instance,
		filters:  append([]datastore.Filter(nil), filters...),
	}
}

func (s *instrumentedDataStore) BeginTransaction() (tx datastore.DataStore, err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "transaction.begin", started, err)
	}()

	tx, err = s.store.BeginTransaction()
	if err != nil {
		return nil, err
	}

	return &instrumentedDataStore{
		store:    tx,
		backend:  s.backend,
		table:    s.table,
		instance: s.instance,
	}, nil
}

func (s *instrumentedDataStore) Commit() (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "transaction.commit", started, err)
	}()
	return s.store.Commit()
}

func (s *instrumentedDataStore) Rollback() (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "transaction.rollback", started, err)
	}()
	return s.store.Rollback()
}

func (s *instrumentedDataStore) IsInTransaction() bool {
	return s.store.IsInTransaction()
}

func (s *instrumentedDataStore) SetDebug(debug bool) {
	s.store.SetDebug(debug)
}

func (s *instrumentedDataStore) Close() (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "close", started, err)
	}()
	return s.store.Close()
}

func (s *instrumentedDataStore) Refresh() (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), s.backend, s.table, "refresh", started, err)
	}()
	return s.store.Refresh()
}

func (s *instrumentedDataStore) AutoIndexStats() datastore.AutoIndexStats {
	provider, ok := s.store.(datastore.AutoIndexStatsProvider)
	if !ok {
		return datastore.AutoIndexStats{
			Supported: false,
			Backend:   s.backend,
		}
	}
	return provider.AutoIndexStats()
}

type instrumentedQueryBuilder struct {
	query    datastore.QueryBuilder
	backend  string
	table    string
	instance any

	filters      []datastore.Filter
	orderBys     []datastore.OrderBy
	limit        int64
	after        []any
	selectFields []string
}

func (q *instrumentedQueryBuilder) OrderBy(options ...datastore.OrderBy) datastore.QueryBuilder {
	q.query = q.query.OrderBy(options...)
	q.orderBys = append([]datastore.OrderBy(nil), options...)
	return q
}

func (q *instrumentedQueryBuilder) Limit(limit int64) datastore.QueryBuilder {
	q.query = q.query.Limit(limit)
	q.limit = limit
	return q
}

func (q *instrumentedQueryBuilder) After(value ...any) datastore.QueryBuilder {
	q.query = q.query.After(value...)
	q.after = append([]any(nil), value...)
	return q
}

func (q *instrumentedQueryBuilder) Select(fields ...string) datastore.QueryBuilder {
	q.query = q.query.Select(fields...)
	q.selectFields = append([]string(nil), fields...)
	return q
}

func (q *instrumentedQueryBuilder) Find() (rows []datastore.Row, err error) {
	started := time.Now()
	defer func() {
		attrs := append(q.attributes(), attribute.Int("datastore.rows", len(rows)))
		recordDatastoreOperation(context.Background(), q.backend, q.table, "query.find", started, err, attrs...)
	}()
	return q.query.Find()
}

func (q *instrumentedQueryBuilder) FindOne() (row datastore.Row, found bool, err error) {
	started := time.Now()
	defer func() {
		rows := 0
		if found {
			rows = 1
		}
		attrs := append(q.attributes(),
			attribute.Bool("datastore.result.found", found),
			attribute.Int("datastore.rows", rows),
		)
		recordDatastoreOperation(context.Background(), q.backend, q.table, "query.find_one", started, err, attrs...)
	}()
	return q.query.FindOne()
}

func (q *instrumentedQueryBuilder) Count() (count int64, err error) {
	started := time.Now()
	defer func() {
		attrs := append(q.attributes(), attribute.Int64("datastore.result.count", count))
		recordDatastoreOperation(context.Background(), q.backend, q.table, "query.count", started, err, attrs...)
	}()
	return q.query.Count()
}

func (q *instrumentedQueryBuilder) Update(obj datastore.Row) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), q.backend, q.table, "query.update", started, err, q.attributes()...)
	}()
	return q.query.Update(obj)
}

func (q *instrumentedQueryBuilder) Upsert(obj datastore.Row) (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), q.backend, q.table, "query.upsert", started, err, q.attributes()...)
	}()
	return q.query.Upsert(obj)
}

func (q *instrumentedQueryBuilder) UpdateFields(fields map[string]interface{}) (err error) {
	started := time.Now()
	defer func() {
		attrs := append(q.attributes(), attribute.Int("datastore.updated_fields", len(fields)))
		recordDatastoreOperation(context.Background(), q.backend, q.table, "query.update_fields", started, err, attrs...)
	}()
	return q.query.UpdateFields(fields)
}

func (q *instrumentedQueryBuilder) Delete() (err error) {
	started := time.Now()
	defer func() {
		recordDatastoreOperation(context.Background(), q.backend, q.table, "query.delete", started, err, q.attributes()...)
	}()
	return q.query.Delete()
}

func (q *instrumentedQueryBuilder) attributes() []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		attribute.Int("datastore.query.filters", len(q.filters)),
		attribute.Int("datastore.query.order_bys", len(q.orderBys)),
		attribute.Bool("datastore.query.has_limit", q.limit > 0),
		attribute.Int("datastore.query.select_fields", len(q.selectFields)),
	}
	if fingerprint := q.queryFingerprint(); fingerprint != "" {
		attrs = append(attrs, attribute.String("datastore.query.fingerprint", fingerprint))
	}
	return attrs
}

func (q *instrumentedQueryBuilder) queryFingerprint() string {
	if q.instance == nil {
		return ""
	}

	afterJSON := ""
	if len(q.after) > 0 {
		if bs, err := json.Marshal(q.after); err == nil {
			afterJSON = string(bs)
		}
	}

	dialect := indexplanner.DialectGeneric
	switch normalizeBackend(q.backend) {
	case "postgres":
		dialect = indexplanner.DialectPostgres
	case "mysql":
		dialect = indexplanner.DialectMySQL
	}

	plan := indexplanner.PlanQuery(q.instance, datastore.Query{
		Filters:   append([]datastore.Filter(nil), q.filters...),
		AfterJson: afterJSON,
		Limit:     q.limit,
		OrderBys:  append([]datastore.OrderBy(nil), q.orderBys...),
	}, indexplanner.PlanOptions{
		Dialect: dialect,
	})

	return plan.ShapeFingerprint
}
