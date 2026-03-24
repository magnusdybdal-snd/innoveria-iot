package dto

// CreateMeasurementTypeRequest is the request body for creating a new measurement type.
type CreateMeasurementTypeRequest struct {
	Slug        string  `json:"slug"         binding:"required"`
	DisplayName string  `json:"display_name" binding:"required"`
	Description *string `json:"description"`
	DefaultUnit *string `json:"default_unit"`
}

// MeasurementTypeResponse is the response body for a single measurement type.
type MeasurementTypeResponse struct {
	Slug        string  `json:"slug"`
	DisplayName string  `json:"display_name"`
	Description *string `json:"description"`
	DefaultUnit *string `json:"default_unit"`
	Deprecated  bool    `json:"deprecated"`
}

// MeasurementTypeListResponse is the response body for a list of measurement types.
type MeasurementTypeListResponse struct {
	TotalCount       int                       `json:"total_count"`
	MeasurementTypes []MeasurementTypeResponse `json:"measurement_types"`
}
