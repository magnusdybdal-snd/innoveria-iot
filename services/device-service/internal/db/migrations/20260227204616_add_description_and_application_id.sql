-- +goose Up
ALTER TYPE "device"."device_status"
    RENAME TO "device_state";

ALTER TABLE "device"."company_config"
    ADD COLUMN "chirpstack_application_id" varchar UNIQUE NOT NULL;

COMMENT ON COLUMN "device"."company_config"."chirpstack_application_id" IS 'One ChirpStack application per company — created during onboarding';

ALTER TABLE "device"."gateway"
    ADD COLUMN "description" varchar;
ALTER TABLE "device"."gateway"
    RENAME COLUMN "status" TO "state";

ALTER TABLE "device"."sensor"
    ADD COLUMN "description" varchar;
ALTER TABLE "device"."sensor"
    RENAME COLUMN "status" TO "state";

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
    RENAME COLUMN "state" TO "status";
ALTER TABLE "device"."sensor"
    DROP COLUMN "description";

ALTER TABLE "device"."gateway"
    RENAME COLUMN "state" TO "status";
ALTER TABLE "device"."gateway"
    DROP COLUMN "description";

ALTER TABLE "device"."company_config"
    DROP COLUMN "chirpstack_application_id";

ALTER TYPE "device"."device_state"
    RENAME TO "device_status";
