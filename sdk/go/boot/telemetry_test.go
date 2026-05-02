package boot

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1backend/1backend/sdk/go/telemetry"
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

func TestNewOptionsLoadsTelemetryByDefault(t *testing.T) {
	var setupCalls int
	var shutdownCalls int
	var gotConfig telemetry.Config
	options := NewOptions("default-telemetry-svc", WithTelemetrySetup(
		func(_ context.Context, cfg telemetry.Config) (telemetry.ShutdownFunc, string, error) {
			setupCalls++
			gotConfig = cfg
			return func(context.Context) error {
				shutdownCalls++
				return nil
			}, "/metrics", nil
		},
	))

	require.NoError(t, options.LoadEnvars())
	require.Equal(t, 1, setupCalls)
	require.Equal(t, "default-telemetry-svc", gotConfig.ServiceName)
	require.Equal(t, "/metrics", options.TelemetryMetricsPath())

	router := mux.NewRouter()
	require.Equal(t, "/default-telemetry-svc/metrics", options.InstrumentRouter(router, "", ""))
	require.Equal(t, 1, setupCalls)

	require.NoError(t, options.ShutdownTelemetry(t.Context()))
	require.Equal(t, 1, shutdownCalls)
}

func TestWithTelemetryDisabledKeepsDefaultSetupOverridable(t *testing.T) {
	var gotConfig telemetry.Config
	options := NewOptions(
		"disabled-telemetry-svc",
		WithTelemetryDisabled(),
		WithTelemetrySetup(func(_ context.Context, cfg telemetry.Config) (telemetry.ShutdownFunc, string, error) {
			gotConfig = cfg
			return func(context.Context) error { return nil }, "", nil
		}),
	)

	require.NoError(t, options.LoadEnvars())
	require.True(t, gotConfig.Disabled)
	require.Empty(t, options.TelemetryMetricsPath())

	router := mux.NewRouter()
	require.Empty(t, options.InstrumentRouter(router, "", ""))
}

func TestInstrumentRouterUsesTelemetryConfigServiceName(t *testing.T) {
	options := &Options{
		Telemetry: telemetry.Config{
			ServiceName: "config-name-svc",
		},
		TelemetrySetup: func(_ context.Context, cfg telemetry.Config) (telemetry.ShutdownFunc, string, error) {
			require.Equal(t, "config-name-svc", cfg.ServiceName)
			return func(context.Context) error { return nil }, "/metrics", nil
		},
	}

	router := mux.NewRouter()
	require.Equal(t, "/config-name-svc/metrics", options.InstrumentRouter(router, "", ""))
}

func TestSetupTelemetryStillWorksAfterBareOptionsLoad(t *testing.T) {
	var setupCalls int
	options := &Options{
		TelemetrySetup: func(_ context.Context, cfg telemetry.Config) (telemetry.ShutdownFunc, string, error) {
			setupCalls++
			require.Equal(t, "late-telemetry-svc", cfg.ServiceName)
			return func(context.Context) error { return nil }, "/metrics", nil
		},
	}

	require.NoError(t, options.LoadEnvars())
	require.Zero(t, setupCalls)

	_, metricsPath, err := options.SetupTelemetry(t.Context(), "late-telemetry-svc")
	require.NoError(t, err)
	require.Equal(t, "/metrics", metricsPath)
	require.Equal(t, 1, setupCalls)
}
