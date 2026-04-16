-- Change production_resource_id from uuid to bigint to match the ERP service's ProductionResource.ID (int64).
-- All existing rows have NULL in this column so there is no data to migrate.
ALTER TABLE device.sensor DROP COLUMN production_resource_id;
ALTER TABLE device.sensor ADD COLUMN production_resource_id bigint;
