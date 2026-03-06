-- Company config (matches collection-service seed for cross-service consistency)
-- chirpstack_application_id: replace with real value before demo
INSERT INTO "device"."company_config" ("company_id", "chirpstack_tenant_id", "chirpstack_application_id")
VALUES ('a0000000-0000-0000-0000-000000000001', '9d878067-58d3-4e3c-962e-f200256131ca', '2a7c4c7e-e38d-4fce-9214-e63839a042c0')
ON CONFLICT DO NOTHING;

-- Gateway (matches chirpstack seed.)
INSERT INTO "device"."gateway" ("gateway_id", "company_id", "gateway_eui", "name", "state")
VALUES ('b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'a000000000000001', 'Dev Gateway', 'ACTIVE')
ON CONFLICT DO NOTHING;

-- Sensors (matches chirpstack seed.)
INSERT INTO "device"."sensor" ("sensor_id", "company_id", "device_eui", "name", "state", "chirpstack_profile_id")
VALUES
('c0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'b000000000000001', 'Dev Sensor 1', 'ACTIVE', 'f0000000-0000-0000-0000-000000000001'),
('c0000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000001', 'b000000000000002', 'Dev Sensor 2', 'ACTIVE', 'f0000000-0000-0000-0000-000000000001'),
('c0000000-0000-0000-0000-000000000003', 'a0000000-0000-0000-0000-000000000001', 'b000000000000003', 'Dev Sensor 3', 'ACTIVE', 'f0000000-0000-0000-0000-000000000001')
ON CONFLICT DO NOTHING;


-- Sensor metrics
INSERT INTO "device"."sensor_metric" ("sensor_id", "measurement_type", "unit")
VALUES
('c0000000-0000-0000-0000-000000000001', 'electric_current', 'A'),
('c0000000-0000-0000-0000-000000000002', 'temperature', 'C'),
('c0000000-0000-0000-0000-000000000003', 'humidity', '%')
ON CONFLICT DO NOTHING;
