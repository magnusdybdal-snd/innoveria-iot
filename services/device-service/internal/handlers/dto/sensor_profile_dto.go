package dto

// SensorProfileResponse TODO(@vinjar): add proper documentation.
type SensorProfileResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`      // LoRaWAN region (EU868)
	MACVersion string `json:"mac_version"` // LoRaWAN version
	VendorId   string `json:"vendor_id"`   // Identification of model producer
	VendorName string `json:"vendor"`      // Vendor name
	// IsCustom bool   `json:"isCustom,omitempty"`
}

// SensorProfileListResponse TODO(@vinjar): add proper documentation.
type SensorProfileListResponse struct {
	TotalCount     int                     `json:"total_count"`
	SensorProfiles []SensorProfileResponse `json:"sensor_profiles"`
}
