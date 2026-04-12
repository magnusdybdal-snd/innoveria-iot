package domain

type WorkCenterType int

const (
	Machine WorkCenterType = iota
	ManualWork
	SubContract
	Pool
	Pick
)

type WorkCenter struct {
	ID int
	// CompanyID
	// FactoryID
	// FactoryAreaID ???
	Number      string
	Description string
	Type        WorkCenterType
}
