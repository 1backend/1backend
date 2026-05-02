package boot

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestOptionsTelemetryInstrumentsServiceRouter(t *testing.T) {
	options := &Options{}
	_, metricsPath, err := options.SetupTelemetry(t.Context(), "boot-test-svc")
	require.NoError(t, err)

	router := mux.NewRouter()
	metricsRoute := options.InstrumentRouter(router, "boot-test-svc", metricsPath)
	require.Equal(t, "/boot-test-svc/metrics", metricsRoute)

	router.HandleFunc("/boot-test-svc/ping", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("pong"))
	}).Methods(http.MethodGet)

	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/boot-test-svc/ping")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.NoError(t, resp.Body.Close())

	resp, err = http.Get(server.URL + metricsRoute)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Contains(t, string(body), "onebackend_http_server_requests_total")
	require.Contains(t, string(body), `service_name="boot-test-svc"`)
	require.Contains(t, string(body), `http_route="/boot-test-svc/ping"`)
}
