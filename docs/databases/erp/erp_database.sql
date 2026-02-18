CREATE SCHEMA "erp";

CREATE TYPE "erp"."resource_status" AS ENUM (
  'ACTIVE',
  'INACTIVE'
);

CREATE TABLE "erp"."company_config" (
  "company_id" uuid PRIMARY KEY,
  "erp_api_key" varchar,
  "erp_base_url" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "erp"."production_resource" (
  "production_resource_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "erp_id" varchar NOT NULL,
  "name" varchar NOT NULL,
  "description" text,
  "status" erp.resource_status NOT NULL DEFAULT 'ACTIVE',
  "factory_area_id" uuid,
  "fetched_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "erp"."order" (
  "order_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "erp_order_id" varchar NOT NULL,
  "order_number" varchar NOT NULL,
  "product_name" varchar,
  "quantity" integer,
  "start_time" timestamptz,
  "end_time" timestamptz,
  "production_resource_id" uuid,
  "status" varchar,
  "fetched_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "erp"."order_operation" (
  "operation_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "order_id" uuid NOT NULL,
  "company_id" uuid NOT NULL,
  "erp_operation_id" varchar,
  "name" varchar,
  "sequence_number" integer,
  "start_time" timestamptz,
  "end_time" timestamptz,
  "production_resource_id" uuid,
  "fetched_at" timestamptz NOT NULL
);

CREATE UNIQUE INDEX "uq_resource_company_erp" ON "erp"."production_resource" ("company_id", "erp_id");

CREATE UNIQUE INDEX "uq_order_company_erp" ON "erp"."order" ("company_id", "erp_order_id");

COMMENT ON COLUMN "erp"."company_config"."company_id" IS 'Same UUID as auth.company — populated during onboarding';

COMMENT ON COLUMN "erp"."company_config"."erp_api_key" IS 'API key or connection string for Monitor ERP';

COMMENT ON COLUMN "erp"."company_config"."erp_base_url" IS 'Base URL for the ERP API — may differ per company';

COMMENT ON COLUMN "erp"."production_resource"."erp_id" IS 'ID from Monitor ERP — the external reference';

COMMENT ON COLUMN "erp"."production_resource"."factory_area_id" IS 'Loose reference to auth.factory_area';

COMMENT ON COLUMN "erp"."production_resource"."fetched_at" IS 'When this record was last synced from ERP';

COMMENT ON COLUMN "erp"."order"."erp_order_id" IS 'Order ID from Monitor ERP';

COMMENT ON COLUMN "erp"."order"."status" IS '[ADDED] e.g. PLANNED, IN_PROGRESS, COMPLETED';

COMMENT ON COLUMN "erp"."order_operation"."sequence_number" IS 'Order of operations within the order';

COMMENT ON COLUMN "erp"."order_operation"."production_resource_id" IS 'Which machine runs this operation';

ALTER TABLE "erp"."order" ADD FOREIGN KEY ("production_resource_id") REFERENCES "erp"."production_resource" ("production_resource_id");

ALTER TABLE "erp"."order_operation" ADD FOREIGN KEY ("order_id") REFERENCES "erp"."order" ("order_id");

ALTER TABLE "erp"."order_operation" ADD FOREIGN KEY ("production_resource_id") REFERENCES "erp"."production_resource" ("production_resource_id");
