package dto

import (
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// OrderReportingResponse represents a single production reporting event in the API response.
type OrderReportingResponse struct {
	ReportingID string    `json:"reporting_id"`
	Timestamp   time.Time `json:"timestamp"`
	Quantity    float64   `json:"quantity"`
	Status      string    `json:"status"`
}

// OrderDetailResponse is the API response shape for a single ERP order.
type OrderDetailResponse struct {
	OrderID        string                   `json:"order_id"`
	ProductName    string                   `json:"product_name"`
	Status         string                   `json:"status"`
	StartTime      time.Time                `json:"start_time"`
	EndTime        time.Time                `json:"end_time"`
	WorkcenterID   string                   `json:"workcenter_id"`
	WorkcenterName string                   `json:"workcenter_name"`
	Reportings     []OrderReportingResponse `json:"reportings"`
}

// MapOrderDetailDomainToDTO converts a domain ERPOrderDetail to its API response shape.
func MapOrderDetailDomainToDTO(order domain.ERPOrderDetail) OrderDetailResponse {
	reportings := make([]OrderReportingResponse, len(order.Reportings))
	for i, r := range order.Reportings {
		reportings[i] = OrderReportingResponse{
			ReportingID: r.ReportingID,
			Timestamp:   r.Timestamp,
			Quantity:    r.Quantity,
			Status:      r.Status,
		}
	}
	return OrderDetailResponse{
		OrderID:        order.OrderID,
		ProductName:    order.ProductName,
		Status:         order.Status,
		StartTime:      order.StartTime,
		EndTime:        order.EndTime,
		WorkcenterID:   order.WorkcenterID,
		WorkcenterName: order.WorkcenterName,
		Reportings:     reportings,
	}
}

// MapOrdersDomainToDTO converts a slice of domain ERPOrderDetail to its API response shape.
func MapOrdersDomainToDTO(orders []domain.ERPOrderDetail) []OrderDetailResponse {
	result := make([]OrderDetailResponse, len(orders))
	for i, o := range orders {
		result[i] = MapOrderDetailDomainToDTO(o)
	}
	return result
}
