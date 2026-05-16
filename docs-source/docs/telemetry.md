---
sidebar_position: 6
tags:
  - telemetry
---

# Telemetry

1Backend uses the Go SDK package `github.com/1backend/1backend/sdk/go/telemetry` for OpenTelemetry metrics and optional OTLP traces.

## 1Backend server

The built-in 1Backend server initializes telemetry by default at process startup, before services and datastores are constructed. No extra configuration is required. It exposes Prometheus-format metrics at `/metrics` by default. Set `OB_OTEL_METRICS_PATH` to change that route, or `OB_OTEL_DISABLED=true` to disable telemetry.

The server records HTTP request counts, response times, response sizes, 4xx/5xx error counts, datastore operation timings, SQL statement timings, automatic-index state, auth helper activity, and service-to-service proxy hops. If `OB_OTEL_TRACES=true` or an OTLP trace endpoint is configured through the standard `OTEL_EXPORTER_OTLP_*` variables, traces are exported through OTLP HTTP.

Every backend service call that flows through 1Backend's service router is represented as a proxy span and receives W3C trace context on the outbound request. Go services built with the 1Backend SDK also use an instrumented HTTP client, so service handler spans, SDK client spans, proxy routing spans, target service handler spans, and datastore spans can appear in one trace. Send OTLP traces to an OpenTelemetry Collector or Grafana Tempo, then view them in Grafana.

## Consuming services

Existing services that already use `infra.NewDataStoreFactory(...)` do not need to change their datastore construction. That factory remains supported and automatically wraps created stores with datastore telemetry. To expose in-process HTTP and datastore metrics, add one telemetry setup call during startup and instrument the service router after routes are registered:

```go
shutdown, metricsPath, err := boot.SetupServiceTelemetry(context.Background(), "basic-svc")
if err != nil {
    return err
}
defer shutdown(context.Background())

dataStoreFactory, err := infra.NewDataStoreFactory(infra.DataStoreConfig{})
if err != nil {
    return err
}

_, err = dataStoreFactory.Create("basicSvcPets", &basic.Pet{})
if err != nil {
    return err
}

router := mux.NewRouter()
// Register service routes on router...

metricsRoute := boot.InstrumentServiceRouter(router, "basic-svc", metricsPath)
log.Println("metrics exposed at", metricsRoute)
```

New services can use `boot.NewOptions(serviceName, ...)` as a convenience. The default options initialize telemetry the first time `LoadEnvars` runs, which happens automatically when a service creates a datastore through `options.NewDataStoreFactory()`.

```go
options := boot.NewOptions("basic-svc", boot.WithSelfUrl(selfURL))
defer options.ShutdownTelemetry(context.Background())

svc, err := basic.NewService(options)
if err != nil {
    return err
}

metricsRoute := options.InstrumentRouter(svc.Router, "", "")
log.Println("metrics exposed at", metricsRoute)
```

With the default metrics path, a service named `basic-svc` exposes `/basic-svc/metrics`, matching the usual `/service-name/endpoint` routing style. The 1Backend proxy can route that endpoint like any other service endpoint, while the service still owns the in-process metrics for its handlers and datastore calls.

Services that use `infra.NewDataStoreFactory(...)` or `options.NewDataStoreFactory()` automatically get datastore instrumentation. Services with a custom datastore can wrap it with `telemetry.InstrumentDataStore(...)`.

### Overrides

Use `boot.WithTelemetryDisabled()` or `OB_OTEL_DISABLED=true` to turn telemetry off for a service using `boot.NewOptions`. Use `boot.WithTelemetryConfig(telemetry.Config{...})` to override the service name, metrics path, version, or disabled flag. Existing services that construct `&boot.Options{}` directly can still call `SetupTelemetry(ctx, serviceName)` explicitly.
