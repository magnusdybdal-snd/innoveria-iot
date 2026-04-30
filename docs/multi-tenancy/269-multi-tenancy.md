# Multi-Tenancy: Enforce Company Data Separation

Users can currently see all data regardless of company. This implements per-company data isolation so a user only sees data belonging to their own company.

**Roles (simplified for now):**
- `PLATFORM_ADMIN` — global admin, manages all companies and users, sees all data
- `USER` — belongs to one company, sees only that company's data

Role-based access control (granular permissions, SUPERUSER etc.) is out of scope and tracked separately.

---

## What's Already in Place

The infrastructure is complete. Services just don't use it yet.

**JWT claims** (issued by auth-service):
```
sub:        user UUID
company_id: company UUID
role:       PLATFORM_ADMIN | USER
```

**API gateway middleware** strips any client-supplied headers and injects trusted values from the validated JWT on every request:
```
X-Auth-User-Id:    <user UUID>
X-Auth-Company-Id: <company UUID>
```

Services receive these headers on every request. They just currently ignore them.

---

## Branching

```
dev
 └── 269-enforce-company-data-seperation  ← main branch for this feature
      ├── multi-tenancy/authctx           ← Issue 1, do first
      ├── multi-tenancy/device-service    ← Issue 2
      ├── multi-tenancy/collection-service
      ├── multi-tenancy/context-service
      ├── multi-tenancy/auth-service
      ├── multi-tenancy/frontend-cleanup
      └── multi-tenancy/frontend-user-mgmt
```

Branch `multi-tenancy` is based off `dev`. All sub-branches merge back into `multi-tenancy`, which merges to `dev` when complete.

> **Note on branch 217:** Branch 217 (voltage/payload schemas) is not yet merged to dev and is ahead of it. The 217 additions are either scoped to existing company-filtered tables (voltage on sensors — works automatically) or are intentionally global/admin vocabulary (payload schemas, measurement types — no company filtering needed). No follow-up multi-tenancy work is required after 217 merges. Minor conflicts will exist in `sensor_repository.go`, `sensor_service.go`, and `sensor_handler.go` — resolvable at merge time.

---

## Issues

### Issue 1 — `pkg/authctx`: shared header extraction helper
**Branch off:** `multi-tenancy` — **do this first, everything depends on it**

Create `pkg/authctx/authctx.go` with a helper that reads `X-Auth-Company-Id` and `X-Auth-User-Id` from a request and returns typed values (UUIDs). Return a clear error if either is missing.

All services use this package instead of reading headers manually.

---

### Issue 2 — device-service: company isolation
**Branch off:** `multi-tenancy`

Two hardcoded company UUIDs need replacing:
- `service/gateway_service.go:130` — `FindAllByCompanyID` called with hardcoded UUID
- `service/sensor_service.go:170` — same

Replace with `authctx` extracted from the request context.

On create/update endpoints, override the `company_id` field in the request body with the value from the header — never trust the client to supply their own company identity.

---

### Issue 3 — collection-service: company ownership on measurement queries
**Branch off:** `multi-tenancy`

`GET /latest?device_eui=...` and `GET /measurements?device_eui=...` return data for any device EUI with no ownership check. The `sensor_measurement` table has a `company_id` column (populated via `tenant_mapping` at ingest time). Add `AND company_id = $n` to the relevant repository queries and pass the value from `authctx`.

---

### Issue 4 — context-service: derive company_id from header
**Branch off:** `multi-tenancy`

`GET /data` currently requires `?company_id=` as a query parameter supplied by the client. Remove this parameter and read `company_id` from `authctx` instead.

---

### Issue 5 — auth-service: user management API
**Branch off:** `multi-tenancy`

New endpoints for the platform admin to manage users:

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/auth/users` | Create a user (name, email, password, role, company_id) |
| `GET` | `/auth/users` | List all users |
| `PATCH` | `/auth/users/:id` | Update user |
| `DELETE` | `/auth/users/:id` | Remove user |

These endpoints must be guarded — only `PLATFORM_ADMIN` can call them.

The `auth.user` and `auth.company` tables already exist with the correct schema.

---

### Issue 6 — frontend: remove client-supplied company_id from API calls
**Branch off:** `multi-tenancy`

The context-service `GET /data` call currently passes `company_id` as a query parameter. Remove it — the backend derives it from the JWT now. Audit other API calls for the same pattern.

---

### Issue 7 — frontend: user management screen
**Branch off:** `multi-tenancy` (depends on Issue 5 being merged first)

Admin-only screen for the platform admin to:
- List all users across all companies
- Create a new user and assign to a company
- Delete a user

---

## Implementation Order

1. **Issue 1** (authctx) — foundation, do first
2. **Issues 2–5** — can be done in parallel after Issue 1
3. **Issue 6** — after Issue 4 is merged
4. **Issue 7** — after Issue 5 is merged
