package domain

type WorkCenterType string

const (
	WorkCenterTypeMachine     WorkCenterType = "machine"
	WorkCenterTypeManualWork  WorkCenterType = "manual_work"
	WorkCenterTypeSubContract WorkCenterType = "sub_contract"
	WorkCenterTypePool        WorkCenterType = "pool"
	WorkCenterTypePick        WorkCenterType = "pick"
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
