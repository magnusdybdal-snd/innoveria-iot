package repository

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

// OrderRepoImpl is the implementation for database operation for orders.
type OrderRepoImpl struct {
	db *dbutil.DB
}

//nolint:godoclint // SQL constant name is intentionally exported for cross-package reuse.
const (
	FindAllOrderByCompanyIDQuery = `
		SELECT id, order_number
		FROM erp."order"
		WHERE company_id = $1
		ORDER BY received_at ASC
	`
	FindOrderByIDQuery = `
		SELECT id, company_id, order_number, part_id, part_description,
			planned_start_date, planned_finish_date,
			actual_start_date, actual_finish_date,
			status, priority, received_at
		FROM erp."order"
		WHERE id = $1 AND company_id = $2
	`
)

// NewOrderRepo initializes a new order repository.
func NewOrderRepo(db *dbutil.DB) *OrderRepoImpl {
	return &OrderRepoImpl{
		db: db,
	}
}

// FindAllByCompanyID retrieves all orders belonging to the given company.
// Returns only id and order number for summary/list views.
func (r *OrderRepoImpl) FindAllByCompanyID(ctx context.Context, companyID string) ([]domain.OrderSummary, error) {
	rows, err := r.db.Pool.Query(ctx, FindAllOrderByCompanyIDQuery, companyID)
	if err != nil {
		return nil, WrapMappedDBError("find all orders by company id", err)
	}
	defer rows.Close()

	out := make([]domain.OrderSummary, 0)

	for rows.Next() {
		var orderSummary domain.OrderSummary
		err := rows.Scan(
			&orderSummary.ID,
			&orderSummary.OrderNumber,
		)
		if err != nil {
			return nil, WrapMappedDBError("scan order summary", err)
		}

		out = append(out, orderSummary)
	}

	if err := rows.Err(); err != nil {
		return nil, WrapMappedDBError("iterate orders", err)
	}

	return out, nil
}

// FindByID retrieves one order based on ID belonging to the given company.
func (r *OrderRepoImpl) FindByID(ctx context.Context, orderID int64, companyID string) (domain.Order, error) {
	var out domain.Order
	err := r.db.Pool.QueryRow(ctx, FindOrderByIDQuery, orderID, companyID).Scan(
		&out.ID,
		&out.CompanyID,
		&out.OrderNumber,
		&out.PartID,
		&out.PartDescription,
		&out.PlannedStartDate,
		&out.PlannedFinishDate,
		&out.ActualStartDate,
		&out.ActualFinishDate,
		&out.Status,
		&out.Priority,
		&out.ReceivedAt,
	)
	if err != nil {
		return domain.Order{}, WrapMappedDBError("find order by id", err)
	}
	return out, nil
}
