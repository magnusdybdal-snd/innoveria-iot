-- Change production_resource_id from uuid to bigint to match the ERP domain model
-- where ProductionResource.ID is an int64 assigned by Monitor ERP.
ALTER TABLE device.sensor
    ALTER COLUMN production_resource_id TYPE BIGINT USING NULL;
