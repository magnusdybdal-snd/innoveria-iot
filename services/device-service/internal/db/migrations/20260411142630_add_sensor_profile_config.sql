-- +goose Up
CREATE TABLE "device"."sensor_profile_config" (
    "chirpstack_profile_id" text PRIMARY KEY,
    "configurable_schema"   boolean NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE "device"."sensor_profile_config";
