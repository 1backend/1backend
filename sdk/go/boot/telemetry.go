/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package boot

import (
	"context"

	"github.com/1backend/1backend/sdk/go/telemetry"
	"github.com/gorilla/mux"
)

// SetupTelemetry initializes the process-wide OpenTelemetry provider for a
// service. Call this before constructing datastores when you want startup
// datastore work to be visible in metrics.
func (o *Options) SetupTelemetry(ctx context.Context, serviceName string) (telemetry.ShutdownFunc, string, error) {
	if err := o.LoadEnvars(); err != nil {
		return nil, "", err
	}
	return telemetry.Setup(ctx, telemetry.Config{
		ServiceName: serviceName,
	})
}

// InstrumentRouter applies HTTP telemetry middleware and registers the
// service metrics route. With the default metrics path, serviceName "basic-svc"
// is exposed at /basic-svc/metrics.
func (o *Options) InstrumentRouter(router *mux.Router, serviceName, metricsPath string) string {
	if router == nil || metricsPath == "" {
		return metricsPath
	}

	metricsRoute := telemetry.ServiceMetricsPath(serviceName, metricsPath)
	router.Use(telemetry.HTTPMiddleware(serviceName))
	telemetry.RegisterMetricsRoute(router, metricsRoute)
	return metricsRoute
}
