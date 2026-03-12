-- Company mapping shared across services (matches device + collection seeds).
INSERT INTO "auth"."company" ("company_id", "name", "address")
VALUES ('a0000000-0000-0000-0000-000000000001', 'Innoveria Dev', 'Gjøvik')
ON CONFLICT DO NOTHING;

-- Factory data for the demo company.
INSERT INTO "auth"."factory" ("factory_id", "company_id", "name", "address")
VALUES
('f1000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001', 'NTNU Gjøvik', 'Gjøvik')
ON CONFLICT DO NOTHING;
