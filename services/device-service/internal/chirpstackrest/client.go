package chirpstackrest

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/device-service/internal/chirpstackrest/dto"
	"innoveria-iot/device-service/internal/config"

	"innoveria-iot/pkg/httpclient"
)

// Client TODO(@vinjar): add proper documentation.
type Client struct {
	baseURL    string
	token      string
	httpClient *httpclient.Client
}

// New TODO(@vinjar): add proper documentation.
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

// GetOneApplication TODO(@vinjar): add proper documentation.
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

// GetAllApplication TODO(@vinjar): add proper documentation.
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

// CreateApplication TODO(@vinjar): add proper documentation.
func (c *Client) CreateApplication(ctx context.Context, body dto.CreateChirpstackApplication) error {
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

// RenameApplication TODO(@vinjar): add proper documentation.
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

// GetAllGateways TODO(@vinjar): add proper documentation.
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

// GetOneGateway returns one chirpstack gateway
// the parameter is gatewayEUI which chirpstack calls gatewayId
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

// CreateGateway TODO(@vinjar): add proper documentation.
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

// RenameGateway TODO(@vinjar): add proper documentation.
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

// DeleteGateway TODO(@vinjar): add proper documentation.
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

// GetAllSensors TODO(@vinjar): add proper documentation.
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

// GetAllSensorProfiles TODO(@vinjar): add proper documentation.
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
