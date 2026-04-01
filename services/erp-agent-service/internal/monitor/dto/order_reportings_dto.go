package dto

import (
	"time"
)

// ManufacturingOrderOperationReporting represents a reporting event
// from Monitor ERP's ManufacturingOrderOperationReportings endpoint.
//
// For sensor correlation, this DTO is event-centric and often the best source of
// "what changed when" at a machine/work center.
//
// Usually required for contextualizing sensor streams:
// - OperationId, WorkCenterId
// - ReportingTimestamp (event time)
// - Type, PreviousNodeStatus (state transition meaning)
// - Quantity / RestQuantity (production progress)
//
// Usually optional at first:
// - Cost/currency fields, comments, reason codes, subcontractor details.
type ManufacturingOrderOperationReporting struct {
	ID int64 `json:"Id"`

	// OperationId links reporting events back to operation-level context.
	OperationId int64   `json:"OperationId"`
	Quantity    float64 `json:"Quantity"`

	// SetupTime *string `json:"SetupTime,omitempty"`
	// UnitTime  *string `json:"UnitTime,omitempty"`

	// Primary join key for mapping event -> machine/sensor source.
	WorkCenterId int64 `json:"WorkCenterId"`

	// EmployeeId          *int64 `json:"EmployeeId,omitempty"`
	// ReportingEmployeeId *int64 `json:"ReportingEmployeeId,omitempty"`

	// WarehouseId int64 `json:"WarehouseId"`

	// ReportingTimestamp is generally the canonical event timestamp for joins.
	ActualReportedDate *time.Time `json:"ActualReportedDate"`

	// PreviousRestQuantity float64 `json:"PreviousRestQuantity"`
	RestQuantity float64 `json:"RestQuantity"`
	// OverReceiveQuantity  float64 `json:"OverReceiveQuantity"`

	// SetupCost                    *float64 `json:"SetupCost,omitempty"`
	// SetupCostCurrencyId          *int64   `json:"SetupCostCurrencyId,omitempty"`
	// SetupCostFactor1             float64  `json:"SetupCostFactor1"`
	// SetupCostFactor2             float64  `json:"SetupCostFactor2"`
	// SetupCostFactor3             float64  `json:"SetupCostFactor3"`
	// SetupCostInForeignCurrency   *float64 `json:"SetupCostInForeignCurrency,omitempty"`
	// SetupCostInForeignCurrencyId *int64   `json:"SetupCostInForeignCurrencyCurrencyId,omitempty"`

	// UnitCost                    *float64 `json:"UnitCost,omitempty"`
	// UnitCostCurrencyId          *int64   `json:"UnitCostCurrencyId,omitempty"`
	// UnitCostFactor1             float64  `json:"UnitCostFactor1"`
	// UnitCostFactor2             float64  `json:"UnitCostFactor2"`
	// UnitCostFactor3             float64  `json:"UnitCostFactor3"`
	// UnitCostInForeignCurrency   *float64 `json:"UnitCostInForeignCurrency,omitempty"`
	// UnitCostInForeignCurrencyId *int64   `json:"UnitCostInForeignCurrencyCurrencyId,omitempty"`

	// IsSubcontractor bool `json:"IsSubcontractor"`
	// Type is critical to interpret event semantics (start/finish/adjust/etc.).
	// Enum mapping is handled in erp-service
	Type int `json:"Type"`
	// SubcontractorRequestedDeliveryDate *time.Time `json:"SubcontractorRequestedDeliveryDate,omitempty"`
	// SubcontractorActualDeliveryDate    *time.Time `json:"SubcontractorActualDeliveryDate,omitempty"`

	// ReportRestQuantitiesOnSubsequentOperations bool `json:"ReportRestQuantitiesOnSubsequentOperations"`
	// IsExportedToManagementAccounting           bool `json:"IsExportedToManagementAccounting"`

	WipLocation *string `json:"WipLocation,omitempty"`
	// PreviousWipLocation *string `json:"PreviousWipLocation,omitempty"`

	// Monitor uses "OfWichTime" in JSON (typo in upstream API name).
	// OfWhichTime *string `json:"OfWichTime,omitempty"` // TimeSpan "HH:MM:SS"

	// CommentId *int64           `json:"CommentId,omitempty"`
	// Comment   *json.RawMessage `json:"Comment,omitempty"`

	// ReasonCodeTimeConsumptionId *int64           `json:"ReasonCodeTimeConsumptionId,omitempty"`
	// ReasonCodeTimeConsumption   *json.RawMessage `json:"ReasonCodeTimeConsumption,omitempty"`

	// PurchaseOrderDeliveryRowId *int64 `json:"PurchaseOrderDeliveryRowId,omitempty"`

	// AutomaticMaterialReporting bool `json:"AutomaticMaterialReporting"`
	// PreviousNodeStatus helps explain the state before this event.
	// Enum mapping is handled in erp-service
	PreviousNodeStatus int `json:"PreviousNodeStatus"`
	// OriginalApprovedQuantityFromDeliveryRow *float64 `json:"OriginalApprovedQuantityFromDeliveryRow,omitempty"`
	// ReportingId                             *int64   `json:"ReportingId,omitempty"`
	// Event time typically used when aligning ERP events to sensor timeseries.
	ReportingTimestamp time.Time `json:"ReportingTimestamp"`

	// RowNumber int `json:"RowNumber"`
}
