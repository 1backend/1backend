package telemetry

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func HTTPClientTransport(serviceName string, base http.RoundTripper) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	if serviceName == "" {
		serviceName = defaultServiceName
	}

	return otelhttp.NewTransport(
		base,
		otelhttp.WithSpanNameFormatter(func(_ string, r *http.Request) string {
			if r == nil {
				return "HTTP"
			}
			return "HTTP " + r.Method
		}),
		otelhttp.WithSpanOptions(
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(attribute.String("service.name", serviceName)),
		),
	)
}

func HTTPClient(serviceName string, base *http.Client) *http.Client {
	if base == nil {
		base = http.DefaultClient
	}

	clone := *base
	clone.Transport = HTTPClientTransport(serviceName, base.Transport)
	return &clone
}
