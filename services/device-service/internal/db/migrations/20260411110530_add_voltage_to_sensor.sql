-- +goose Up
ALTER TABLE "device"."sensor"
    ADD COLUMN "voltage" integer;


-- +goose Down
ALTER TABLE "device"."sensor"
    DROP COLUMN "voltage";