package models

type TokenB2B2CResponseDTO struct {
	ResponseCode           string      `json:"responseCode"`
	ResponseMessage        string      `json:"responseMessage"`
	AccessToken            string      `json:"accessToken"`
	TokenType              string      `json:"tokenType"`
	AccessTokenExpiryTime  string      `json:"accessTokenExpiryTime"`
	RefreshToken           string      `json:"refreshToken"`
	RefreshTokenExpiryTime string      `json:"refreshTokenExpiryTime"`
	AdditionalInfo         interface{} `json:"additionalInfo,omitempty"`
}
