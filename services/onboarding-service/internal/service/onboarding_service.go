// Package service implements the business logic for the onboarding service.
package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/onboarding-service/internal/domain"
)

// OnboardingServiceImpl implements domain.OnboardingService.
type OnboardingServiceImpl struct {
	authClient       domain.AuthClient
	deviceClient     domain.DeviceClient
	collectionClient domain.CollectionClient
}

// NewOnboardingService creates a new OnboardingServiceImpl.
func NewOnboardingService(
	authClient domain.AuthClient,
	deviceClient domain.DeviceClient,
	collectionClient domain.CollectionClient,
) *OnboardingServiceImpl {
	return &OnboardingServiceImpl{
		authClient:       authClient,
		deviceClient:     deviceClient,
		collectionClient: collectionClient,
	}
}

// CreateCompany runs the company onboarding SAGA:
//  1. Create company in auth service → get companyID
//  2. Create Chirpstack tenant+application in device service → get tenantID
//  3. Store companyID-tenantID mapping in collection service
//
// Compensating transactions are run on failure to keep all services in sync.
func (s *OnboardingServiceImpl) CreateCompany(ctx context.Context, company domain.Company) error {
	// Step 1: create company in auth service
	companyID, err := s.authClient.CreateCompany(ctx, company)
	if err != nil {
		return fmt.Errorf("create company: auth service: %w", err)
	}

	// Step 2: create chirpstack config in device service
	tenantID, err := s.deviceClient.CreateCompanyConfig(ctx, companyID)
	if err != nil {
		// Compensate: delete company from auth service
		if compErr := s.authClient.DeleteCompany(ctx, companyID); compErr != nil {
			slog.Error("saga compensation failed: could not delete company from auth service after device service failure",
				"companyID", companyID, "error", compErr)
		}
		return fmt.Errorf("create company: device service: %w", err)
	}

	// Step 3: store companyID-tenantID mapping in collection service
	if err := s.collectionClient.CreateCompanyConfig(ctx, companyID, tenantID); err != nil {
		// Compensate: delete device config and company from auth service
		if compErr := s.deviceClient.DeleteCompanyConfig(ctx, companyID); compErr != nil {
			slog.Error("saga compensation failed: could not delete company config from device service after collection service failure",
				"companyID", companyID, "error", compErr)
		}
		if compErr := s.authClient.DeleteCompany(ctx, companyID); compErr != nil {
			slog.Error("saga compensation failed: could not delete company from auth service after collection service failure",
				"companyID", companyID, "error", compErr)
		}
		return fmt.Errorf("create company: collection service: %w", err)
	}

	slog.Info("successfully onboarded company", "companyID", companyID)
	return nil
}
