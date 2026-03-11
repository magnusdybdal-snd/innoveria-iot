package dto

import "innoveria-iot/auth-service/internal/domain"

// FactoryResponse is the response payload for factory operations.
type FactoryResponse struct {
	Id        string `json:"id"`
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// FactoryListResponse wraps a list of factories with total count.
type FactoryListResponse struct {
	TotalCount int               `json:"total_count"`
	Factories  []FactoryResponse `json:"factories"`
}

// CreateNewFactory is the request payload for creating a factory.
type CreateNewFactory struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}

// MapCreateFactoryToDomain maps a create-factory DTO to domain Factory.
func MapCreateFactoryToDomain(from CreateNewFactory) domain.Factory {
	return domain.Factory{
		CompanyID: from.CompanyID,
		Name:      from.Name,
		Address:   from.Address,
	}
}

// MapFactoryFromDomain maps domain Factory to a FactoryResponse DTO.
func MapFactoryFromDomain(from domain.Factory) FactoryResponse {
	return FactoryResponse{
		Id:        from.ID,
		CompanyID: from.CompanyID,
		Name:      from.Name,
		Address:   from.Address,
		CreatedAt: from.CreatedAt,
		UpdatedAt: from.UpdatedAt,
	}
}

// MapFactoryListFromDomain maps domain factories to list response DTO.
func MapFactoryListFromDomain(from []domain.Factory) FactoryListResponse {
	tot := len(from)
	factories := make([]FactoryResponse, 0, tot)

	for _, factory := range from {
		factories = append(factories, MapFactoryFromDomain(factory))
	}

	return FactoryListResponse{
		TotalCount: tot,
		Factories:  factories,
	}
}
