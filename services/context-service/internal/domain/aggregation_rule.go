package domain

import (
	"context"
	"slices"
	"time"
)

// ValidAggregationMethods is the single source of truth for allowed aggregation strategies.
var ValidAggregationMethods = []string{"AVG", "SUM", "MAX", "MIN"}

// IsValidAggregationMethod reports whether the given method is one of the allowed aggregation strategies.
func IsValidAggregationMethod(method string) bool {
	return slices.Contains(ValidAggregationMethods, method)
}

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
	Create(ctx context.Context, rule AggregationRule) (string, error)
}

// RuleService defines business logic operations for aggregation rules.
type RuleService interface {
	GetRules(ctx context.Context, companyID string) ([]AggregationRule, error)
	CreateRule(ctx context.Context, rule AggregationRule) (string, error)
}
