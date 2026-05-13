package sqlstore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/indexplanner"
)

const autoIndexLockTimeout = 2 * time.Minute

func (s *SQLStore) initAutoIndexes(instance any) error {
	if s.autoIndexes == nil {
		return nil
	}

	if indexable, ok := instance.(datastore.Indexable); ok {
		for _, index := range indexable.Indexes() {
			plan := indexplanner.ManualIndexPlan(index)
			s.autoIndexes.RegisterManual(plan, manualIndexName(s.tableName, index))
		}
	}

	if !s.autoIndexesOn || s.driverName != DriverPostGRES {
		return nil
	}

	rows, err := s.db.Query(
		"SELECT indexname, indexdef FROM pg_indexes WHERE schemaname = ANY(current_schemas(false)) AND tablename = $1 AND indexname LIKE $2",
		s.tableName,
		s.autoIndexPrefix()+"%",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name     string
			indexDef string
		)
		if err := rows.Scan(&name, &indexDef); err != nil {
			return err
		}
		fingerprint := strings.TrimPrefix(name, s.autoIndexPrefix())
		s.autoIndexes.RegisterDiscovered(fingerprint, detectedIndexMethod(indexDef), name)
	}

	return rows.Err()
}

func (s *SQLStore) autoIndexPrefix() string {
	return fmt.Sprintf("%s_autoidx_", s.tableName)
}

func (s *SQLStore) autoIndexName(plan indexplanner.PlannedIndex) string {
	return s.autoIndexPrefix() + plan.Fingerprint
}

func manualIndexName(tableName string, index datastore.Index) string {
	return fmt.Sprintf("%s_%s_idx", tableName, strings.Join(index.Fields, "_"))
}

func detectedIndexMethod(indexDef string) string {
	lower := strings.ToLower(indexDef)
	if strings.Contains(lower, " using gin ") {
		return indexplanner.MethodGIN
	}
	return indexplanner.MethodBTree
}

func (s *SQLStore) observePlan(plan indexplanner.QueryPlan) {
	if s.autoIndexes == nil || !s.autoIndexesOn {
		return
	}

	result := s.autoIndexes.Observe(plan)
	if s.driverName != DriverPostGRES {
		return
	}

	for _, candidate := range result.Candidates {
		name := s.autoIndexName(candidate)
		if !s.autoIndexes.MarkPending(candidate, name) {
			continue
		}
		go s.createAutoIndex(candidate, name)
	}
}

func (s *SQLStore) createAutoIndex(plan indexplanner.PlannedIndex, name string) {
	if s.autoIndexMu != nil {
		s.autoIndexMu.Lock()
		defer s.autoIndexMu.Unlock()
	}

	releaseLock, err := s.acquireAutoIndexLock()
	if err != nil {
		s.autoIndexes.MarkFailed(plan, name, err)
		return
	}
	defer releaseLock()

	if plan.OperatorClass == "gin_trgm_ops" {
		if err := s.ensurePGTrgm(); err != nil {
			s.autoIndexes.MarkFailed(plan, name, err)
			return
		}
	}

	query, err := s.buildCreateIndexStatement(plan, name)
	if err != nil {
		s.autoIndexes.MarkFailed(plan, name, err)
		return
	}

	if _, err := s.db.Exec(query); err != nil {
		s.autoIndexes.MarkFailed(plan, name, err)
		return
	}

	s.autoIndexes.MarkCreated(plan, name)
}

func (s *SQLStore) acquireAutoIndexLock() (func(), error) {
	if s.autoIndexLock == nil {
		return func() {}, nil
	}

	key := s.autoIndexLockKey()
	ctx, cancel := context.WithTimeout(context.Background(), autoIndexLockTimeout)
	defer cancel()

	if err := s.autoIndexLock.Acquire(ctx, key); err != nil {
		return nil, err
	}

	return func() {
		_ = s.autoIndexLock.Release(context.Background(), key)
	}, nil
}

func (s *SQLStore) autoIndexLockKey() string {
	return fmt.Sprintf("sqlstore:autoindex:%s", s.tableName)
}

func (s *SQLStore) ensurePGTrgm() error {
	s.pgTrgmOnce.Do(func() {
		_, s.pgTrgmErr = s.db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;")
	})
	return s.pgTrgmErr
}

func (s *SQLStore) buildCreateIndexStatement(plan indexplanner.PlannedIndex, name string) (string, error) {
	parts := make([]string, 0, len(plan.Parts))
	for _, part := range plan.Parts {
		expr := s.indexExpression(part)
		if expr == "" {
			return "", fmt.Errorf("could not compile index expression for %q", part.Field)
		}

		switch plan.Method {
		case indexplanner.MethodBTree:
			if part.Desc {
				expr += " DESC"
			}
		case indexplanner.MethodGIN:
			if plan.OperatorClass != "" {
				expr += " " + plan.OperatorClass
			}
		default:
			return "", fmt.Errorf("unsupported planned index method %q", plan.Method)
		}

		parts = append(parts, expr)
	}

	switch plan.Method {
	case indexplanner.MethodBTree:
		return fmt.Sprintf(
			"CREATE INDEX CONCURRENTLY IF NOT EXISTS %s ON %s (%s);",
			name,
			s.tableName,
			strings.Join(parts, ", "),
		), nil
	case indexplanner.MethodGIN:
		return fmt.Sprintf(
			"CREATE INDEX CONCURRENTLY IF NOT EXISTS %s ON %s USING gin (%s);",
			name,
			s.tableName,
			strings.Join(parts, ", "),
		), nil
	default:
		return "", fmt.Errorf("unsupported planned index method %q", plan.Method)
	}
}

func (s *SQLStore) indexExpression(part indexplanner.IndexPart) string {
	field := part.Field
	if len(part.Path) > 0 {
		field += "." + strings.Join(part.Path, ".")
	}
	if part.Cast != "" {
		return s.fieldName(field, part.Cast)
	}
	return s.fieldName(field)
}

func (q *SQLQueryBuilder) observeExecution() {
	if q.store == nil {
		return
	}
	dialect := indexplanner.DialectGeneric
	switch q.store.driverName {
	case DriverPostGRES:
		dialect = indexplanner.DialectPostgres
	case DriverMySQL:
		dialect = indexplanner.DialectMySQL
	}
	q.store.observePlan(indexplanner.PlanQuery(q.store.instance, q.queryDefinition(), indexplanner.PlanOptions{
		Dialect: dialect,
	}))
}

func (q *SQLQueryBuilder) queryDefinition() datastore.Query {
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
