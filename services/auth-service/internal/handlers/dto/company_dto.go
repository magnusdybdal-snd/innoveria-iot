// Package dto converts dto to domain and vice versa
package dto

import "innoveria-iot/auth-service/internal/domain"

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
