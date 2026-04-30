// Package dto converts dto to domain and vice versa
package dto

import (
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// CompanyResponse is response payload for response
type CompanyResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CompanyListResponse wraps a slice of CompanyResponse with a total count
type CompanyListResponse struct {
	TotalCount int               `json:"total_count"`
	Companies  []CompanyResponse `json:"companies"`
}

// ERPAgentTokenResponse is response payload when issuing a new ERP agent token.
type ERPAgentTokenResponse struct {
	ERPAgentToken string `json:"erp_agent_token"`
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
		CreatedAt: from.CreatedAt.Format(time.RFC3339),
		UpdatedAt: from.UpdatedAt.Format(time.RFC3339),
	}
}

// MapERPAgentTokenResponse maps a raw token string to response DTO.
func MapERPAgentTokenResponse(token string) ERPAgentTokenResponse {
	return ERPAgentTokenResponse{ERPAgentToken: token}
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
