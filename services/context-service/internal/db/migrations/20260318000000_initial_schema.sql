-- +goose Up
CREATE SCHEMA IF NOT EXISTS "context";

CREATE TABLE "context"."aggregation_rule" (
  rule_id               UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  company_id            UUID NOT NULL,
  name                  VARCHAR NOT NULL,
  context_type          VARCHAR NOT NULL,
  measurement_type      VARCHAR NOT NULL,
  aggregation_method    VARCHAR NOT NULL CHECK (aggregation_method IN ('AVG', 'SUM', 'MAX', 'MIN')),
  time_bucket_minutes   INTEGER NOT NULL,
  is_active             BOOLEAN NOT NULL DEFAULT TRUE,
  created_at            TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at            TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (company_id, context_type)
);

CREATE TABLE "context"."context_data" (
  context_id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  company_id             UUID NOT NULL,
  context_type           VARCHAR NOT NULL,
  order_id               UUID,
  production_resource_id UUID,
  device_eui             VARCHAR NOT NULL,
  value                  DOUBLE PRECISION NOT NULL,
  unit                   VARCHAR NOT NULL,
  period_start           TIMESTAMPTZ NOT NULL,
  period_end             TIMESTAMPTZ NOT NULL,
  calculated_at          TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (company_id, device_eui, context_type, period_start, period_end)
);

CREATE INDEX idx_context_data_company_type_calc
  ON "context".context_data (company_id, context_type, calculated_at);

-- +goose Down
DROP SCHEMA "context" CASCADE;
