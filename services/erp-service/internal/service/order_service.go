package service

import (
	"context"

	"innoveria-iot/erp-service/internal/domain"
)

// OrderServiceImpl coordinates order-related use cases across repositories.
type OrderServiceImpl struct {
	orderRepo              domain.OrderRepo
	orderOperationRepo     domain.OrderOperationRepo
	orderReportRepo        domain.OrderReportRepo
	productionResourceRepo domain.ProductionResourceRepo
}

// NewOrderService wires repository dependencies for OrderServiceImpl.
func NewOrderService(
	orderRepo domain.OrderRepo,
	orderOperationRepo domain.OrderOperationRepo,
	orderReportRepo domain.OrderReportRepo,
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

// GetOne returns an aggregated view for one order.
func (s *OrderServiceImpl) GetOne(ctx context.Context, orderID int64, companyID string) (domain.OrderAggregate, error) {
	orderData, err := s.orderRepo.FindByID(ctx, orderID, companyID)
	if err != nil {
		return domain.OrderAggregate{}, err
	}

	orderOperations, err := s.orderOperationRepo.FindAllByOrderID(ctx, orderID, companyID)
	if err != nil {
		return domain.OrderAggregate{}, err
	}

	operationsWithReports := make([]domain.OrderOperationWithReports, 0, len(orderOperations))
	resources := make([]domain.ProductionResource, 0, len(orderOperations))
	resourceSeen := make(map[int64]struct{}, len(orderOperations)) // struct carries no data. So only 8bytes from the int (+ map overhead), better than bool check
	for _, operation := range orderOperations {
		reports, err := s.orderReportRepo.FindAllByOrderOperationID(ctx, operation.ID, companyID)
		if err != nil {
			return domain.OrderAggregate{}, err
		}

		resource, err := s.productionResourceRepo.FindByID(ctx, operation.ProductionResourceID, companyID)
		if err != nil {
			return domain.OrderAggregate{}, err
		}

		operationsWithReports = append(operationsWithReports, domain.OrderOperationWithReports{
			Operation: operation,
			Reports:   reports,
		})

		if _, exists := resourceSeen[resource.ID]; !exists {
			resources = append(resources, resource)
			resourceSeen[resource.ID] = struct{}{} // apply empty body
		}
	}

	return domain.OrderAggregate{
		Order:               orderData,
		Operations:          operationsWithReports,
		ProductionResources: resources,
	}, nil
}
