CREATE SCHEMA "auth";

CREATE TYPE "auth"."role_type" AS ENUM (
  'FACTORY_WORKER',
  'FACTORY_SUPERUSER',
  'ADMIN'
);

CREATE TABLE "auth"."company" (
  "company_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "address" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "auth"."factory_area" (
  "area_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
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

ALTER TABLE "auth"."factory_area" ADD FOREIGN KEY ("company_id") REFERENCES "auth"."company" ("company_id");

ALTER TABLE "auth"."user" ADD FOREIGN KEY ("company_id") REFERENCES "auth"."company" ("company_id");
