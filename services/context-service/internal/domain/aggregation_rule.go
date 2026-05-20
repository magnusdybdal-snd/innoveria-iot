package domain

import (
	"context"
	"time"
)

// AggregationRule is the domain model for aggregation rules
type AggregationRule struct {
	ID                string
	CompanyID         string
	Name              string
	ContextType       string
	MeasurementType   string
	AggregationMethod AggMethod
	TimeBucketMinutes int
	IsActive          bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// AggMethod represents the type of aggregation to perform on measurements for a given context rule.
type AggMethod string

const (
	// Avg aggregates by average.
	Avg AggMethod = "AVG"
	// Sum aggregates by sum.
	Sum AggMethod = "SUM"
	// Max aggregates by maximum.
	Max AggMethod = "MAX"
	// Min aggregates by minimum.
	Min AggMethod = "MIN"
)

// IsValid checks if the AggMethod value is one of the allowed aggregation methods.
func (a AggMethod) IsValid() bool {
	return a == Avg || a == Sum || a == Max || a == Min
}

// RuleRepository defines persistence operations for aggregation rules.
type RuleRepository interface {
	GetByCompanyID(ctx context.Context, companyID string) ([]AggregationRule, error)
	GetByID(ctx context.Context, ruleID string) (AggregationRule, error)
	Create(ctx context.Context, rule AggregationRule) (string, error)
	DeleteByID(ctx context.Context, ruleID string) error
}

// RuleService defines business logic operations for aggregation rules.
type RuleService interface {
	GetRules(ctx context.Context, companyID string) ([]AggregationRule, error)
	CreateRule(ctx context.Context, rule AggregationRule) (string, error)
	DeleteRule(ctx context.Context, ruleID string) error
}
