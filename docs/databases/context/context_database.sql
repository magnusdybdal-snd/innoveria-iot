CREATE SCHEMA "context";

CREATE TABLE "context"."context_data" (
  "context_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "context_type" varchar NOT NULL,
  "order_id" uuid,
  "production_resource_id" uuid,
  "device_eui" varchar,
  "value" double NOT NULL,
  "unit" varchar NOT NULL,
  "period_start" timestamptz,
  "period_end" timestamptz,
  "calculated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "context"."dashboard_config" (
  "config_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "user_id" uuid,
  "name" varchar NOT NULL,
  "config_json" jsonb NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "context"."aggregation_rule" (
  "rule_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "context_type" varchar NOT NULL,
  "measurement_type" varchar NOT NULL,
  "aggregation_method" varchar NOT NULL,
  "time_bucket_minutes" integer NOT NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX "idx_context_company_type_time" ON "context"."context_data" ("company_id", "context_type", "calculated_at");

COMMENT ON COLUMN "context"."context_data"."context_type" IS 'e.g. nitrogen_per_order, temperature_avg_hourly';

COMMENT ON COLUMN "context"."context_data"."order_id" IS 'Loose reference to erp.order — the order this context relates to';

COMMENT ON COLUMN "context"."context_data"."production_resource_id" IS 'Loose ref to erp.production_resource';

COMMENT ON COLUMN "context"."context_data"."device_eui" IS 'Which sensor produced the underlying data';

COMMENT ON COLUMN "context"."context_data"."period_start" IS '[ADDED] Start of the period this value covers';

COMMENT ON COLUMN "context"."context_data"."period_end" IS '[ADDED] End of the period this value covers';

COMMENT ON COLUMN "context"."dashboard_config"."user_id" IS 'Loose ref to auth.user — null means company-wide default';

COMMENT ON COLUMN "context"."dashboard_config"."config_json" IS 'Widget layout, selected sensors, time ranges, etc.';

COMMENT ON COLUMN "context"."aggregation_rule"."context_type" IS 'What type of context_data this rule produces';

COMMENT ON COLUMN "context"."aggregation_rule"."measurement_type" IS 'Which sensor measurement type to aggregate';

COMMENT ON COLUMN "context"."aggregation_rule"."aggregation_method" IS 'e.g. SUM, AVG, MAX, MIN';

COMMENT ON COLUMN "context"."aggregation_rule"."time_bucket_minutes" IS 'e.g. 60 for hourly, 1440 for daily';
