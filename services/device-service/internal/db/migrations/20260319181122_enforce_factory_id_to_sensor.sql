-- +goose Up

-- Make factory_area_id NOT NULL — every sensor must be assigned to a factory area.
-- Backfill existing rows with a dev placeholder before adding the constraint.
UPDATE "device"."sensor"
    SET "factory_area_id" = 'f2000000-0000-0000-0000-000000000001'
    WHERE "factory_area_id" IS NULL;

ALTER TABLE "device"."sensor"
    ALTER COLUMN "factory_area_id" SET NOT NULL;

COMMENT ON COLUMN "device"."sensor"."factory_area_id" IS 'Reference to the factory area to which the sensor belongs.';

-- +goose Down

ALTER TABLE "device"."sensor"
    ALTER COLUMN "factory_area_id" DROP NOT NULL;
