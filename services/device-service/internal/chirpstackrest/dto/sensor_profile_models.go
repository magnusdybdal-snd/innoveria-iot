package dto

import "time"

// DeviceProfileListResponse is the paginated response from the Chirpstack GET /api/device-profiles endpoint.
type DeviceProfileListResponse struct {
	Result     []DeviceProfile `json:"result"`
	TotalCount int             `json:"totalCount"`
}

// DeviceProfile represents a single LoRaWAN device profile as returned by the Chirpstack API.
type DeviceProfile struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	DeviceID          string    `json:"deviceId"`
	DeviceName        string    `json:"deviceName"`
	FirmwareVersion   string    `json:"firmwareVersion"`
	MACVersion        string    `json:"macVersion"`
	RegParamsRevision string    `json:"regParamsRevision"`
	Region            string    `json:"region"`
	SupportsClassB    bool      `json:"supportsClassB"`
	SupportsClassC    bool      `json:"supportsClassC"`
	SupportsOtaa      bool      `json:"supportsOtaa"`
	VendorID          string    `json:"vendorId"`
	VendorName        string    `json:"vendorName"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
