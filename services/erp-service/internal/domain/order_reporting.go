package domain

import "time"

type OrderReportType int

const (
	Regular OrderReportType = iota
	SendToSubcontractor
	ReceiveFromSubcontractor
	CancelRest
	MaterialOnly
	SubcontractorInvoicePrice
	RecordingTerminal
	AdjustRecording
	UndoRegular
	UndoRecordingTerminal
	UndoAdjustRecording
	SubcontractorPosteriorReport
	SubcontractorCostsManualReport
	PickWorkCenter
)

type OrderStatus int

const (
	NotInitialized OrderStatus = iota
	Registrered
	Printed
	Started
	Finished
	PostCalculated
	Delivered
	_                      // 7 (unused)
	_                      // 8 (unused)
	Historical OrderStatus = 9
)

type OrderReporting struct {
	ID                 int
	OperationID        int // refrenced by OrderOperation
	Quantity           float64
	WorkCenterID       int // refrenced by workcenter
	ActualReportedDate *time.Time
	RestQuantity       float64
	Type               OrderReportType
	WipLocation        *string
	PreviousNodeStatus OrderStatus
	ReportingTimestamp time.Time
}
