# Issue 1: Remove Internal-Only Endpoints from API Gateway

## Instructions for the implementing agent

Read this entire document before writing a single line of code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**. This document describes exactly what to change; any deviation should be questioned, not silently resolved.

---

## Background

The API gateway (`services/api-gateway`) is the public-facing entry point. It proxies requests from the frontend and external callers to internal microservices.

Some endpoints on internal microservices are **only ever called by other internal services** (service-to-service). They must never be reachable by external clients. Currently, some of these are accidentally exposed through the gateway's proxy configuration.

The fix here is purely in the **API gateway's router** — remove specific route strings from the proxy allow-lists. No handler code changes are needed.

---

## What to change

### File: `services/api-gateway/internal/server/router.go`

`RegisterProxyService` takes a list of path suffixes it will forward. Remove the paths listed below from those lists.

#### Auth service — remove `/companies`

The current call:
```go
handlers.RegisterProxyService(mux, AUTHENTICATION_ROUTE, "auth-service", cfg.AuthSvcURL, []string{
    "/companies",        // <-- REMOVE THIS
    "/factories",
    "/factory-areas",
    "/login",
    "/refresh",
    "/me",
    "/users",
})
```

`POST /companies` and `DELETE /companies/{id}` on auth-service are only called by the onboarding-service internally. The onboarding-service is already separately registered on the gateway under `/api/v1/onboarding/company`. Regular users must never be able to POST or DELETE companies directly.

> **Note:** `GET /companies` and `GET /companies/{id}` also live under the `/companies` prefix. Once `/companies` is removed from the gateway, those GET endpoints will also become unreachable from outside. That is intentional for now — `GetAllCompanies` and `GetOneCompany` will get their own admin role checks in Issue 2, and can be re-exposed if needed later. Check with the team if this changes scope.

#### Collection service — remove `/measurements` (but keep `/latest`)

The current call:
```go
handlers.RegisterProxyService(mux, COLLECTION_ROUTE, "collection-service", cfg.CollSvcURL, []string{
    "/latest",
    "/measurements",   // <-- REMOVE THIS
})
```

`GET /measurements` (by time range) is used by the context-service internally. The frontend uses only `/latest`. After this change, the time-range measurement endpoint is no longer reachable from the gateway. (Issue 5 covers adding a proper internal port to collection-service so that context-service can still reach it.)

#### Collection service — `PostTenantMapping` (`POST /company-config`)

`POST /api/v1/collection/company-config` is not currently in the gateway proxy list (check the router — if it is not present, no change is needed here). If it is present, remove it. This endpoint is called only by onboarding-service.

#### Device service — `PostCompanyConfig` and `DeleteCompanyConfig` (`/company-config`)

Check whether `/company-config` is in the device service proxy list. If it is, remove it. These handlers are only called by onboarding-service internally.

```go
handlers.RegisterProxyService(mux, DEVICE_ROUTE, "device-service", cfg.DeviceSvcURL, []string{
    "/gateways",
    "/sensors",
    "/sensor-profiles",
    "/sensor-groups",
    "/measurement-types",
    "/payload-schema",
    // "/company-config"  <-- remove if present
})
```

---

## How `RegisterProxyService` works

Before making changes, read `services/api-gateway/internal/handlers/` to understand how `RegisterProxyService` matches paths. It is a prefix-based proxy — removing `/companies` removes ALL paths that start with `/companies`, not just an exact match. Confirm this understanding before proceeding and ask if the behavior is different from what is described here.

---

## What NOT to change

- Do not modify any microservice handler code in this issue.
- Do not add auth checks in this issue.
- Do not change `docker-compose.yml` or `docker-compose.prod.yml`.

---

## Acceptance criteria

- `POST /api/v1/auth/companies` returns 404 (not proxied) from the gateway.
- `DELETE /api/v1/auth/companies/{id}` returns 404 from the gateway.
- `GET /api/v1/collection/measurements` returns 404 from the gateway.
- `POST /api/v1/onboarding/company` still works (onboarding-service proxy untouched).
- `GET /api/v1/auth/factories`, `/factory-areas`, `/login`, `/refresh`, `/me`, `/users` still work.
- `GET /api/v1/collection/latest` still works.
