-- +goose Up
CREATE SCHEMA "auth";

CREATE TYPE "auth"."role_type" AS ENUM (
  'FACTORY_WORKER',
  'FACTORY_SUPERUSER',
  'PLATFORM_ADMIN' -- Innoveria role, for helping with onboarding
);

CREATE TABLE "auth"."company" (
  "company_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "address" varchar NOT NULL, -- Business address
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "auth"."factory" (
  "factory_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "address" varchar, -- Geographical
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
    
  CONSTRAINT fk_factory_company
    FOREIGN KEY (company_id)
    REFERENCES auth.company(company_id)
    ON DELETE CASCADE
);

CREATE TABLE "auth"."factory_area" (
  "area_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "factory_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),

  CONSTRAINT fk_factory_area_factory
    FOREIGN KEY (factory_id)
    REFERENCES auth.factory(factory_id)
    ON DELETE CASCADE
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
  "updated_at" timestamptz NOT NULL DEFAULT (now()),

  CONSTRAINT fk_user_company
    FOREIGN KEY (company_id)
    REFERENCES auth.company(company_id)
    ON DELETE CASCADE
);

CREATE TABLE "auth"."refresh_token" (
  "token_id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "token_hash" text NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "revoked_at" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "device_info" text,
  "ip_address" inet,
    
  CONSTRAINT fk_refresh_token_user
    FOREIGN KEY (user_id)
    REFERENCES auth.user(user_id)
    ON DELETE CASCADE
);

CREATE INDEX idx_factory_company_id
  ON auth.factory(company_id);

CREATE INDEX idx_factory_area_factory_id
  ON auth.factory_area(factory_id);

CREATE INDEX idx_user_company_id
  ON auth."user"(company_id);

CREATE INDEX idx_refresh_token_user_id
  ON auth.refresh_token(user_id);

-- +goose Down
DROP SCHEMA "auth" CASCADE;
