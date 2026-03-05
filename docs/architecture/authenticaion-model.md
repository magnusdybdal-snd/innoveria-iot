# Authentication model notes

## What exists in the current auth database

Based on `docs/databases/auth/auth_database.sql`:

- Schema: `auth`
- Core entity: `company`
- Tenancy grouping inside company: `factory_area`
- Identity table: `user`
  - Includes `company_id` directly on the user (single-company model in current schema)
  - Stores `email`, `password_hash`, `role`, and `last_logged_in`
- Session persistence: `refresh_token`
  - Token material is stored as `token_hash` (not plaintext token)
  - Includes expiration (`expires_at`), revocation (`revoked_at`), and device/network metadata (`device_info`, `ip_address`)

## Role model (current schema)

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

This is a role-based model with role value stored directly on `auth.user.role`.


## Relationships and boundaries

- `auth.factory_area.company_id -> auth.company.company_id`
- `auth.user.company_id -> auth.company.company_id`

Interpretation:

- Company is the tenant boundary.
- Users currently belong to exactly one company (via direct foreign key).
- Factory areas are sub-units inside a company.

## Authentication/session flow (from architecture diagram)

From `docs/architecture/authentication-data-flow.md`:

1. Frontend request reaches API Gateway.
2. Gateway validates JWT.
3. Gateway resolves active company context (`active_company_id`).
4. Gateway performs RBAC authorization check using `(user_id, company_id, permission)`.
5. Gateway injects trusted context headers:
   - `X-User-Id`
   - `X-Company-Id`
   - `X-Roles`
   - `X-Permissions`
   - `X-Request-Id`
6. Downstream services (Auth/Device/Collection) consume trusted headers instead of re-parsing identity.

## Auth service responsibilities

- Handlers:
  - `POST /auth/login`
  - `GET /auth/me`
  - `POST /auth/switch-company`
  - `POST /companies` (Platform Admin)
- Service layer:
  - Auth logic (login/profile/context)
  - Company onboarding

## Important model observation (schema vs diagram)

The architecture diagram references repository/domain concepts not present in the current SQL file:

- `user_company_memberships`
- `roles`
- `permissions`

This suggests a target/expanded authorization design (membership + granular permissions), while the current DB schema is still a simpler model (single `user.company_id` + enum role).

## Practical summary

- **AuthN:** JWT + refresh-token based sessions.
- **Tenant context:** Company-scoped (`company_id`), resolved at gateway.
- **AuthZ:** RBAC at gateway; currently backed by enum role in DB, with diagram indicating future/extended permission model.
- **Operational traceability:** Header-injected request context and refresh-token metadata support auditing and security operations.
