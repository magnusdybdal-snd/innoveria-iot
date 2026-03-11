-- Company mapping shared across services (matches device + collection seeds).
INSERT INTO "auth"."company" ("company_id", "name", "address")
VALUES ('a0000000-0000-0000-0000-000000000001', 'Innoveria Dev', 'Gjøvik')
ON CONFLICT DO NOTHING;

-- Factory data for the demo company.
INSERT INTO "auth"."factory" ("factory_id", "company_id", "name", "address")
VALUES
('f1000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'NTNU Gjøvik', 'Gjøvik')
ON CONFLICT DO NOTHING;

-- Factory areas for assigning gateways/sensors in downstream services.
INSERT INTO "auth"."factory_area" ("area_id", "factory_id", "name", "description")
VALUES
('fa000000-0000-0000-0000-000000000001', 'f1000000-0000-0000-0000-000000000001', 'Labben', 'Gjøvik'),
ON CONFLICT DO NOTHING;
