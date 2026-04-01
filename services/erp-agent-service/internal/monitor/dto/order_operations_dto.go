package dto

import "time"

// ManufacturingOrderOperation represents a single step (operation)
// in a manufacturing order process from Monitor ERP
// (ManufacturingOrderOperations endpoint).
//
// For sensor context enrichment, the fields below are usually enough:
// - ManufacturingOrderId / OperationNumber (business identity)
// - WorkCenterId (machine mapping key)
// - PlannedStartDate / PlannedFinishDate and ActualStartDate / ActualFinishDate (time window)
// - Status / WorkshopOperationStatus (running vs done)
//
// Most cost/accounting fields and staffing factors are optional for first-pass
// sensor correlation and can be ignored until needed by downstream analytics.
type ManufacturingOrderOperation struct {
	ID int64 `json:"Id"`

	Priority *int `json:"Priority,omitempty"`

	// Primary join key to map operation -> physical machine/sensor source.
	WorkCenterId int64 `json:"WorkCenterId"`
	PartId       int64 `json:"PartId"`

	PlannedQuantity  float64 `json:"PlannedQuantity"`
	ReportedQuantity float64 `json:"ReportedQuantity"`
	RestQuantity     float64 `json:"RestQuantity"`

	// TimeSpan values come as "HH:MM:SS" from Monitor.
	ReportedSetupTime string `json:"ReportedSetupTime"`
	ReportedUnitTime  string `json:"ReportedUnitTime"`

	RejectedQuantity float64 `json:"RejectedQuantity"`
	WipLocation      *string `json:"WipLocation,omitempty"`

	ReportNumber int64 `json:"ReportNumber"`

	InstructionCommentId *int64 `json:"InstructionCommentId,omitempty"`

	// Stable operation identity used to group related events.
	ManufacturingOrderId     int64 `json:"ManufacturingOrderId"`
	ManufacturingOrderNodeId int64 `json:"ManufacturingOrderNodeId"`

	OperationNumber int `json:"OperationNumber"`

	PlannedSetupTime string `json:"PlannedSetupTime"`
	PlannedUnitTime  string `json:"PlannedUnitTime"`

	// Planned interval used as coarse expected time window.
	PlannedStartDate  time.Time `json:"PlannedStartDate"`
	PlannedFinishDate time.Time `json:"PlannedFinishDate"`

	// Actual interval is better than planned when available.
	ActualStartDate  *time.Time `json:"ActualStartDate,omitempty"`
	ActualFinishDate *time.Time `json:"ActualFinishDate,omitempty"`

	OperationRowId *int64 `json:"OperationRowId,omitempty"`

	// Keep both status fields; they often indicate different lifecycle states.
	Status                  int `json:"Status"`
	WorkshopOperationStatus int `json:"WorkshopOperationStatus"`

	BundleId    *int64  `json:"BundleId,omitempty"`
	Description *string `json:"Description,omitempty"`

	ClearedStatus  bool    `json:"ClearedStatus"`
	OnPriorityPlan bool    `json:"OnPriorityPlan"`
	FixedLeadTime  *string `json:"FixedLeadTime,omitempty"`

	// Optional for sensor-context use; mostly planning/cost semantics.
	UnitStaffingFactor  *float64 `json:"UnitStaffingFactor,omitempty"`
	SetupStaffingFactor *float64 `json:"SetupStaffingFactor,omitempty"`
}
