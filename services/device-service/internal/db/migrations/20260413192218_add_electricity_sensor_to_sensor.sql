-- +goose Up
ALTER TABLE "device"."sensor"
    ADD COLUMN "electricity_sensor" boolean NOT NULL DEFAULT false;


-- +goose Down
ALTER TABLE "device"."sensor"
    DROP COLUMN "electricity_sensor";
