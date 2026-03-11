package dto

// CreateTenantMappingRequest is the request body for POST /collection/company-config.
type CreateTenantMappingRequest struct {
	CompanyID string `json:"company_id"`
	TenantID  string `json:"tenant_id"`
}
