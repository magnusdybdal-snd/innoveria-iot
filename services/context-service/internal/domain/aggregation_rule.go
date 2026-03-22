package domain

import (
	"context"
	"time"
)

// AggregationRule is the domain model for affregation rules
type AggregationRule struct {
	ID                string
	CompanyID         string
	Name              string
	ContextType       string
	MeasurementType   string
	AggregationMethod string
	TimeBucketMinutes int
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// RuleRepository defines persistence operations for aggregation rules.
type RuleRepository interface {
	GetByCompanyID(ctx context.Context, companyID string) ([]AggregationRule, error)
	GetByID(ctx context.Context, ruleID string) (AggregationRule, error)
}

// RuleService defines business logic operations for aggregation rules.
type RuleService interface {
	GetRules(ctx context.Context, companyID string) ([]AggregationRule, error)
}
