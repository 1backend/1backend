package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestClientFactoryUsesTelemetryHTTPClient(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	defer func() {
		require.NoError(t, provider.Shutdown(t.Context()))
	}()

	var gotTraceparent string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotTraceparent = r.Header.Get("Traceparent")
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	ctx, span := otel.Tracer("test").Start(t.Context(), "service-handler")
	defer span.End()

	apiClient := NewApiClientFactory(server.URL).Client()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/user-svc/public-key", nil)
	require.NoError(t, err)

	resp, err := apiClient.GetConfig().HTTPClient.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	require.NotEmpty(t, gotTraceparent)

	span.End()
	spans := recorder.Ended()
	require.Len(t, spans, 2)
	require.Equal(t, "HTTP GET", spans[0].Name())
	require.Equal(t, "service-handler", spans[1].Name())
}
