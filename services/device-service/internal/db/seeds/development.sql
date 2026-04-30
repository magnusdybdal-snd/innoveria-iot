-- Company config (matches collection-service seed for cross-service consistency)
-- chirpstack_application_id: replace with real value before demo
INSERT INTO "device"."company_config" ("company_id", "chirpstack_tenant_id", "chirpstack_application_id")
VALUES ('a0000000-0000-0000-0000-000000000001', 'd0000000-0000-0000-0000-000000000001', 'e0000000-0000-0000-0000-000000000001')
ON CONFLICT DO NOTHING;

-- Gateway (matches chirpstack seed.)
INSERT INTO "device"."gateway" ("gateway_id", "company_id", "gateway_eui", "name", "state", "factory_id", "factory_area_id")
VALUES ('b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'a000000000000001', 'Dev Gateway', 'ACTIVE', 'f1000000-0000-0000-0000-000000000001', 'a1000000-0000-0000-0000-000000000001')
ON CONFLICT DO NOTHING;

-- Measurement types (vocabulary required by sensor_metric and payload_schema FKs)
INSERT INTO "device"."measurement_type" ("slug", "display_name", "description", "default_unit")
VALUES
('temperature',       'Temperature',        'Ambient or surface temperature',    '°C'),
('humidity',          'Relative Humidity',  'Relative humidity percentage',      '%'),
('electric_current',  'Electric Current',   'RMS current draw',                  'A'),
('voltage',           'Voltage',            'Line or supply voltage',            'V')
ON CONFLICT DO NOTHING;

-- Sensors (matches chirpstack seed.)
-- Sensor 1: profile A, has sensor_metric rows  → tests sensor-level override path
-- Sensor 2: profile A, no sensor_metric rows   → tests fallback to payload_schema
INSERT INTO "device"."sensor" ("sensor_id", "company_id", "device_eui", "app_key", "factory_id", "factory_area_id", "name", "state", "chirpstack_profile_id", "global_chirpstack_profile_id")
VALUES
('c0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'b000000000000001', '00000000000000000000000000000001', 'f1000000-0000-0000-0000-000000000001', 'a1000000-0000-0000-0000-000000000001', 'Dev Sensor 1 (override)',  'ACTIVE', 'f0000000-0000-0000-0000-000000000001', 'f0000000-0000-0000-0000-000000000001'),
('c0000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000001', 'b000000000000002', '00000000000000000000000000000002', 'f1000000-0000-0000-0000-000000000001', 'a1000000-0000-0000-0000-000000000002', 'Dev Sensor 2 (fallback)',  'ACTIVE', 'f0000000-0000-0000-0000-000000000001', 'f0000000-0000-0000-0000-000000000001')
ON CONFLICT DO NOTHING;

-- Payload schema for profile A (f0000000-0000-0000-0000-000000000001)
-- Fully labeled — used by sensor 2 fallback path
INSERT INTO "device"."payload_schema" ("chirpstack_profile_id", "payload_key", "measurement_type", "unit")
VALUES
('f0000000-0000-0000-0000-000000000001', 'temperature', 'temperature',      '°C'),
('f0000000-0000-0000-0000-000000000001', 'humidity',    'humidity',         '%')
ON CONFLICT DO NOTHING;

-- Sensor metrics for sensor 1 only (override path)
-- payload_key maps to the raw JSON field names sent by the simulator
INSERT INTO "device"."sensor_metric" ("sensor_id", "payload_key", "measurement_type", "unit")
VALUES
('c0000000-0000-0000-0000-000000000001', 'temperature', 'temperature', '°C'),
('c0000000-0000-0000-0000-000000000001', 'humidity',    'humidity',    '%')
ON CONFLICT DO NOTHING;
