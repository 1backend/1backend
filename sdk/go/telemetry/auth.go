package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	AuthOperationTokenRefresh    = "token_refresh"
	AuthOperationPermissionCheck = "permission_check"
	AuthOperationTokenExchange   = "token_exchange"
	AuthOperationPublicKeyFetch  = "public_key_fetch"

	AuthCacheTokenRefresh  = "token_refresh"
	AuthCachePermission    = "permission"
	AuthCacheTokenExchange = "token_exchange"
)

func RecordAuthOperation(ctx context.Context, operation, result, source string, started time.Time, err error) {
	ensureInstruments()
	if ctx == nil {
		ctx = context.Background()
	}
	if result == "" {
		result = "unknown"
	}
	if source == "" {
		source = "unknown"
	}

	attrs := []attribute.KeyValue{
		attribute.String("auth.operation", operation),
		attribute.String("auth.result", result),
		attribute.String("auth.source", source),
	}

	authOperations.Add(ctx, 1, metric.WithAttributes(attrs...))
	authOperationDuration.Record(ctx, time.Since(started).Seconds(), metric.WithAttributes(attrs...))
	if err != nil {
		authOperationErrors.Add(ctx, 1, metric.WithAttributes(append(attrs, errorAttributes(err)...)...))
	}
}

func RecordAuthCache(ctx context.Context, cache, result string) {
	ensureInstruments()
	if ctx == nil {
		ctx = context.Background()
	}
	if result == "" {
		result = "unknown"
	}

	authCacheEvents.Add(ctx, 1, metric.WithAttributes(
		attribute.String("auth.cache", cache),
		attribute.String("auth.cache.result", result),
	))
}
