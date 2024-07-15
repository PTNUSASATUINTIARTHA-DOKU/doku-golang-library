package controllers

import (
	tokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

var TokenServices services.TokenServices

type TokenController struct{}

func (tc TokenController) GetTokenB2B(privateKey string, clientId string, isProduction bool) tokenModels.TokenB2BResponseDTO {
	var xtimestamp = TokenServices.GenerateTimestamp()
	var signature, _ = TokenServices.CreateSignature(privateKey, clientId, xtimestamp)
	var createTokenB2BRequestDTO tokenModels.TokenB2BRequestDTO = TokenServices.CreateTokenB2BRequestDTO(signature, xtimestamp, clientId)
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
