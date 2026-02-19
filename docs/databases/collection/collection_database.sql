CREATE SCHEMA "collection";

CREATE TABLE "collection"."tenant_mapping" (
  "chirpstack_tenant_id" varchar PRIMARY KEY,
  "company_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "collection"."sensor_measurement" (
  "measurement_id" uuid DEFAULT (gen_random_uuid()),
  "device_eui" varchar NOT NULL,
  "timestamp" timestamptz NOT NULL,
  "payload" jsonb NOT NULL,
  "company_id" uuid NOT NULL
);

SELECT create_hypertable('collection.sensor_measurement', 'timestamp');

CREATE INDEX "idx_measurement_time_device" ON "collection"."sensor_measurement" ("timestamp", "device_eui");

CREATE INDEX "idx_measurement_company_time" ON "collection"."sensor_measurement" ("company_id", "timestamp");

CREATE INDEX "idx_measurement_payload" ON "collection"."sensor_measurement" USING GIN ("payload");

COMMENT ON COLUMN "collection"."tenant_mapping"."chirpstack_tenant_id" IS 'Tenant ID from ChirpStack MQTT messages';

COMMENT ON COLUMN "collection"."tenant_mapping"."company_id" IS 'Our internal company ID';

COMMENT ON TABLE "collection"."sensor_measurement" IS 'Hypertable: SELECT create_hypertable(''sensor_measurement'', ''timestamp'');';

COMMENT ON COLUMN "collection"."sensor_measurement"."measurement_id" IS 'Not PK — TimescaleDB uses (timestamp, device_eui) for partitioning';

COMMENT ON COLUMN "collection"."sensor_measurement"."device_eui" IS 'Identifies the sensor — matches device.sensor.device_eui';

COMMENT ON COLUMN "collection"."sensor_measurement"."timestamp" IS 'Hypertable partition key';

COMMENT ON COLUMN "collection"."sensor_measurement"."company_id" IS 'Mapped from chirpstackTenantId via collection.tenant_mapping at ingest time';
