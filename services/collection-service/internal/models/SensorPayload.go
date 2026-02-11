package models

type SensorPayload struct {
	Data            string       `json:"data"` // base64 encoded
}
