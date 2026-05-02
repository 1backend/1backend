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
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/1backend/1backend/sdk/go/datastore/indexplanner"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const instrumentationName = "github.com/1backend/1backend/sdk/go/telemetry"

var (
	instrumentOnce sync.Once

	datastoreOperations        metric.Int64Counter
	datastoreOperationErrors   metric.Int64Counter
	datastoreOperationDuration metric.Float64Histogram

	sqlStatements        metric.Int64Counter
	sqlStatementErrors   metric.Int64Counter
	sqlStatementDuration metric.Float64Histogram

	autoIndexSupported metric.Int64ObservableGauge
	autoIndexShapes    metric.Int64ObservableGauge
	autoIndexInfo      metric.Int64ObservableGauge

	autoIndexMu     sync.RWMutex
	autoIndexStores = map[string]autoIndexStore{}
)

type autoIndexStore struct {
	backend  string
	table    string
	provider datastore.AutoIndexStatsProvider
}

func ensureInstruments() {
	instrumentOnce.Do(func() {
		meter := otel.Meter(instrumentationName)

		datastoreOperations, _ = meter.Int64Counter(
			"onebackend.datastore.operations",
			metric.WithDescription("Number of datastore operations executed."),
			metric.WithUnit("{operation}"),
		)
		datastoreOperationErrors, _ = meter.Int64Counter(
			"onebackend.datastore.operation.errors",
			metric.WithDescription("Number of datastore operations that returned an error."),
			metric.WithUnit("{error}"),
		)
		datastoreOperationDuration, _ = meter.Float64Histogram(
			"onebackend.datastore.operation.duration",
			metric.WithDescription("Duration of datastore operations."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5),
		)

		sqlStatements, _ = meter.Int64Counter(
			"onebackend.sql.statements",
			metric.WithDescription("Number of SQL statements executed by datastores."),
			metric.WithUnit("{statement}"),
		)
		sqlStatementErrors, _ = meter.Int64Counter(
			"onebackend.sql.statement.errors",
			metric.WithDescription("Number of SQL statements that returned an error."),
			metric.WithUnit("{error}"),
		)
		sqlStatementDuration, _ = meter.Float64Histogram(
			"onebackend.sql.statement.duration",
			metric.WithDescription("Duration of SQL statements executed by datastores."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5),
		)

		autoIndexSupported, _ = meter.Int64ObservableGauge(
			"onebackend.datastore.autoindex.supported",
			metric.WithDescription("Whether automatic indexing is supported for a datastore."),
			metric.WithUnit("1"),
		)
		autoIndexShapes, _ = meter.Int64ObservableGauge(
			"onebackend.datastore.autoindex.shape.hits",
			metric.WithDescription("Observed query shape hit counts used by the auto-index planner."),
			metric.WithUnit("{hit}"),
		)
		autoIndexInfo, _ = meter.Int64ObservableGauge(
			"onebackend.datastore.autoindex.info",
			metric.WithDescription("Auto-index entries currently known by the datastore. Value is always 1; labels describe the index."),
			metric.WithUnit("1"),
		)

		_, _ = meter.RegisterCallback(
			observeAutoIndexes,
			autoIndexSupported,
			autoIndexShapes,
			autoIndexInfo,
		)
	})
}

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

func RegisterAutoIndexStatsProvider(backend, table string, provider any) {
	statsProvider, ok := provider.(datastore.AutoIndexStatsProvider)
	if !ok || statsProvider == nil {
		return
	}

	ensureInstruments()

	backend = normalizeBackend(backend)
	key := backend + "/" + table

	autoIndexMu.Lock()
	defer autoIndexMu.Unlock()
	autoIndexStores[key] = autoIndexStore{
		backend:  backend,
		table:    table,
		provider: statsProvider,
	}
}

func RecordSQLStatement(ctx context.Context, backend, table, statement string, started time.Time, err error) {
	ensureInstruments()
	if ctx == nil {
		ctx = context.Background()
	}

	attrs := []attribute.KeyValue{
		attribute.String("db.system", normalizeBackend(backend)),
		attribute.String("db.collection.name", table),
		attribute.String("db.operation.name", SQLStatementOperation(statement)),
	}

	sqlStatements.Add(ctx, 1, metric.WithAttributes(attrs...))
	sqlStatementDuration.Record(ctx, time.Since(started).Seconds(), metric.WithAttributes(attrs...))
	if err != nil {
		sqlStatementErrors.Add(ctx, 1, metric.WithAttributes(append(attrs, errorAttributes(err)...)...))
	}
}

