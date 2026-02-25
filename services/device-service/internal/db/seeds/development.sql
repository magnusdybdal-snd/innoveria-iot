-- Company config (matches collection-service seed for cross-service consistency)
INSERT INTO "device"."company_config" ("company_id", "chirpstack_tenant_id")
VALUES ('a0000000-0000-0000-0000-000000000001', 'NTNU-12345')
ON CONFLICT DO NOTHING;

-- Gateway
INSERT INTO "device"."gateway" ("gateway_id", "company_id", "gateway_eui", "chirpstack_gateway_id", "name")
VALUES ('b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'AA:BB:CC:DD:EE:FF:00:01', 'mock-gw-001',
'Dev Gateway')
ON CONFLICT DO NOTHING;

-- Sensors
INSERT INTO "device"."sensor" ("sensor_id", "company_id", "gateway_id", "device_eui", "chirpstack_device_id", "name")
VALUES
('c0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001',
'AA:BB:CC:DD:EE:FF:01:01', 'mock-dev-001', 'Climate Sensor A'),
('c0000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000001', 'b0000000-0000-0000-0000-000000000001',
'AA:BB:CC:DD:EE:FF:01:02', 'mock-dev-002', 'Current Sensor A')
ON CONFLICT DO NOTHING;

-- Sensor metrics
INSERT INTO "device"."sensor_metric" ("sensor_id", "measurement_type", "unit")
VALUES
('c0000000-0000-0000-0000-000000000001', 'temperature', '°C'),
('c0000000-0000-0000-0000-000000000001', 'humidity',    '%RH'),
('c0000000-0000-0000-0000-000000000002', 'electric_current', 'A')
ON CONFLICT DO NOTHING;