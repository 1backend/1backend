package telemetry

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestAuthMetricsAreScrapeable(t *testing.T) {
	setupTestTelemetry(t)

	router := mux.NewRouter()
	RegisterMetricsRoute(router, ServiceMetricsPath("auth-metrics-test-svc", "/metrics"))
	server := httptest.NewServer(router)
	defer server.Close()

	RecordAuthOperation(t.Context(), AuthOperationTokenRefresh, "success", "network", time.Now().Add(-10*time.Millisecond), nil)
	RecordAuthOperation(t.Context(), AuthOperationTokenRefresh, "error", "network", time.Now().Add(-10*time.Millisecond), assertErr{})
	RecordAuthCache(t.Context(), AuthCacheTokenRefresh, "hit")
	RecordAuthCache(t.Context(), AuthCachePermission, "miss")

	resp, err := http.Get(server.URL + "/auth-metrics-test-svc/metrics")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	body := string(bodyBytes)

	require.Contains(t, body, "onebackend_auth_operations_total")
	require.Contains(t, body, "onebackend_auth_operation_errors_total")
	require.Contains(t, body, "onebackend_auth_operation_duration_seconds")
	require.Contains(t, body, "onebackend_auth_cache_events_total")
	require.Contains(t, body, `auth_operation="token_refresh"`)
	require.Contains(t, body, `auth_cache="permission"`)
	require.Contains(t, body, `auth_cache_result="miss"`)
}

type assertErr struct{}

func (assertErr) Error() string {
	return "assert error"
}
