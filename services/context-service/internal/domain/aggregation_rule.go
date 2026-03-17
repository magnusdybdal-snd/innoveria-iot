package domain

import "time"

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
