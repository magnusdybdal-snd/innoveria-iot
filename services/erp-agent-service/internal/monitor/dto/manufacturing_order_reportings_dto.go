package dto

import "time"

// ManufacturingOrderOperationReporting represents a reporting event
// for a manufacturing operation in Monitor ERP.
//
// This is the MOST important entity for:
// - timestamps
// - machine state
// - order tracking
// - production analytics
//
// Each record = one "event" or "report" from a work center.
type ManufacturingOrderOperationReporting struct {
	ID int64 `json:"Id"`

	// --- ORDER / OPERATION ---
	ManufacturingOrderId     int64  `json:"ManufacturingOrderId"`
	ManufacturingOrderNumber string `json:"ManufacturingOrderNumber"`
	OperationNumber          int    `json:"OperationNumber"`
	ReportingId              *int64 `json:"ReportingId,omitempty"`

	// --- WORK CENTER ---
	WorkCenterId int64 `json:"WorkCenterId"`

	// --- TIME / TIMESTAMPS ---
	ReportingTimestamp time.Time  `json:"ReportingTimestamp"`
	ActualReportedDate *time.Time `json:"ActualReportedDate,omitempty"`

	// --- QUANTITIES ---
	ReportedQuantity         float64  `json:"ReportedQuantity"`
	RestQuantity             float64  `json:"RestQuantity"`
	PreviousRestQuantity     float64  `json:"PreviousRestQuantity"`
	OverReceiveQuantity      float64  `json:"OverReceiveQuantity"`
	OriginalApprovedQuantity *float64 `json:"OriginalApprovedQuantityFromDeliveryRow,omitempty"`

	// --- STATE / TYPE ---
	Type int `json:"Type"` // enum (VERY important)

	// --- STATUS ---
	PreviousNodeStatus int `json:"PreviousNodeStatus"` // enum

	// --- EMPLOYEES ---
	EmployeeId          *int64 `json:"EmployeeId,omitempty"`
	ReportingEmployeeId *int64 `json:"ReportingEmployeeId,omitempty"`

	// --- WAREHOUSE / LOCATION ---
	WarehouseId         int64  `json:"WarehouseId"`
	WipLocation         string `json:"WipLocation"`
	PreviousWipLocation string `json:"PreviousWipLocation"`

	// --- TIME DETAILS ---
	OfWhichTime *string `json:"OfWichTime,omitempty"` // duration (TimeSpan)

	// --- COSTS ---
	SetupCost           *float64 `json:"SetupCost,omitempty"`
	SetupCostCurrencyId *int64   `json:"SetupCostCurrencyId,omitempty"`
	SetupCostFactor1    float64  `json:"SetupCostFactor1"`
	SetupCostFactor2    float64  `json:"SetupCostFactor2"`
	SetupCostFactor3    float64  `json:"SetupCostFactor3"`

	UnitCost           *float64 `json:"UnitCost,omitempty"`
	UnitCostCurrencyId *int64   `json:"UnitCostCurrencyId,omitempty"`
	UnitCostFactor1    float64  `json:"UnitCostFactor1"`
	UnitCostFactor2    float64  `json:"UnitCostFactor2"`
	UnitCostFactor3    float64  `json:"UnitCostFactor3"`

	// --- SUBCONTRACTING ---
	IsSubcontractor                    bool       `json:"IsSubcontractor"`
	SubcontractorRequestedDeliveryDate *time.Time `json:"SubcontractorRequestedDeliveryDate,omitempty"`
	SubcontractorActualDeliveryDate    *time.Time `json:"SubcontractorActualDeliveryDate,omitempty"`

	// --- FLAGS ---
	ReportRestQuantitiesOnSubsequentOperations bool `json:"ReportRestQuantitiesOnSubsequentOperations"`
	IsExportedToManagementAccounting           bool `json:"IsExportedToManagementAccounting"`
	AutomaticMaterialReporting                 bool `json:"AutomaticMaterialReporting"`

	// --- COMMENTS / REASON ---
	CommentId *int64  `json:"CommentId,omitempty"`
	Comment   *string `json:"Comment,omitempty"`

	ReasonCodeTimeConsumptionId *int64 `json:"ReasonCodeTimeConsumptionId,omitempty"`

	// --- PURCHASING ---
	PurchaseOrderDeliveryRowId *int64 `json:"PurchaseOrderDeliveryRowId,omitempty"`

	// --- META ---
	RowNumber int `json:"RowNumber"`
}
