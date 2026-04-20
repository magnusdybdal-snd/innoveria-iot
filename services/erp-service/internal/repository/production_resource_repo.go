package repository

import (
	"context"
	"fmt"

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
)

// NewProductionResourceRepo initializes a new production repository.
func NewProductionResourceRepo(db *dbutil.DB) *ProductionResourceRepoImpl {
	return &ProductionResourceRepoImpl{
		db: db,
	}
}

// FindAllProductionResource retreives all production resources belonging to the given company
// Ordered by when it was recevied by the agent service
func (r *ProductionResourceRepoImpl) FindAllProductionResource(ctx context.Context, companyID string) ([]domain.ProductionResource, error) {
	rows, err := r.db.Pool.Query(ctx, FindAllProductionResourceByCompanyIDQuery, companyID)
	if err != nil {
		return nil, fmt.Errorf("find all production resource by company id: %w", err)
	}
	defer rows.Close()

	var out []domain.ProductionResource
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
			return nil, fmt.Errorf("scan production resource: %w", err)
		}

		out = append(out, prodRes)
	}

	if err != nil {
		return nil, fmt.Errorf("find all production resource by id: %w", err)
	}

	return out, nil
}
