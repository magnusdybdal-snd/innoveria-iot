package service

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
)

// OrderServiceImpl coordinates order-related use cases across repositories.
type OrderServiceImpl struct {
	orderRepo              domain.OrderRepo
	orderOperationRepo     domain.OrderOperationRepo
	orderReportRepo        domain.OrderOperationRepo
	productionResourceRepo domain.ProductionResourceRepo
}

// NewOrderService wires repository dependencies for OrderServiceImpl.
func NewOrderService(
	orderRepo domain.OrderRepo,
	orderOperationRepo domain.OrderOperationRepo,
	orderReportRepo domain.OrderOperationRepo,
	productionResourceRepo domain.ProductionResourceRepo,
) *OrderServiceImpl {
	return &OrderServiceImpl{
		orderRepo:              orderRepo,
		orderOperationRepo:     orderOperationRepo,
		orderReportRepo:        orderReportRepo,
		productionResourceRepo: productionResourceRepo,
	}
}

// GetOrderSummary returns lightweight order rows for a single company.
func (s *OrderServiceImpl) GetOrderSummary(ctx context.Context, companyID string) ([]domain.OrderSummary, error) {
	data, err := s.orderRepo.FindAllByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
