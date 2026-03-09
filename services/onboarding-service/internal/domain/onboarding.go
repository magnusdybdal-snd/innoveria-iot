// Package domain defines the core types and interfaces for the onboarding service.
package domain

import "context"

// Company holds the data needed to onboard a new company.
type Company struct {
	Name    string
	Address string
}

// OnboardingService defines the SAGA orchestration for onboarding new companies.
type OnboardingService interface {
	CreateCompany(ctx context.Context, company Company) error
}

// AuthClient defines the operations this service needs from the auth service.
type AuthClient interface {
	CreateCompany(ctx context.Context, company Company) (companyID string, err error)
	DeleteCompany(ctx context.Context, companyID string) error
}

// DeviceClient defines the operations this service needs from the device service.
// CreateCompanyConfig creates a Chirpstack tenant and application for the company,
// stores the companyID-applicationID mapping internally, and returns the tenantID.
type DeviceClient interface {
	CreateCompanyConfig(ctx context.Context, companyID string) (tenantID string, err error)
	DeleteCompanyConfig(ctx context.Context, companyID string) error
}

// CollectionClient defines the operations this service needs from the collection service.
type CollectionClient interface {
	CreateCompanyConfig(ctx context.Context, companyID string, tenantID string) error
}
