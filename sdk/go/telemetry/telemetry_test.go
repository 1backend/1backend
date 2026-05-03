package telemetry

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/datastore"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

var testSetupOnce sync.Once
var testSetupErr error

func setupTestTelemetry(t *testing.T) {
	t.Helper()
	testSetupOnce.Do(func() {
		_, _, testSetupErr = Setup(t.Context(), Config{
			ServiceName: "telemetry-test-svc",
		})
	})
	require.NoError(t, testSetupErr)
}

func TestServiceMetricsPath(t *testing.T) {
	require.Equal(t, "/basic-svc/metrics", ServiceMetricsPath("basic-svc", ""))
	require.Equal(t, "/basic-svc/metrics", ServiceMetricsPath("/basic-svc/", "/metrics"))
	require.Equal(t, "/custom", ServiceMetricsPath("basic-svc", "/custom"))
	require.Equal(t, "/metrics", ServiceMetricsPath("", "/metrics"))
}

func TestHTTPAndDatastoreMetricsAreScrapeable(t *testing.T) {
	setupTestTelemetry(t)

	router := mux.NewRouter()
	router.Use(HTTPMiddleware("telemetry-test-svc"))
	RegisterMetricsRoute(router, ServiceMetricsPath("telemetry-test-svc", "/metrics"))
	router.HandleFunc("/telemetry-test-svc/ok", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}).Methods(http.MethodGet)
	router.HandleFunc("/telemetry-test-svc/fail", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "failed", http.StatusInternalServerError)
	}).Methods(http.MethodGet)

	server := httptest.NewServer(router)
	defer server.Close()

	requireHTTPStatus(t, server.URL+"/telemetry-test-svc/ok", http.StatusOK)
	requireHTTPStatus(t, server.URL+"/telemetry-test-svc/fail", http.StatusInternalServerError)

	store := InstrumentDataStore("localstore", "pets", testTelemetryRow{}, &testTelemetryStore{})
	require.NoError(t, store.Create(testTelemetryRow{ID: "pet_1"}))
	_, err := store.Query(datastore.Equals([]string{"name"}, "Ada")).Find()
	require.NoError(t, err)
	RecordSQLStatementTarget(t.Context(), "postgres", "pets", SQLTargetReadReplica, "SELECT * FROM pets", time.Now(), nil)

	body := getBody(t, server.URL+"/telemetry-test-svc/metrics")
	require.Contains(t, body, "onebackend_http_server_requests_total")
	require.Contains(t, body, "onebackend_http_server_errors_total")
	require.Contains(t, body, `http_route="/telemetry-test-svc/ok"`)
	require.Contains(t, body, `http_route="/telemetry-test-svc/fail"`)
	require.Contains(t, body, `service_name="telemetry-test-svc"`)
	require.Contains(t, body, "onebackend_datastore_operations_total")
	require.Contains(t, body, "onebackend_datastore_operation_duration_seconds")
	require.Contains(t, body, "onebackend_datastore_autoindex_shape_hits")
	require.Contains(t, body, "onebackend_sql_statements_total")
	require.Contains(t, body, `db_collection_name="pets"`)
	require.Contains(t, body, `db_operation_name="select"`)
	require.Contains(t, body, `datastore_sql_target="read_replica"`)
}

func requireHTTPStatus(t *testing.T, url string, status int) {
	t.Helper()
	resp, err := http.Get(url)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, status, resp.StatusCode)
}

func getBody(t *testing.T, url string) string {
	t.Helper()
	resp, err := http.Get(url)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(body)
}

type testTelemetryRow struct {
	ID string
}

func (r testTelemetryRow) GetId() string {
	return r.ID
}

type testTelemetryStore struct{}

func (s *testTelemetryStore) Create(datastore.Row) error {
	return nil
}

func (s *testTelemetryStore) CreateMany([]datastore.Row) error {
	return nil
}

func (s *testTelemetryStore) Upsert(datastore.Row, ...datastore.UpsertOption) error {
	return nil
}

func (s *testTelemetryStore) UpsertMany([]datastore.Row, ...datastore.UpsertOption) error {
	return nil
}

func (s *testTelemetryStore) Patch(string, map[string]any) error {
	return nil
}

func (s *testTelemetryStore) PatchMany([]datastore.Patch) error {
	return nil
}

func (s *testTelemetryStore) Query(...datastore.Filter) datastore.QueryBuilder {
	return &testTelemetryQuery{}
}

func (s *testTelemetryStore) BeginTransaction() (datastore.DataStore, error) {
	return s, nil
}

func (s *testTelemetryStore) Commit() error {
	return nil
}

func (s *testTelemetryStore) Rollback() error {
	return nil
}

func (s *testTelemetryStore) IsInTransaction() bool {
	return false
}

func (s *testTelemetryStore) SetDebug(bool) {
}

func (s *testTelemetryStore) Close() error {
	return nil
}

func (s *testTelemetryStore) Refresh() error {
	return nil
}

func (s *testTelemetryStore) AutoIndexStats() datastore.AutoIndexStats {
	return datastore.AutoIndexStats{
		Supported: true,
		Backend:   "localstore",
		Shapes: []datastore.AutoIndexShapeStats{
			{
				Fingerprint: "shape-test",
				Hits:        3,
				Eligible:    true,
			},
		},
		Indexes: []datastore.AutoIndexEntry{
			{
				Fingerprint: "shape-test",
				Kind:        datastore.AutoIndexKindAuto,
				Status:      datastore.AutoIndexStatusCreated,
				Method:      "btree",
				Name:        "idx_test",
			},
		},
	}
}

type testTelemetryQuery struct{}

func (q *testTelemetryQuery) OrderBy(...datastore.OrderBy) datastore.QueryBuilder {
	return q
}

func (q *testTelemetryQuery) Limit(int64) datastore.QueryBuilder {
	return q
}

func (q *testTelemetryQuery) After(...any) datastore.QueryBuilder {
	return q
}

func (q *testTelemetryQuery) Select(...string) datastore.QueryBuilder {
	return q
}

func (q *testTelemetryQuery) Find() ([]datastore.Row, error) {
	return []datastore.Row{testTelemetryRow{ID: "pet_1"}}, nil
}

func (q *testTelemetryQuery) FindOne() (datastore.Row, bool, error) {
	return testTelemetryRow{ID: "pet_1"}, true, nil
}

func (q *testTelemetryQuery) Count() (int64, error) {
	return 1, nil
}

func (q *testTelemetryQuery) Update(datastore.Row) error {
	return nil
}

func (q *testTelemetryQuery) Upsert(datastore.Row) error {
	return nil
}

func (q *testTelemetryQuery) UpdateFields(map[string]interface{}) error {
	return nil
}

func (q *testTelemetryQuery) Delete() error {
	return nil
}
