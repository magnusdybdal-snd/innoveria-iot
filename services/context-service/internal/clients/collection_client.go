// Package clients accesses other clients
package clients

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/clients/mappers"
	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// CollectionClient is an HTTP client for the collection service.
type CollectionClient struct {
	baseURL string
	client  *httpclient.Client
}

// NewCollectionClient creates a new CollectionClient targeting the given base URL.
func NewCollectionClient(baseURL string) *CollectionClient {
	return &CollectionClient{
		baseURL: baseURL,
		client:  httpclient.New(),
	}
}

// GetMeasurements fetches sensor measurements for a device over a time range.
func (c *CollectionClient) GetMeasurements(ctx context.Context, companyID, userID, role, deviceEUI string, from, to time.Time) ([]domain.MeasurementReading, error) {
	url := fmt.Sprintf("%s/api/v1/collection/measurements?device_eui=%s&from=%s&to=%s",
		c.baseURL, deviceEUI, from.Format(time.RFC3339), to.Format(time.RFC3339))

	resp, err := httpclient.DoRequest[[]dto.MeasurementResponse](
		c.client, ctx, url, http.MethodGet, nil, map[string]string{
			"X-Auth-Company-Id": companyID,
			"X-Auth-User-Id":    userID,
			"X-Auth-Role":       role,
		},
	)
	if err != nil {
		return nil, err
	}

	readings := make([]domain.MeasurementReading, len(resp))
	for idx, measurement := range resp {
		readings[idx] = mappers.ToMeasurementReading(measurement)
	}
	return readings, nil
}
