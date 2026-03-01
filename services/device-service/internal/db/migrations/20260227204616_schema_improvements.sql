-- +goose Up

-- Rename device status enum to device_state for clarity. 
ALTER TYPE "device"."device_status"
    RENAME TO "device_state";

-- company config: add chirpstack application id (one application per company, created during onboarding)
-- Add with a temporary default so existing rows are not rejected, then drop the default
-- so future inserts are required to supply a real value.
ALTER TABLE "device"."company_config"
    ADD COLUMN "chirpstack_application_id" varchar UNIQUE NOT NULL DEFAULT 'dev-placeholder';
ALTER TABLE "device"."company_config"
    ALTER COLUMN "chirpstack_application_id" DROP DEFAULT;

COMMENT ON COLUMN "device"."company_config"."chirpstack_application_id" IS 'One ChirpStack application per company — created during onboarding';

-- gateway: add description, rename status to state.
ALTER TABLE "device"."gateway"
    ADD COLUMN "description" varchar;
ALTER TABLE "device"."gateway"
    RENAME COLUMN "status" TO "state";
ALTER TABLE "device"."gateway"
    DROP COLUMN "chirpstack_gateway_id";

-- sensor: add description, rename status to state, add chirpstack profile id.
ALTER TABLE "device"."sensor"
    ADD COLUMN "description" varchar;
ALTER TABLE "device"."sensor"
    RENAME COLUMN "status" TO "state";
-- Same pattern: temporary default for existing rows, then drop it.
ALTER TABLE "device"."sensor"
    ADD COLUMN "chirpstack_profile_id" varchar NOT NULL DEFAULT 'dev-placeholder';
ALTER TABLE "device"."sensor"
    ALTER COLUMN "chirpstack_profile_id" DROP DEFAULT;
ALTER TABLE "device"."sensor"
    DROP COLUMN "chirpstack_device_id";

COMMENT ON COLUMN "device"."sensor"."chirpstack_profile_id" IS 'The ChirpStack profile ID associated with this sensor, selected during sensor registration';

-- sensor: remove FK constraint and drop gateway_id (gateways are infrastructure)
ALTER TABLE "device"."sensor"
    DROP CONSTRAINT "sensor_gateway_id_fkey";
ALTER TABLE "device"."sensor"
    DROP COLUMN "gateway_id";

-- Adding on delete cascade to sensor_metric, so when a sensor is deleted, its metrics are also deleted
ALTER TABLE "device"."sensor_metric"
    DROP CONSTRAINT "sensor_metric_sensor_id_fkey";

ALTER TABLE "device"."sensor_metric"
    ADD CONSTRAINT "sensor_metric_sensor_id_fkey" 
    FOREIGN KEY ("sensor_id") 
    REFERENCES "device"."sensor"("sensor_id") 
    ON DELETE CASCADE;

-- +goose Down

-- Revert sensor_metric foreign key to no cascade
ALTER TABLE "device"."sensor_metric"
    DROP CONSTRAINT "sensor_metric_sensor_id_fkey";

ALTER TABLE "device"."sensor_metric"
    ADD CONSTRAINT "sensor_metric_sensor_id_fkey"
    FOREIGN KEY ("sensor_id")
    REFERENCES "device"."sensor"("sensor_id");

ALTER TABLE "device"."sensor"
    ADD COLUMN "gateway_id" uuid;
ALTER TABLE "device"."sensor"
    ADD CONSTRAINT "sensor_gateway_id_fkey"
    FOREIGN KEY ("gateway_id")
    REFERENCES "device"."gateway"("gateway_id");


ALTER TABLE "device"."sensor"
    ADD COLUMN "chirpstack_device_id" varchar;
ALTER TABLE "device"."sensor"
    DROP COLUMN "chirpstack_profile_id";
ALTER TABLE "device"."sensor"
    RENAME COLUMN "state" TO "status";
ALTER TABLE "device"."sensor"
    DROP COLUMN "description";

ALTER TABLE "device"."gateway"
    ADD COLUMN "chirpstack_gateway_id" varchar;
ALTER TABLE "device"."gateway"
    RENAME COLUMN "state" TO "status";
ALTER TABLE "device"."gateway"
    DROP COLUMN "description";

ALTER TABLE "device"."company_config"
    DROP COLUMN "chirpstack_application_id";

ALTER TYPE "device"."device_state"
    RENAME TO "device_status";
