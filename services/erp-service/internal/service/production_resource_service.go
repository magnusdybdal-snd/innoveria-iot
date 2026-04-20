package service

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
)

// ProductionResourceSvcImpl provides application-level operations
// for production resource read use-cases.
type ProductionResourceSvcImpl struct {
	repo domain.ProductionResourceRepo
}

// NewProductionResourceSvc creates a new production resource service.
func NewProductionResourceSvc(repo domain.ProductionResourceRepo) *ProductionResourceSvcImpl {
	return &ProductionResourceSvcImpl{
		repo: repo,
	}
}

// GetAll retrieves production resources for one company.
func (s *ProductionResourceSvcImpl) GetAll(ctx context.Context, companyID string) ([]domain.ProductionResource, error) {
	data, err := s.repo.FindAllProductionResource(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
