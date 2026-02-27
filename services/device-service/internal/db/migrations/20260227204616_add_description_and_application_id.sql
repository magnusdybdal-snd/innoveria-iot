-- +goose Up
ALTER TYPE "device"."device_status"
    RENAME TO "device_state";

ALTER TABLE "device"."company_config"
    ADD COLUMN "chirpstack_application_id" varchar UNIQUE NOT NULL;
ALTER TABLE "device"."company_config"
    COMMENT ON COLUMN "chirpstack_application_id" IS 'One ChirpStack application per company — created during onboarding';

ALTER TABLE "device"."gateway"
    ADD COLUMN "description" varchar;
ALTER TABLE "device"."gateway"
    RENAME COLUMN "status" TO "state";

ALTER TABLE "device"."sensor"
    ADD COLUMN "description" varchar;
ALTER TABLE "device"."sensor"
    RENAME COLUMN "status" TO "state";


-- +goose Down
ALTER TABLE "device"."sensor"
    RENAME COLUMN "state" TO "status";
ALTER TABLE "device"."sensor"
    DROP COLUMN "description";

ALTER TABLE "device"."gateway"
    RENAME COLUMN "state" TO "status";
ALTER TABLE "device"."gateway"
    DROP COLUMN "description";

ALTER TABLE "device"."company_config"
    DROP COLUMN "chirpstack_application_id";

ALTER TYPE "device"."device_state"
    RENAME TO "device_status";
