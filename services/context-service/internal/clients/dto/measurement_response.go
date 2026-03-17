// Package dto contains data transfer objects for external service responses.
package dto

import "time"

// MeasurementResponse is the response shape returned by the collection service measurements endpoint.
type MeasurementResponse struct {
	DeviceEUI string         `json:"device_eui"`
	Timestamp time.Time      `json:"timestamp"`
	Payload   map[string]any `json:"payload"`
	CompanyID string         `json:"company_id"`
}
