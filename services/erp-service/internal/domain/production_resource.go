package domain

import "time"

type WorkCenterType string

const (
	WorkCenterTypeMachine     WorkCenterType = "machine"
	WorkCenterTypeManualWork  WorkCenterType = "manual_work"
	WorkCenterTypeSubContract WorkCenterType = "sub_contract"
	WorkCenterTypePool        WorkCenterType = "pool"
	WorkCenterTypePick        WorkCenterType = "pick"
)

type ProductionResource struct {
	ID            string
	CompanyID     string
	FactoryID     string
	FactoryAreaID string
	Number        string // human readable machine name
	Description   string
	Type          WorkCenterType
	ReceivedAt    time.Time
}
