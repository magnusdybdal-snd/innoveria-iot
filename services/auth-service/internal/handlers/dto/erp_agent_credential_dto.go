package dto

import (
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// ERPAgentCredentialDetailsResponse is the response payload for credential reads.
type ERPAgentCredentialDetailsResponse struct {
	CompanyID  string  `json:"company_id"`
	KeyID      string  `json:"key_id"`
	SecretHash string  `json:"secret_hash"`
	CreatedAt  string  `json:"created_at"`
	RotatedAt  *string `json:"rotated_at,omitempty"`
	RevokedAt  *string `json:"revoked_at,omitempty"`
}

// MapERPAgentCredentialFromDomain maps domain credential to response DTO.
func MapERPAgentCredentialFromDomain(from domain.ERPAgentCredential) ERPAgentCredentialDetailsResponse {
	resp := ERPAgentCredentialDetailsResponse{
		CompanyID:  from.CompanyID,
		KeyID:      from.KeyID,
		SecretHash: from.SecretHash,
		CreatedAt:  from.CreatedAt.Format(time.RFC3339),
	}

	if from.RotatedAt != nil {
		rotatedAt := from.RotatedAt.Format(time.RFC3339)
		resp.RotatedAt = &rotatedAt
	}

	if from.RevokedAt != nil {
		revokedAt := from.RevokedAt.Format(time.RFC3339)
		resp.RevokedAt = &revokedAt
	}

	return resp
}
