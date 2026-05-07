package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// ExportServiceImpl implements domain.ExportService.
type ExportServiceImpl struct {
	measurementSvc domain.MeasurementService
	deviceClient   domain.DeviceClient
}

// NewExportService creates a new ExportServiceImpl.
func NewExportService(measurementSvc domain.MeasurementService, deviceClient domain.DeviceClient) *ExportServiceImpl {
	return &ExportServiceImpl{
		measurementSvc: measurementSvc,
		deviceClient:   deviceClient,
	}
}

// GetExportData fetches measurements and sensor metric mappings, then returns structured
// export data ready for CSV rendering. Configured metric keys appear first (with typed
// headers), followed by any unconfigured payload keys found in the measurements.
func (s *ExportServiceImpl) GetExportData(ctx context.Context, companyID, deviceEUI string, from, to time.Time) (domain.ExportData, error) {
	metrics, err := s.deviceClient.GetSensorMetrics(ctx, deviceEUI)
	if err != nil {
		var httpErr *httpclient.HTTPError
		// 404 means no metrics found, we still want to proceed just using payload keys
		if !errors.As(err, &httpErr) || httpErr.StatusCode != http.StatusNotFound {
			return domain.ExportData{}, fmt.Errorf("get sensor metrics: %w", err)
		}
	}

	measurements, err := s.measurementSvc.GetByTimeRange(ctx, companyID, deviceEUI, from, to)
	if err != nil {
		return domain.ExportData{}, fmt.Errorf("get measurements: %w", err)
	}

	metricHeader := make(map[string]string, len(metrics))
	for _, m := range metrics {
		header := m.MeasurementType
		if m.Unit != nil && *m.Unit != "" {
			header += " (" + *m.Unit + ")"
		}
		metricHeader[m.PayloadKey] = header
	}

	seen := make(map[string]bool, len(metrics))
	columns := make([]domain.ExportColumn, 0, len(metrics))
	for _, m := range metrics {
		seen[m.PayloadKey] = true
		columns = append(columns, domain.ExportColumn{PayloadKey: m.PayloadKey, Header: metricHeader[m.PayloadKey]})
	}
	for _, meas := range measurements {
		for key := range meas.Payload {
			if !seen[key] {
				seen[key] = true
				columns = append(columns, domain.ExportColumn{PayloadKey: key, Header: key})
			}
		}
	}

	return domain.ExportData{Columns: columns, Measurements: measurements}, nil
}
