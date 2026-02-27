package chirpstackrest

import "time"

// Chirpstack v4 model for device profiles (sensor profiles)
type DeviceProfileListResponse struct {
	Result     []DeviceProfile `json:"result"`
	TotalCount int             `json:"totalCount"`
}

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
