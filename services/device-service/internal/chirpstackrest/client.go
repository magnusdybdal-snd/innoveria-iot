package chirpstackrest

import (
	"net/http"
	"net/url"

	"innoveria-iot/device-service/internal/config"

	"innoveria-iot/pkg/httpclient"
)

type Client struct {
	baseURL    *url.URL
	token      string
	httpClient *http.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		baseURL:    &cfg.ChirpstackURL,
		token:      cfg.ChirpstackSecret,
		httpClient: httpclient.New(),
	}
}

func (c *Client) GetApplications() {

}
