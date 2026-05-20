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

-- Dev device profile (required before devices)
INSERT INTO device_profile (
    id, tenant_id, created_at, updated_at,
    name, description, region, mac_version, reg_params_revision,
    adr_algorithm_id, payload_codec_runtime, payload_codec_script,
    uplink_interval, device_status_req_interval,
    supports_otaa, supports_class_b, supports_class_c,
    flush_queue_on_activate, auto_detect_measurements, allow_roaming,
    rx1_delay, firmware_version, vendor_profile_id,
    supported_uplink_data_rates, tags, measurements, app_layer_params
)
VALUES (
    'f0000000-0000-0000-0000-000000000001',
    'd0000000-0000-0000-0000-000000000001',
    NOW(), NOW(),
    'Dev Sensor Profile', '', 'EU868', '1.0.3', 'A',
    'default', 'NONE', '',
    3600, 1,
    true, false, false,
    true, false, false,
    0, '', 0,
    '{}', '{}', '{}',
    '{"ts003_f_port": 202, "ts004_f_port": 201, "ts005_f_port": 200, "ts003_version": null, "ts004_version": null, "ts005_version": null}'
)
ON CONFLICT (id) DO NOTHING;

-- Dev gateway
INSERT INTO gateway (
    gateway_id, tenant_id, created_at, updated_at,
    name, description, latitude, longitude, altitude,
    stats_interval_secs, tags, properties
)
VALUES (
    decode('a000000000000001', 'hex'),
    'd0000000-0000-0000-0000-000000000001',
    NOW(), NOW(),
    'Dev Gateway 1', '', 0.0, 0.0, 0.0,
    30, '{}', '{}'
)
ON CONFLICT (gateway_id) DO NOTHING;

-- Dev sensors
INSERT INTO device (
    dev_eui, application_id, device_profile_id, created_at, updated_at,
    name, description, external_power_source, enabled_class,
    skip_fcnt_check, is_disabled, tags, variables,
    join_eui, app_layer_params, f_cnt_up
)
VALUES
    (decode('b000000000000001', 'hex'), 'e0000000-0000-0000-0000-000000000001', 'f0000000-0000-0000-0000-000000000001', NOW(), NOW(), 'Dev Sensor 1', '', false, 'A', false, false, '{}', '{}', decode('0000000000000000', 'hex'), '{}', 0),
    (decode('b000000000000002', 'hex'), 'e0000000-0000-0000-0000-000000000001', 'f0000000-0000-0000-0000-000000000001', NOW(), NOW(), 'Dev Sensor 2', '', false, 'A', false, false, '{}', '{}', decode('0000000000000000', 'hex'), '{}', 0),
    (decode('b000000000000003', 'hex'), 'e0000000-0000-0000-0000-000000000001', 'f0000000-0000-0000-0000-000000000001', NOW(), NOW(), 'Dev Sensor 3', '', false, 'A', false, false, '{}', '{}', decode('0000000000000000', 'hex'), '{}', 0),
    (decode('b000000000000004', 'hex'), 'e0000000-0000-0000-0000-000000000001', 'f0000000-0000-0000-0000-000000000001', NOW(), NOW(), 'Dev Sensor 4', '', false, 'A', false, false, '{}', '{}', decode('0000000000000000', 'hex'), '{}', 0)
ON CONFLICT (dev_eui) DO NOTHING;
