package fileservice

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const fileInstrumentationName = "github.com/1backend/1backend/server/internal/services/file"

var (
	fileTelemetryOnce sync.Once

	fileDownloadServeRequests metric.Int64Counter
	fileDownloadServeBytes    metric.Int64Histogram
	fileDownloadFetches       metric.Int64Counter
	fileDownloadFetchDuration metric.Float64Histogram
	fileDownloadRecovery      metric.Int64Counter
	fileDownloadGCSBackfill   metric.Int64Counter
	fileDownloadGCSDuration   metric.Float64Histogram
	fileDownloadStorageOps    metric.Int64Counter
	fileDownloadStorageBytes  metric.Int64Histogram
	fileDownloadStorageTime   metric.Float64Histogram
)

func ensureFileTelemetry() {
	fileTelemetryOnce.Do(func() {
		meter := otel.Meter(fileInstrumentationName)

		fileDownloadServeRequests, _ = meter.Int64Counter(
			"onebackend.file.download.serve.requests",
			metric.WithDescription("Number of file download serve attempts."),
			metric.WithUnit("{request}"),
		)
		fileDownloadServeBytes, _ = meter.Int64Histogram(
			"onebackend.file.download.serve.bytes",
			metric.WithDescription("File download response size."),
			metric.WithUnit("By"),
			metric.WithExplicitBucketBoundaries(1024, 16*1024, 64*1024, 256*1024, 1024*1024, 8*1024*1024, 64*1024*1024),
		)
		fileDownloadFetches, _ = meter.Int64Counter(
			"onebackend.file.download.fetches",
			metric.WithDescription("Number of upstream internet download fetch attempts."),
			metric.WithUnit("{fetch}"),
		)
		fileDownloadFetchDuration, _ = meter.Float64Histogram(
			"onebackend.file.download.fetch.duration",
			metric.WithDescription("Duration of upstream internet download fetches."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30, 60, 120),
		)
		fileDownloadRecovery, _ = meter.Int64Counter(
			"onebackend.file.download.recovery.attempts",
			metric.WithDescription("Number of file download recovery path attempts."),
			metric.WithUnit("{attempt}"),
		)
		fileDownloadGCSBackfill, _ = meter.Int64Counter(
			"onebackend.file.download.gcs.backfill.checks",
			metric.WithDescription("Sampled checks that completed downloads are present in GCS, including repairs."),
			metric.WithUnit("{check}"),
		)
		fileDownloadGCSDuration, _ = meter.Float64Histogram(
			"onebackend.file.download.gcs.backfill.duration",
			metric.WithDescription("Duration of sampled GCS presence checks and repairs for downloaded files."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30),
		)
		fileDownloadStorageOps, _ = meter.Int64Counter(
			"onebackend.file.download.storage.operations",
			metric.WithDescription("Number of file download storage backend operations."),
			metric.WithUnit("{operation}"),
		)
		fileDownloadStorageBytes, _ = meter.Int64Histogram(
			"onebackend.file.download.storage.bytes",
			metric.WithDescription("Bytes read from or written to the file download storage backend."),
			metric.WithUnit("By"),
			metric.WithExplicitBucketBoundaries(1024, 16*1024, 64*1024, 256*1024, 1024*1024, 8*1024*1024, 64*1024*1024),
		)
		fileDownloadStorageTime, _ = meter.Float64Histogram(
			"onebackend.file.download.storage.operation.duration",
			metric.WithDescription("Duration of file download storage backend operations."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 30),
		)
	})
}

func startFileSpan(ctx context.Context, name string, attrs ...attribute.KeyValue) (context.Context, trace.Span) {
	return otel.Tracer(fileInstrumentationName).Start(ctx, name, trace.WithAttributes(attrs...))
}

func recordFileDownloadServe(ctx context.Context, route string, result string, bytes int64) {
	ensureFileTelemetry()
	attrs := metric.WithAttributes(
		attribute.String("route", route),
		attribute.String("result", result),
	)
	fileDownloadServeRequests.Add(ctx, 1, attrs)
	if bytes >= 0 {
		fileDownloadServeBytes.Record(ctx, bytes, attrs)
	}
}

func recordFileDownloadFetch(ctx context.Context, result string, resumed bool, duration time.Duration) {
	ensureFileTelemetry()
	attrs := metric.WithAttributes(
		attribute.String("result", result),
		attribute.Bool("resumed", resumed),
	)
	fileDownloadFetches.Add(ctx, 1, attrs)
	fileDownloadFetchDuration.Record(ctx, duration.Seconds(), attrs)
}

func recordFileDownloadRecovery(ctx context.Context, path string, result string) {
	ensureFileTelemetry()
	fileDownloadRecovery.Add(ctx, 1, metric.WithAttributes(
		attribute.String("path", path),
		attribute.String("result", result),
	))
}

func recordFileDownloadGCSBackfill(ctx context.Context, result string, duration time.Duration) {
	ensureFileTelemetry()
	attrs := metric.WithAttributes(attribute.String("result", result))
	fileDownloadGCSBackfill.Add(ctx, 1, attrs)
	fileDownloadGCSDuration.Record(ctx, duration.Seconds(), attrs)
}

func recordFileDownloadStorageOperation(
	ctx context.Context,
	operation string,
	backend string,
	result string,
	bytes int64,
	duration time.Duration,
) {
	ensureFileTelemetry()
	attrs := metric.WithAttributes(
		attribute.String("operation", operation),
		attribute.String("backend", backend),
		attribute.String("result", result),
	)
	fileDownloadStorageOps.Add(ctx, 1, attrs)
	fileDownloadStorageTime.Record(ctx, duration.Seconds(), attrs)
	if bytes >= 0 {
		fileDownloadStorageBytes.Record(ctx, bytes, attrs)
	}
}
