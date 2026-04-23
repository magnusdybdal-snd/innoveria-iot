package dto

import "time"

// WorkCenterType describes the category of a production resource.
type WorkCenterType string

//nolint:godoclint // SQL constant name is intentionally exported for cross-package reuse.
const (
	WorkCenterTypeMachine     WorkCenterType = "machine"
	WorkCenterTypeManualWork  WorkCenterType = "manual_work"
	WorkCenterTypeSubContract WorkCenterType = "sub_contract"
	WorkCenterTypePool        WorkCenterType = "pool"
	WorkCenterTypePick        WorkCenterType = "pick"
)

// ProductionResource is the shared API contract for production resources.
type ProductionResource struct {
	ID          int64          `json:"ID"`
	Number      string         `json:"Number"`
	Description string         `json:"Description"`
	Type        WorkCenterType `json:"Type"`
	ReceivedAt  time.Time      `json:"ReceivedAt"`
}
