# Authentication model notes

## What exists in the current auth database

Based on `docs/databases/auth/auth_database.sql`:

- Schema: `auth`
- Core tenant entity: `company`
- Factory hierarchy inside company:
  - `factory` (belongs to `company`)
  - `factory_area` (belongs to `factory`)
- Identity table: `user`
  - Includes `company_id` directly on the user (single-company model in current schema)
  - Stores `email`, `password_hash`, `role`, and `last_logged_in`
- Session persistence: `refresh_token`
  - Token material is stored as `token_hash` (not plaintext token)
  - Includes expiration (`expires_at`), revocation (`revoked_at`)
- Authorization tables:
  - `permission` (permission catalog, keyed by `permission_key`)
  - `role_permission` (maps `auth.role_type` -> `permission_id`)

## Role and permission model (current schema)

`auth.role_type` enum defines three roles:

- `FACTORY_WORKER`
- `FACTORY_SUPERUSER`
- `PLATFORM_ADMIN` (Innoveria support/onboarding role)

Role intent (project-specific):

- `FACTORY_WORKER`: Regular factory user for day-to-day usage (view assigned areas/devices, view measurements, basic actions).
- `FACTORY_SUPERUSER`: Company power user (manage users/areas, configure devices/gateways, broader configuration within the company).
- `PLATFORM_ADMIN`: Innoveria/internal onboarding/support role (cross-company).
  - Can add/configure gateways and sensors.
  - Can see sensor data and ERP data needed for support.
  - Must NOT be able to access sensitive/customer-confidential data beyond what is required (least privilege).

This schema uses role + permission mapping:

- `auth.user.role` stores the role enum on each user.
- `auth.permission` defines permission keys (for example `gateway:create`, `sensor:read`, `user:manage`, `factory:manage`).
- `auth.role_permission` assigns permissions to each role.

Seeded role-to-permission behavior in the SQL:

- `FACTORY_WORKER`: read-only access to gateways/sensors/measurements.
- `FACTORY_SUPERUSER`: read/write for gateways/sensors, plus `user:manage` and `factory:manage`.
- `PLATFORM_ADMIN`: gateway/sensor create-update-read and `measurement:read` (does not include `user:manage`/`factory:manage` in current seed).


## Relationships and boundaries

- `auth.factory.company_id -> auth.company.company_id`
- `auth.factory_area.factory_id -> auth.factory.factory_id`
- `auth.user.company_id -> auth.company.company_id`
- `auth.role_permission.permission_id -> auth.permission.permission_id`

Interpretation:

- Company is the tenant boundary.
- Users currently belong to exactly one company (via direct foreign key).
- Factories are sub-units inside a company.
- Factory areas are sub-units inside a factory.
- Authorization checks can be role-derived via `role_permission`.

## Authentication/session flow (from architecture diagram)

From `docs/architecture/authentication-data-flow.md`:

1. Frontend request reaches API Gateway.
2. Login (`POST /auth/login`) is called without an access JWT.
   - Auth Service validates credentials.
   - Auth Service returns a short-lived access JWT for API calls.
   - Auth Service also returns a long-lived refresh token and stores only its hash in `auth.refresh_token.token_hash` (plus `expires_at` / `revoked_at`).
3. For protected API calls, client sends `Authorization: Bearer <access_jwt>`.
4. Gateway validates the access JWT.
5. Gateway resolves active company context (`active_company_id`).
6. Gateway performs RBAC authorization check using `(user_id, company_id, permission)`.
7. Gateway injects trusted context headers:
   - `X-User-Id`
   - `X-Company-Id`
   - `X-Roles`
   - `X-Permissions`
   - `X-Request-Id`
8. Downstream services (Auth/Device/Collection) consume trusted headers instead of re-parsing identity.
9. When the access JWT expires, client calls an auth endpoint (commonly `POST /auth/refresh`) with the refresh token to obtain a new access JWT.
   - Auth Service hashes the presented refresh token and matches it against `auth.refresh_token.token_hash` and checks `expires_at` / `revoked_at`.

## Auth service responsibilities

- Handlers:
  - `POST /auth/login`
  - `GET /auth/me`
  - `POST /auth/switch-company` (Platform Admin)
  - `POST /companies` (Platform Admin)
- Service layer:
  - Auth logic (login/profile/context)
  - Company onboarding

Why `POST /auth/switch-company` exists for Innoveria support users:

- It allows a platform admin to act in a selected company context.
- In practice this implies issuing a company-context token rather than relying on one global token.

## Practical summary

- **AuthN:** JWT + refresh-token based sessions.
- **Tenant context:** Company-scoped (`company_id`), resolved at gateway.
- **Domain structure:** `company -> factory -> factory_area`.
- **AuthZ:** RBAC at gateway, backed by role enum (`auth.user.role`) plus permission mapping (`auth.permission`, `auth.role_permission`).
- **Operational traceability:** Header-injected request context and refresh-token metadata support auditing and security operations.
