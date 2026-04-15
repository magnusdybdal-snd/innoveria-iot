-- +goose Up
CREATE SCHEMA IF NOT EXISTS "erp";

CREATE TYPE "erp"."order_status" AS ENUM (
  'not_initialized',
  'registered',
  'printed',
  'started',
  'finished',
  'post_calculated',
  'delivered',
  'historical'
);

CREATE TYPE "erp"."operation_status" AS ENUM (
  'none',
  'started',
  'partially_shipped',
  'fully_shipped',
  'partially_reported',
  'finished'
);

CREATE TYPE "erp"."work_center_type" AS ENUM (
  'machine',
  'manual_work',
  'sub_contract',
  'pool',
  'pick'
);

CREATE TYPE "erp"."order_report_type" AS ENUM (
  'regular',
  'send_to_subcontractor',
  'receive_from_subcontractor',
  'cancel_rest',
  'material_only',
  'subcontractor_invoice_price',
  'recording_terminal',
  'adjust_recording',
  'undo_regular',
  'undo_recording_terminal',
  'undo_adjust_recording',
  'subcontractor_posterior_report',
  'subcontractor_costs_manual_report',
  'pick_work_center'
);

CREATE TABLE IF NOT EXISTS "erp"."production_resource" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "number" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "type" "erp"."work_center_type" NOT NULL,
  "received_at" TIMESTAMPTZ NOT NULL,
  PRIMARY KEY ("company_id", "id")
);

CREATE TABLE IF NOT EXISTS "erp"."order" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "order_number" TEXT NOT NULL,
  "part_id" TEXT NOT NULL,
  "part_description" TEXT NOT NULL,
  "planned_start_date" TIMESTAMPTZ NOT NULL,
  "planned_finish_date" TIMESTAMPTZ NOT NULL,
  "actual_start_date" TIMESTAMPTZ,
  "actual_finish_date" TIMESTAMPTZ,
  "status" "erp"."order_status" NOT NULL,
  "priority" INTEGER NOT NULL,
  "received_at" TIMESTAMPTZ NOT NULL,
  PRIMARY KEY ("company_id", "id")
);

CREATE TABLE IF NOT EXISTS "erp"."order_operation" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "production_resource_id" BIGINT NOT NULL,
  "order_id" BIGINT NOT NULL,
  "planned_start_date" TIMESTAMPTZ NOT NULL,
  "planned_finish_date" TIMESTAMPTZ NOT NULL,
  "actual_start_date" TIMESTAMPTZ,
  "actual_finish_date" TIMESTAMPTZ,
  "status" "erp"."operation_status" NOT NULL,
  "production_resource_status" "erp"."operation_status" NOT NULL,
  "received_at" TIMESTAMPTZ NOT NULL,
  PRIMARY KEY ("company_id", "id"),
  CONSTRAINT "fk_order_operation_order"
    FOREIGN KEY ("company_id", "order_id") REFERENCES "erp"."order" ("company_id", "id"),
  CONSTRAINT "fk_order_operation_resource"
    FOREIGN KEY ("company_id", "production_resource_id") REFERENCES "erp"."production_resource" ("company_id", "id")
);

CREATE TABLE IF NOT EXISTS "erp"."order_report" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "order_operation_id" BIGINT NOT NULL,
  "production_resource_id" BIGINT NOT NULL,
  "quantity" DOUBLE PRECISION NOT NULL,
  "rest_quantity" DOUBLE PRECISION NOT NULL,
  "type" "erp"."order_report_type" NOT NULL,
  "reporting_timestamp" TIMESTAMPTZ NOT NULL,
  "actual_reported_date" TIMESTAMPTZ,
  "received_at" TIMESTAMPTZ NOT NULL,
  PRIMARY KEY ("company_id", "id"),
  CONSTRAINT "fk_order_report_order_operation"
    FOREIGN KEY ("company_id", "order_operation_id") REFERENCES "erp"."order_operation" ("company_id", "id"),
  CONSTRAINT "fk_order_report_resource"
    FOREIGN KEY ("company_id", "production_resource_id") REFERENCES "erp"."production_resource" ("company_id", "id")
);

CREATE INDEX IF NOT EXISTS "idx_order_company_received"
  ON "erp"."order" ("company_id", "received_at" DESC);

CREATE INDEX IF NOT EXISTS "idx_order_operation_company_received"
  ON "erp"."order_operation" ("company_id", "received_at" DESC);

CREATE INDEX IF NOT EXISTS "idx_order_report_company_received"
  ON "erp"."order_report" ("company_id", "received_at" DESC);

CREATE INDEX IF NOT EXISTS "idx_production_resource_company_received"
  ON "erp"."production_resource" ("company_id", "received_at" DESC);

-- +goose Down
DROP SCHEMA IF EXISTS "erp" CASCADE;
