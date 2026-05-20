package dto

import "time"

// BucketResponse represents a single time bucket in the context data API response.
type BucketResponse struct {
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
	Value       float64   `json:"value"`
}

// ContextDataResponse is the API response shape for a single device's computed context data.
// TotalValue is the aggregate across the full time window; Buckets breaks it into time slices.
// The frontend decides which representation to render.
type ContextDataResponse struct {
	DeviceEUI    string           `json:"device_eui"`
	CompanyID    string           `json:"company_id"`
	ContextType  string           `json:"context_type"`
	Unit         string           `json:"unit"`
	PeriodStart  time.Time        `json:"period_start"`
	PeriodEnd    time.Time        `json:"period_end"`
	TotalValue   float64          `json:"total_value"`
	Buckets      []BucketResponse `json:"buckets"`
	CalculatedAt time.Time        `json:"calculated_at"`
}
