package mappers

import (
	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/domain"
)

// ToERPOrder converts an ERP order summary DTO to its domain representation.
func ToERPOrder(r dto.ERPOrderResponse) domain.ERPOrder {
	return domain.ERPOrder{
		OrderID:        r.OrderID,
		ProductName:    r.ProductName,
		Status:         r.Status,
		StartTime:      r.StartTime,
		EndTime:        r.EndTime,
		WorkcenterID:   r.WorkcenterID,
		WorkcenterName: r.WorkcenterName,
	}
}

// ToERPOrderDetail converts an ERP order detail DTO (including reportings) to its domain representation.
func ToERPOrderDetail(r dto.ERPOrderDetailResponse) domain.ERPOrderDetail {
	reportings := make([]domain.ERPOrderReporting, len(r.Reportings))
	for i, rep := range r.Reportings {
		reportings[i] = domain.ERPOrderReporting{
			ReportingID: rep.ReportingID,
			Timestamp:   rep.Timestamp,
			Quantity:    rep.Quantity,
			Status:      rep.Status,
		}
	}
	return domain.ERPOrderDetail{
		ERPOrder:   ToERPOrder(r.ERPOrderResponse),
		Reportings: reportings,
	}
}
