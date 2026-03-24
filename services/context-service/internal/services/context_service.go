// Package services contains context-service business logic.
package services

import (
	"context"
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// ContextServiceImpl implements context use cases for context service.
type ContextServiceImpl struct {
	collectionClient domain.CollectionClient
}

// NewContextServiceImpl creates a new ContextServiceImpl instance.
func NewContextServiceImpl(collectionClient domain.CollectionClient) *ContextServiceImpl {
	return &ContextServiceImpl{
		collectionClient: collectionClient,
	}
}

// GetContextData fetches raw measurements from collection-service and returns computed context data.
func (s *ContextServiceImpl) GetContextData(
	ctx context.Context,
	companyID string,
	deviceEUI string,
	contextType string,
	from, to time.Time,
) (domain.ContextData, error) {

	// Fetch raw measurements from collection-service
	_, err := s.collectionClient.GetMeasurements(ctx, deviceEUI, from, to)
	if err != nil {
		return domain.ContextData{}, err
	}

	// TODO: implement aggregation logic (for example AVG)

	// Return the result
	return domain.ContextData{
		CompanyID:   companyID,
		DeviceEUI:   deviceEUI,
		ContextType: contextType,
		PeriodStart: from,
		PeriodEnd:   to,
		// todo: value: computed result goes here
	}, nil
}
