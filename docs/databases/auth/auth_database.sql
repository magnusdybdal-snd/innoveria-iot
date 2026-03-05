CREATE SCHEMA "auth";

CREATE TYPE "auth"."role_type" AS ENUM (
  'FACTORY_WORKER',
  'FACTORY_SUPERUSER',
  'PLATFORM_ADMIN' -- Innoveria role, for helping with onboarding
);

CREATE TABLE "auth"."company" (
  "company_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "address" varchar, -- Business address
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "auth"."factory" (
  "factory_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "address" varchar, -- Geographical
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "auth"."factory_area" (
  "area_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "factory_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "auth"."user" (
  "user_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "role" auth.role_type NOT NULL,
  "last_logged_in" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "auth"."refresh_token" (
  "token_id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "token_hash" text NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "revoked_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "device_info" text,
  "ip_address" inet
);

CREATE TABLE "auth"."permission" (
  "permission_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "permission_key" varchar UNIQUE NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "auth"."role_permission" (
  "role" auth.role_type NOT NULL,
  "permission_id" uuid NOT NULL,
  PRIMARY KEY ("role", "permission_id")
);

-- How the permission per role is set

ALTER TABLE "auth"."factory" ADD FOREIGN KEY ("company_id") REFERENCES "auth"."company" ("company_id");

ALTER TABLE "auth"."factory_area" ADD FOREIGN KEY ("factory_id") REFERENCES "auth"."factory" ("factory_id");

ALTER TABLE "auth"."user" ADD FOREIGN KEY ("company_id") REFERENCES "auth"."company" ("company_id");

ALTER TABLE "auth"."role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "auth"."permission" ("permission_id");

INSERT INTO "auth"."permission" ("permission_key", "description") VALUES
  ('gateway:create', 'Create gateways'),
  ('gateway:update', 'Update gateway configuration'),
  ('gateway:read', 'View gateway details and status'),
  ('sensor:create', 'Create sensors'),
  ('sensor:update', 'Update sensor configuration'),
  ('sensor:read', 'View sensor details and telemetry'),
  ('measurement:read', 'View measurements/time-series data'),
  ('user:manage', 'Create/update/deactivate users in company'), -- for super user
  ('factory:manage', 'Manage factories and factory areas');
  -- ERP handling not ready yet, but platform admin, cannot read this

INSERT INTO "auth"."role_permission" ("role", "permission_id")
SELECT 'FACTORY_WORKER', p.permission_id
FROM "auth"."permission" p
WHERE p.permission_key IN ('gateway:read', 'sensor:read', 'measurement:read'); -- incase someone should not read this data

INSERT INTO "auth"."role_permission" ("role", "permission_id")
SELECT 'FACTORY_SUPERUSER', p.permission_id
FROM "auth"."permission" p
WHERE p.permission_key IN (
  'gateway:create', 'gateway:update', 'gateway:read',
  'sensor:create', 'sensor:update', 'sensor:read',
  'measurement:read', 'user:manage', 'factory:manage'
);

INSERT INTO "auth"."role_permission" ("role", "permission_id")
SELECT 'PLATFORM_ADMIN', p.permission_id
FROM "auth"."permission" p
WHERE p.permission_key IN (
  'gateway:create', 'gateway:update', 'gateway:read',
  'sensor:create', 'sensor:update', 'sensor:read',
  'measurement:read'
);
