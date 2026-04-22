package repository

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

// OrderReportRepoImpl is the implementation for database operation for order reports.
type OrderReportRepoImpl struct {
	db *dbutil.DB
}

//nolint:godoclint // SQL constant name is intentionally exported for cross-package reuse.
const (
	FindAllOrderReportByCompanyIDQuery = `
		SELECT company_id, id, order_operation_id, production_resource_id,
			quantity, rest_quantity, type, reporting_timestamp,
			actual_reported_date, received_at
		FROM erp.order_report
		WHERE company_id = $1
		ORDER BY reporting_timestamp ASC
	`
	FindAllOrderReportByOrderOperationIDQuery = `
		SELECT company_id, id, order_operation_id, production_resource_id,
			quantity, rest_quantity, type, reporting_timestamp,
			actual_reported_date, received_at
		FROM erp.order_report
		WHERE order_operation_id = $1 AND company_id = $2
		ORDER BY reporting_timestamp ASC
	`
	FindOrderReportByIDQuery = `
		SELECT company_id, id, order_operation_id, production_resource_id,
			quantity, rest_quantity, type, reporting_timestamp,
			actual_reported_date, received_at
		FROM erp.order_report
		WHERE id = $1 AND company_id = $2
	`
)

// NewOrderReportRepo initializes a new order report repository.
func NewOrderReportRepo(db *dbutil.DB) *OrderReportRepoImpl {
	return &OrderReportRepoImpl{
		db: db,
	}
}

// FindAllByCompanyID retrieves all order reports belonging to the given company.
func (r *OrderReportRepoImpl) FindAllByCompanyID(ctx context.Context, companyID string) ([]domain.OrderReport, error) {
	rows, err := r.db.Pool.Query(ctx, FindAllOrderReportByCompanyIDQuery, companyID)
	if err != nil {
		return nil, WrapMappedDBError("find all order reports by company id", err)
	}
	defer rows.Close()

	out := make([]domain.OrderReport, 0)

	for rows.Next() {
		var report domain.OrderReport
		err := rows.Scan(
			&report.CompanyID,
			&report.ID,
			&report.OrderOperationID,
			&report.ProductionResourceID,
			&report.Quantity,
			&report.RestQuantity,
			&report.Type,
			&report.ReportingTimestamp,
			&report.ActualReportedDate,
			&report.ReceivedAt,
		)
		if err != nil {
			return nil, WrapMappedDBError("scan order report", err)
		}

		out = append(out, report)
	}

	if err := rows.Err(); err != nil {
		return nil, WrapMappedDBError("iterate order reports", err)
	}

	return out, nil
}

// FindAllByOrderOperationID retrieves all order reports for a specific order operation.
func (r *OrderReportRepoImpl) FindAllByOrderOperationID(ctx context.Context, orderOperationID int64, companyID string) ([]domain.OrderReport, error) {
	rows, err := r.db.Pool.Query(ctx, FindAllOrderReportByOrderOperationIDQuery, orderOperationID, companyID)
	if err != nil {
		return nil, WrapMappedDBError("find all order reports by order operation id", err)
	}
	defer rows.Close()

	out := make([]domain.OrderReport, 0)

	for rows.Next() {
		var report domain.OrderReport
		err := rows.Scan(
			&report.CompanyID,
			&report.ID,
			&report.OrderOperationID,
			&report.ProductionResourceID,
			&report.Quantity,
			&report.RestQuantity,
			&report.Type,
			&report.ReportingTimestamp,
			&report.ActualReportedDate,
			&report.ReceivedAt,
		)
		if err != nil {
			return nil, WrapMappedDBError("scan order report", err)
		}

		out = append(out, report)
	}

	if err := rows.Err(); err != nil {
		return nil, WrapMappedDBError("iterate order reports", err)
	}

	return out, nil
}

// FindByID retrieves one order report based on ID belonging to the given company.
func (r *OrderReportRepoImpl) FindByID(ctx context.Context, orderReportID int64, companyID string) (domain.OrderReport, error) {
	var out domain.OrderReport
	err := r.db.Pool.QueryRow(ctx, FindOrderReportByIDQuery, orderReportID, companyID).Scan(
		&out.CompanyID,
		&out.ID,
		&out.OrderOperationID,
		&out.ProductionResourceID,
		&out.Quantity,
		&out.RestQuantity,
		&out.Type,
		&out.ReportingTimestamp,
		&out.ActualReportedDate,
		&out.ReceivedAt,
	)
	if err != nil {
		return domain.OrderReport{}, WrapMappedDBError("find order report by id", err)
	}
	return out, nil
}
