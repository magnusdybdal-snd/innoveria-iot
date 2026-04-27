package dto

import "time"

/**
	This file will be omited if/when'pkg' will take over
	contract responsibilities between services.
**/

// ERPOrderSummaryResponse is the slim JSON shape returned by GET /api/v1/erp/orders.
type ERPOrderSummaryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ERPProductionResourceResponse is the JSON shape for a work center returned
// by the erp-service GET /orders/{id} endpoint.
type ERPProductionResourceResponse struct {
	ID          int64  `json:"id"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// ERPOrderReportResponse is the JSON shape for a single production reporting
// event nested under its parent operation.
type ERPOrderReportResponse struct {
	ID                 int64      `json:"id"`
	Quantity           float64    `json:"quantity"`
	RestQuantity       float64    `json:"rest_quantity"`
	Type               string     `json:"type"`
	ReportingTimestamp time.Time  `json:"reporting_timestamp"`
	ActualReportedDate *time.Time `json:"actual_reported_date"`
}

// ERPOrderOperationResponse is the JSON shape for a single manufacturing
// operation, including its associated production resource and reporting events.
type ERPOrderOperationResponse struct {
	ID                       int64                         `json:"id"`
	ProductionResource       ERPProductionResourceResponse `json:"production_resource"`
	PlannedStartDate         time.Time                     `json:"planned_start_date"`
	PlannedFinishDate        time.Time                     `json:"planned_finish_date"`
	ActualStartDate          *time.Time                    `json:"actual_start_date"`
	ActualFinishDate         *time.Time                    `json:"actual_finish_date"`
	Status                   string                        `json:"status"`
	ProductionResourceStatus string                        `json:"production_resource_status"`
	Reports                  []ERPOrderReportResponse      `json:"reports"`
}

// ERPOrderResponse is the JSON shape returned by GET /api/v1/erp/orders/{id}.
type ERPOrderResponse struct {
	ID                int64                       `json:"id"`
	OrderNumber       string                      `json:"order_number"`
	PartID            string                      `json:"part_id"`
	PartDescription   string                      `json:"part_description"`
	PlannedStartDate  time.Time                   `json:"planned_start_date"`
	PlannedFinishDate time.Time                   `json:"planned_finish_date"`
	ActualStartDate   *time.Time                  `json:"actual_start_date"`
	ActualFinishDate  *time.Time                  `json:"actual_finish_date"`
	Status            string                      `json:"status"`
	Priority          int                         `json:"priority"`
	Operations        []ERPOrderOperationResponse `json:"operations"`
	ReceivedAt        time.Time                   `json:"received_at"`
}
