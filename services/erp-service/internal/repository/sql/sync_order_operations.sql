WITH picked AS (
	SELECT company_id, id, production_resource_id, order_id, planned_start_date,
		planned_finish_date, actual_start_date, actual_finish_date, status,
		production_resource_status, received_at
	FROM erp_raw.order_operation r
	WHERE sync_status IN ('pending', 'deferred')
	  AND (next_retry_at IS NULL OR next_retry_at <= now())
	ORDER BY received_at
	LIMIT $1
	FOR UPDATE SKIP LOCKED
),
failed_enum AS (
	UPDATE erp_raw.order_operation r
	SET sync_status = 'failed',
		sync_error = 'invalid operation_status or production_resource_status',
		retry_count = r.retry_count + 1,
		next_retry_at = NULL
	FROM picked p
	WHERE r.company_id = p.company_id
	  AND r.id = p.id
	  AND NOT (
		p.status = ANY(enum_range(NULL::erp.operation_status)::text[])
		AND p.production_resource_status = ANY(enum_range(NULL::erp.operation_status)::text[])
	  )
	RETURNING r.company_id, r.id
),
deferred_dependency AS (
	UPDATE erp_raw.order_operation r
	SET sync_status = CASE
			WHEN r.retry_count + 1 >= 24 THEN 'failed'
			ELSE 'deferred'
		END,
		sync_error = CASE
			WHEN r.retry_count + 1 >= 24 THEN 'failed: missing order or production_resource dependency after max retries'
			ELSE 'deferred: missing order or production_resource dependency'
		END,
		retry_count = r.retry_count + 1,
		next_retry_at = CASE
			WHEN r.retry_count + 1 >= 24 THEN NULL
			ELSE now() + make_interval(secs => LEAST(300, 5 * (2 ^ LEAST(r.retry_count, 6))))
		END
	FROM picked p
	WHERE r.company_id = p.company_id
	  AND r.id = p.id
      AND NOT EXISTS (
        SELECT 1
        FROM failed_enum fe
        WHERE fe.company_id = p.company_id
            AND fe.id = p.id
      )
	  AND p.status = ANY(enum_range(NULL::erp.operation_status)::text[])
	  AND p.production_resource_status = ANY(enum_range(NULL::erp.operation_status)::text[])
	  AND (
		NOT EXISTS (
			SELECT 1
			FROM erp."order" o
			WHERE o.company_id = p.company_id
			  AND o.id = p.order_id
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
	WHERE p.status = ANY(enum_range(NULL::erp.operation_status)::text[])
	  AND p.production_resource_status = ANY(enum_range(NULL::erp.operation_status)::text[])
	  AND EXISTS (
		SELECT 1
		FROM erp."order" o
		WHERE o.company_id = p.company_id
		  AND o.id = p.order_id
	  )
	  AND EXISTS (
		SELECT 1
		FROM erp.production_resource pr
		WHERE pr.company_id = p.company_id
		  AND pr.id = p.production_resource_id
	  )
),
upserted AS (
	INSERT INTO erp.order_operation (
		company_id,
		id,
		production_resource_id,
		order_id,
		planned_start_date,
		planned_finish_date,
		actual_start_date,
		actual_finish_date,
		status,
		production_resource_status,
		received_at
	)
	SELECT
		company_id,
		id,
		production_resource_id,
		order_id,
		planned_start_date,
		planned_finish_date,
		actual_start_date,
		actual_finish_date,
		status::erp.operation_status,
		production_resource_status::erp.operation_status,
		received_at
	FROM ready
	ON CONFLICT (company_id, id) DO UPDATE SET
		production_resource_id = EXCLUDED.production_resource_id,
		order_id = EXCLUDED.order_id,
		planned_start_date = EXCLUDED.planned_start_date,
		planned_finish_date = EXCLUDED.planned_finish_date,
		actual_start_date = EXCLUDED.actual_start_date,
		actual_finish_date = EXCLUDED.actual_finish_date,
		status = EXCLUDED.status,
		production_resource_status = EXCLUDED.production_resource_status,
		received_at = EXCLUDED.received_at
	RETURNING company_id, id
)
UPDATE erp_raw.order_operation r
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
