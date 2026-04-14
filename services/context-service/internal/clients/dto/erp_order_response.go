package dto

import "time"

// ERPOrderResponse is the JSON shape returned by GET /api/v1/erp/orders.
//
// TODO: adjust field names to match the real ERP endpoint once it is implemented.
type ERPOrderResponse struct {
	OrderID        string    `json:"order_id"`
	ProductName    string    `json:"product_name"`
	Status         string    `json:"status"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	WorkcenterID   string    `json:"workcenter_id"`
	WorkcenterName string    `json:"workcenter_name"`
}

// ERPOrderReportingResponse is a single reporting event as returned by the ERP service.
//
// TODO: adjust field names to match the real ERP endpoint once it is implemented.
type ERPOrderReportingResponse struct {
	ReportingID string    `json:"reporting_id"`
	Timestamp   time.Time `json:"timestamp"`
	Quantity    float64   `json:"quantity"`
	Status      string    `json:"status"`
}

// ERPOrderDetailResponse is the JSON shape returned by GET /api/v1/erp/orders/{id}.
//
// TODO: adjust field names to match the real ERP endpoint once it is implemented.
type ERPOrderDetailResponse struct {
	ERPOrderResponse
	Reportings []ERPOrderReportingResponse `json:"reportings"`
}
