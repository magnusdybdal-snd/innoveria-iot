// Package dto defines data transfer objects for the onboarding service HTTP layer.
package dto

import "innoveria-iot/onboarding-service/internal/domain"

// CreateCompanyRequest is the request body for POST /onboarding/company.
type CreateCompanyRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// MapToDomain maps a CreateCompanyRequest to a domain.Company.
func MapToDomain(req CreateCompanyRequest) domain.Company {
	return domain.Company{
		Name:    req.Name,
		Address: req.Address,
	}
}
