package proxyservice

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	openapi "github.com/1backend/1backend/clients/go"
	"github.com/1backend/1backend/sdk/go/telemetry"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestRouteBackendPropagatesTraceContextAndRecordsProxySpan(t *testing.T) {
	recorder := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(recorder))
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	defer func() {
		require.NoError(t, provider.Shutdown(t.Context()))
	}()

	traceparentCh := make(chan string, 1)
	mockBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceparentCh <- r.Header.Get("Traceparent")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success": true}`))
	}))
	defer mockBackend.Close()

	proxy := &ProxyService{
		started: true,
		httpClient: &http.Client{
			Transport: telemetry.HTTPClientTransport("1backend-proxy-test", nil),
		},
	}
	proxy.instanceCache.Store("target-svc", cacheEntry{
		instances: []openapi.RegistrySvcInstance{{
			Url:    mockBackend.URL,
			Status: openapi.InstanceStatusHealthy,
		}},
		expiry: time.Now().Add(time.Minute),
	})

	req := httptest.NewRequest(http.MethodGet, "http://1backend.test/target-svc/traced", nil)
	rr := httptest.NewRecorder()

	status, err := proxy.routeBackend(rr, req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, status)
	require.Equal(t, http.StatusOK, rr.Code)

	select {
	case traceparent := <-traceparentCh:
		require.NotEmpty(t, traceparent)
	case <-time.After(2 * time.Second):
		t.Fatal("mock backend did not receive proxied request")
	}

	var foundProxySpan bool
	for _, span := range recorder.Ended() {
		if span.Name() == "proxy.route_backend" {
			foundProxySpan = true
			require.Equal(t, "1backend", attrValue(span.Attributes(), "onebackend.proxy.service"))
			require.Equal(t, "target-svc", attrValue(span.Attributes(), "onebackend.target_service"))
			require.Equal(t, int64(http.StatusOK), attrValue(span.Attributes(), "http.response.status_code"))
			require.Equal(t, "hit", attrValue(span.Attributes(), "onebackend.proxy.instance_cache"))
		}
	}
	require.True(t, foundProxySpan, "expected proxy.route_backend span")
}

func attrValue(attrs []attribute.KeyValue, key string) any {
	for _, attr := range attrs {
		if string(attr.Key) == key {
			return attr.Value.AsInterface()
		}
	}
	return nil
}
