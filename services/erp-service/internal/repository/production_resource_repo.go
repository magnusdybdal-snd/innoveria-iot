package repository

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

// ProductionResourceRepoImpl is the implementation for database operation for production resources.
type ProductionResourceRepoImpl struct {
	db *dbutil.DB
}

//nolint:godoclint // SQL constant name is intentionally exported for cross-package reuse.
const (
	FindAllProductionResourceByCompanyIDQuery = `
		SELECT company_id, id, number, description, type, received_at
		FROM erp.production_resource
		WHERE company_id = $1
		ORDER BY received_at ASC
	`
	FindProductionResourceByIDQuery = `
		SELECT company_id, id, number, description, type, received_at
		FROM erp.production_resource
		WHERE id = $1 AND company_id = $2
	`
)

// NewProductionResourceRepo initializes a new production repository.
func NewProductionResourceRepo(db *dbutil.DB) *ProductionResourceRepoImpl {
	return &ProductionResourceRepoImpl{
		db: db,
	}
}

// FindAllByCompanyID retrieves all production resources belonging to the given company.
// Ordered by when it was received by the agent service.
func (r *ProductionResourceRepoImpl) FindAllByCompanyID(ctx context.Context, companyID string) ([]domain.ProductionResource, error) {
	rows, err := r.db.Pool.Query(ctx, FindAllProductionResourceByCompanyIDQuery, companyID)
	if err != nil {
		return nil, WrapMappedDBError("find all production resource by company id", err)
	}
	defer rows.Close()

	out := make([]domain.ProductionResource, 0) // return a empty slice and not nil

	for rows.Next() {
		var prodRes domain.ProductionResource
		err := rows.Scan(
			&prodRes.CompanyID,
			&prodRes.ID,
			&prodRes.Number,
			&prodRes.Description,
			&prodRes.Type,
			&prodRes.ReceivedAt,
		)
		if err != nil {
			return nil, WrapMappedDBError("scan production resource", err)
		}

		out = append(out, prodRes)
	}

	if err := rows.Err(); err != nil {
		return nil, WrapMappedDBError("iterate production resources", err)
	}

	return out, nil
}

// FindByID retrieves one production resource based on ID belonging to the given company.
func (r *ProductionResourceRepoImpl) FindByID(ctx context.Context, productionResourceID int64, companyID string) (domain.ProductionResource, error) {
	var out domain.ProductionResource
	err := r.db.Pool.QueryRow(ctx, FindProductionResourceByIDQuery, productionResourceID, companyID).Scan(
		&out.CompanyID,
		&out.ID,
		&out.Number,
		&out.Description,
		&out.Type,
		&out.ReceivedAt,
	)
	if err != nil {
		return domain.ProductionResource{}, WrapMappedDBError("find production resource by id", err)
	}
	return out, nil
}
