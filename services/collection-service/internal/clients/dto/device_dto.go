// Package dto defines the data transfer objects for collection-service client responses.
package dto

// SensorMetricResponse mirrors the device-service SensorMetricResponse shape.
type SensorMetricResponse struct {
	PayloadKey      string  `json:"payload_key"`
	MeasurementType string  `json:"measurement_type"`
	Unit            *string `json:"unit"`
}

// SensorMetricListResponse mirrors the device-service SensorMetricListResponse shape.
type SensorMetricListResponse struct {
	TotalCount int                    `json:"total_count"`
	Metrics    []SensorMetricResponse `json:"metrics"`
}
