-- +goose Up
CREATE INDEX idx_measurement_company_device_time
    ON collection.sensor_measurement (company_id, device_eui, timestamp DESC);

-- +goose Down
DROP INDEX IF EXISTS collection.idx_measurement_company_device_time;
