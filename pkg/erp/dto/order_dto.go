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
	ID          int64  `json:"ID"`
	OrderNumber string `json:"OrderNumber"`
}

// Order is the shared API contract for one manufacturing order.
type Order struct {
	ID                int64       `json:"ID"`
	CompanyID         string      `json:"CompanyID"`
	OrderNumber       string      `json:"OrderNumber"`
	PartID            string      `json:"PartID"`
	PartDescription   string      `json:"PartDescription"`
	PlannedStartDate  time.Time   `json:"PlannedStartDate"`
	PlannedFinishDate time.Time   `json:"PlannedFinishDate"`
	ActualStartDate   *time.Time  `json:"ActualStartDate"`
	ActualFinishDate  *time.Time  `json:"ActualFinishDate"`
	Status            OrderStatus `json:"Status"`
	Priority          int         `json:"Priority"`
	ReceivedAt        time.Time   `json:"ReceivedAt"`
}

// OrderReport is the shared API contract for one operation report event.
type OrderReport struct {
	ID                   int64           `json:"ID"`
	OrderOperationID     int64           `json:"OrderOperationID"`
	ProductionResourceID int64           `json:"ProductionResourceID"`
	Quantity             float64         `json:"Quantity"`
	RestQuantity         float64         `json:"RestQuantity"`
	Type                 OrderReportType `json:"Type"`
	ReportingTimestamp   time.Time       `json:"ReportingTimestamp"`
	ActualReportedDate   *time.Time      `json:"ActualReportedDate"`
	ReceivedAt           time.Time       `json:"ReceivedAt"`
}

// OrderOperationWithReports is one operation plus related reports/resource.
type OrderOperationWithReports struct {
	ID                       int64           `json:"ID"`
	ProductionResourceID     int64           `json:"ProductionResourceID"`
	OrderID                  int64           `json:"OrderID"`
	PlannedStartDate         time.Time       `json:"PlannedStartDate"`
	PlannedFinishDate        time.Time       `json:"PlannedFinishDate"`
	ActualStartDate          *time.Time      `json:"ActualStartDate"`
	ActualFinishDate         *time.Time      `json:"ActualFinishDate"`
	Status                   OperationStatus `json:"Status"`
	ProductionResourceStatus OperationStatus `json:"ProductionResourceStatus"`
	ReceivedAt               time.Time       `json:"ReceivedAt"`
	Reports                  []OrderReport   `json:"Reports"`
}

// OrderAggregate is the shared API contract for one aggregated order view.
type OrderAggregate struct {
	Order               []Order                     `json:"Order"`
	Operations          []OrderOperationWithReports `json:"Operations"`
	ProductionResources []ProductionResource        `json:"ProductionResources"`
}
