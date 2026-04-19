-- +goose Up
CREATE SCHEMA IF NOT EXISTS "erp_raw";

CREATE TABLE IF NOT EXISTS "erp_raw"."production_resource" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "number" TEXT NOT NULL,
  "description" TEXT NOT NULL,
  "type" TEXT NOT NULL,
  "received_at" TIMESTAMPTZ NOT NULL,
  "ingested_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "last_seen_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "sync_status" TEXT NOT NULL DEFAULT 'pending',
  "sync_error" TEXT,
  "retry_count" INTEGER NOT NULL DEFAULT 0,
  "next_retry_at" TIMESTAMPTZ,
  "last_synced_at" TIMESTAMPTZ,
  PRIMARY KEY ("company_id", "id")
);

CREATE TABLE IF NOT EXISTS "erp_raw"."order" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "order_number" TEXT NOT NULL,
  "part_id" TEXT NOT NULL,
  "part_description" TEXT NOT NULL,
  "planned_start_date" TIMESTAMPTZ NOT NULL,
  "planned_finish_date" TIMESTAMPTZ NOT NULL,
  "actual_start_date" TIMESTAMPTZ,
  "actual_finish_date" TIMESTAMPTZ,
  "status" TEXT NOT NULL,
  "priority" INTEGER NOT NULL,
  "received_at" TIMESTAMPTZ NOT NULL,
  "ingested_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "last_seen_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "sync_status" TEXT NOT NULL DEFAULT 'pending',
  "sync_error" TEXT,
  "retry_count" INTEGER NOT NULL DEFAULT 0,
  "next_retry_at" TIMESTAMPTZ,
  "last_synced_at" TIMESTAMPTZ,
  PRIMARY KEY ("company_id", "id")
);

CREATE TABLE IF NOT EXISTS "erp_raw"."order_operation" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "production_resource_id" BIGINT NOT NULL,
  "order_id" BIGINT NOT NULL,
  "planned_start_date" TIMESTAMPTZ NOT NULL,
  "planned_finish_date" TIMESTAMPTZ NOT NULL,
  "actual_start_date" TIMESTAMPTZ,
  "actual_finish_date" TIMESTAMPTZ,
  "status" TEXT NOT NULL,
  "production_resource_status" TEXT NOT NULL,
  "received_at" TIMESTAMPTZ NOT NULL,
  "ingested_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "last_seen_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "sync_status" TEXT NOT NULL DEFAULT 'pending',
  "sync_error" TEXT,
  "retry_count" INTEGER NOT NULL DEFAULT 0,
  "next_retry_at" TIMESTAMPTZ,
  "last_synced_at" TIMESTAMPTZ,
  PRIMARY KEY ("company_id", "id")
);

CREATE TABLE IF NOT EXISTS "erp_raw"."order_report" (
  "company_id" TEXT NOT NULL,
  "id" BIGINT NOT NULL,
  "order_operation_id" BIGINT NOT NULL,
  "production_resource_id" BIGINT NOT NULL,
  "quantity" DOUBLE PRECISION NOT NULL,
  "rest_quantity" DOUBLE PRECISION NOT NULL,
  "type" TEXT NOT NULL,
  "reporting_timestamp" TIMESTAMPTZ NOT NULL,
  "actual_reported_date" TIMESTAMPTZ,
  "received_at" TIMESTAMPTZ NOT NULL,
  "ingested_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "last_seen_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  "sync_status" TEXT NOT NULL DEFAULT 'pending',
  "sync_error" TEXT,
  "retry_count" INTEGER NOT NULL DEFAULT 0,
  "next_retry_at" TIMESTAMPTZ,
  "last_synced_at" TIMESTAMPTZ,
  PRIMARY KEY ("company_id", "id")
);

CREATE INDEX IF NOT EXISTS "idx_raw_order_sync_queue"
  ON "erp_raw"."order" ("sync_status", "next_retry_at", "received_at");

CREATE INDEX IF NOT EXISTS "idx_raw_order_operation_sync_queue"
  ON "erp_raw"."order_operation" ("sync_status", "next_retry_at", "received_at");

CREATE INDEX IF NOT EXISTS "idx_raw_order_report_sync_queue"
  ON "erp_raw"."order_report" ("sync_status", "next_retry_at", "received_at");

CREATE INDEX IF NOT EXISTS "idx_raw_production_resource_sync_queue"
  ON "erp_raw"."production_resource" ("sync_status", "next_retry_at", "received_at");

-- +goose Down
DROP SCHEMA IF EXISTS "erp_raw" CASCADE;
