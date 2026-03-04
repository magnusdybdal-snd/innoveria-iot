-- Insert mock tenant mapping that reflects the tenant ID from simulator
INSERT INTO "collection"."tenant_mapping" ("chirpstack_tenant_id", "company_id")
VALUES ('d0000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000001')
ON CONFLICT DO NOTHING;