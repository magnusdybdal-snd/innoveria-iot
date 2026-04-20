package service

import "innoveria-iot/erp-service/internal/domain"

// ProductionResourceSvcImpl provides application-level operations
// for production resource read use-cases.
type ProductionResourceSvcImpl struct {
	repo domain.ProductionResourceRepo
}

// ProductionResourceSvc creates a new production resource service.
func ProductionResourceSvc(repo domain.ProductionResourceRepo) *ProductionResourceSvcImpl {
	return &ProductionResourceSvcImpl{
		repo: repo,
	}
}

// GetProductionResource retrieves production resources for one company.
func GetProductionResource(companyID string) ([]domain.ProductionResourceRepo, error) {
	return nil, nil
}
