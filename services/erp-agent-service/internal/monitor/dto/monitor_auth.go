// Package dto contains  models for the Monitor erp REST API.
package dto

// AuthBody is the authentication body used by monitor erp
type AuthBody struct {
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	ForceRelogin bool   `json:"ForceRelogin"`
}
