package domain

import "time"

type OperationStatus string

const (
	OperationStatusNone              OperationStatus = "none"
	OperationStatusStarted           OperationStatus = "started"
	OperationStatusPartiallyShipped  OperationStatus = "partially_shipped"
	OperationStatusFullyShipped      OperationStatus = "fully_shipped"
	OperationStatusPartiallyReported OperationStatus = "partially_reported"
	OperationStatusFinished          OperationStatus = "finished"
)

type OrderOperation struct {
	ID                       string
	CompanyID                string
	FactoryID                string
	ProductionResourceID     string
	ReportedQuantity         float64
	RestQuantity             float64
	OrderId                  int64 // referenced by Order, which is NOT retreived.
	OperationNumber          int   // Represent the Order. Order represent the whole process
	PlannedStartDate         time.Time
	PlannedFinishDate        time.Time
	ActualStartDate          *time.Time
	ActualFinishDate         *time.Time
	Status                   OperationStatus // Whats planned to happen
	ProductionResourceStatus OperationStatus // What the Machine is doing right now
	Description              *string
	ReceivedAt               time.Time
}
