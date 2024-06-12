package controllers

import (
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

var TokenService services.TokenService

type TokenController struct{}

func (tc TokenController) GetTokenB2B(privateKey string, clientId string, isProduction bool) models.TokenB2BResponseDTO {
	var xtimestamp = TokenService.GenerateTimestamp()
	var signature, _ = TokenService.CreateSignature(privateKey, clientId, xtimestamp)
	var createTokenB2BRequestDTO models.TokenB2BRequestDTO = TokenService.CreateTokenB2BRequestDTO(signature, xtimestamp, clientId)
	return TokenService.CreateTokenB2B(createTokenB2BRequestDTO, isProduction)
}
