// Package dto converts dto to domain and vice versa
package dto

import (
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// CompanyResponse is response payload for response
type CompanyResponse struct {
	CompanyID string    `json:"company_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateNewCompany is the request payload for creating a company.
type CreateNewCompany struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// MapCreateCompanyToDomain maps a create-company DTO into a domain Company.
func MapCreateCompanyToDomain(from CreateNewCompany) domain.Company {
	return domain.Company{
		Name:    from.Name,
		Address: from.Address,
	}
}

// MapCompanyFromDomain maps a domain Company to a CompanyResponse DTO.
func MapCompanyFromDomain(from domain.Company) CompanyResponse {
	return CompanyResponse{
		CompanyID: from.CompanyID,
		Name:      from.Name,
		Address:   from.Address,
		CreatedAt: from.CreatedAt,
		UpdatedAt: from.UpdatedAt,
	}
}
