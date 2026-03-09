package chirpstackrest

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/config"

	"innoveria-iot/pkg/httpclient"
)

// Client is an HTTP client for the Chirpstack REST API.
type Client struct {
	baseURL    string
	token      string
	httpClient *httpclient.Client
}

// New creates a new Chirpstack Client from the given configuration.
func New(cfg config.Config) *Client {
	return &Client{
		baseURL:    cfg.ChirpstackURL,
		token:      cfg.ChirpstackSecret,
		httpClient: httpclient.New(),
	}
}

// TODO: Add tennant authentication, so add tennantID as query

/*
	Application requests (sensor groups)
*/

// GetOneApplication retrieves a single Chirpstack application by its ID.
func (c *Client) GetOneApplication(ctx context.Context, applicationId string) (dto.ChirpstackApplication, error) {
	url := fmt.Sprintf("%s/api/applications/%s", c.baseURL, applicationId)

	resp, err := httpclient.DoRequest[dto.ChirpstackApplication](
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
		return dto.ChirpstackApplication{}, handleChirpstackError(err)
	}

	return resp, nil
}

// GetAllApplication retrieves a paginated list of Chirpstack applications up to the given limit.
func (c *Client) GetAllApplication(ctx context.Context, limit int) (dto.ChirpstackApplicationList, error) {
	url := fmt.Sprintf("%s/api/applications?limit=%d", c.baseURL, limit)

	resp, err := httpclient.DoRequest[dto.ChirpstackApplicationList](
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
		return dto.ChirpstackApplicationList{}, handleChirpstackError(err)
	}

	return resp, nil
}

// CreateApplication registers a new application in Chirpstack.
func (c *Client) CreateApplication(ctx context.Context, body dto.CreateChirpstackApplication) (string, error) {
	url := fmt.Sprintf("%s/api/applications", c.baseURL)

	resp, err := httpclient.DoRequest[dto.ChirpstackApplicationCreateResponse](
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
		return "", handleChirpstackError(err)
	}

	return resp.ID, nil
}

// CreateTenant registers a new tenant in Chirpstack.
func (c *Client) CreateTenant(ctx context.Context, body dto.CreateChirpstackApplication) (string, error) {
	url := fmt.Sprintf("%s/api/tenants", c.baseURL)

	resp, err := httpclient.DoRequest[dto.ChirpstackTenantCreateResponse](
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
		return "", handleChirpstackError(err)
	}

	return resp.ID, nil
}

// DeleteTenant deletes a tenant from Chirpstack
func (c *Client) DeleteTenant(ctx context.Context, tenantID string) error {
	url := fmt.Sprintf("%s/api/tenants/%s", c.baseURL, tenantID)
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

// RenameApplication updates an existing Chirpstack application.
func (c *Client) RenameApplication(ctx context.Context, body dto.ChirpstackApplication) error {
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

// GetAllGateways retrieves a paginated list of gateways from Chirpstack up to the given limit.
func (c *Client) GetAllGateways(ctx context.Context, limit int) (dto.ChirpstackGatewayList, error) {
	// Chirpstack needs a limit to send the correct response
	url := fmt.Sprintf("%s/api/gateways?limit=%d", c.baseURL, limit)

	resp, err := httpclient.DoRequest[dto.ChirpstackGatewayList](
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
		return dto.ChirpstackGatewayList{}, handleChirpstackError(err)
	}

	return resp, nil
}

// GetOneGateway retrieves a single gateway from Chirpstack by its EUI.
// Note: Chirpstack refers to the EUI as gatewayId.
func (c *Client) GetOneGateway(ctx context.Context, gatewayEUI string) (dto.ChirpstackGateway, error) {
	url := fmt.Sprintf("%s/api/gateways/%s", c.baseURL, gatewayEUI)

	resp, err := httpclient.DoRequest[dto.ChirpstackGateway](
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
		return dto.ChirpstackGateway{}, handleChirpstackError(err)
	}

	return resp, nil
}

// CreateGateway registers a new gateway in Chirpstack.
func (c *Client) CreateGateway(ctx context.Context, body dto.CreateChirpstackGatewayRequest) error {
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

// RenameGateway updates the name of an existing Chirpstack gateway, identified by its EUI.
func (c *Client) RenameGateway(ctx context.Context, body dto.CreateChirpstackGatewayRequest) error {
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

// DeleteGateway removes a gateway from Chirpstack by its EUI.
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

// GetAllSensors retrieves a paginated list of devices in a Chirpstack application up to the given limit.
func (c *Client) GetAllSensors(ctx context.Context, limit int, applicationID string) (dto.ChirpstackSensorList, error) {
	url := fmt.Sprintf("%s/api/devices?limit=%d&applicationId=%s", c.baseURL, limit, applicationID)
	resp, err := httpclient.DoRequest[dto.ChirpstackSensorList](
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
		return dto.ChirpstackSensorList{}, handleChirpstackError(err)
	}

	return resp, nil
}

// GetOneSensor retrieves a single device from Chirpstack by its EUI.
// Note: Chirpstack refers to the EUI as deviceId.
func (c *Client) GetOneSensor(ctx context.Context, deviceEUI string) (dto.ChirpstackSensor, error) {
	url := fmt.Sprintf("%s/api/devices/%s", c.baseURL, deviceEUI)

	resp, err := httpclient.DoRequest[dto.ChirpstackSensor](
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
		return dto.ChirpstackSensor{}, handleChirpstackError(err)
	}

	return resp, nil
}

// CreateSensor registers a new device in Chirpstack.
func (c *Client) CreateSensor(ctx context.Context, body dto.ChirpstackSensorRequest) error {
	url := fmt.Sprintf("%s/api/devices", c.baseURL)
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

// UpdateSensor updates a device's name, description and device profile in Chirpstack.
// DeviceEUI identifies the device and cannot be changed.
func (c *Client) UpdateSensor(ctx context.Context, body dto.ChirpstackSensorRequest) error {
	url := fmt.Sprintf("%s/api/devices/%s", c.baseURL, body.DeviceEUI)
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

// DeleteSensor removes a device from Chirpstack by its EUI.
func (c *Client) DeleteSensor(ctx context.Context, deviceEUI string) error {
	url := fmt.Sprintf("%s/api/devices/%s", c.baseURL, deviceEUI)
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

// GetAllSensorProfiles retrieves a paginated list of device profiles from Chirpstack up to the given limit.
func (c *Client) GetAllSensorProfiles(ctx context.Context, limit int) (dto.DeviceProfileListResponse, error) {
	url := fmt.Sprintf("%s/api/device-profiles?limit=%d", c.baseURL, limit)
	resp, err := httpclient.DoRequest[dto.DeviceProfileListResponse](
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
		return dto.DeviceProfileListResponse{}, handleChirpstackError(err)
	}

	return resp, nil
}
