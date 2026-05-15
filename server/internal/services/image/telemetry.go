package imageservice

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const imageInstrumentationName = "github.com/1backend/1backend/server/internal/services/image"

var (
	imageTelemetryOnce sync.Once

	imageServeRequests     metric.Int64Counter
	imageTransforms        metric.Int64Counter
	imageTransformDuration metric.Float64Histogram
	imageServedBytes       metric.Int64Histogram
	imageFileOpens         metric.Int64Counter
)

func ensureImageTelemetry() {
	imageTelemetryOnce.Do(func() {
		meter := otel.Meter(imageInstrumentationName)

		imageServeRequests, _ = meter.Int64Counter(
			"onebackend.image.serve.requests",
			metric.WithDescription("Number of image service serve attempts."),
			metric.WithUnit("{request}"),
		)
		imageTransforms, _ = meter.Int64Counter(
			"onebackend.image.transforms",
			metric.WithDescription("Number of image transform attempts."),
			metric.WithUnit("{transform}"),
		)
		imageTransformDuration, _ = meter.Float64Histogram(
			"onebackend.image.transform.duration",
			metric.WithDescription("Duration of image cache miss transforms."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10),
		)
		imageServedBytes, _ = meter.Int64Histogram(
			"onebackend.image.serve.bytes",
			metric.WithDescription("Image service response size."),
			metric.WithUnit("By"),
			metric.WithExplicitBucketBoundaries(1024, 8*1024, 32*1024, 128*1024, 512*1024, 2*1024*1024, 8*1024*1024),
		)
		imageFileOpens, _ = meter.Int64Counter(
			"onebackend.image.file.opens",
			metric.WithDescription("Number of upstream file service opens made by image service."),
			metric.WithUnit("{open}"),
		)
	})
}

func startImageSpan(ctx context.Context, name string, attrs ...attribute.KeyValue) (context.Context, trace.Span) {
	return otel.Tracer(imageInstrumentationName).Start(ctx, name, trace.WithAttributes(attrs...))
}

func recordImageServe(ctx context.Context, source string, cache string, result string, bytes int64, params *imageParams) {
	ensureImageTelemetry()
	attrs := metric.WithAttributes(
		attribute.String("source", source),
		attribute.String("cache", cache),
		attribute.String("result", result),
		attribute.String("width_bucket", imageDimensionBucket(params.Width)),
		attribute.String("height_bucket", imageDimensionBucket(params.Height)),
		attribute.String("requested_format", imageRequestedFormatLabel(params.RequestedFormat)),
	)
	imageServeRequests.Add(ctx, 1, attrs)
	if bytes >= 0 {
		imageServedBytes.Record(ctx, bytes, attrs)
	}
}

func imageDimensionBucket(value int) string {
	switch {
	case value <= 0:
		return "auto"
	case value <= 320:
		return "001-320"
	case value <= 480:
		return "321-480"
	case value <= 720:
		return "481-720"
	case value <= 960:
		return "721-960"
	case value <= 1200:
		return "961-1200"
	case value <= 1600:
		return "1201-1600"
	default:
		return "1601+"
	}
}

func imageRequestedFormatLabel(format string) string {
	if format == "" {
		return "default"
	}
	return format
}

func recordImageTransform(ctx context.Context, source string, result string, duration time.Duration) {
	ensureImageTelemetry()
	attrs := metric.WithAttributes(
		attribute.String("source", source),
		attribute.String("result", result),
	)
	imageTransforms.Add(ctx, 1, attrs)
	imageTransformDuration.Record(ctx, duration.Seconds(), attrs)
}

func recordImageFileOpen(ctx context.Context, source string, result string) {
	ensureImageTelemetry()
	imageFileOpens.Add(ctx, 1, metric.WithAttributes(
		attribute.String("source", source),
		attribute.String("result", result),
	))
}
