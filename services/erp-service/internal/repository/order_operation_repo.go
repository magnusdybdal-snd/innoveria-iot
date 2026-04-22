package repository

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

// OrderOperationRepoImpl is the implementation for database operation for order operations.
type OrderOperationRepoImpl struct {
	db *dbutil.DB
}

//nolint:godoclint // SQL constant name is intentionally exported for cross-package reuse.
const (
	FindAllOrderOperationByCompanyIDQuery = `
		SELECT company_id, id, production_resource_id, order_id,
			planned_start_date, planned_finish_date,
			actual_start_date, actual_finish_date,
			status, production_resource_status, received_at
		FROM erp.order_operation
		WHERE company_id = $1
		ORDER BY planned_start_date ASC
	`
	FindAllOrderOperationByOrderIDQuery = `
		SELECT company_id, id, production_resource_id, order_id,
			planned_start_date, planned_finish_date,
			actual_start_date, actual_finish_date,
			status, production_resource_status, received_at
		FROM erp.order_operation
		WHERE order_id = $1 AND company_id = $2
		ORDER BY planned_start_date ASC
	`
	FindOrderOperationByIDQuery = `
		SELECT company_id, id, production_resource_id, order_id,
			planned_start_date, planned_finish_date,
			actual_start_date, actual_finish_date,
			status, production_resource_status, received_at
		FROM erp.order_operation
		WHERE id = $1 AND company_id = $2
	`
)

// NewOrderOperationRepo initializes a new order operation repository.
func NewOrderOperationRepo(db *dbutil.DB) *OrderOperationRepoImpl {
	return &OrderOperationRepoImpl{
		db: db,
	}
}

// FindAllByCompanyID retrieves all order operations belonging to the given company.
// Ordered by planned start date.
func (r *OrderOperationRepoImpl) FindAllByCompanyID(ctx context.Context, companyID string) ([]domain.OrderOperation, error) {
	rows, err := r.db.Pool.Query(ctx, FindAllOrderOperationByCompanyIDQuery, companyID)
	if err != nil {
		return nil, WrapMappedDBError("find all order operations by company id", err)
	}
	defer rows.Close()

	out := make([]domain.OrderOperation, 0)

	for rows.Next() {
		var op domain.OrderOperation
		err := rows.Scan(
			&op.CompanyID,
			&op.ID,
			&op.ProductionResourceID,
			&op.OrderID,
			&op.PlannedStartDate,
			&op.PlannedFinishDate,
			&op.ActualStartDate,
			&op.ActualFinishDate,
			&op.Status,
			&op.ProductionResourceStatus,
			&op.ReceivedAt,
		)
		if err != nil {
			return nil, WrapMappedDBError("scan order operation", err)
		}

		out = append(out, op)
	}

	if err := rows.Err(); err != nil {
		return nil, WrapMappedDBError("iterate order operations", err)
	}

	return out, nil
}

// FindAllByOrderID retrieves all order operations for a specific order.
func (r *OrderOperationRepoImpl) FindAllByOrderID(ctx context.Context, orderID int64, companyID string) ([]domain.OrderOperation, error) {
	rows, err := r.db.Pool.Query(ctx, FindAllOrderOperationByOrderIDQuery, orderID, companyID)
	if err != nil {
		return nil, WrapMappedDBError("find all order operations by order id", err)
	}
	defer rows.Close()

	out := make([]domain.OrderOperation, 0)

	for rows.Next() {
		var op domain.OrderOperation
		err := rows.Scan(
			&op.CompanyID,
			&op.ID,
			&op.ProductionResourceID,
			&op.OrderID,
			&op.PlannedStartDate,
			&op.PlannedFinishDate,
			&op.ActualStartDate,
			&op.ActualFinishDate,
			&op.Status,
			&op.ProductionResourceStatus,
			&op.ReceivedAt,
		)
		if err != nil {
			return nil, WrapMappedDBError("scan order operation", err)
		}

		out = append(out, op)
	}

	if err := rows.Err(); err != nil {
		return nil, WrapMappedDBError("iterate order operations", err)
	}

	return out, nil
}

// FindByID retrieves one order operation based on ID belonging to the given company.
func (r *OrderOperationRepoImpl) FindByID(ctx context.Context, orderOperationID int64, companyID string) (domain.OrderOperation, error) {
	var out domain.OrderOperation
	err := r.db.Pool.QueryRow(ctx, FindOrderOperationByIDQuery, orderOperationID, companyID).Scan(
		&out.CompanyID,
		&out.ID,
		&out.ProductionResourceID,
		&out.OrderID,
		&out.PlannedStartDate,
		&out.PlannedFinishDate,
		&out.ActualStartDate,
		&out.ActualFinishDate,
		&out.Status,
		&out.ProductionResourceStatus,
		&out.ReceivedAt,
	)
	if err != nil {
		return domain.OrderOperation{}, WrapMappedDBError("find order operation by id", err)
	}
	return out, nil
}
