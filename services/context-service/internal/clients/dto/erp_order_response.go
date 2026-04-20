package dto

import "time"

// ERPOrderSummaryResponse is the slim JSON shape returned by GET /api/v1/erp/orders.
type ERPOrderSummaryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ERPProductionResourceResponse is the JSON shape for a work center returned
// by the erp-service GET orders/{id} endpoint.
//
// TODO: adjust field names to match the real erp-service GET endpoint once it is implemented.
type ERPProductionResourceResponse struct {
	ID          int64  `json:"id"`
	Number      string `json:"number"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// ERPOrderOperationResponse is the JSON shape for a single manufacturing
// operation, including its associated production resource.
//
// TODO: adjust field names to match the real erp-service GET endpoint once it is implemented.
type ERPOrderOperationResponse struct {
	ID                       int64                         `json:"id"`
	ProductionResource       ERPProductionResourceResponse `json:"production_resource"`
	PlannedStartDate         time.Time                     `json:"planned_start_date"`
	PlannedFinishDate        time.Time                     `json:"planned_finish_date"`
	ActualStartDate          *time.Time                    `json:"actual_start_date"`
	ActualFinishDate         *time.Time                    `json:"actual_finish_date"`
	Status                   string                        `json:"status"`
	ProductionResourceStatus string                        `json:"production_resource_status"`
}

// ERPOrderResponse is the JSON shape returned by GET /api/v1/erp/orders,
// with operations and production resources nested inline.
//
// TODO: adjust field names to match the real erp-service GET endpoint once it is implemented.
type ERPOrderResponse struct {
	ID                int64                       `json:"id"`
	OrderNumber       string                      `json:"order_number"`
	PartDescription   string                      `json:"part_description"`
	PlannedStartDate  time.Time                   `json:"planned_start_date"`
	PlannedFinishDate time.Time                   `json:"planned_finish_date"`
	ActualStartDate   *time.Time                  `json:"actual_start_date"`
	ActualFinishDate  *time.Time                  `json:"actual_finish_date"`
	Status            string                      `json:"status"`
	Priority          int                         `json:"priority"`
	Operations        []ERPOrderOperationResponse `json:"operations"`
}
