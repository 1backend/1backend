package telemetry

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/1backend/1backend/sdk/go/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	instrumentationName = "github.com/1backend/1backend/server/internal/telemetry"
	defaultServiceName  = "1backend"
	defaultMetricsPath  = "/metrics"
)

type Config struct {
	ServiceName    string
	ServiceVersion string
	MetricsPath    string
	Disabled       bool
}

type ShutdownFunc func(context.Context) error

func Setup(ctx context.Context, cfg Config) (ShutdownFunc, string, error) {
	if cfg.Disabled || strings.EqualFold(os.Getenv("OB_OTEL_DISABLED"), "true") {
		return func(context.Context) error { return nil }, "", nil
	}

	if cfg.ServiceName == "" {
		cfg.ServiceName = envOrDefault("OTEL_SERVICE_NAME", defaultServiceName)
	}
	if cfg.ServiceVersion == "" {
		cfg.ServiceVersion = os.Getenv("OB_VERSION")
	}
	if cfg.MetricsPath == "" {
		cfg.MetricsPath = envOrDefault("OB_OTEL_METRICS_PATH", defaultMetricsPath)
	}
	cfg.MetricsPath = normalizeMetricsPath(cfg.MetricsPath)

	res, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", cfg.ServiceName),
			attribute.String("service.version", cfg.ServiceVersion),
			attribute.String("deployment.environment", os.Getenv("OB_ENV")),
		),
	)
	if err != nil {
		return nil, "", fmt.Errorf("create telemetry resource: %w", err)
	}

	promExporter, err := prometheus.New()
	if err != nil {
		return nil, "", fmt.Errorf("create prometheus exporter: %w", err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(promExporter),
	)
	otel.SetMeterProvider(meterProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	var traceProvider *trace.TracerProvider
	if tracesEnabled() {
		traceExporter, err := otlptracehttp.New(ctx)
		if err != nil {
			return nil, "", fmt.Errorf("create OTLP trace exporter: %w", err)
		}

		traceProvider = trace.NewTracerProvider(
			trace.WithResource(res),
			trace.WithBatcher(traceExporter),
			trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(traceSampleRatio()))),
		)
		otel.SetTracerProvider(traceProvider)
		logger.Info("OpenTelemetry OTLP traces enabled")
	}

	return func(ctx context.Context) error {
		var errs []error
		if traceProvider != nil {
			errs = append(errs, traceProvider.Shutdown(ctx))
		}
		errs = append(errs, meterProvider.Shutdown(ctx))
		return errors.Join(errs...)
	}, cfg.MetricsPath, nil
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func normalizeMetricsPath(path string) string {
	path = strings.TrimSpace(path)
	if path == "" {
		return defaultMetricsPath
	}
	if !strings.HasPrefix(path, "/") {
		return "/" + path
	}
	return path
}

func tracesEnabled() bool {
	if endpoint := os.Getenv("OB_OTEL_EXPORTER_OTLP_ENDPOINT"); endpoint != "" && os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") == "" {
		_ = os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", endpoint)
	}

	if strings.EqualFold(os.Getenv("OB_OTEL_TRACES"), "true") {
		return true
	}
	if strings.Contains(strings.ToLower(os.Getenv("OTEL_TRACES_EXPORTER")), "otlp") {
		return true
	}
	return os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT") != "" ||
		os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT") != ""
}

func traceSampleRatio() float64 {
	raw := os.Getenv("OB_OTEL_TRACE_SAMPLE_RATIO")
	if raw == "" {
		raw = os.Getenv("OTEL_TRACES_SAMPLER_ARG")
	}
	if raw == "" {
		return 1
	}

	ratio, err := strconv.ParseFloat(raw, 64)
	if err != nil || ratio < 0 || ratio > 1 {
		return 1
	}
	return ratio
}
