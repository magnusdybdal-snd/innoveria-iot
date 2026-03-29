package dto

// PayloadTagsResponse is the response payload for the payload-tags endpoint
type PayloadTagsResponse struct {
	Keys []string `json:"keys"`
}