# Issue 2: Admin Role Checks — Auth Service and Onboarding Service

## Instructions for the implementing agent

Read this entire document before writing any code. If anything is ambiguous or you need to make a judgment call, **ask first — do not assume**.

---

## Background

`pkg/authctx` provides a shared pattern for extracting the authenticated user from an HTTP request and checking their role:

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

`IsAdmin()` returns `true` when the user has the `PLATFORM_ADMIN` role. See `pkg/authctx/authctx.go` for the definition.

The existing pattern with both an authctx check and an `IsAdmin()` guard is fully implemented in `services/auth-service/internal/handlers/user_handlers.go`. Read that file to understand the exact idiom before writing any code.

---

## What to change

### auth-service: `GetAllCompanies` and `GetOneCompany`

**File:** `services/auth-service/internal/handlers/company_handlers.go`

Both handlers currently have no authctx extraction at all. Add the standard two-step guard (extract authctx → check `IsAdmin()`) at the top of each handler's inner function, before any service call.

- `GetAllCompanies` — lists all companies across the platform. Only a platform admin should see this.
- `GetOneCompany` — retrieves a single company by ID. Only a platform admin should see this.

`PostCompany` and `DeleteCompany` in the same file are being **removed from the gateway** (Issue 1) and called only by onboarding-service internally, so they do NOT need authctx checks — leave them as-is.

### onboarding-service: `PostCompany`

**File:** `services/onboarding-service/internal/handlers/` (find the company handler)

`PostCompany` on onboarding-service is the user-facing endpoint for creating a new company. It is exposed on the gateway under `/api/v1/onboarding/company`. Add the standard authctx + `IsAdmin()` guard.

Read the existing onboarding handler to understand the import path for `authctx` in this service before adding the import.

---

## Swagger annotations

After adding the checks, update the Swagger annotations for each changed handler to add:

```
// @Failure 401
// @Failure 403
```

if those lines are not already present.

---

## What NOT to change

- Do not change `GetAllFactories`, `GetOneFactory`, `PostFactory`, `DeleteFactory` — those already use `authctx.FromRequest` for company scoping and are not gated to admin.
- Do not change factory area handlers.
- Do not change `PostLogin`, `PostRefresh`, `GetMe` — those are public or user-scoped auth endpoints.
- Do not touch `docker-compose.yml` or the gateway router.

---

## Acceptance criteria

- `GET /api/v1/auth/companies` with a non-admin JWT returns 403.
- `GET /api/v1/auth/companies` with no JWT returns 401.
- `GET /api/v1/auth/companies` with an admin JWT returns 200 and the company list.
- `GET /api/v1/auth/companies/{id}` enforces the same rules.
- `POST /api/v1/onboarding/company` with a non-admin JWT returns 403.
- `POST /api/v1/onboarding/company` with an admin JWT succeeds (creates company).
