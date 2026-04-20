WITH picked AS (
	SELECT company_id, id, number, description, type, received_at
	FROM erp_raw.production_resource r
	WHERE sync_status IN ('pending', 'deferred')
	  AND (next_retry_at IS NULL OR next_retry_at <= now())
	ORDER BY received_at
	LIMIT $1
	FOR UPDATE SKIP LOCKED
),
failed_enum AS (
	UPDATE erp_raw.production_resource r
	SET sync_status = 'failed',
		sync_error = 'invalid work_center_type',
		retry_count = r.retry_count + 1,
		next_retry_at = NULL
	FROM picked p
	WHERE r.company_id = p.company_id
	  AND r.id = p.id
	  AND NOT (p.type = ANY(enum_range(NULL::erp.work_center_type)::text[]))
	RETURNING r.company_id, r.id
),
ready AS (
	SELECT p.*
	FROM picked p
	WHERE p.type = ANY(enum_range(NULL::erp.work_center_type)::text[])
),
upserted AS (
	INSERT INTO erp.production_resource (
		company_id,
		id,
		number,
		description,
		type,
		received_at
	)
	SELECT
		company_id,
		id,
		number,
		description,
		type::erp.work_center_type,
		received_at
	FROM ready
	ON CONFLICT (company_id, id) DO UPDATE SET
		number = EXCLUDED.number,
		description = EXCLUDED.description,
		type = EXCLUDED.type,
		received_at = EXCLUDED.received_at
	RETURNING company_id, id
)
UPDATE erp_raw.production_resource r
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
