package dto

// DeviceSensorResponse mirrors the device-service SensorResponse shape.
type DeviceSensorResponse struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	DeviceEUI          string `json:"device_eui"`
	ProductionResource *int64 `json:"production_resource"`
	ElectricitySensor  bool   `json:"electricity_sensor"`
	Voltage            *int   `json:"voltage"`
}

// DeviceSensorListResponse mirrors the device-service SensorListResponse shape.
type DeviceSensorListResponse struct {
	TotalCount int                    `json:"total_count"`
	Sensors    []DeviceSensorResponse `json:"sensors"`
}

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
