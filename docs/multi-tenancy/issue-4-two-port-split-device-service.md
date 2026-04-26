# Issue 4: Two-Port Split — Device Service

## Instructions for the implementing agent

Read this entire document before writing any code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**. This is a structural change across multiple files; understanding the full picture before starting is important.

---

## Problem

`GetSensors` and `GetSensorMetrics` on device-service are currently on the same router as all other device endpoints. Once Issue 3 adds admin role checks to several handlers, we need `GetSensors` and `GetSensorMetrics` to remain accessible to other internal services (specifically context-service) **without requiring auth**.

The solution is a **two-port split**: run a second HTTP server on an internal port (9090) that only exposes the endpoints internal services need. The external port (8080) keeps all the user-facing, auth-gated endpoints. Docker container networking ensures the internal port is only reachable by other containers, not from outside the Docker network.

---

## Architecture overview

```
External (port 8080)         Internal (port 9090)
────────────────────         ────────────────────
All existing routes          GET /api/v1/device/sensors
+ auth enforcement           GET /api/v1/device/sensors/{eui}/metrics
```

Context-service calls the internal URL directly (e.g. `http://device-service:9090`) — not through the API gateway.

---

## Files to change

### 1. `services/device-service/internal/config/config.go`

Add `InternalAddr string` to the `Config` struct.

In `Load()`, add:
```go
InternalAddr: ":" + env.Get("INTERNAL_PORT", "9090"),
```

### 2. `services/device-service/internal/server/router.go`

Add a new function `NewInternalRouter` that registers only the two endpoints context-service needs:

```go
func NewInternalRouter(
    sensorSvc domain.SensorService,
    sensorMetricSvc domain.SensorMetricService,
) *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET "+SENSOR_ROUTE, handlers.GetSensors(sensorSvc))
    mux.HandleFunc("GET "+SENSOR_METRICS_ROUTE, handlers.GetSensorMetrics(sensorMetricSvc))
    return mux
}
```

Use the existing route constants from `consts.go` (or wherever `SENSOR_ROUTE` and `SENSOR_METRICS_ROUTE` are defined — check the file). Do not duplicate the string literals.

### 3. `services/device-service/internal/server/server.go`

Run a second HTTP server alongside the existing one. After building the main `mux` and `server`, add:

```go
internalMux := NewInternalRouter(sensorSvc, sensorMetricSvc)
internalServer := &http.Server{
    Addr:              cfg.InternalAddr,
    Handler:           internalMux,
    ReadHeaderTimeout: 5 * time.Second,
    IdleTimeout:       120 * time.Second,
}
go func() {
    slog.Info("device-service internal server listening", "addr", cfg.InternalAddr)
    if err := internalServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
        slog.Error("internal server error", "err", err)
    }
}()
```

Shut down the internal server alongside the main one in the existing graceful shutdown block. Read the existing `Run()` function carefully before modifying it — you must not break the existing shutdown logic.

### 4. `docker-compose.yml`

In the `device-service` service block, add the internal port env var:

```yaml
environment:
  - GO_ENV=development
  - ENABLE_SWAGGER=true
  - INTERNAL_PORT=9090   # <-- add this
```

Do **not** add a `ports:` mapping for 9090. The internal port must not be exposed to the Docker host — only other containers can reach it via the Docker internal network.

### 5. `docker-compose.prod.yml`

Apply the same environment variable addition to the `device-service` block in the production compose file.

---

## Context-service changes

Context-service does not currently have a device client. When a device client is added in the future (or as part of this issue if needed), it must target the **internal URL** (`http://device-service:9090`), not the external one.

If context-service already has a device client by the time this issue is implemented, check `services/context-service/internal/clients/` and `services/context-service/internal/config/config.go`.

If a device client needs to be created:
- Add `DeviceSvcInternalURL string` to the context-service `Config` struct.
- In `Load()`: `DeviceSvcInternalURL: env.Get("DEVICE_SERVICE_INTERNAL", "http://device-service:9090")`
- Add `DEVICE_SERVICE_INTERNAL=http://device-service:9090` to context-service's environment in both `docker-compose.yml` and `docker-compose.prod.yml`.
- Create `services/context-service/internal/clients/device_client.go` following the same pattern as `collection_client.go` in the same directory.

**Ask before implementing the device client** — confirm whether this work belongs in this issue or a separate one.

---

## What NOT to change

- Do not add auth checks to `GetSensors` or `GetSensorMetrics` on the external router — they will be removed from there if the external-router version should be admin-gated. Confirm with the team.
- Do not expose port 9090 in `ports:` in docker-compose — it must stay internal-only.
- Do not change the collection-service (that is Issue 5).

---

## Acceptance criteria

- `GET http://device-service:9090/api/v1/device/sensors` (from within Docker network) returns sensor data without auth.
- `GET http://localhost:8083/api/v1/device/sensors` (external port, from host) still requires auth.
- The service starts cleanly with two listeners logged in the startup output.
- Graceful shutdown stops both servers.
