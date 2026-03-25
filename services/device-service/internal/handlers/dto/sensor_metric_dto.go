package dto

// SensorMetricRequest represents a single sensor metric submitted by the tenant operator.
type SensorMetricRequest struct {
	PayloadKey      string  `json:"payload_key"      binding:"required"`
	MeasurementType string  `json:"measurement_type" binding:"required"`
	Unit            *string `json:"unit"`
}

// UpsertSensorMetricsRequest is the request body for saving metrics for a sensor.
type UpsertSensorMetricsRequest struct {
	Metrics []SensorMetricRequest `json:"metrics" binding:"required"`
}

// SensorMetricResponse is the response body for a single sensor metric.
type SensorMetricResponse struct {
	PayloadKey      string  `json:"payload_key"`
	MeasurementType string  `json:"measurement_type"`
	Unit            *string `json:"unit"`
}

// SensorMetricListResponse is the response body for a list of sensor metrics.
type SensorMetricListResponse struct {
	TotalCount int                    `json:"total_count"`
	Metrics    []SensorMetricResponse `json:"metrics"`
}
