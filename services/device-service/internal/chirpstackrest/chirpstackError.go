package chirpstackrest

type ChirpstackError struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details"`
}

type ErrorDetail struct {
	Type  string            `json:"@type"`
	Props map[string]string `json:"-"`
}
