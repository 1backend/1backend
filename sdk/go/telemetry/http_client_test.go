package telemetry

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

func TestHTTPClientTransportCreatesClientSpanAndInjectsTraceContext(t *testing.T) {
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

	ctx, parent := otel.Tracer("test").Start(t.Context(), "caller")
	defer parent.End()

	client := &http.Client{
		Transport: HTTPClientTransport("caller-svc", nil),
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/target-svc/ping", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	require.NotEmpty(t, gotTraceparent)

	parent.End()
	spans := recorder.Ended()
	require.Len(t, spans, 2)
	require.Equal(t, "HTTP GET", spans[0].Name())
	require.Equal(t, "caller", spans[1].Name())
	require.Equal(t, spans[1].SpanContext().TraceID(), spans[0].SpanContext().TraceID())
	require.Equal(t, spans[1].SpanContext().SpanID(), spans[0].Parent().SpanID())
}
