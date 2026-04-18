package domain

import "time"

// OperationStatus describes operation progress at plan or resource level.
type OperationStatus string

const (
	// OperationStatusNone means no operation progress is reported.
	OperationStatusNone OperationStatus = "none"
	// OperationStatusStarted means operation has started.
	OperationStatusStarted OperationStatus = "started"
	// OperationStatusPartiallyShipped means operation output is partially shipped.
	OperationStatusPartiallyShipped OperationStatus = "partially_shipped"
	// OperationStatusFullyShipped means operation output is fully shipped.
	OperationStatusFullyShipped OperationStatus = "fully_shipped"
	// OperationStatusPartiallyReported means operation is partially reported.
	OperationStatusPartiallyReported OperationStatus = "partially_reported"
	// OperationStatusFinished means operation is finished.
	OperationStatusFinished OperationStatus = "finished"
)

// OrderOperation represents one manufacturing operation from Monitor ERP.
type OrderOperation struct {
	ID                   int64
	CompanyID            string
	ProductionResourceID int64
	// ReportedQuantity         float64
	// RestQuantity             float64
	OrderID int64 // referenced by Order
	// OperationNumber          int   // Retreived in Get Order

	// Gives time interval per production resource
	PlannedStartDate  time.Time
	PlannedFinishDate time.Time
	ActualStartDate   *time.Time
	ActualFinishDate  *time.Time

	Status                   OperationStatus // Whats planned to happen
	ProductionResourceStatus OperationStatus // What the Machine is doing right now
	// Description              *string         // Monitor erp description
	ReceivedAt time.Time
}
