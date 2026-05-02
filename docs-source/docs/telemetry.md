---
sidebar_position: 6
tags:
  - telemetry
---

# Telemetry

1Backend uses the Go SDK package `github.com/1backend/1backend/sdk/go/telemetry` for OpenTelemetry metrics and optional OTLP traces.

## 1Backend server

The built-in 1Backend server initializes telemetry at process startup, before services and datastores are constructed. It exposes Prometheus-format metrics at `/metrics` by default. Set `OB_OTEL_METRICS_PATH` to change that route, or `OB_OTEL_DISABLED=true` to disable telemetry.

The server records HTTP request counts, response times, response sizes, 4xx/5xx error counts, datastore operation timings, SQL statement timings, and automatic-index state. If `OB_OTEL_TRACES=true` or an OTLP trace endpoint is configured through the standard `OTEL_EXPORTER_OTLP_*` variables, traces are exported through OTLP HTTP.

## Consuming services

External services should initialize telemetry once per process before constructing datastores, then instrument their router after routes are registered:

```go
options := &boot.Options{SelfUrl: selfURL}

shutdown, metricsPath, err := options.SetupTelemetry(context.Background(), "basic-svc")
if err != nil {
    return err
}
defer shutdown(context.Background())

svc, err := basic.NewService(options)
if err != nil {
    return err
}

metricsRoute := options.InstrumentRouter(svc.Router, "basic-svc", metricsPath)
log.Println("metrics exposed at", metricsRoute)
```

With the default metrics path, a service named `basic-svc` exposes `/basic-svc/metrics`, matching the usual `/service-name/endpoint` routing style. The 1Backend proxy can route that endpoint like any other service endpoint, while the service still owns the in-process metrics for its handlers and datastore calls.

Services that use `options.NewDataStoreFactory()` or `infra.NewDataStoreFactory(...)` automatically get datastore instrumentation. Services with a custom datastore can wrap it with `telemetry.InstrumentDataStore(...)`.
