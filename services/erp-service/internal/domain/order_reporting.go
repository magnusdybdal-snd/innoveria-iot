package domain

import "time"

type OrderReportType string

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

type OrderStatus string

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

type OrderReporting struct {
	ID                 int
	OperationID        int // referenced by OrderOperation
	Quantity           float64
	WorkCenterID       int // referenced by workcenter
	ActualReportedDate *time.Time
	RestQuantity       float64
	Type               OrderReportType
	WipLocation        *string
	PreviousNodeStatus OrderStatus
	ReportingTimestamp time.Time
}
