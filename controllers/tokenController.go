package controllers

import (
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

var TokenServices services.TokenServices

type TokenController struct{}

func (tc TokenController) GetTokenB2B(privateKey string, clientId string, isProduction bool) models.TokenB2BResponseDTO {
	var xtimestamp = TokenServices.GenerateTimestamp()
	var signature, _ = TokenServices.CreateSignature(privateKey, clientId, xtimestamp)
	var createTokenB2BRequestDTO models.TokenB2BRequestDTO = TokenServices.CreateTokenB2BRequestDTO(signature, xtimestamp, clientId)
	return TokenServices.CreateTokenB2B(createTokenB2BRequestDTO, isProduction)
}

func (tc TokenController) IsTokenInvalid(tokenB2B string, tokenExpiresIn int, tokenGeneratedTimestamp string) bool {
	if TokenServices.IsTokenEmpty(tokenB2B) {
		return true
	} else {
		if TokenServices.IsTokenExpired(tokenExpiresIn, tokenGeneratedTimestamp) {
			return true
		} else {
			return false
		}
	}
}
