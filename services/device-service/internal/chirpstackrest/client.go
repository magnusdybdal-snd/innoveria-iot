package chirpstackrest

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/device-service/internal/config"

	"innoveria-iot/pkg/httpclient"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *httpclient.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		baseURL:    cfg.ChirpstackURL,
		token:      cfg.ChirpstackSecret,
		httpClient: httpclient.New(),
	}
}

// Returns all gateways in chirpstack
// TODO: Add authentication for tennatns
func (c *Client) GetAllGateways(ctx context.Context, limit int) (ChirpstackGatewayList, error) {

	// Chirpstack needs a limit to send the correct response
	url := fmt.Sprintf("%s/api/gateways?limit=%d", c.baseURL, limit)

	resp, err := httpclient.DoRequest[ChirpstackGatewayList](
		c.httpClient,
		ctx,
		url,
		http.MethodGet,
		nil,
		map[string]string{
			"Authorization": "Bearer " + c.token,
		},
	)
	if err != nil {
		return ChirpstackGatewayList{}, err
	}

	return resp, nil
}

// Returns all sensors in chirpstack
// TODO: Add authentication for tennatns
func (c *Client) GetAllSensors(ctx context.Context, limit int, applicationID string) (ChirpstackSensorList, error) {
	url := fmt.Sprintf("%s/api/devices?limit=%d&applicationId=%s", c.baseURL, limit, applicationID)
	resp, err := httpclient.DoRequest[ChirpstackSensorList](
		c.httpClient,
		ctx,
		url,
		http.MethodGet,
		nil,
		map[string]string{
			"Authorization": "Bearer " + c.token,
		},
	)
	if err != nil {
		return ChirpstackSensorList{}, err
	}

	return resp, nil
}
