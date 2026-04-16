WITH picked AS (
	SELECT company_id, id, order_operation_id, production_resource_id, quantity,
		rest_quantity, type, reporting_timestamp, actual_reported_date, received_at
	FROM erp_raw.order_report r
	WHERE sync_status IN ('pending', 'deferred')
	  AND (next_retry_at IS NULL OR next_retry_at <= now())
	ORDER BY received_at
	LIMIT $1
	FOR UPDATE SKIP LOCKED
),
failed_enum AS (
	UPDATE erp_raw.order_report r
	SET sync_status = 'failed',
		sync_error = 'invalid order_report_type',
		retry_count = r.retry_count + 1,
		next_retry_at = NULL
	FROM picked p
	WHERE r.company_id = p.company_id
	  AND r.id = p.id
	  AND NOT (p.type = ANY(enum_range(NULL::erp.order_report_type)::text[]))
	RETURNING r.company_id, r.id
),
deferred_dependency AS (
	UPDATE erp_raw.order_report r
	SET sync_status = CASE
			WHEN r.retry_count + 1 >= 24 THEN 'failed'
			ELSE 'deferred'
		END,
		sync_error = CASE
			WHEN r.retry_count + 1 >= 24 THEN 'failed: missing order_operation or production_resource dependency after max retries'
			ELSE 'deferred: missing order_operation or production_resource dependency'
		END,
		retry_count = r.retry_count + 1,
		next_retry_at = CASE
			WHEN r.retry_count + 1 >= 24 THEN NULL
			ELSE now() + make_interval(secs => LEAST(300, 5 * (2 ^ LEAST(r.retry_count, 6))))
		END
	FROM picked p
	WHERE r.company_id = p.company_id
	  AND r.id = p.id
	  AND p.type = ANY(enum_range(NULL::erp.order_report_type)::text[])
	  AND (
		NOT EXISTS (
			SELECT 1
			FROM erp.order_operation op
			WHERE op.company_id = p.company_id
			  AND op.id = p.order_operation_id
		)
		OR NOT EXISTS (
			SELECT 1
			FROM erp.production_resource pr
			WHERE pr.company_id = p.company_id
			  AND pr.id = p.production_resource_id
		)
	  )
	RETURNING r.company_id, r.id
),
ready AS (
	SELECT p.*
	FROM picked p
	WHERE p.type = ANY(enum_range(NULL::erp.order_report_type)::text[])
	  AND EXISTS (
		SELECT 1
		FROM erp.order_operation op
		WHERE op.company_id = p.company_id
		  AND op.id = p.order_operation_id
	  )
	  AND EXISTS (
		SELECT 1
		FROM erp.production_resource pr
		WHERE pr.company_id = p.company_id
		  AND pr.id = p.production_resource_id
	  )
),
upserted AS (
	INSERT INTO erp.order_report (
		company_id,
		id,
		order_operation_id,
		production_resource_id,
		quantity,
		rest_quantity,
		type,
		reporting_timestamp,
		actual_reported_date,
		received_at
	)
	SELECT
		company_id,
		id,
		order_operation_id,
		production_resource_id,
		quantity,
		rest_quantity,
		type::erp.order_report_type,
		reporting_timestamp,
		actual_reported_date,
		received_at
	FROM ready
	ON CONFLICT (company_id, id) DO UPDATE SET
		order_operation_id = EXCLUDED.order_operation_id,
		production_resource_id = EXCLUDED.production_resource_id,
		quantity = EXCLUDED.quantity,
		rest_quantity = EXCLUDED.rest_quantity,
		type = EXCLUDED.type,
		reporting_timestamp = EXCLUDED.reporting_timestamp,
		actual_reported_date = EXCLUDED.actual_reported_date,
		received_at = EXCLUDED.received_at
	RETURNING company_id, id
)
UPDATE erp_raw.order_report r
SET sync_status = 'synced',
	sync_error = NULL,
	retry_count = 0,
	next_retry_at = NULL,
	last_synced_at = now()
WHERE EXISTS (
	SELECT 1 FROM upserted u
	WHERE u.company_id = r.company_id
	  AND u.id = r.id
)
