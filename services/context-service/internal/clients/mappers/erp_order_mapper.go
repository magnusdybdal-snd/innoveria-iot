package mappers

import (
	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/domain"
)

// ToERPProductionResource converts a production resource DTO to its domain representation.
func ToERPProductionResource(r dto.ERPProductionResourceResponse) domain.ERPProductionResource {
	return domain.ERPProductionResource{
		ID:          r.ID,
		Number:      r.Number,
		Description: r.Description,
		Type:        r.Type,
	}
}

// ToERPOrderOperation converts an order operation DTO to its domain representation.
func ToERPOrderOperation(r dto.ERPOrderOperationResponse) domain.ERPOrderOperation {
	return domain.ERPOrderOperation{
		ID:                       r.ID,
		ProductionResource:       ToERPProductionResource(r.ProductionResource),
		PlannedStartDate:         r.PlannedStartDate,
		PlannedFinishDate:        r.PlannedFinishDate,
		ActualStartDate:          r.ActualStartDate,
		ActualFinishDate:         r.ActualFinishDate,
		Status:                   r.Status,
		ProductionResourceStatus: r.ProductionResourceStatus,
	}
}

// ToERPOrder converts an ERP order DTO (with nested operations) to its domain representation.
func ToERPOrder(r dto.ERPOrderResponse) domain.ERPOrder {
	ops := make([]domain.ERPOrderOperation, len(r.Operations))
	for i, op := range r.Operations {
		ops[i] = ToERPOrderOperation(op)
	}
	return domain.ERPOrder{
		ID:                r.ID,
		OrderNumber:       r.OrderNumber,
		PartDescription:   r.PartDescription,
		PlannedStartDate:  r.PlannedStartDate,
		PlannedFinishDate: r.PlannedFinishDate,
		ActualStartDate:   r.ActualStartDate,
		ActualFinishDate:  r.ActualFinishDate,
		Status:            r.Status,
		Priority:          r.Priority,
		Operations:        ops,
	}
}
