/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func recordDatastoreOperation(ctx context.Context, backend, table, operation string, started time.Time, err error, extra ...attribute.KeyValue) {
	ensureInstruments()
	if ctx == nil {
		ctx = context.Background()
	}

	attrs := []attribute.KeyValue{
		attribute.String("datastore.backend", normalizeBackend(backend)),
		attribute.String("db.collection.name", table),
		attribute.String("db.operation.name", operation),
	}
	attrs = append(attrs, extra...)

	datastoreOperations.Add(ctx, 1, metric.WithAttributes(attrs...))
	datastoreOperationDuration.Record(ctx, time.Since(started).Seconds(), metric.WithAttributes(attrs...))
	if err != nil {
		datastoreOperationErrors.Add(ctx, 1, metric.WithAttributes(append(attrs, errorAttributes(err)...)...))
	}
}
