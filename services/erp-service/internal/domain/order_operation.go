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
	ID                      int
	CompanyID               int
	FactoryID               int
	ProductionResource      int
	ReportedQuantity        float64
	RestQuantity            float64
	OrderId                 int64 // referenced by Order, which is NOT retreived.
	OperationNumber         int   // Represent the Order. Order represent the whole process
	PlannedStartDate        time.Time
	PlannedFinishDate       time.Time
	ActualStartDate         *time.Time
	ActualFinishDate        *time.Time
	Status                  OperationStatus
	WorkshopOperationStatus OperationStatus
	Description             *string
	ReceivedAt              time.Time
}
