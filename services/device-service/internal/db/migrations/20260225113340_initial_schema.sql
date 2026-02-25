-- +goose Up                                                                                                                          CREATE SCHEMA "device";

CREATE TYPE "device"."device_status" AS ENUM (
  'ACTIVE',
  'INACTIVE'
);

CREATE TABLE "device"."company_config" (
  "company_id"           uuid PRIMARY KEY,
  "chirpstack_tenant_id" varchar UNIQUE NOT NULL,
  "created_at"           timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "device"."company_config"."company_id" IS 'Same UUID as auth.company — populated during onboarding';
COMMENT ON COLUMN "device"."company_config"."chirpstack_tenant_id" IS 'Created by Device Service via ChirpStack API';

CREATE TABLE "device"."gateway" (
  "gateway_id"            uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id"            uuid NOT NULL,
  "gateway_eui"           varchar UNIQUE NOT NULL,
  "chirpstack_gateway_id" varchar,
  "name"                  varchar NOT NULL,
  "status"                device.device_status NOT NULL DEFAULT 'ACTIVE',
  "factory_area_id"       uuid,
  "created_at"            timestamptz NOT NULL DEFAULT (now()),
  "updated_at"            timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "device"."gateway"."gateway_eui" IS 'Shared key with ChirpStack';
COMMENT ON COLUMN "device"."gateway"."chirpstack_gateway_id" IS 'ChirpStack internal ID for API calls';
COMMENT ON COLUMN "device"."gateway"."factory_area_id" IS 'References auth.factory_area — loose cross-service ref';

CREATE TABLE "device"."sensor" (
  "sensor_id"             uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "company_id"            uuid NOT NULL,
  "gateway_id"            uuid REFERENCES "device"."gateway"("gateway_id"),
  "device_eui"            varchar UNIQUE NOT NULL,
  "chirpstack_device_id"  varchar,
  "production_resource_id" uuid,
  "factory_area_id"       uuid,
  "name"                  varchar NOT NULL,
  "status"                device.device_status NOT NULL DEFAULT 'ACTIVE',
  "created_at"            timestamptz NOT NULL DEFAULT (now()),
  "updated_at"            timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "device"."sensor"."company_id" IS 'Tenant isolation — references auth.company but no FK across DBs';
COMMENT ON COLUMN "device"."sensor"."gateway_id" IS 'Gateway this sensor is administered under — nullable until assigned';
COMMENT ON COLUMN "device"."sensor"."device_eui" IS 'Shared key with ChirpStack. Bound to sensor hardware';
COMMENT ON COLUMN "device"."sensor"."chirpstack_device_id" IS 'ChirpStack internal ID for API calls';
COMMENT ON COLUMN "device"."sensor"."production_resource_id" IS 'References erp.production_resource — references the machine sensor
is installed on';
COMMENT ON COLUMN "device"."sensor"."factory_area_id" IS 'References auth.factory_area — loose cross-service ref';

CREATE TABLE "device"."sensor_metric" (
  "sensor_id"        uuid NOT NULL REFERENCES "device"."sensor"("sensor_id"),
  "measurement_type" varchar NOT NULL,
  "unit"             varchar,
  PRIMARY KEY ("sensor_id", "measurement_type")
);

COMMENT ON COLUMN "device"."sensor_metric"."measurement_type" IS 'e.g. temperature, humidity, electric_current';
COMMENT ON COLUMN "device"."sensor_metric"."unit" IS 'e.g. °C, %RH, A — useful for display and validation';

-- +goose Down
DROP SCHEMA "device" CASCADE;