-- Company config (matches collection-service seed for cross-service consistency)
-- chirpstack_application_id: replace with real value before demo
INSERT INTO "device"."company_config" ("company_id", "chirpstack_tenant_id", "chirpstack_application_id")
VALUES ('a0000000-0000-0000-0000-000000000001', '9d878067-58d3-4e3c-962e-f200256131ca', '2a7c4c7e-e38d-4fce-9214-e63839a042c0')
ON CONFLICT DO NOTHING;

-- Gateway
INSERT INTO "device"."gateway" ("gateway_id", "company_id", "gateway_eui", "name", "state")
VALUES ('b0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'ac1f09fffe1f628b', 'Dev Gateway', 'ACTIVE')
ON CONFLICT DO NOTHING;

-- Sensors
-- chirpstack_profile_id: replace with real value before demo
INSERT INTO "device"."sensor" ("sensor_id", "company_id", "device_eui", "name", "state", "chirpstack_profile_id")
VALUES
('c0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'f62ccf710469ad3b', 'Stromsensor Homag',  'ACTIVE', 'b2eff42b-4cde-4d6b-bc9b-bdb410ae8a8e'),
ON CONFLICT DO NOTHING;

-- Sensor metrics
INSERT INTO "device"."sensor_metric" ("sensor_id", "measurement_type", "unit")
VALUES
('c0000000-0000-0000-0000-000000000001', 'electric_current', 'A'),
ON CONFLICT DO NOTHING;
