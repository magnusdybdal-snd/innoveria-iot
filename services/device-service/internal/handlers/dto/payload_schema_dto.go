package dto

// DiscoverPayloadKeysRequest is the request body for discovering payload keys for a profile.
type DiscoverPayloadKeysRequest struct {
	PayloadKeys []string `json:"payload_keys" binding:"required"`
}

// PayloadSchemaLabelRequest represents a single payload key label submitted by the admin.
type PayloadSchemaLabelRequest struct {
	PayloadKey      string  `json:"payload_key"      binding:"required"`
	MeasurementType string  `json:"measurement_type" binding:"required"`
	Unit            *string `json:"unit"`
}

// SavePayloadSchemaLabelsRequest is the request body for saving labels for a profile.
type SavePayloadSchemaLabelsRequest struct {
	Labels []PayloadSchemaLabelRequest `json:"labels" binding:"required"`
}

// PayloadSchemaResponse is the response body for a single payload schema row.
type PayloadSchemaResponse struct {
	ID                  string  `json:"id"`
	ChirpstackProfileID string  `json:"chirpstack_profile_id"`
	PayloadKey          string  `json:"payload_key"`
	MeasurementType     *string `json:"measurement_type"`
	Unit                *string `json:"unit"`
}

// PayloadSchemaListResponse is the response body for a list of payload schema rows.
type PayloadSchemaListResponse struct {
	TotalCount int                     `json:"total_count"`
	Schemas    []PayloadSchemaResponse `json:"schemas"`
}

// DraftProfilesResponse is the response body for the admin draft badge endpoint.
type DraftProfilesResponse struct {
	TotalCount int      `json:"total_count"`
	ProfileIDs []string `json:"profile_ids"`
}
