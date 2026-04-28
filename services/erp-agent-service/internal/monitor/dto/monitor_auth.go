// Package dto contains  models for the Monitor erp REST API.
// The rest of monitor erp models can be found in pkg/
package dto

// AuthBody is the authentication body used by monitor erp
type AuthBody struct {
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	ForceRelogin bool   `json:"ForceRelogin"`
}
