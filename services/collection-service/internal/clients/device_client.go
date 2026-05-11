// Package clients provides HTTP clients for communicating with internal service endpoints.
package clients

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"innoveria-iot/collection-service/internal/clients/dto"
	"innoveria-iot/collection-service/internal/clients/mappers"
	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// DeviceClient is an HTTP client for the device service internal port.
type DeviceClient struct {
	internalBaseURL string
	client          *httpclient.Client
}

// NewDeviceClient creates a DeviceClient targeting the unauthenticated internal port (9090).
func NewDeviceClient(internalBaseURL string) *DeviceClient {
	return &DeviceClient{
		internalBaseURL: internalBaseURL,
		client:          httpclient.New(),
	}
}

// GetSensorMetrics fetches the payload-key → measurement type/unit mappings for a sensor.
func (c *DeviceClient) GetSensorMetrics(ctx context.Context, deviceEUI string) ([]domain.SensorMetric, error) {
	url := fmt.Sprintf("%s/api/v1/device/sensors/%s/metrics", c.internalBaseURL, url.PathEscape(deviceEUI))

	resp, err := httpclient.DoRequest[dto.SensorMetricListResponse](c.client, ctx, url, http.MethodGet, nil, nil)
	if err != nil {
		var httpErr *httpclient.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get sensor metrics for %s: %w", deviceEUI, err)
	}

	// TotalCount is not used; this assumes the device service returns all metrics in one page.
	// If the device service ever introduces pagination, this must be updated to fetch all pages.
	metrics := make([]domain.SensorMetric, len(resp.Metrics))
	for i, m := range resp.Metrics {
		metrics[i] = mappers.ToSensorMetric(m)
	}
	return metrics, nil
}
