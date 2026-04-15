-- +goose Up

-- Remove any existing draft rows before enforcing NOT NULL.
-- Safe to run unconditionally: the draft concept was never merged to dev or deployed to prod,
-- so no real data exists with measurement_type IS NULL.
DELETE FROM device.payload_schema WHERE measurement_type IS NULL;

ALTER TABLE device.payload_schema
    ALTER COLUMN measurement_type SET NOT NULL;

COMMENT ON TABLE device.payload_schema IS 'Profile-level mapping from raw payload keys to canonical measurement types.';
COMMENT ON COLUMN device.payload_schema.measurement_type IS 'Canonical measurement type slug — references device.measurement_type(slug).';

-- +goose Down

ALTER TABLE device.payload_schema
    ALTER COLUMN measurement_type DROP NOT NULL;

COMMENT ON TABLE device.payload_schema IS 'Profile-level mapping from raw payload keys to canonical measurement types. Rows with measurement_type = NULL are drafts awaiting admin labeling.';
COMMENT ON COLUMN device.payload_schema.measurement_type IS 'NULL = draft (discovered but not yet labeled by admin)';
