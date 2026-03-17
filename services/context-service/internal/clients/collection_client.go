// Package clients accesses other clients
package clients

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"innoveria-iot/context-service/internal/clients/dto"
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
func (c *CollectionClient) GetMeasurements(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.MeasurementReading, error) {
	url := fmt.Sprintf("%s/api/v1/collection/measurements?device_eui=%s&from=%s&to=%s",
		c.baseURL, deviceEUI, from.Format(time.RFC3339), to.Format(time.RFC3339))

	resp, err := httpclient.DoRequest[[]dto.MeasurementResponse](
		c.client, ctx, url, http.MethodGet, nil, nil,
	)
	if err != nil {
		return nil, err
	}

	readings := make([]domain.MeasurementReading, len(resp))
	for i, m := range resp {
		readings[i] = domain.MeasurementReading{
			DeviceEUI: m.DeviceEUI,
			Timestamp: m.Timestamp,
			Payload:   m.Payload,
			CompanyID: m.CompanyID,
		}
	}
	return readings, nil
}
