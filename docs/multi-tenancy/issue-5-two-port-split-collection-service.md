# Issue 5: Two-Port Split — Collection Service

## Instructions for the implementing agent

Read this entire document before writing any code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**.

---

## Problem

`GET /measurements` (by time range) on collection-service is called by context-service internally to fetch raw sensor readings. If we add auth enforcement to this handler (so the frontend cannot query arbitrary time ranges without company scoping), we would break context-service, which calls it without auth.

The solution is the same two-port pattern used in Issue 4: run a second HTTP server on an internal port (9090) that exposes `GET /measurements` without auth. The external port (8080) keeps the user-facing endpoints (`/latest`).

Issue 1 removes `GET /measurements` from the API gateway's proxy list, so it will no longer be reachable from outside after that change. This issue ensures context-service can still reach it via the internal port.

---

## Architecture overview

```
External (port 8080)              Internal (port 9090)
────────────────────              ────────────────────
GET /latest                       GET /api/v1/collection/measurements
(user-facing, auth-gated)         (internal, no auth required)
```

Context-service calls the internal URL directly (`http://collection-service:9090`), not through the API gateway.

---

## Files to change

### 1. `services/collection-service/internal/config/config.go`

Add `InternalAddr string` to the `Config` struct.

In `Load()`, add:
```go
InternalAddr: ":" + env.Get("INTERNAL_PORT", "9090"),
```

### 2. `services/collection-service/internal/server/router.go`

Add a new `NewInternalRouter` function that registers only the measurements-by-time-range handler:

```go
func NewInternalRouter(svc domain.MeasurementService) *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("GET "+MEASUREMENTS_BY_TIME, handlers.HandleMeasurementsByTimeRange(svc))
    return mux
}
```

Use the existing `MEASUREMENTS_BY_TIME` constant (check `consts.go` or wherever constants are defined in collection-service). Do not duplicate the string literal.

### 3. `services/collection-service/internal/server/server.go`

Run a second HTTP server alongside the existing one. After the existing `mux` and `server` are set up, add:

```go
internalMux := NewInternalRouter(svc)
internalServer := &http.Server{
    Addr:              cfg.InternalAddr,
    Handler:           internalMux,
    ReadHeaderTimeout: 5 * time.Second,
    IdleTimeout:       120 * time.Second,
}
go func() {
    slog.Info("collection-service internal server listening", "addr", cfg.InternalAddr)
    if err := internalServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
        slog.Error("internal server error", "err", err)
    }
}()
```

Shut down the internal server in the existing graceful shutdown block. Read `Run()` carefully before modifying — the shutdown logic already handles `serverErrors` and `shutdown` channels; integrate without breaking that structure.

### 4. `services/context-service/internal/config/config.go`

The context-service `Config` struct currently has `CollectionSvcURL`. This URL is used by `services/context-service/internal/clients/collection_client.go` to call `GET /measurements`.

Change the config to use a dedicated internal URL:

```go
CollectionSvcInternalURL: env.Get("COLLECTION_SERVICE_INTERNAL", "http://collection-service:9090"),
```

You may keep `CollectionSvcURL` for other purposes or rename it — confirm whether collection-service is called at any other path from context-service before removing the existing field.

### 5. `services/context-service/internal/clients/collection_client.go`

Update `NewCollectionClient` to accept the internal base URL (already done if you renamed the config field). Verify the URL the client builds points to port 9090 in the default case.

### 6. `services/context-service/internal/server/server.go`

Update the line that constructs the collection client to pass `cfg.CollectionSvcInternalURL` (or whatever you named it in step 4).

### 7. `docker-compose.yml`

**collection-service** — add the internal port env var:
```yaml
environment:
  - GO_ENV=development
  - ENABLE_SWAGGER=true
  - INTERNAL_PORT=9090   # <-- add
```

Do **not** add a `ports:` mapping for 9090.

**context-service** — add the internal URL env var:
```yaml
environment:
  - GO_ENV=development
  - COLLECTION_SERVICE_INTERNAL=http://collection-service:9090   # <-- add
```

### 8. `docker-compose.prod.yml`

Apply the same additions to both `collection-service` and `context-service` in the production compose file.

---

## What NOT to change

- `GET /latest` stays on the external router — it is user-facing and will get auth enforcement separately.
- `PostTenantMapping` (`POST /company-config`) stays on the external router but is removed from the gateway in Issue 1 — do not move it to the internal router here.
- Do not expose port 9090 in `ports:` in docker-compose.

---

## Acceptance criteria

- `GET http://collection-service:9090/api/v1/collection/measurements?...` (from within Docker network) returns measurement data without auth.
- `GET /api/v1/collection/measurements` via the API gateway returns 404 (after Issue 1 is applied).
- `GET /api/v1/collection/latest` via the API gateway still works.
- Context-service can still compute context data end-to-end.
- Both servers start and stop cleanly.
