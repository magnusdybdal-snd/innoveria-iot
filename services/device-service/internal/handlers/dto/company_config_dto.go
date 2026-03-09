package dto

// CreateCompanyConfigRequest is the request body for POST /device/company-config.
type CreateCompanyConfigRequest struct {
	CompanyID string `json:"company_id"`
	Name      string `json:"name"`
}

// CreateCompanyConfigResponse is the response body for POST /device/company-config.
type CreateCompanyConfigResponse struct {
	TenantID string `json:"tenant_id"`
}
