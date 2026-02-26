package chirpstackrest

import (
	"context"
	"encoding/json"
	"errors"
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

func (c *Client) CreateGateway(ctx context.Context, body CreateChirpstackGatewayRequest) error {
	url := fmt.Sprintf("%s/api/gateways", c.baseURL)
	resp, err := httpclient.DoRaw(
		c.httpClient,
		ctx,
		url,
		http.MethodPost,
		body,
		map[string]string{
			"Authorization": "Bearer " + c.token,
		},
	)

	if err != nil {
		var httpErr *httpclient.HTTPError
		if errors.As(err, &httpErr) {
			var apiErr ChirpstackError
			if json.Unmarshal(httpErr.Body, &apiErr) == nil {
				return fmt.Errorf("chirpstack error: %s, (code=%d)", apiErr.Message, apiErr.Code)
			}
			return fmt.Errorf("chirpstack error: %s", string(httpErr.Body))
		}
		return err
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	return nil
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
