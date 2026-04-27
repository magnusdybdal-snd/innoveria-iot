-- Production resources (must be inserted before orders and operations)
INSERT INTO "erp"."production_resource" ("company_id", "id", "number", "description", "type", "received_at")
VALUES
('a0000000-0000-0000-0000-000000000001', 1, 'WC-101', 'Welding Station 1',    'machine', '2026-04-14 05:00:00+00'),
('a0000000-0000-0000-0000-000000000001', 2, 'WC-102', 'Press Station 2',      'machine', '2026-04-14 13:00:00+00'),
('a0000000-0000-0000-0000-000000000001', 4, 'WC-104', 'Electricity Monitor 1','machine', '2026-04-14 05:00:00+00'),
('a0000000-0000-0000-0000-000000000001', 6, 'WC-106', 'Cutting Station 3',    'machine', '2026-04-14 05:00:00+00')
ON CONFLICT DO NOTHING;

-- Orders
INSERT INTO "erp"."order" (
    "company_id", "id", "order_number", "part_id", "part_description",
    "planned_start_date", "planned_finish_date",
    "actual_start_date", "actual_finish_date",
    "status", "priority", "received_at"
)
VALUES
(
    'a0000000-0000-0000-0000-000000000001', 1, 'MO-2026-001', 'PART-001', 'Steel Frame A',
    '2026-04-14 06:00:00+00', '2026-04-14 14:00:00+00',
    '2026-04-14 06:12:00+00', '2026-04-14 13:55:00+00',
    'finished', 1, '2026-04-14 05:00:00+00'
),
(
    'a0000000-0000-0000-0000-000000000001', 2, 'MO-2026-002', 'PART-002', 'Aluminium Panel B',
    '2026-04-14 14:00:00+00', '2026-04-14 22:00:00+00',
    '2026-04-14 14:05:00+00', '2026-04-14 21:50:00+00',
    'finished', 2, '2026-04-14 13:00:00+00'
),
-- Order 3: live sensor demo — window covers 30 days before service startup
(
    'a0000000-0000-0000-0000-000000000001', 3, 'MO-2026-003', 'PART-003', 'Live Sensor Demo',
    NOW() - INTERVAL '30 days', NOW(),
    NOW() - INTERVAL '30 days', NOW(),
    'finished', 1, NOW() - INTERVAL '30 days'
)
ON CONFLICT DO NOTHING;

-- Order operations
INSERT INTO "erp"."order_operation" (
    "company_id", "id", "production_resource_id", "order_id",
    "planned_start_date", "planned_finish_date",
    "actual_start_date", "actual_finish_date",
    "status", "production_resource_status", "received_at"
)
VALUES
(
    'a0000000-0000-0000-0000-000000000001', 10, 1, 1,
    '2026-04-14 06:00:00+00', '2026-04-14 14:00:00+00',
    '2026-04-14 06:12:00+00', '2026-04-14 13:55:00+00',
    'finished', 'finished', '2026-04-14 05:00:00+00'
),
(
    'a0000000-0000-0000-0000-000000000001', 20, 2, 2,
    '2026-04-14 14:00:00+00', '2026-04-14 22:00:00+00',
    '2026-04-14 14:05:00+00', '2026-04-14 21:50:00+00',
    'finished', 'finished', '2026-04-14 13:00:00+00'
),
-- Op 30: WC-101 on order 3 — sensors 1 & 2 are mapped to resource ID 1
(
    'a0000000-0000-0000-0000-000000000001', 30, 1, 3,
    NOW() - INTERVAL '30 days', NOW(),
    NOW() - INTERVAL '30 days', NOW(),
    'finished', 'finished', NOW() - INTERVAL '30 days'
),
-- Op 31: WC-106 on order 3 — no sensors mapped (demos degraded state)
(
    'a0000000-0000-0000-0000-000000000001', 31, 6, 3,
    NOW() - INTERVAL '30 days', NOW(),
    NOW() - INTERVAL '30 days', NOW(),
    'finished', 'finished', NOW() - INTERVAL '30 days'
),
-- Op 32: WC-104 on order 3 — electricity sensor (Dev Sensor 4, EUI b000000000000004)
(
    'a0000000-0000-0000-0000-000000000001', 32, 4, 3,
    NOW() - INTERVAL '30 days', NOW(),
    NOW() - INTERVAL '30 days', NOW(),
    'finished', 'finished', NOW() - INTERVAL '30 days'
)
ON CONFLICT DO NOTHING;

-- Order reports
INSERT INTO "erp"."order_report" (
    "company_id", "id", "order_operation_id", "production_resource_id",
    "quantity", "rest_quantity", "type", "reporting_timestamp", "actual_reported_date", "received_at"
)
VALUES
(
    'a0000000-0000-0000-0000-000000000001', 100, 10, 1,
    8.0, 0.0, 'regular',
    '2026-04-14 13:55:00+00', '2026-04-14 14:00:00+00', '2026-04-14 14:00:00+00'
),
(
    'a0000000-0000-0000-0000-000000000001', 200, 20, 2,
    12.5, 0.0, 'regular',
    '2026-04-14 21:50:00+00', '2026-04-14 22:00:00+00', '2026-04-14 22:00:00+00'
),
(
    'a0000000-0000-0000-0000-000000000001', 300, 30, 1,
    5.0, 10.0, 'regular',
    NOW() - INTERVAL '22 days', NOW() - INTERVAL '22 days' + INTERVAL '5 minutes',
    NOW() - INTERVAL '22 days'
),
(
    'a0000000-0000-0000-0000-000000000001', 301, 30, 1,
    7.0, 3.0, 'regular',
    NOW() - INTERVAL '10 days', NOW() - INTERVAL '10 days' + INTERVAL '3 minutes',
    NOW() - INTERVAL '10 days'
),
(
    'a0000000-0000-0000-0000-000000000001', 310, 31, 6,
    4.0, 3.0, 'regular',
    NOW() - INTERVAL '8 days', NOW() - INTERVAL '8 days' + INTERVAL '2 minutes',
    NOW() - INTERVAL '8 days'
),
-- Report 320: electricity monitor operation
(
    'a0000000-0000-0000-0000-000000000001', 320, 32, 4,
    1.0, 0.0, 'regular',
    NOW() - INTERVAL '15 days', NOW() - INTERVAL '15 days' + INTERVAL '2 minutes',
    NOW() - INTERVAL '15 days'
)
ON CONFLICT DO NOTHING;
