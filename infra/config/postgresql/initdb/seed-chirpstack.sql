INSERT INTO tenant (id, created_at, updated_at, name, description,
    can_have_gateways, max_gateway_count, max_device_count,
    private_gateways_up, private_gateways_down, tags)
VALUES ('d0000000-0000-0000-0000-000000000001', NOW(), NOW(),
    'Innoveria Dev', '', true, 0, 0, false, false, '{}')
ON CONFLICT (id) DO NOTHING;

INSERT INTO application (id, tenant_id, created_at, updated_at, name,
description, mqtt_tls_cert, tags)
VALUES ('e0000000-0000-0000-0000-000000000001',
    'd0000000-0000-0000-0000-000000000001',
    NOW(), NOW(), 'Innoveria Dev App', '', '\x', '{}')
ON CONFLICT (id) DO NOTHING;
