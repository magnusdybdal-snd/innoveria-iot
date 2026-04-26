# Issue 3: Admin Role Checks — Device Service

## Instructions for the implementing agent

Read this entire document before writing any code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**.

---

## Background

Device-service handlers that need authctx-scoped data (company ID) already use this pattern — see `services/device-service/internal/handlers/gateway_handler.go` and `sensor_handler.go` for examples.

The admin role check pattern adds one more step after extracting authctx:

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

Read `pkg/authctx/authctx.go` for the `IsAdmin()` definition before starting.

---

## Handlers that need the admin guard

### `services/device-service/internal/handlers/measurement_type_handlers.go`

| Handler | Action |
|---|---|
| `GetAllMeasurementTypes` | Add `authctx.FromRequest` + `IsAdmin()` check |
| `PostMeasurementType` | Add `authctx.FromRequest` + `IsAdmin()` check |
| `PatchDeprecateMeasurementType` | Add `authctx.FromRequest` + `IsAdmin()` check |
| `GetMeasurementTypes` | **No change** — this returns company-scoped measurement types; it already works correctly (or if it has no authctx, add `FromRequest` for company scoping but NOT the admin check) |

Before touching `GetMeasurementTypes`, check whether it already calls `authctx.FromRequest`. If it does not, add the extract + company scoping, but no admin check. If it already does both, leave it alone. Ask if you are unsure.

### `services/device-service/internal/handlers/payload_schema_handlers.go`

All four handlers need `authctx.FromRequest` + `IsAdmin()`:

| Handler | Notes |
|---|---|
| `GetPayloadSchemaDrafts` | Admin only |
| `GetPayloadSchemaByProfile` | Admin only |
| `PostDiscoverPayloadKeys` | Admin only |
| `PutPayloadSchemaLabels` | Admin only |

### `services/device-service/internal/handlers/sensor_metric_handlers.go`

| Handler | Action |
|---|---|
| `PutSensorMetrics` | Add `authctx.FromRequest` + `IsAdmin()` check |
| `GetSensorMetrics` | **No change here** — this will move to the internal router in Issue 4 |

---

## Swagger annotations

For each handler that gains a new admin check, add the following lines to the Swagger comment block if they are not already present:

```
// @Failure 401
// @Failure 403
```

---

## What NOT to change

- `GetGateways`, `PostGateway`, `PatchGateway`, `DeleteGateway` — already have authctx for company scoping; not admin-gated.
- `GetSensors`, `PostSensor`, `PatchSensor`, `DeleteSensor` — already have authctx; not admin-gated.
- `GetAllSensorProfiles` — check current state, but this should not be admin-gated.
- `GetSensorMetrics` — leave for Issue 4 (moves to internal router).
- `PostCompanyConfig`, `DeleteCompanyConfig` — internal-only, no authctx needed (handled by Issue 1).

---

## Acceptance criteria

- `GET /api/v1/device/payload-schema` with a non-admin JWT returns 403.
- `POST /api/v1/device/measurement-types` with a non-admin JWT returns 403.
- `PUT /api/v1/device/sensors/{eui}/metrics` with a non-admin JWT returns 403.
- All above endpoints return 401 with no JWT.
- `GET /api/v1/device/measurement-types` (company-scoped) still works for regular users.
- `GET /api/v1/device/sensors` still works for regular users.
