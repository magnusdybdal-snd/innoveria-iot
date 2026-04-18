package dto

// WorkCenter represents a production resource in Monitor ERP.
//
// A work center is a physical or logical unit where manufacturing operations
// are performed. It can represent machines, workstations, or groups of workers.
//
// Work centers define capacity, availability, and staffing, and are referenced
// by manufacturing operations to indicate where the work is executed.
type WorkCenter struct {
	ID          int64  `json:"Id"`
	Number      string `json:"Number"`      // Unique identifier / code
	Description string `json:"Description"` // Name of the work center
	// OperationDescription string `json:"OperationDescription"` // Description used in operations

	// See documentation for enum mapping, (handled in erp service)
	Type int `json:"Type"` // Type of work center (machine, labor, etc.)

	// DepartmentID int64 `json:"DepartmentId"`
	// WarehouseID  int64 `json:"WarehouseId"`

	// Capacity & planning
	// BasicTime             string  `json:"BasicTime"`             // Capacity per day (TimeSpan)
	// NumberOfPlanningUnits float64 `json:"NumberOfPlanningUnits"` // Machines or people
	// NumberOfFlows         float64 `json:"NumberOfFlows"`         // Parallel capacity per order

	// Efficiency & simulation
	// AvailabilityFactor float64 `json:"AvailabilityFactor"` // % uptime
	// SimulationFactor   float64 `json:"SimulationFactor"`   // % used in simulations

	// Staffing
	// UnitStaffFactor  float64 `json:"UnitStaffFactor"`  // Staffing for runtime
	// SetupStaffFactor float64 `json:"SetupStaffFactor"` // Staffing for setup

	// Time precision (e.g. minutes, hours)
	// TimePrecision int `json:"TimePrecision"`

	// Optional / expandable fields
	// CommentID *int64 `json:"CommentId,omitempty"`
}
