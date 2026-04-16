package repository

import (
	"context"
	"fmt"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

const (
	syncProductionResourcesQuery = `
	WITH picked AS (
		SELECT company_id, id, number, description, type, received_at
		FROM erp_raw.production_resource
		WHERE sync_status = 'pending'
		  AND (next_retry_at IS NULL OR next_retry_at <= now())
		  AND type = ANY(enum_range(NULL::erp.work_center_type)::text[])
		ORDER BY received_at
		LIMIT $1
		FOR UPDATE SKIP LOCKED
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
		FROM picked
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
		last_synced_at = now()
	WHERE EXISTS (
		SELECT 1 FROM upserted u
		WHERE u.company_id = r.company_id
		  AND u.id = r.id
	)
	`

	syncOrdersQuery = `
	WITH picked AS (
		SELECT company_id, id, order_number, part_id, part_description, planned_start_date,
			planned_finish_date, actual_start_date, actual_finish_date, status, priority, received_at
		FROM erp_raw."order"
		WHERE sync_status = 'pending'
		  AND (next_retry_at IS NULL OR next_retry_at <= now())
		  AND status = ANY(enum_range(NULL::erp.order_status)::text[])
		ORDER BY received_at
		LIMIT $1
		FOR UPDATE SKIP LOCKED
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
		FROM picked
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
		last_synced_at = now()
	WHERE EXISTS (
		SELECT 1 FROM upserted u
		WHERE u.company_id = r.company_id
		  AND u.id = r.id
	)
	`

	syncOrderOperationsQuery = `
	WITH picked AS (
		SELECT company_id, id, production_resource_id, order_id, planned_start_date,
			planned_finish_date, actual_start_date, actual_finish_date, status,
			production_resource_status, received_at
		FROM erp_raw.order_operation r
		WHERE sync_status = 'pending'
		  AND (next_retry_at IS NULL OR next_retry_at <= now())
		  AND status = ANY(enum_range(NULL::erp.operation_status)::text[])
		  AND production_resource_status = ANY(enum_range(NULL::erp.operation_status)::text[])
		  AND EXISTS (
			SELECT 1
			FROM erp."order" o
			WHERE o.company_id = r.company_id
			  AND o.id = r.order_id
		  )
		  AND EXISTS (
			SELECT 1
			FROM erp.production_resource pr
			WHERE pr.company_id = r.company_id
			  AND pr.id = r.production_resource_id
		  )
		ORDER BY received_at
		LIMIT $1
		FOR UPDATE SKIP LOCKED
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
		FROM picked
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
		last_synced_at = now()
	WHERE EXISTS (
		SELECT 1 FROM upserted u
		WHERE u.company_id = r.company_id
		  AND u.id = r.id
	)
	`

	syncOrderReportsQuery = `
	WITH picked AS (
		SELECT company_id, id, order_operation_id, production_resource_id, quantity,
			rest_quantity, type, reporting_timestamp, actual_reported_date, received_at
		FROM erp_raw.order_report r
		WHERE sync_status = 'pending'
		  AND (next_retry_at IS NULL OR next_retry_at <= now())
		  AND type = ANY(enum_range(NULL::erp.order_report_type)::text[])
		  AND EXISTS (
			SELECT 1
			FROM erp.order_operation op
			WHERE op.company_id = r.company_id
			  AND op.id = r.order_operation_id
		  )
		  AND EXISTS (
			SELECT 1
			FROM erp.production_resource pr
			WHERE pr.company_id = r.company_id
			  AND pr.id = r.production_resource_id
		  )
		ORDER BY received_at
		LIMIT $1
		FOR UPDATE SKIP LOCKED
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
		FROM picked
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
		last_synced_at = now()
	WHERE EXISTS (
		SELECT 1 FROM upserted u
		WHERE u.company_id = r.company_id
		  AND u.id = r.id
	)
	`
)

// ReconcileRepoImpl implements reconcile persistence operations.
type ReconcileRepoImpl struct {
	db *dbutil.DB
}

// NewReconcileRepo creates a new reconcile repository.
func NewReconcileRepo(db *dbutil.DB) *ReconcileRepoImpl {
	return &ReconcileRepoImpl{db: db}
}

// SyncProductionResources syncs raw production resources into curated tables.
func (r *ReconcileRepoImpl) SyncProductionResources(ctx context.Context, batchSize int) (int64, error) {
	res, err := r.db.Pool.Exec(ctx, syncProductionResourcesQuery, batchSize)
	if err != nil {
		return 0, fmt.Errorf("sync production resources: %w: %w", domain.ErrDatabase, err)
	}
	return res.RowsAffected(), nil
}

// SyncOrders syncs raw orders into curated tables.
func (r *ReconcileRepoImpl) SyncOrders(ctx context.Context, batchSize int) (int64, error) {
	res, err := r.db.Pool.Exec(ctx, syncOrdersQuery, batchSize)
	if err != nil {
		return 0, fmt.Errorf("sync orders: %w: %w", domain.ErrDatabase, err)
	}
	return res.RowsAffected(), nil
}

// SyncOrderOperations syncs raw order operations into curated tables.
func (r *ReconcileRepoImpl) SyncOrderOperations(ctx context.Context, batchSize int) (int64, error) {
	res, err := r.db.Pool.Exec(ctx, syncOrderOperationsQuery, batchSize)
	if err != nil {
		return 0, fmt.Errorf("sync order operations: %w: %w", domain.ErrDatabase, err)
	}
	return res.RowsAffected(), nil
}

// SyncOrderReports syncs raw order reports into curated tables.
func (r *ReconcileRepoImpl) SyncOrderReports(ctx context.Context, batchSize int) (int64, error) {
	res, err := r.db.Pool.Exec(ctx, syncOrderReportsQuery, batchSize)
	if err != nil {
		return 0, fmt.Errorf("sync order reports: %w: %w", domain.ErrDatabase, err)
	}
	return res.RowsAffected(), nil
}
