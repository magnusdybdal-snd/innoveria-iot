// Package dto converts dto to domain and vice versa
package dto

import (
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// CompanyResponse is response payload for response
type CompanyResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CompanyListResponse wraps a slice of CompanyResponse with a total count
type CompanyListResponse struct {
	TotalCount int               `json:"total_count"`
	Companies  []CompanyResponse `json:"companies"`
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
		ID:        from.ID,
		Name:      from.Name,
		Address:   from.Address,
		CreatedAt: from.CreatedAt,
		UpdatedAt: from.UpdatedAt,
	}
}

// MapCompanyListFromDomain maps domain companies to a company list response DTO.
func MapCompanyListFromDomain(from []domain.Company) CompanyListResponse {
	companies := make([]CompanyResponse, 0, len(from))
	for _, company := range from {
		companies = append(companies, MapCompanyFromDomain(company))
	}

	return CompanyListResponse{
		TotalCount: len(companies),
		Companies:  companies,
	}
}
