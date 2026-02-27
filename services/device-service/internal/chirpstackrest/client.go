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

// TODO: Add tennant authentication, so add tennantID as query

/*
	Application requests (factory area)
*/

// Returns all availabe application (factory areas)
func (c *Client) GetOneApplication(ctx context.Context, applicationId string) (ChirpstackApplication, error) {
	url := fmt.Sprintf("%s/api/applications/%s", c.baseURL, applicationId)

	resp, err := httpclient.DoRequest[ChirpstackApplication](
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
		return ChirpstackApplication{}, handleChirpstackError(err)
	}

	return resp, nil
}

// Returns all chirpstack applications
func (c *Client) GetAllApplication(ctx context.Context, limit int) (ChirpstackApplicationList, error) {
	url := fmt.Sprintf("%s/api/applications?limit=%d", c.baseURL, limit)

	resp, err := httpclient.DoRequest[ChirpstackApplicationList](
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
		return ChirpstackApplicationList{}, handleChirpstackError(err)
	}

	return resp, nil
}

// Creates a new application in chirpstack
func (c *Client) CreateApplication(ctx context.Context, body ChirpstackApplicationList) error {
	url := fmt.Sprintf("%s/api/applications", c.baseURL)

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
		return handleChirpstackError(err)
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

// Renames the application name
func (c *Client) RenameApplication(ctx context.Context, body ChirpstackApplication) error {
	url := fmt.Sprintf("%s/api/applications/%s", c.baseURL, body.ID)

	resp, err := httpclient.DoRaw(
		c.httpClient,
		ctx,
		url,
		http.MethodPut,
		body,
		map[string]string{
			"Authorization": "Bearer " + c.token,
		},
	)
	if err != nil {
		return handleChirpstackError(err)
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

/*
	Gateway requests
*/

// Returns all gateways in chirpstack
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
		return ChirpstackGatewayList{}, handleChirpstackError(err)
	}

	return resp, nil
}

// GetOneGateway returns one chirpstack gateway
// the parameter is gatewayEUI which chirpstack calls gatewayId
func (c *Client) GetOneGateway(ctx context.Context, gatewayEUI string) (ChirpstackGateway, error) {
	url := fmt.Sprintf("%s/api/gateways/%s", c.baseURL, gatewayEUI)

	resp, err := httpclient.DoRequest[ChirpstackGateway](
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
		return ChirpstackGateway{}, handleChirpstackError(err)
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
		return handleChirpstackError(err)
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	return nil
}

func (c *Client) RenameGateway(ctx context.Context, body CreateChirpstackGatewayRequest) error {
	url := fmt.Sprintf("%s/api/gateways/%s", c.baseURL, body.GatewayEUI)
	resp, err := httpclient.DoRaw(
		c.httpClient,
		ctx,
		url,
		http.MethodPut,
		body,
		map[string]string{
			"Authorization": "Bearer " + c.token,
		},
	)
	if err != nil {
		return handleChirpstackError(err)
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteGateway(ctx context.Context, gatewayEUI string) error {
	url := fmt.Sprintf("%s/api/gateways/%s", c.baseURL, gatewayEUI)
	resp, err := httpclient.DoRaw(
		c.httpClient,
		ctx,
		url,
		http.MethodDelete,
		nil,
		map[string]string{
			"Authorization": "Bearer " + c.token,
		},
	)
	if err != nil {
		return handleChirpstackError(err)
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

/*
	Sensor Requests
*/
// Returns all sensors in chirpstack
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
		return ChirpstackSensorList{}, handleChirpstackError(err)
	}

	return resp, nil
}

func (c *Client) GetAllSensorProfiles(ctx context.Context) (DeviceProfileListResponse, error) {
	url := fmt.Sprintf("%s/api/device-profiles?limit=10", c.baseURL)
	resp, err := httpclient.DoRequest[DeviceProfileListResponse](
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
		return DeviceProfileListResponse{}, handleChirpstackError(err)
	}

	return resp, nil
}
