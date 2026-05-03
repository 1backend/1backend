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
	"strings"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	SQLTargetPrimary     = "primary"
	SQLTargetReadReplica = "read_replica"
)

// RecordSQLStatement records low-level SQL statement metrics without storing
// raw SQL text as a metric label.
func RecordSQLStatement(ctx context.Context, backend, table, statement string, started time.Time, err error) {
	RecordSQLStatementTarget(ctx, backend, table, SQLTargetPrimary, statement, started, err)
}

// RecordSQLStatementTarget records low-level SQL statement metrics tagged with
// the datastore connection target that executed the statement.
func RecordSQLStatementTarget(ctx context.Context, backend, table, target, statement string, started time.Time, err error) {
	ensureInstruments()
	if ctx == nil {
		ctx = context.Background()
	}
	if strings.TrimSpace(target) == "" {
		target = SQLTargetPrimary
	}

	attrs := []attribute.KeyValue{
		attribute.String("db.system", normalizeBackend(backend)),
		attribute.String("db.collection.name", table),
		attribute.String("db.operation.name", SQLStatementOperation(statement)),
		attribute.String("datastore.sql.target", target),
	}

	sqlStatements.Add(ctx, 1, metric.WithAttributes(attrs...))
	sqlStatementDuration.Record(ctx, time.Since(started).Seconds(), metric.WithAttributes(attrs...))
	if err != nil {
		sqlStatementErrors.Add(ctx, 1, metric.WithAttributes(append(attrs, errorAttributes(err)...)...))
	}
}

// SQLStatementOperation returns a low-cardinality operation name from a SQL
// statement, for example "select" or "create.index".
func SQLStatementOperation(statement string) string {
	parts := strings.Fields(strings.TrimSpace(statement))
	if len(parts) == 0 {
		return "unknown"
	}

	first := strings.ToLower(parts[0])
	if len(parts) > 1 {
		switch first {
		case "alter", "create", "drop":
			return first + "." + strings.ToLower(parts[1])
		}
	}

	return first
}
