// Package dto defines data transfer objects used in the context service for representing aggregation rules in API responses.
package dto

import "time"

// AggregationRuleResponse represents the response structure for an aggregation rule in the context service API.
type AggregationRuleResponse struct {
	ID                string    `json:"id"`
	CompanyID         string    `json:"company_id"`
	Name              string    `json:"name"`
	ContextType       string    `json:"context_type"`
	MeasurementType   string    `json:"measurement_type"`
	AggregationMethod string    `json:"aggregation_method"`
	TimeBucketMinutes int       `json:"time_bucket_minutes"`
	IsActive          bool      `json:"is_active"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
