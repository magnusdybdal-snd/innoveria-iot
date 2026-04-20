WITH picked AS (
	SELECT company_id, id, order_number, part_id, part_description, planned_start_date,
		planned_finish_date, actual_start_date, actual_finish_date, status, priority, received_at
	FROM erp_raw."order" r
	WHERE sync_status IN ('pending', 'deferred')
	  AND (next_retry_at IS NULL OR next_retry_at <= now())
	ORDER BY received_at
	LIMIT $1
	FOR UPDATE SKIP LOCKED
),
failed_enum AS (
	UPDATE erp_raw."order" r
	SET sync_status = 'failed',
		sync_error = 'invalid order_status',
		retry_count = r.retry_count + 1,
		next_retry_at = NULL
	FROM picked p
	WHERE r.company_id = p.company_id
	  AND r.id = p.id
	  AND NOT (p.status = ANY(enum_range(NULL::erp.order_status)::text[]))
	RETURNING r.company_id, r.id
),
ready AS (
	SELECT p.*
	FROM picked p
	WHERE p.status = ANY(enum_range(NULL::erp.order_status)::text[])
),
upserted AS (
	INSERT INTO erp."order" (
		company_id,
		id,
		order_number,
		part_id,
		part_description,
		planned_start_date,
		planned_finish_date,
		actual_start_date,
		actual_finish_date,
		status,
		priority,
		received_at
	)
	SELECT
		company_id,
		id,
		order_number,
		part_id,
		part_description,
		planned_start_date,
		planned_finish_date,
		actual_start_date,
		actual_finish_date,
		status::erp.order_status,
		priority,
		received_at
	FROM ready
	ON CONFLICT (company_id, id) DO UPDATE SET
		order_number = EXCLUDED.order_number,
		part_id = EXCLUDED.part_id,
		part_description = EXCLUDED.part_description,
		planned_start_date = EXCLUDED.planned_start_date,
		planned_finish_date = EXCLUDED.planned_finish_date,
		actual_start_date = EXCLUDED.actual_start_date,
		actual_finish_date = EXCLUDED.actual_finish_date,
		status = EXCLUDED.status,
		priority = EXCLUDED.priority,
		received_at = EXCLUDED.received_at
	RETURNING company_id, id
)
UPDATE erp_raw."order" r
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
