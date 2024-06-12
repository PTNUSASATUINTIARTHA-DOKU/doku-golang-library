package models

type TokenB2BRequestDTO struct {
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
	ClientID  string `json:"clientId"`
	GrantType string `json:"grantType"`
}
