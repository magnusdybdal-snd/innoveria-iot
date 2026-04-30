package clients

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/clients/mappers"
	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// DeviceClient is an HTTP client for the device service.
type DeviceClient struct {
	baseURL         string
	internalBaseURL string
	client          *httpclient.Client
}

// NewDeviceClient creates a new DeviceClient. baseURL targets the authenticated
// public port (8080); internalBaseURL targets the unauthenticated internal port (9090).
func NewDeviceClient(baseURL, internalBaseURL string) *DeviceClient {
	return &DeviceClient{
		baseURL:         baseURL,
		internalBaseURL: internalBaseURL,
		client:          httpclient.New(),
	}
}

// GetSensorsByProductionResourceID fetches all sensors assigned to the given
// production resource ID, scoped to the given company.
func (c *DeviceClient) GetSensorsByProductionResourceID(ctx context.Context, companyID, userID, role, productionResourceID string) ([]domain.DeviceSensor, error) {
	url := fmt.Sprintf("%s/api/v1/device/sensors?production_resource_id=%s", c.baseURL, productionResourceID)

	resp, err := httpclient.DoRequest[dto.DeviceSensorListResponse](
		c.client, ctx, url, http.MethodGet, nil, map[string]string{
			"X-Auth-Company-Id": companyID,
			"X-Auth-User-Id":    userID,
			"X-Auth-Role":       role,
		},
	)
	if err != nil {
		var httpErr *httpclient.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	sensors := make([]domain.DeviceSensor, len(resp.Sensors))
	for i, s := range resp.Sensors {
		sensors[i] = mappers.ToDeviceSensor(s)
	}
	return sensors, nil
}

// GetSensorMetrics fetches the payload-key → measurement type/unit mappings
// for the given sensor EUI. Uses the internal port (no auth required).
func (c *DeviceClient) GetSensorMetrics(ctx context.Context, deviceEUI string) ([]domain.SensorMetric, error) {
	url := fmt.Sprintf("%s/api/v1/device/sensors/%s/metrics", c.internalBaseURL, deviceEUI)

	resp, err := httpclient.DoRequest[dto.SensorMetricListResponse](
		c.client, ctx, url, http.MethodGet, nil, nil,
	)
	if err != nil {
		return nil, err
	}

	metrics := make([]domain.SensorMetric, len(resp.Metrics))
	for i, m := range resp.Metrics {
		metrics[i] = mappers.ToSensorMetric(m)
	}
	return metrics, nil
}
