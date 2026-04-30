-- +goose Up

ALTER TABLE device.sensor
    ADD COLUMN global_chirpstack_profile_id varchar;

-- Best-effort backfill: existing sensors get chirpstack_profile_id as a stand-in.
-- These rows should be manually corrected by looking up the original global profile
-- IDs in Chirpstack for any sensors created before this migration.
UPDATE device.sensor
    SET global_chirpstack_profile_id = chirpstack_profile_id
    WHERE global_chirpstack_profile_id IS NULL;

ALTER TABLE device.sensor
    ALTER COLUMN global_chirpstack_profile_id SET NOT NULL;

-- +goose Down

ALTER TABLE device.sensor
    DROP COLUMN global_chirpstack_profile_id;
