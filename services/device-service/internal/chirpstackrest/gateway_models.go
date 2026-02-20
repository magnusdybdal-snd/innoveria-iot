package chirpstackrest

import "time"

// Chirpstack v4 model for Gateway
type ChirpstackGatewayList struct {
	Result     []ChirpstackGateway `json:"result"`
	TotalCount int                 `json:"totalCount"`
}

type ChirpstackGateway struct {
	GatewayID   string `json:"gatewayId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TenantID    string `json:"tenantId"`

	State string `json:"state"`

	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	LastSeenAt time.Time `json:"lastSeenAt"`

	Location   ChirpstackGatewayLocation `json:"location"`
	Properties map[string]string         `json:"properties"`
}

type ChirpstackGatewayLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Accuracy  float64 `json:"accuracy"`
	Source    string  `json:"source"`
}
