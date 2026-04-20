package domain

import (
	"context"
	"time"
)

// WorkCenterType describes the category of a production resource.
type WorkCenterType string

const (
	// WorkCenterTypeMachine identifies machine-based work centers.
	WorkCenterTypeMachine WorkCenterType = "machine"
	// WorkCenterTypeManualWork identifies manual labor work centers.
	WorkCenterTypeManualWork WorkCenterType = "manual_work"
	// WorkCenterTypeSubContract identifies subcontractor work centers.
	WorkCenterTypeSubContract WorkCenterType = "sub_contract"
	// WorkCenterTypePool identifies pooled capacity work centers.
	WorkCenterTypePool WorkCenterType = "pool"
	// WorkCenterTypePick identifies pick/warehouse work centers.
	WorkCenterTypePick WorkCenterType = "pick"
)

// ProductionResource represents one work center from Monitor ERP.
type ProductionResource struct {
	ID          int64
	CompanyID   string
	Number      string // human readable machine name
	Description string // Monitor erp description
	Type        WorkCenterType
	ReceivedAt  time.Time
}

// ProductionResourceRepo defines the retrive operation
type ProductionResourceRepo interface {
	FindAllProductionResource(ctx context.Context, companyID string) ([]ProductionResource, error)
}

// ProductionResourceSvc defines the business logic for a production resource
type ProductionResourceSvc interface {
	GetAll(ctx context.Context, companyID string) ([]ProductionResource, error)
}
