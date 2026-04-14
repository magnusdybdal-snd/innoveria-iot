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

type OrderReport struct {
	ID                   int64
	CompanyID            string
	OrderOperationID     int // referenced by OrderOperation id, (Not order)
	ProductionResourceID string
	Quantity             float64
	RestQuantity         float64
	Type                 OrderReportType
	ReportingTimestamp   time.Time
	ActualReportedDate   *time.Time
	ReceivedAt           time.Time
}
