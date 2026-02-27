package dto

type SensorProfileResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`     // LoRaWAN region (EU868)
	MACVersion string `json:"macVersion"` // LoRaWAN version
	VendorId   string `json:"vendorId"`   // Identification of model producer
	VendorName string `json:"vendor"`     // Vendor name
	// IsCustom bool   `json:"isCustom,omitempty"`
}

type SensorProfileListResponse struct {
	TotalCount     int                     `json:"total_count"`
	SensorProfiles []SensorProfileResponse `json:"sensor_profiles"`
}
