package domain

import "time"

// OrderReportType describes the semantic type of an operation reporting event.
type OrderReportType string

const (
	// OrderReportTypeRegular is a normal production reporting event.
	OrderReportTypeRegular OrderReportType = "regular"
	// OrderReportTypeSendToSubcontractor sends work to a subcontractor.
	OrderReportTypeSendToSubcontractor OrderReportType = "send_to_subcontractor"
	// OrderReportTypeReceiveFromSubcontractor receives work from a subcontractor.
	OrderReportTypeReceiveFromSubcontractor OrderReportType = "receive_from_subcontractor"
	// OrderReportTypeCancelRest cancels the remaining quantity.
	OrderReportTypeCancelRest OrderReportType = "cancel_rest"
	// OrderReportTypeMaterialOnly reports material without production quantity.
	OrderReportTypeMaterialOnly OrderReportType = "material_only"
	// OrderReportTypeSubcontractorInvoicePrice reports subcontractor invoice pricing.
	OrderReportTypeSubcontractorInvoicePrice OrderReportType = "subcontractor_invoice_price"
	// OrderReportTypeRecordingTerminal is a terminal-based report event.
	OrderReportTypeRecordingTerminal OrderReportType = "recording_terminal"
	// OrderReportTypeAdjustRecording adjusts a prior recording.
	OrderReportTypeAdjustRecording OrderReportType = "adjust_recording"
	// OrderReportTypeUndoRegular undoes a regular report.
	OrderReportTypeUndoRegular OrderReportType = "undo_regular"
	// OrderReportTypeUndoRecordingTerminal undoes a terminal report.
	OrderReportTypeUndoRecordingTerminal OrderReportType = "undo_recording_terminal"
	// OrderReportTypeUndoAdjustRecording undoes an adjusted recording.
	OrderReportTypeUndoAdjustRecording OrderReportType = "undo_adjust_recording"
	// OrderReportTypeSubcontractorPosteriorReport is a posterior subcontractor report.
	OrderReportTypeSubcontractorPosteriorReport OrderReportType = "subcontractor_posterior_report"
	// OrderReportTypeSubcontractorCostsManual reports manual subcontractor costs.
	OrderReportTypeSubcontractorCostsManual OrderReportType = "subcontractor_costs_manual_report"
	// OrderReportTypePickWorkCenter reports activity at a pick work center.
	OrderReportTypePickWorkCenter OrderReportType = "pick_work_center"
)

// OrderReport represents one operation reporting event from Monitor ERP.
type OrderReport struct {
	ID                   int64
	CompanyID            string
	OrderOperationID     int64 // referenced by OrderOperation id, (Not order)
	ProductionResourceID int64
	Quantity             float64
	RestQuantity         float64
	Type                 OrderReportType
	ReportingTimestamp   time.Time
	ActualReportedDate   *time.Time
	ReceivedAt           time.Time
}
