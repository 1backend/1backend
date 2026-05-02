/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package telemetry

import (
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

var (
	httpInstrumentOnce sync.Once

	httpRequests     metric.Int64Counter
	httpErrors       metric.Int64Counter
	httpDuration     metric.Float64Histogram
	httpResponseSize metric.Int64Histogram
)

func RegisterMetricsRoute(router *mux.Router, path string) {
	if router == nil || strings.EqualFold(os.Getenv("OB_OTEL_DISABLED"), "true") {
		return
	}
	path = normalizeMetricsPath(path)
	router.Handle(path, promhttp.Handler()).Methods(http.MethodGet)
}

// ServiceMetricsPath returns a metrics route that matches the usual
// /service-name/endpoint routing style. The default path for "basic-svc" is
// /basic-svc/metrics; explicit non-default paths are left unchanged.
func ServiceMetricsPath(serviceName, metricsPath string) string {
	metricsPath = normalizeMetricsPath(metricsPath)
	serviceName = strings.Trim(strings.TrimSpace(serviceName), "/")
	if serviceName == "" || metricsPath != defaultMetricsPath {
		return metricsPath
	}
	return "/" + serviceName + metricsPath
}

func HTTPMiddleware(serviceName string) mux.MiddlewareFunc {
	if serviceName == "" {
		serviceName = defaultServiceName
	}
	ensureHTTPInstruments()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			spanName := r.Method + " " + routeTemplate(r)
			ctx, span := otel.Tracer(instrumentationName).Start(
				ctx,
				spanName,
				trace.WithSpanKind(trace.SpanKindServer),
				trace.WithAttributes(spanRequestAttributes(r, serviceName)...),
			)
			defer span.End()

			r = r.WithContext(ctx)
			metrics := httpsnoop.CaptureMetrics(next, w, r)
			statusCode := metrics.Code
			if statusCode == 0 {
				statusCode = http.StatusOK
			}

			route := routeTemplate(r)
			span.SetName(r.Method + " " + route)
			spanAttrs := append(spanRequestAttributes(r, serviceName),
				attribute.String("http.route", route),
				attribute.Int("http.response.status_code", statusCode),
			)
			if metrics.Written >= 0 {
				spanAttrs = append(spanAttrs, attribute.Int64("http.response.body.size", metrics.Written))
			}
			span.SetAttributes(spanAttrs...)
			if statusCode >= http.StatusInternalServerError {
				span.SetStatus(codes.Error, http.StatusText(statusCode))
			}

			metricAttrs := metricRequestAttributes(r, serviceName, route, statusCode)
			httpRequests.Add(ctx, 1, metric.WithAttributes(metricAttrs...))
			httpDuration.Record(ctx, metrics.Duration.Seconds(), metric.WithAttributes(metricAttrs...))
			if metrics.Written >= 0 {
				httpResponseSize.Record(ctx, metrics.Written, metric.WithAttributes(metricAttrs...))
			}
			if statusCode >= http.StatusBadRequest {
				httpErrors.Add(ctx, 1, metric.WithAttributes(append(metricAttrs, errorStatusAttributes(statusCode)...)...))
			}
		})
	}
}

func ensureHTTPInstruments() {
	httpInstrumentOnce.Do(func() {
		meter := otel.Meter(instrumentationName)

		httpRequests, _ = meter.Int64Counter(
			"onebackend.http.server.requests",
			metric.WithDescription("Number of HTTP requests handled by 1Backend."),
			metric.WithUnit("{request}"),
		)
		httpErrors, _ = meter.Int64Counter(
			"onebackend.http.server.errors",
			metric.WithDescription("Number of HTTP requests with a 4xx or 5xx response."),
			metric.WithUnit("{error}"),
		)
		httpDuration, _ = meter.Float64Histogram(
			"onebackend.http.server.request.duration",
			metric.WithDescription("Duration of HTTP requests handled by 1Backend."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10),
		)
		httpResponseSize, _ = meter.Int64Histogram(
			"onebackend.http.server.response.size",
			metric.WithDescription("HTTP response body size written by 1Backend."),
			metric.WithUnit("By"),
		)
	})
}

func metricRequestAttributes(r *http.Request, serviceName, route string, statusCode int) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("service.name", serviceName),
		attribute.String("http.request.method", r.Method),
		attribute.String("http.route", route),
		attribute.Int("http.response.status_code", statusCode),
		attribute.String("network.protocol.name", protocolName(r)),
		attribute.String("network.protocol.version", protocolVersion(r)),
		attribute.String("onebackend.service", oneBackendServiceName(r)),
	}
}

func spanRequestAttributes(r *http.Request, serviceName string) []attribute.KeyValue {
	attrs := []attribute.KeyValue{
		attribute.String("service.name", serviceName),
		attribute.String("http.request.method", r.Method),
		attribute.String("url.scheme", requestScheme(r)),
		attribute.String("url.path", r.URL.Path),
		attribute.String("network.protocol.name", protocolName(r)),
		attribute.String("network.protocol.version", protocolVersion(r)),
		attribute.String("onebackend.service", oneBackendServiceName(r)),
	}

	if r.Host != "" {
		host, port, err := net.SplitHostPort(r.Host)
		if err == nil {
			attrs = append(attrs, attribute.String("server.address", host))
			if p, parseErr := strconv.Atoi(port); parseErr == nil {
				attrs = append(attrs, attribute.Int("server.port", p))
			}
		} else {
			attrs = append(attrs, attribute.String("server.address", r.Host))
		}
	}
	if r.RemoteAddr != "" {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			attrs = append(attrs, attribute.String("client.address", host))
		}
	}
	if userAgent := r.UserAgent(); userAgent != "" {
		attrs = append(attrs, attribute.String("user_agent.original", userAgent))
	}

	return attrs
}

func routeTemplate(r *http.Request) string {
	if route := mux.CurrentRoute(r); route != nil {
		if path, err := route.GetPathTemplate(); err == nil && path != "" {
			return path
		}
	}
	if r.URL != nil && r.URL.Path != "" {
		return r.URL.Path
	}
	return "unknown"
}

func oneBackendServiceName(r *http.Request) string {
	route := routeTemplate(r)
	parts := strings.Split(strings.Trim(route, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		return "root"
	}
	return parts[0]
}

func requestScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	if scheme := r.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	return "http"
}

func protocolName(r *http.Request) string {
	if strings.HasPrefix(strings.ToLower(r.Proto), "http/") {
		return "http"
	}
	return strings.ToLower(r.Proto)
}

func protocolVersion(r *http.Request) string {
	parts := strings.SplitN(r.Proto, "/", 2)
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func errorStatusAttributes(statusCode int) []attribute.KeyValue {
	statusClass := "4xx"
	if statusCode >= http.StatusInternalServerError {
		statusClass = "5xx"
	}
	return []attribute.KeyValue{
		attribute.String("error.type", statusClass),
	}
}