func SQLStatementOperation(statement string) string {
	parts := strings.Fields(strings.TrimSpace(statement))
	if len(parts) == 0 {
		return "unknown"
	}

	first := strings.ToLower(parts[0])
	if len(parts) > 1 {
		switch first {
		case "alter", "create", "drop":
			return first + "." + strings.ToLower(parts[1])
		}
	}

	return first
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

func recordDatastoreOperation(ctx context.Context, backend, table, operation string, started time.Time, err error, extra ...attribute.KeyValue) {
	ensureInstruments()
	if ctx == nil {
		ctx = context.Background()
	}

	attrs := []attribute.KeyValue{
		attribute.String("datastore.backend", normalizeBackend(backend)),
		attribute.String("db.collection.name", table),
		attribute.String("db.operation.name", operation),
	}
	attrs = append(attrs, extra...)

	datastoreOperations.Add(ctx, 1, metric.WithAttributes(attrs...))
	datastoreOperationDuration.Record(ctx, time.Since(started).Seconds(), metric.WithAttributes(attrs...))
	if err != nil {
		datastoreOperationErrors.Add(ctx, 1, metric.WithAttributes(append(attrs, errorAttributes(err)...)...))
	}
}

func observeAutoIndexes(_ context.Context, observer metric.Observer) error {
	autoIndexMu.RLock()
	stores := make([]autoIndexStore, 0, len(autoIndexStores))
	for _, store := range autoIndexStores {
		stores = append(stores, store)
	}
	autoIndexMu.RUnlock()

	for _, store := range stores {
		stats := store.provider.AutoIndexStats()
		baseAttrs := []attribute.KeyValue{
			attribute.String("datastore.backend", normalizeBackend(store.backend)),
			attribute.String("db.collection.name", store.table),
			attribute.String("datastore.autoindex.backend", stats.Backend),
		}

		supported := int64(0)
		if stats.Supported {
			supported = 1
		}
		observer.ObserveInt64(autoIndexSupported, supported, metric.WithAttributes(baseAttrs...))

		for _, shape := range stats.Shapes {
			attrs := append([]attribute.KeyValue{}, baseAttrs...)
			attrs = append(attrs,
				attribute.String("datastore.query.fingerprint", shape.Fingerprint),
				attribute.Bool("datastore.autoindex.eligible", shape.Eligible),
				attribute.String("datastore.autoindex.reason", shape.Reason),
			)
			observer.ObserveInt64(autoIndexShapes, int64(shape.Hits), metric.WithAttributes(attrs...))
		}

		for _, index := range stats.Indexes {
			attrs := append([]attribute.KeyValue{}, baseAttrs...)
			attrs = append(attrs,
				attribute.String("datastore.index.fingerprint", index.Fingerprint),
				attribute.String("datastore.index.kind", string(index.Kind)),
				attribute.String("datastore.index.status", string(index.Status)),
				attribute.String("datastore.index.method", index.Method),
				attribute.String("datastore.index.name", index.Name),
			)
			observer.ObserveInt64(autoIndexInfo, 1, metric.WithAttributes(attrs...))
		}
	}

	return nil
}

func errorAttributes(err error) []attribute.KeyValue {
	if err == nil {
		return nil
	}
	return []attribute.KeyValue{
		attribute.String("error.type", errorType(err)),
	}
}

func errorType(err error) string {
	if err == nil {
		return ""
	}

	t := strings.TrimPrefix(reflect.TypeOf(err).String(), "*")
	t = strings.TrimPrefix(t, "github.com/1backend/1backend/")
	t = strings.TrimPrefix(t, "github.com/pkg/errors.")
	t = strings.TrimPrefix(t, "errors.")
	t = strings.TrimPrefix(t, "fmt.")
	if t == "" {
		return "error"
	}
	return t
}

func normalizeBackend(backend string) string {
	if strings.TrimSpace(backend) == "" {
		return "localstore"
	}
	return strings.TrimSpace(strings.ToLower(backend))
}
