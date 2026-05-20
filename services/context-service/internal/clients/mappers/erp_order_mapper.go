package mappers

import (
	"innoveria-iot/context-service/internal/domain"
	erpdto "innoveria-iot/pkg/erp/dto"
)

// ToERPProductionResource converts a shared ProductionResource DTO to its domain representation.
func ToERPProductionResource(r erpdto.ProductionResource) domain.ERPProductionResource {
	return domain.ERPProductionResource{
		ID:          r.ID,
		Number:      r.Number,
		Description: r.Description,
		Type:        string(r.Type),
	}
}

// ToERPOrderSummary converts a shared OrderSummary DTO to its domain representation.
// OrderNumber is used as the human-readable Name shown in order pickers.
func ToERPOrderSummary(r erpdto.OrderSummary) domain.ERPOrderSummary {
	return domain.ERPOrderSummary{
		ID:   r.ID,
		Name: r.OrderNumber,
	}
}

// ToERPOrder converts an OrderAggregate from the erp-service into the domain ERPOrder.
// Production resources arrive as a flat list and are joined into each operation by ID.
func ToERPOrder(r erpdto.OrderAggregate) domain.ERPOrder {
	prMap := make(map[int64]erpdto.ProductionResource, len(r.ProductionResources))
	for _, pr := range r.ProductionResources {
		prMap[pr.ID] = pr
	}

	ops := make([]domain.ERPOrderOperation, len(r.Operations))
	for i, op := range r.Operations {
		ops[i] = toERPOrderOperation(op, prMap[op.ProductionResourceID])
	}

	return domain.ERPOrder{
		ID:                r.Order.ID,
		OrderNumber:       r.Order.OrderNumber,
		PartID:            r.Order.PartID,
		PartDescription:   r.Order.PartDescription,
		PlannedStartDate:  r.Order.PlannedStartDate,
		PlannedFinishDate: r.Order.PlannedFinishDate,
		ActualStartDate:   r.Order.ActualStartDate,
		ActualFinishDate:  r.Order.ActualFinishDate,
		Status:            string(r.Order.Status),
		Priority:          r.Order.Priority,
		Operations:        ops,
		ReceivedAt:        r.Order.ReceivedAt,
	}
}

func toERPOrderOperation(op erpdto.OrderOperationWithReports, pr erpdto.ProductionResource) domain.ERPOrderOperation {
	reports := make([]domain.ERPOrderReport, len(op.Reports))
	for i, rep := range op.Reports {
		reports[i] = toERPOrderReport(rep)
	}
	return domain.ERPOrderOperation{
		ID: op.ID,
		ProductionResource: domain.ERPProductionResource{
			ID:          pr.ID,
			Number:      pr.Number,
			Description: pr.Description,
			Type:        string(pr.Type),
		},
		PlannedStartDate:         op.PlannedStartDate,
		PlannedFinishDate:        op.PlannedFinishDate,
		ActualStartDate:          op.ActualStartDate,
		ActualFinishDate:         op.ActualFinishDate,
		Status:                   string(op.Status),
		ProductionResourceStatus: string(op.ProductionResourceStatus),
		Reports:                  reports,
	}
}

func toERPOrderReport(r erpdto.OrderReport) domain.ERPOrderReport {
	return domain.ERPOrderReport{
		ID:                 r.ID,
		Quantity:           r.Quantity,
		RestQuantity:       r.RestQuantity,
		Type:               string(r.Type),
		ReportingTimestamp: r.ReportingTimestamp,
		ActualReportedDate: r.ActualReportedDate,
	}
}
