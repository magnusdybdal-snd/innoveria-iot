-- +goose Up

CREATE TABLE "device"."payload_schema" (
  "id"                    uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "chirpstack_profile_id" text NOT NULL,
  "payload_key"           text NOT NULL,
  "measurement_type"      text REFERENCES "device"."measurement_type"("slug"),
  "unit"                  text,
  UNIQUE ("chirpstack_profile_id", "payload_key")
);

COMMENT ON TABLE "device"."payload_schema" IS 'Profile-level mapping from raw payload keys to canonical measurement types. Rows with measurement_type = NULL are drafts awaiting admin labeling.';
COMMENT ON COLUMN "device"."payload_schema"."chirpstack_profile_id" IS 'References the ChirpStack device profile ID stored on device.sensor';
COMMENT ON COLUMN "device"."payload_schema"."measurement_type" IS 'NULL = draft (discovered but not yet labeled by admin)';

-- +goose Down

DROP TABLE "device"."payload_schema";
