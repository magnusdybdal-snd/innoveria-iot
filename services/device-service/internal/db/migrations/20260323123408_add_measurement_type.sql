-- +goose Up

CREATE TABLE "device"."measurement_type" (
  "slug"         text PRIMARY KEY,
  "display_name" text NOT NULL,
  "description"  text,
  "default_unit" text,
  "deprecated"   boolean NOT NULL DEFAULT false
);

COMMENT ON COLUMN "device"."measurement_type"."slug" IS 'Immutable identifier used as the cross-service contract. e.g. accumulated_current, nitrogen_ppm';
COMMENT ON COLUMN "device"."measurement_type"."deprecated" IS 'Retired types are hidden from dropdowns but never deleted — existing references remain valid';

-- +goose Down

DROP TABLE "device"."measurement_type";
