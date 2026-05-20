# Issue 4: Two-Port Split — Device Service

## Instructions for the implementing agent

Read this entire document before writing any code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**. This is a structural change across multiple files; understanding the full picture before starting is important.

---

## Problem

`GetSensorMetrics` on device-service is needed by two callers with conflicting auth requirements:

- **Frontend (admin)** calls it through the API gateway with a valid admin JWT — it should be admin-gated.
- **Context-service** calls it directly (service-to-service, bypassing the gateway) with no JWT — it cannot send an admin JWT.

The solution is a **two-port split**: run a second HTTP server on an internal port (9090) that exposes `GetSensorMetrics` without auth. The external port (8080) keeps `GetSensorMetrics` with an admin guard for the frontend. Docker container networking ensures the internal port is only reachable by other containers, not from outside the Docker network.

`GetSensors` does **not** need the internal router. It is already company-scoped via `authctx` and has no admin guard — context-service can call it on port 8080 by forwarding `X-Auth-Company-Id` (that is a context-service change, handled separately by the context-service team).

---

## Architecture overview

```
External (port 8080)                    Internal (port 9090)
────────────────────                    ────────────────────
All existing routes                     GET /api/v1/device/sensors/{eui}/metrics
GetSensorMetrics → admin-gated          (no auth, internal only)
GetSensors       → company-scoped
```

Context-service calls metrics on the internal URL (`http://device-service:9090`) and sensors on the external URL (`http://device-service:8080`) with `X-Auth-Company-Id` forwarded.

---

## Files to change

### 1. `services/device-service/internal/handlers/sensor_metric_handlers.go`

Add `authctx.FromRequest` + `IsAdmin()` guard to `GetSensorMetrics` on the **external** router. This is the same pattern used in Issue 3.

```go
auth, err := authctx.FromRequest(r)
if err != nil {
    json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
    return
}
if !auth.IsAdmin() {
    json.HandleError(w, http.StatusForbidden, nil, "forbidden")
    return
}
```

Also add `@Failure 401` and `@Failure 403` to the Swagger comment block.

> The handler registered on the **internal** router is the same `GetSensorMetrics` function — it just won't have the guard because it is registered on a different mux that is not exposed externally. No separate handler function is needed.

### 2. `services/device-service/internal/config/config.go`

Add `InternalAddr string` to the `Config` struct.

In `Load()`, add:
```go
InternalAddr: ":" + env.Get("INTERNAL_PORT", "9090"),
```

### 3. `services/device-service/internal/server/router.go`

Add a new function `NewInternalRouter` that registers only the metrics endpoint:

```go
func NewInternalRouter(sensorMetricSvc domain.SensorMetricService) *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET "+SENSOR_METRICS_ROUTE, handlers.GetSensorMetrics(sensorMetricSvc))
    return mux
}
```

Use the existing route constant from `consts.go`. Do not duplicate the string literal.

### 4. `services/device-service/internal/server/server.go`

Run a second HTTP server alongside the existing one. After building the main `mux` and `server`, add:

```go
internalMux := NewInternalRouter(sensorMetricSvc)
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

### 5. `docker-compose.yml`

In the `device-service` service block, add the internal port env var:

```yaml
environment:
  - GO_ENV=development
  - ENABLE_SWAGGER=true
  - INTERNAL_PORT=9090   # <-- add this
```

Do **not** add a `ports:` mapping for 9090. The internal port must not be exposed to the Docker host — only other containers can reach it via the Docker internal network.

### 6. `docker-compose.prod.yml`

Apply the same environment variable addition to the `device-service` block in the production compose file.

---

## Context-service changes (handled by context-service team)

Context-service already has a device client at `services/context-service/internal/clients/device_client.go` that calls two device-service endpoints:

1. `GET /api/v1/device/sensors?production_resource_id={id}` — stays on port 8080. The context-service team must forward `X-Auth-Company-Id` from the incoming request context when making this call, so device-service returns only that company's sensors.

2. `GET /api/v1/device/sensors/{eui}/metrics` — must move to port 9090. The context-service team must:
   - Add `DeviceSvcInternalURL string` to the context-service `Config` struct
   - In `Load()`: `DeviceSvcInternalURL: env.Get("DEVICE_SERVICE_INTERNAL", "http://device-service:9090")`
   - Update the device client to use the internal URL for the metrics call
   - Add `DEVICE_SERVICE_INTERNAL=http://device-service:9090` to context-service's environment in both `docker-compose.yml` and `docker-compose.prod.yml`

---

## What NOT to change

- `GetSensors` does not go on the internal router — it is already company-scoped and does not require an admin guard.
- Do not expose port 9090 in `ports:` in docker-compose.
- Do not change the collection-service (that is Issue 5).

---

## Acceptance criteria

- `GET http://device-service:9090/api/v1/device/sensors/{eui}/metrics` (from within Docker network) returns metrics without auth.
- `GET /api/v1/device/sensors/{eui}/metrics` via the API gateway returns 403 for a non-admin JWT and 401 with no JWT.
- `GET /api/v1/device/sensors` via the API gateway still works for regular users.
- The service starts cleanly with two listeners logged in the startup output.
- Graceful shutdown stops both servers.
