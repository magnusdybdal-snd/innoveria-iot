-- +goose Up
CREATE SCHEMA IF NOT EXISTS "erp";

-- +goose Down
DROP SCHEMA IF EXISTS "erp" CASCADE;
