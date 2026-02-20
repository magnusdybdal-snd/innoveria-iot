package chirpstackrest

import (
	"net/url"

	"innoveria-iot/device-service/internal/config"

	"innoveria-iot/pkg/httpclient"
)

type Client struct {
	baseURL    *url.URL
	token      string
	httpClient *httpclient.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		baseURL:    &cfg.ChirpstackURL,
		token:      cfg.ChirpstackSecret,
		httpClient: httpclient.New(),
	}
}

func GetSensorStatus(devEUI string)
