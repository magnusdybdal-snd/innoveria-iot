-- +goose Up

ALTER TABLE "device"."sensor_metric"
    DROP CONSTRAINT "sensor_metric_pkey";

ALTER TABLE "device"."sensor_metric"
    ADD COLUMN "id" uuid NOT NULL DEFAULT gen_random_uuid();

ALTER TABLE "device"."sensor_metric"
    ADD PRIMARY KEY ("id");

ALTER TABLE "device"."sensor_metric"
    ADD COLUMN "payload_key" text NOT NULL DEFAULT '';

ALTER TABLE "device"."sensor_metric"
    ALTER COLUMN "payload_key" DROP DEFAULT;

ALTER TABLE "device"."sensor_metric"
    ADD CONSTRAINT "sensor_metric_sensor_id_payload_key_key"
    UNIQUE ("sensor_id", "payload_key");

ALTER TABLE "device"."sensor_metric"
    ADD CONSTRAINT "sensor_metric_measurement_type_fkey"
    FOREIGN KEY ("measurement_type")
    REFERENCES "device"."measurement_type"("slug");

COMMENT ON COLUMN "device"."sensor_metric"."payload_key" IS 'The actual key in the raw JSON payload, e.g. "bus1", "AccAmp". Used to extract the correct field for aggregation.';

-- +goose Down

ALTER TABLE "device"."sensor_metric"
    DROP CONSTRAINT "sensor_metric_measurement_type_fkey";

ALTER TABLE "device"."sensor_metric"
    DROP CONSTRAINT "sensor_metric_sensor_id_payload_key_key";

ALTER TABLE "device"."sensor_metric"
    DROP COLUMN "payload_key";

ALTER TABLE "device"."sensor_metric"
    DROP CONSTRAINT "sensor_metric_pkey";

ALTER TABLE "device"."sensor_metric"
    DROP COLUMN "id";

ALTER TABLE "device"."sensor_metric"
    ADD PRIMARY KEY ("sensor_id", "measurement_type");
