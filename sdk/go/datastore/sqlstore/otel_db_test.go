package sqlstore

import (
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1backend/1backend/sdk/go/telemetry"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestInstrumentDBLabelsSQLTarget(t *testing.T) {
	_, _, err := telemetry.Setup(t.Context(), telemetry.Config{
		ServiceName: "sqlstore-test-svc",
	})
	require.NoError(t, err)

	router := mux.NewRouter()
	telemetry.RegisterMetricsRoute(router, telemetry.ServiceMetricsPath("sqlstore-test-svc", "/metrics"))
	server := httptest.NewServer(router)
	defer server.Close()

	db := instrumentDB(DriverPostGRES, "pets", telemetry.SQLTargetReadReplica, &otelTargetTestDB{tableName: "pets"})
	rows, err := db.Query("SELECT * FROM pets")
	require.NoError(t, err)
	require.Nil(t, rows)

	body := getSQLStoreMetricBody(t, server.URL+"/sqlstore-test-svc/metrics")
	require.Contains(t, body, "onebackend_sql_statements_total")
	require.Contains(t, body, `db_collection_name="pets"`)
	require.Contains(t, body, `db_operation_name="select"`)
	require.Contains(t, body, `datastore_sql_target="read_replica"`)
}

func getSQLStoreMetricBody(t *testing.T, url string) string {
	t.Helper()

	resp, err := http.Get(url)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(body)
}

type otelTargetTestDB struct {
	tableName string
}

func (db *otelTargetTestDB) Close() error {
	return nil
}

func (db *otelTargetTestDB) SetDebug(bool) {
}

func (db *otelTargetTestDB) SkipExec(bool) {
}

func (db *otelTargetTestDB) Tablename() string {
	return db.tableName
}

func (db *otelTargetTestDB) Query(string, ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (db *otelTargetTestDB) QueryRow(string, ...interface{}) *sql.Row {
	return &sql.Row{}
}

func (db *otelTargetTestDB) Exec(string, ...interface{}) (sql.Result, error) {
	return otelTargetTestResult{}, nil
}

func (db *otelTargetTestDB) Prepare(string) (*sql.Stmt, error) {
	return nil, nil
}

func (db *otelTargetTestDB) Begin() (Tx, error) {
	return &otelTargetTestTx{}, nil
}

type otelTargetTestTx struct{}

func (tx *otelTargetTestTx) Query(string, ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func (tx *otelTargetTestTx) QueryRow(string, ...interface{}) *sql.Row {
	return &sql.Row{}
}

func (tx *otelTargetTestTx) Exec(string, ...interface{}) (sql.Result, error) {
	return otelTargetTestResult{}, nil
}

func (tx *otelTargetTestTx) Prepare(string) (*sql.Stmt, error) {
	return nil, nil
}

func (tx *otelTargetTestTx) Rollback() error {
	return nil
}

func (tx *otelTargetTestTx) Commit() error {
	return nil
}

type otelTargetTestResult struct{}

func (r otelTargetTestResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (r otelTargetTestResult) RowsAffected() (int64, error) {
	return 0, nil
}
