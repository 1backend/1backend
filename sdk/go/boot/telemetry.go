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
	if serviceName != "" {
		o.ServiceName = serviceName
	}
	if err := o.loadEnvars(); err != nil {
		return nil, "", err
	}
	return o.ensureTelemetry(ctx, true)
}

// InstrumentRouter applies HTTP telemetry middleware and registers the
// service metrics route. With the default metrics path, serviceName "basic-svc"
// is exposed at /basic-svc/metrics.
func (o *Options) InstrumentRouter(router *mux.Router, serviceName, metricsPath string) string {
	if router == nil {
		return ""
	}
	if serviceName == "" {
		serviceName = o.ServiceName
	}
	if serviceName == "" {
		serviceName = o.Telemetry.ServiceName
	}
	if !o.telemetryInitialized && serviceName != "" {
		_, _, _ = o.ensureTelemetry(context.Background(), false)
	}
	if metricsPath == "" {
		metricsPath = o.telemetryMetricsPath
	}
	if metricsPath == "" {
		return ""
	}

	metricsRoute := telemetry.ServiceMetricsPath(serviceName, metricsPath)
	router.Use(telemetry.HTTPMiddleware(serviceName))
	telemetry.RegisterMetricsRoute(router, metricsRoute)
	return metricsRoute
}

// TelemetryMetricsPath returns the process metrics path chosen during
// telemetry setup. It is empty when telemetry is disabled or not configured.
func (o *Options) TelemetryMetricsPath() string {
	return o.telemetryMetricsPath
}

// ShutdownTelemetry flushes and closes telemetry providers configured through
// boot options. It is safe to call when telemetry is disabled.
func (o *Options) ShutdownTelemetry(ctx context.Context) error {
	if o.telemetryShutdown == nil {
		return nil
	}
	return o.telemetryShutdown(ctx)
}

func (o *Options) ensureTelemetry(ctx context.Context, force bool) (telemetry.ShutdownFunc, string, error) {
	if o.telemetryInitialized {
		return o.telemetryShutdown, o.telemetryMetricsPath, nil
	}

	cfg := o.Telemetry
	if cfg.ServiceName == "" {
		cfg.ServiceName = o.ServiceName
	}
	if !force && cfg.ServiceName == "" {
		return func(context.Context) error { return nil }, "", nil
	}

	setup := o.TelemetrySetup
	if setup == nil {
		setup = telemetry.Setup
	}

	shutdown, metricsPath, err := setup(ctx, cfg)
	if err != nil {
		return nil, "", err
	}
	if shutdown == nil {
		shutdown = func(context.Context) error { return nil }
	}

	o.Telemetry = cfg
	o.telemetryShutdown = shutdown
	o.telemetryMetricsPath = metricsPath
	o.telemetryInitialized = true
	return shutdown, metricsPath, nil
}
