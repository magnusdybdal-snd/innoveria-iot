package dto

import "time"

// ManufacturingOrder represents the HIGH-LEVEL context of a production job.
//
// This struct is NOT meant to mirror the full Monitor ERP model.
// Instead, it is a REDUCED, DOMAIN-FOCUSED version tailored for:
//
// - Sensor data correlation
// - Production tracking
// - Time alignment (planned vs actual)
// - Progress analysis
//
// Think of this as:
// "What is being produced, how much, and when?"
type ManufacturingOrder struct {
	// --- Identity ---

	ID          int    // Internal unique identifier (primary key from Monitor)
	OrderNumber string // Human-readable order number (useful for UI/logging)

	// --- Product context ---

	PartID          string // What is being produced (critical for sensor grouping)
	PartDescription string // Optional: human-readable product name

	// --- Quantity / progress ---

	PlannedQuantity  float64 // Total quantity planned for production
	ReportedQuantity float64 // Quantity already produced (actual output)
	RestQuantity     float64 // Remaining quantity (Planned - Reported)

	// --- Time context ---

	PlannedStartDate  time.Time // When production was supposed to start
	PlannedFinishDate time.Time // When production is expected to finish

	ActualStartDate  *time.Time // When production ACTUALLY started (nil if not started)
	ActualFinishDate *time.Time // When production ACTUALLY finished (nil if ongoing)

	// --- Status / lifecycle ---

	Status   int // High-level progress state (NotStarted, Started, Finished, etc.)
	Priority int // Optional: useful for scheduling or filtering important orders

	// =========================================================
	//  NOT NEEDED FOR SENSOR / CONTEXT PLATFORM (COMMENTED OUT)
	// =========================================================

	// --- ERP / financial / administrative ---

	// CustomerID        string
	// ProjectID         string
	// OrderType         int
	// OrderCategory     string
	// Cost              float64
	// Revenue           float64

	// --- Deep ERP structure ---

	// Nodes             []Node
	// Materials         []Material
	// Reservations      []Reservation

	// --- Redundant if you already fetch separately ---

	// Operations        []ManufacturingOrderOperation // Extract this in ManufacturingOrder

	// --- Misc / rarely useful ---

	// CreatedDate       time.Time
	// ModifiedDate      time.Time
	// CreatedBy         string
	// ChangedBy         string
}
