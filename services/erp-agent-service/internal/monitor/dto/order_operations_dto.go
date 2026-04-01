package dto

import "time"

// ManufacturingOrderOperation represents a single step (operation)
// in a manufacturing order process.
//
// This defines HOW a product is made (routing).
// Each operation is typically executed at a WorkCenter.
type ManufacturingOrderOperation struct {
	ID int64 `json:"Id"`

	// --- ORDER RELATION ---
	ManufacturingOrderId     int64 `json:"ManufacturingOrderId"`
	ManufacturingOrderNodeId int64 `json:"ManufacturingOrderNodeId"`

	// --- OPERATION ---
	OperationNumber int    `json:"OperationNumber"`
	Description     string `json:"Description"`

	// --- WORK CENTER ---
	WorkCenterId int64 `json:"WorkCenterId"`

	// --- PART ---
	PartId int64 `json:"PartId"`

	// --- QUANTITIES ---
	PlannedQuantity  float64 `json:"PlannedQuantity"`
	ReportedQuantity float64 `json:"ReportedQuantity"`
	RestQuantity     float64 `json:"RestQuantity"`
	RejectedQuantity float64 `json:"RejectedQuantity"`

	// --- TIME (PLANNED) ---
	PlannedSetupTime time.Duration `json:"PlannedSetupTime"`
	PlannedUnitTime  time.Duration `json:"PlannedUnitTime"`

	PlannedStartDate  time.Time `json:"PlannedStartDate"`
	PlannedFinishDate time.Time `json:"PlannedFinishDate"`

	// --- TIME (ACTUAL) ---
	ActualStartDate  *time.Time `json:"ActualStartDate,omitempty"`
	ActualFinishDate *time.Time `json:"ActualFinishDate,omitempty"`

	ReportedSetupTime time.Duration `json:"ReportedSetupTime"`
	ReportedUnitTime  time.Duration `json:"ReportedUnitTime"`

	// --- STATUS ---
	Status                  int `json:"Status"` // Started, Finished, etc.
	WorkshopOperationStatus int `json:"WorkshopOperationStatus"`

	// --- LOCATION ---
	WipLocation string `json:"WipLocation"`

	// --- REPORTING ---
	ReportNumber int64 `json:"ReportNumber"`

	// --- OPTIONAL ---
	OperationRowId *int64 `json:"OperationRowId,omitempty"`
	BundleId       *int64 `json:"BundleId,omitempty"`

	// --- FLAGS ---
	ClearedStatus  bool `json:"ClearedStatus"`
	OnPriorityPlan bool `json:"OnPriorityPlan"`

	// --- STAFFING ---
	UnitStaffingFactor  *float64 `json:"UnitStaffingFactor,omitempty"`
	SetupStaffingFactor *float64 `json:"SetupStaffingFactor,omitempty"`
}
