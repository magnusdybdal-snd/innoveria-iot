CREATE EXTENSION IF NOT EXISTS timescaledb;

CREATE TABLE device_measurements (
    time        TIMESTAMPTZ       NOT NULL,
    device_id   TEXT              NOT NULL,
    temperature DOUBLE PRECISION,
    humidity    DOUBLE PRECISION,
    battery     DOUBLE PRECISION
);

CREATE INDEX ON device_measurements (device_id, time DESC);
