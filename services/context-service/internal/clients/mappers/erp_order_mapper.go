package mappers

import (
	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/domain"
)

// ToERPOrderSummary converts a slim order summary DTO to its domain representation.
func ToERPOrderSummary(r dto.ERPOrderSummaryResponse) domain.ERPOrderSummary {
	return domain.ERPOrderSummary{
		ID:   r.ID,
		Name: r.Name,
	}
}

// ToERPProductionResource converts a production resource DTO to its domain representation.
func ToERPProductionResource(r dto.ERPProductionResourceResponse) domain.ERPProductionResource {
	return domain.ERPProductionResource{
		ID:          r.ID,
		Number:      r.Number,
		Description: r.Description,
		Type:        r.Type,
	}
}

// ToERPOrderReport converts an order report DTO to its domain representation.
func ToERPOrderReport(r dto.ERPOrderReportResponse) domain.ERPOrderReport {
	return domain.ERPOrderReport{
		ID:                 r.ID,
		Quantity:           r.Quantity,
		RestQuantity:       r.RestQuantity,
		Type:               r.Type,
		ReportingTimestamp: r.ReportingTimestamp,
		ActualReportedDate: r.ActualReportedDate,
	}
}

// ToERPOrderOperation converts an order operation DTO to its domain representation.
func ToERPOrderOperation(r dto.ERPOrderOperationResponse) domain.ERPOrderOperation {
	reports := make([]domain.ERPOrderReport, len(r.Reports))
	for i, rep := range r.Reports {
		reports[i] = ToERPOrderReport(rep)
	}
	return domain.ERPOrderOperation{
		ID:                       r.ID,
		ProductionResource:       ToERPProductionResource(r.ProductionResource),
		PlannedStartDate:         r.PlannedStartDate,
		PlannedFinishDate:        r.PlannedFinishDate,
		ActualStartDate:          r.ActualStartDate,
		ActualFinishDate:         r.ActualFinishDate,
		Status:                   r.Status,
		ProductionResourceStatus: r.ProductionResourceStatus,
		Reports:                  reports,
	}
}

// ToERPOrder converts an ERP order DTO (with nested operations and reports) to its domain representation.
func ToERPOrder(r dto.ERPOrderResponse) domain.ERPOrder {
	ops := make([]domain.ERPOrderOperation, len(r.Operations))
	for i, op := range r.Operations {
		ops[i] = ToERPOrderOperation(op)
	}
	return domain.ERPOrder{
		ID:                r.ID,
		OrderNumber:       r.OrderNumber,
		PartID:            r.PartID,
		PartDescription:   r.PartDescription,
		PlannedStartDate:  r.PlannedStartDate,
		PlannedFinishDate: r.PlannedFinishDate,
		ActualStartDate:   r.ActualStartDate,
		ActualFinishDate:  r.ActualFinishDate,
		Status:            r.Status,
		Priority:          r.Priority,
		Operations:        ops,
		ReceivedAt:        r.ReceivedAt,
	}
}
