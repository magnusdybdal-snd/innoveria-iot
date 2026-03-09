package clients

import (
	"context"
	"fmt"

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
func (c *DeviceClient) CreateCompanyConfig(_ context.Context, _ string, _ string) (string, error) {
	// TODO: implement
	return "", fmt.Errorf("not implemented")
}

// DeleteCompanyConfig calls the device service to delete the company's Chirpstack config (compensating transaction).
func (c *DeviceClient) DeleteCompanyConfig(_ context.Context, _ string) error {
	// TODO: implement
	return fmt.Errorf("not implemented")
}
