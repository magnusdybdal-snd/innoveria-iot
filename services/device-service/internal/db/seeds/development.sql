-- Company config (matches collection-service seed for cross-service consistency)
-- chirpstack_application_id: replace with real value before demo
INSERT INTO "device"."company_config" ("company_id", "chirpstack_tenant_id", "chirpstack_application_id")
VALUES ('a0000000-0000-0000-0000-000000000001', 'NTNU-12345', 'mock-app-001')
ON CONFLICT DO NOTHING;

-- Gateway
INSERT INTO "device"."gateway" ("gateway_id", "company_id", "gateway_eui", "name", "state")
VALUES ('b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'AA:BB:CC:DD:EE:FF:00:01', 'Dev Gateway', 'ACTIVE')
ON CONFLICT DO NOTHING;

-- Sensors
-- chirpstack_profile_id: replace with real value before demo
INSERT INTO "device"."sensor" ("sensor_id", "company_id", "device_eui", "name", "state", "chirpstack_profile_id")
VALUES
('c0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'AA:BB:CC:DD:EE:FF:01:01', 'Climate Sensor A',  'ACTIVE', 'mock-profile-001'),
('c0000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000001', 'AA:BB:CC:DD:EE:FF:01:02', 'Current Sensor A',  'ACTIVE', 'mock-profile-001')
ON CONFLICT DO NOTHING;

-- Sensor metrics
INSERT INTO "device"."sensor_metric" ("sensor_id", "measurement_type", "unit")
VALUES
('c0000000-0000-0000-0000-000000000001', 'temperature',      '°C'),
('c0000000-0000-0000-0000-000000000001', 'humidity',         '%RH'),
('c0000000-0000-0000-0000-000000000002', 'electric_current', 'A')
ON CONFLICT DO NOTHING;
