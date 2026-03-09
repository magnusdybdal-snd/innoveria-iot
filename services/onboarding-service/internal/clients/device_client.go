package clients

import (
	"context"
	"fmt"
	"net/http"
)

// DeviceClient is an HTTP client for the device service.
type DeviceClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewDeviceClient creates a new DeviceClient targeting the given base URL.
func NewDeviceClient(baseURL string) *DeviceClient {
	return &DeviceClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// CreateCompanyConfig calls the device service to create a Chirpstack tenant and application
// for the given company. Returns the Chirpstack tenantID.
func (c *DeviceClient) CreateCompanyConfig(_ context.Context, _ string) (string, error) {
	// TODO: implement
	return "", fmt.Errorf("not implemented")
}

// DeleteCompanyConfig calls the device service to delete the company's Chirpstack config (compensating transaction).
func (c *DeviceClient) DeleteCompanyConfig(_ context.Context, _ string) error {
	// TODO: implement
	return fmt.Errorf("not implemented")
}
