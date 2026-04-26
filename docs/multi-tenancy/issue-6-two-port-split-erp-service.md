# Issue 6: Two-Port Split — ERP Service

## Instructions for the implementing agent

Read this entire document before writing any code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**. This issue is assigned to a specific colleague; confirm scope with them before starting.

---

## Context

This issue is scoped to the **ERP service** (`services/erp-service`). At the time of writing, the ERP service router (`services/erp-service/internal/server/router.go`) is largely a placeholder — it only has a root handler. The actual production-resource endpoints and order-related endpoints are expected to be added by the colleague working on this service.

**This document describes the pattern to follow when those endpoints are added.** Do not implement the ERP service's business logic here — that is separate work. This issue is about ensuring the structural two-port pattern is in place when that work lands.

---

## Problem

ERP service has two categories of endpoints:

1. **User-facing** (exposed via API gateway): e.g. `GET /production-resources` — regular users need to query production resources from the frontend. These must go through the gateway and be auth-gated (company scoping, optionally admin check).

2. **Internal-only** (called by other services): Order-related endpoints (e.g. for context-service or other internal consumers) — these must not be reachable from outside the Docker network.

Running both categories on one port forces a choice: either expose everything (including internal endpoints) to the gateway, or hide everything. The two-port split solves this cleanly.

---

## Pattern to follow (same as Issues 4 and 5)

```
External (port 8080)              Internal (port 9090)
────────────────────              ────────────────────
GET /production-resources         Order handlers
(user-facing, auth required)      (internal, no auth required)
```

---

## Files to change

### 1. `services/erp-service/internal/config/config.go`

Add `InternalAddr string` to the `Config` struct and load it:

```go
InternalAddr: ":" + env.Get("INTERNAL_PORT", "9090"),
```

### 2. `services/erp-service/internal/server/router.go`

Register user-facing endpoints on the existing `NewRouter`. Create a new `NewInternalRouter` function that registers only the internal endpoints (order handlers and anything else that should not be externally reachable).

The exact handler functions will be determined by the colleague implementing ERP. This document describes the structural requirement: **internal-only endpoints must go in `NewInternalRouter`, not in `NewRouter`**.

### 3. `services/erp-service/internal/server/server.go`

Follow the same pattern as Issues 4 and 5: start a second `http.Server` on `cfg.InternalAddr` with `NewInternalRouter(...)` as its handler. Shut it down alongside the main server in the graceful shutdown block.

### 4. `docker-compose.yml`

In the `erp-service` service block:
```yaml
environment:
  - GO_ENV=development
  - INTERNAL_PORT=9090   # <-- add
```

Do **not** add a `ports:` mapping for 9090.

### 5. `docker-compose.prod.yml`

Same addition to `erp-service` environment.

---

## Auth requirements for user-facing endpoints

### `GET /production-resources`

This endpoint is queried by authenticated users from the frontend. It must:
1. Call `authctx.FromRequest(r)` to extract the user's company ID.
2. Use `auth.CompanyID` to scope results (only return resources belonging to the user's company).
3. Do NOT require `IsAdmin()` — regular company users need this.

Read `services/device-service/internal/handlers/gateway_handler.go:GetGateways` as a reference for the company-scoped-but-not-admin pattern.

---

## If context-service calls ERP service

If context-service (or another internal service) needs to call ERP's internal endpoints, follow the same client pattern used in context-service for the collection client (`services/context-service/internal/clients/collection_client.go`). The client must target `http://erp-service:9090` (the internal port), not `http://erp-service:8080`.

Add `ERP_SERVICE_INTERNAL=http://erp-service:9090` to the calling service's environment in docker-compose (both dev and prod).

---

## What NOT to change

- Do not implement ERP business logic in this issue — this is structural scaffolding only.
- Do not expose port 9090 in `ports:`.
- Do not touch any other service's code unless adding a client for erp-service (discuss first).

---

## Acceptance criteria

- External port (8080) serves only user-facing endpoints through the API gateway.
- Internal port (9090) serves internal-only endpoints; not mapped in `ports:`.
- Both servers start and stop cleanly.
- `GET /api/v1/erp/production-resources` is accessible via the gateway for authenticated users.
- Order endpoints (or other internal endpoints) are **not** reachable via the gateway.
