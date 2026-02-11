package models

// Chirpstack v4 MQTT uplink event
type ChirpstackUpEvent struct {
	DeduplicationID string       `json:"deduplicationId"`
	Time            string       `json:"time"`
	DeviceInfo      DeviceInfo   `json:"deviceInfo"`
	DevAddr         string       `json:"devAddr"`
	ADR             bool         `json:"adr"`
	DR              int          `json:"dr"`
	FCnt            int          `json:"fCnt"`
	FPort           int          `json:"fPort"`
	Confirmed       bool         `json:"confirmed"`
	Data            string       `json:"data"` // base64 encoded
	RxInfo          []RxInfo     `json:"rxInfo"`
	TxInfo          TxInfo       `json:"txInfo"`
	RegionConfigID  string       `json:"regionConfigId"`
}


type DeviceInfo struct {
	TenantID         string            `json:"tenantId"`
	TenantName       string            `json:"tenantName"`
	ApplicationID    string            `json:"applicationId"`
	ApplicationName  string            `json:"applicationName"`
	DeviceProfileID  string            `json:"deviceProfileId"`
	DeviceProfileName string           `json:"deviceProfileName"`
	DeviceName       string            `json:"deviceName"`
	DevEUI           string            `json:"devEui"`
	DeviceClass      string            `json:"deviceClassEnabled"`
	Tags             map[string]string `json:"tags"`
}


type RxInfo struct {
	GatewayID string     `json:"gatewayId"`
	UplinkID  uint32     `json:"uplinkId"`
	NSTime    string     `json:"nsTime"`
	RSSI      int        `json:"rssi"`
	SNR       float64    `json:"snr"`
	Channel   int        `json:"channel"`
	RFChain   int        `json:"rfChain"`
	Location  Location   `json:"location"`
	Context   string     `json:"context"`
}


type Location struct {
	Source string `json:"source"`
}

type TxInfo struct {
	Frequency  int64      `json:"frequency"`
	Modulation Modulation `json:"modulation"`
}


type Modulation struct {
	LoRa LoRaModulation `json:"lora"`
}

type LoRaModulation struct {
	Bandwidth       int    `json:"bandwidth"`
	SpreadingFactor int    `json:"spreadingFactor"`
	CodeRate        string `json:"codeRate"`
}

