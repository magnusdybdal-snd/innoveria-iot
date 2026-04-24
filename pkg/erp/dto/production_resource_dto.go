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
	ID          int64          `json:"id"`
	Number      string         `json:"number"`
	Description string         `json:"description"`
	Type        WorkCenterType `json:"type"`
	ReceivedAt  time.Time      `json:"received_at"`
}
