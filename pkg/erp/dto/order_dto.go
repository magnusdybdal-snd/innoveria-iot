package dto

import "time"

// OrderStatus describes the lifecycle state of a manufacturing order.
type OrderStatus string

//nolint:godoclint // constant name is intentionally exported for cross-package reuse.
const (
	OrderStatusNotInitialized OrderStatus = "not_initialized"
	OrderStatusRegistered     OrderStatus = "registered"
	OrderStatusPrinted        OrderStatus = "printed"
	OrderStatusStarted        OrderStatus = "started"
	OrderStatusFinished       OrderStatus = "finished"
	OrderStatusPostCalculated OrderStatus = "post_calculated"
	OrderStatusDelivered      OrderStatus = "delivered"
	OrderStatusHistorical     OrderStatus = "historical"
)

// OperationStatus describes operation progress at plan or resource level.
type OperationStatus string

//nolint:godoclint // constant names are intentionally exported for shared DTO contracts.
const (
	OperationStatusNone              OperationStatus = "none"
	OperationStatusStarted           OperationStatus = "started"
	OperationStatusPartiallyShipped  OperationStatus = "partially_shipped"
	OperationStatusFullyShipped      OperationStatus = "fully_shipped"
	OperationStatusPartiallyReported OperationStatus = "partially_reported"
	OperationStatusFinished          OperationStatus = "finished"
)

// OrderReportType describes the semantic type of an operation reporting event.
type OrderReportType string

//nolint:godoclint,lll // Constant names are intentionally exported and explicit for shared contracts.
const (
	OrderReportTypeRegular                      OrderReportType = "regular"
	OrderReportTypeSendToSubcontractor          OrderReportType = "send_to_subcontractor"
	OrderReportTypeReceiveFromSubcontractor     OrderReportType = "receive_from_subcontractor"
	OrderReportTypeCancelRest                   OrderReportType = "cancel_rest"
	OrderReportTypeMaterialOnly                 OrderReportType = "material_only"
	OrderReportTypeSubcontractorInvoicePrice    OrderReportType = "subcontractor_invoice_price"
	OrderReportTypeRecordingTerminal            OrderReportType = "recording_terminal"
	OrderReportTypeAdjustRecording              OrderReportType = "adjust_recording"
	OrderReportTypeUndoRegular                  OrderReportType = "undo_regular"
	OrderReportTypeUndoRecordingTerminal        OrderReportType = "undo_recording_terminal"
	OrderReportTypeUndoAdjustRecording          OrderReportType = "undo_adjust_recording"
	OrderReportTypeSubcontractorPosteriorReport OrderReportType = "subcontractor_posterior_report"
	OrderReportTypeSubcontractorCostsManual     OrderReportType = "subcontractor_costs_manual_report"
	OrderReportTypePickWorkCenter               OrderReportType = "pick_work_center"
)

// OrderSummary is the list contract for selecting orders by ID/number.
type OrderSummary struct {
	ID          int64  `json:"id"`
	OrderNumber string `json:"order_number"`
}

// Order is the shared API contract for one manufacturing order.
type Order struct {
	ID                int64       `json:"id"`
	CompanyID         string      `json:"company_id"`
	OrderNumber       string      `json:"order_number"`
	PartID            string      `json:"part_id"`
	PartDescription   string      `json:"part_description"`
	PlannedStartDate  time.Time   `json:"planned_start_date"`
	PlannedFinishDate time.Time   `json:"planned_finish_date"`
	ActualStartDate   *time.Time  `json:"actual_start_date"`
	ActualFinishDate  *time.Time  `json:"actual_finish_date"`
	Status            OrderStatus `json:"status"`
	Priority          int         `json:"priority"`
	ReceivedAt        time.Time   `json:"received_at"`
}

// OrderReport is the shared API contract for one operation report event.
type OrderReport struct {
	ID                   int64           `json:"id"`
	OrderOperationID     int64           `json:"order_operation_id"`
	ProductionResourceID int64           `json:"production_resource_id"`
	Quantity             float64         `json:"quantity"`
	RestQuantity         float64         `json:"rest_quantity"`
	Type                 OrderReportType `json:"type"`
	ReportingTimestamp   time.Time       `json:"reporting_timestamp"`
	ActualReportedDate   *time.Time      `json:"actual_reported_date"`
	ReceivedAt           time.Time       `json:"received_at"`
}

// OrderOperationWithReports is one operation plus related reports/resource.
type OrderOperationWithReports struct {
	ID                       int64           `json:"id"`
	ProductionResourceID     int64           `json:"production_resource_id"`
	OrderID                  int64           `json:"order_id"`
	PlannedStartDate         time.Time       `json:"planned_start_date"`
	PlannedFinishDate        time.Time       `json:"planned_finish_date"`
	ActualStartDate          *time.Time      `json:"actual_start_date"`
	ActualFinishDate         *time.Time      `json:"actual_finish_date"`
	Status                   OperationStatus `json:"status"`
	ProductionResourceStatus OperationStatus `json:"production_resource_status"`
	ReceivedAt               time.Time       `json:"received_at"`
	Reports                  []OrderReport   `json:"reports"`
}

// OrderAggregate is the shared API contract for one aggregated order view.
type OrderAggregate struct {
	Order               Order                       `json:"order"`
	Operations          []OrderOperationWithReports `json:"operations"`
	ProductionResources []ProductionResource        `json:"production_resources"`
}
