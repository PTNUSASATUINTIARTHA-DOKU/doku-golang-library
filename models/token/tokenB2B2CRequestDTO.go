package models

type TokenB2B2CRequestDTO struct {
	GrantType      string      `json:"grantType"`
	AuthCode       string      `json:"authCode"`
	RefreshToken   string      `json:"refreshToken"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
}
