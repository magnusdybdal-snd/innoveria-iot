INSERT INTO "context"."aggregation_rule"
   (company_id, name, context_type, measurement_type, aggregation_method, time_bucket_minutes)
   VALUES
('a0000000-0000-0000-0000-000000000001', 'Hourly avg temperature', 'temperature_avg_hourly', 'temperature',
  'AVG', 60),
    ('a0000000-0000-0000-0000-000000000001', 'Daily total nitrogen',   'nitrogen_total_daily',   'nitrogen',
  'SUM', 1440)
  ON CONFLICT (company_id, context_type) DO NOTHING;