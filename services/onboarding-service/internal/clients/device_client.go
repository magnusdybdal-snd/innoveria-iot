package clients

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/onboarding-service/internal/clients/dto"
	"innoveria-iot/pkg/httpclient"
)

// DeviceClient is an HTTP client for the device service.
type DeviceClient struct {
	baseURL string
	client  *httpclient.Client
}

// NewDeviceClient creates a new DeviceClient targeting the given base URL.
func NewDeviceClient(baseURL string) *DeviceClient {
	return &DeviceClient{
		baseURL: baseURL,
		client:  httpclient.New(),
	}
}

// CreateCompanyConfig calls the device service to create a Chirpstack tenant and application
// for the given company. Returns the Chirpstack tenantID.
func (c *DeviceClient) CreateCompanyConfig(ctx context.Context, companyID string, name string) (string, error) {
	resp, err := httpclient.DoRequest[dto.CreateCompanyConfigResponse](
		c.client,
		ctx,
		c.baseURL+"/api/v1/device/company-config",
		http.MethodPost,
		dto.CreateCompanyConfigRequest{CompanyID: companyID, Name: name},
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("device client create company config: %w", err)
	}

	return resp.TenantID, nil
}

// DeleteCompanyConfig calls the device service to delete the company's Chirpstack config (compensating transaction).
func (c *DeviceClient) DeleteCompanyConfig(ctx context.Context, companyID string) error {
	resp, err := httpclient.DoRaw(
		c.client,
		ctx,
		fmt.Sprintf("%s/api/v1/device/company-config/%s", c.baseURL, companyID),
		http.MethodDelete,
		nil,
		nil,
	)
	if err != nil {
		return fmt.Errorf("device client delete company config: %w", err)
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}
