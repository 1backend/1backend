/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package telemetry

import (
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const instrumentationName = "github.com/1backend/1backend/sdk/go/telemetry"

var (
	instrumentOnce sync.Once

	datastoreOperations        metric.Int64Counter
	datastoreOperationErrors   metric.Int64Counter
	datastoreOperationDuration metric.Float64Histogram

	sqlStatements        metric.Int64Counter
	sqlStatementErrors   metric.Int64Counter
	sqlStatementDuration metric.Float64Histogram

	autoIndexSupported metric.Int64ObservableGauge
	autoIndexShapes    metric.Int64ObservableGauge
	autoIndexInfo      metric.Int64ObservableGauge
)

func ensureInstruments() {
	instrumentOnce.Do(func() {
		meter := otel.Meter(instrumentationName)

		datastoreOperations, _ = meter.Int64Counter(
			"onebackend.datastore.operations",
			metric.WithDescription("Number of datastore operations executed."),
			metric.WithUnit("{operation}"),
		)
		datastoreOperationErrors, _ = meter.Int64Counter(
			"onebackend.datastore.operation.errors",
			metric.WithDescription("Number of datastore operations that returned an error."),
			metric.WithUnit("{error}"),
		)
		datastoreOperationDuration, _ = meter.Float64Histogram(
			"onebackend.datastore.operation.duration",
			metric.WithDescription("Duration of datastore operations."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5),
		)

		sqlStatements, _ = meter.Int64Counter(
			"onebackend.sql.statements",
			metric.WithDescription("Number of SQL statements executed by datastores."),
			metric.WithUnit("{statement}"),
		)
		sqlStatementErrors, _ = meter.Int64Counter(
			"onebackend.sql.statement.errors",
			metric.WithDescription("Number of SQL statements that returned an error."),
			metric.WithUnit("{error}"),
		)
		sqlStatementDuration, _ = meter.Float64Histogram(
			"onebackend.sql.statement.duration",
			metric.WithDescription("Duration of SQL statements executed by datastores."),
			metric.WithUnit("s"),
			metric.WithExplicitBucketBoundaries(0.0005, 0.001, 0.0025, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5),
		)

		autoIndexSupported, _ = meter.Int64ObservableGauge(
			"onebackend.datastore.autoindex.supported",
			metric.WithDescription("Whether automatic indexing is supported for a datastore."),
			metric.WithUnit("1"),
		)
		autoIndexShapes, _ = meter.Int64ObservableGauge(
			"onebackend.datastore.autoindex.shape.hits",
			metric.WithDescription("Observed query shape hit counts used by the auto-index planner."),
			metric.WithUnit("{hit}"),
		)
		autoIndexInfo, _ = meter.Int64ObservableGauge(
			"onebackend.datastore.autoindex.info",
			metric.WithDescription("Auto-index entries currently known by the datastore. Value is always 1; labels describe the index."),
			metric.WithUnit("1"),
		)

		_, _ = meter.RegisterCallback(
			observeAutoIndexes,
			autoIndexSupported,
			autoIndexShapes,
			autoIndexInfo,
		)
	})
}
