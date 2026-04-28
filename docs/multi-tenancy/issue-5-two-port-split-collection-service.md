# Issue 5: Company Scoping — Collection Service

## Instructions for the implementing agent

Read this entire document before writing any code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**.

---

## Problem

`GET /measurements` and `GET /latest` on collection-service query by `device_eui` only — there is no `company_id` filter on the read queries. `company_id` is stored on every measurement row (resolved from the tenant mapping at ingest time) but is unused for read scoping.

A two-port split is **not** needed here. Context-service calls collection-service directly (service-to-service) and can forward `X-Auth-Company-Id` from its own request context, just as a regular authenticated caller would. Adding `AND company_id = $n` to the queries and wiring up `authctx` in the handlers is sufficient — and also adds defence-in-depth so that even a caller who knows another company's device EUI cannot retrieve their data.

---

## Architecture overview

```
External (port 8080)
────────────────────
GET /latest          → requires X-Auth-Company-Id, scoped to company
GET /measurements    → requires X-Auth-Company-Id, scoped to company
```

Context-service calls both endpoints directly at `http://collection-service:8080` and forwards `X-Auth-Company-Id`. No internal port needed.

---

## Files to change

### 1. `services/collection-service/internal/repository/measurement_repository.go`

Add `AND company_id = $n` to both read queries.

`findLatestQuery`:
```sql
SELECT device_eui, timestamp, payload, company_id
FROM collection.sensor_measurement
WHERE device_eui = $1
  AND company_id = $2
ORDER BY timestamp DESC
LIMIT 1
```

`findByTimeRangeQuery`:
```sql
SELECT device_eui, timestamp, payload, company_id
FROM collection.sensor_measurement
WHERE device_eui = $1
  AND company_id = $2
  AND timestamp >= $3
  AND timestamp <= $4
ORDER BY timestamp ASC
```

Update the repository method signatures to accept `companyID string` and pass it as the new parameter.

### 2. `services/collection-service/internal/domain/measurement.go` (or wherever the service interface is defined)

Update `MeasurementService` interface methods to accept `companyID string`:

```go
GetLatest(ctx context.Context, companyID string, deviceEUI string) (SensorMeasurement, error)
GetByTimeRange(ctx context.Context, companyID string, deviceEUI string, from, to time.Time) ([]SensorMeasurement, error)
```

### 3. `services/collection-service/internal/service/measurement_service.go`

Update the service methods to accept and pass `companyID` through to the repository.

### 4. `services/collection-service/internal/handlers/latest.go`

Extract `companyID` from `authctx` and pass it to the service:

```go
auth, err := authctx.FromRequest(r)
if err != nil {
    json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
    return
}
```

Pass `auth.CompanyID` to `svc.GetLatest(...)`.

### 5. `services/collection-service/internal/handlers/by_time_frame.go`

Same pattern — extract `authctx` and pass `auth.CompanyID` to `svc.GetByTimeRange(...)`.

---

## Context-service changes (handled by context-service team)

Context-service's `CollectionClient` at `services/context-service/internal/clients/collection_client.go` currently calls `GET /measurements` with no headers. The context-service team must forward `X-Auth-Company-Id` on this call so collection-service can scope the query.

The company ID is available in the context-service request context (set by the gateway from the caller's JWT) — it just needs to be extracted and forwarded as a header on outbound calls to collection-service.

No URL changes are needed — collection-service stays on port 8080.

---

## What NOT to change

- No two-port split needed for collection-service.
- `GET /latest` and `GET /measurements` both stay on the external router at port 8080.
- Do not change docker-compose for collection-service as part of this issue.
- `PostTenantMapping` is internal-only and unaffected.

---

## Acceptance criteria

- `GET /api/v1/collection/latest?device_eui=X` with a valid JWT returns only that company's data.
- `GET /api/v1/collection/latest?device_eui=X` with no JWT returns 401.
- `GET /api/v1/collection/measurements?device_eui=X&from=...&to=...` with a valid JWT returns only that company's data.
- A request with a valid JWT from company A cannot retrieve measurements for a device belonging to company B, even if the device EUI is known.
- Context-service can still compute context data end-to-end after the context-service team adds header forwarding.
