-- +goose Up

-- Add application key and factory area ID to sensor for better integration with ChirpStack and factory context.
ALTER TABLE "device"."sensor"
    ADD COLUMN "app_key" varchar NOT NULL DEFAULT 'dev-placeholder';
ALTER TABLE "device"."sensor"
    ALTER COLUMN "app_key" DROP DEFAULT;

COMMENT ON COLUMN "device"."sensor"."app_key" IS 'LoRaWAN OTAA AppKey - used to authenticate the sensor with the LoRaWAN network.';

ALTER TABLE "device"."sensor"
    ADD COLUMN "factory_id" uuid NOT NULL DEFAULT 'f1000000-0000-0000-0000-000000000001';
ALTER TABLE "device"."sensor"
    ALTER COLUMN "factory_id" DROP DEFAULT;

COMMENT ON COLUMN "device"."sensor"."factory_id" IS 'Reference to the factory to which the sensor belongs.';

-- +goose Down
ALTER TABLE "device"."sensor"
    DROP COLUMN "app_key";
ALTER TABLE "device"."sensor"
    DROP COLUMN "factory_id";
