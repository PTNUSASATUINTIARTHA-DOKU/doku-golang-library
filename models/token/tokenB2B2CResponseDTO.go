package models

type TokenB2B2CResponseDTO struct {
	ResponseCode           string      `json:"responseCode"`
	ResponseMessage        string      `json:"responseMessage"`
	AccessToken            string      `json:"accessToken,omitempty"`
	TokenType              string      `json:"tokenType,omitempty"`
	AccessTokenExpiryTime  string      `json:"accessTokenExpiryTime,omitempty"`
	RefreshToken           string      `json:"refreshToken,omitempty"`
	RefreshTokenExpiryTime string      `json:"refreshTokenExpiryTime,omitempty"`
	AdditionalInfo         interface{} `json:"additionalInfo,omitempty"`
}
