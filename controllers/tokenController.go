package controllers

import (
	"net/http"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
	tokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	notificationTokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/token"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

type TokenControllerInterface interface {
	GetTokenB2B(privateKey string, clientId string, isProduction bool) tokenModels.TokenB2BResponseDTO
	IsTokenInvalid(tokenB2B string, tokenExpiresIn int, tokenGeneratedTimestamp string) bool
	ValidateTokenB2B(requestTokenB2B string, publicKey string) bool
	ValidateSignature(request *http.Request, privateKey string, clientId string) bool
	GenerateTokenB2B(expiredIn int, issuer string, privateKey string, clientId string) notificationTokenModels.NotificationTokenDTO
	GenerateInvalidSignatureResponse() notificationTokenModels.NotificationTokenDTO
	DoGenerateRequestHeader(privateKey string, clientId string, tokenB2B string) createVaModels.RequestHeaderDTO
}

var TokenServices services.TokenServices
var SnapUtils utils.SnapUtils

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

func (tc TokenController) ValidateTokenB2B(requestTokenB2B string, publicKey string) bool {
	return TokenServices.ValidateTokenB2B(requestTokenB2B, publicKey)
}

func (tc TokenController) ValidateSignature(request *http.Request, privateKey string, clientId string) bool {
	timestamp := request.Header.Get("x-timestamp")
	requestSignature := request.Header.Get("x-signature")
	var newSignature, _ = TokenServices.CreateSignature(privateKey, clientId, timestamp)
	return tokenServices.CompareSignature(requestSignature, newSignature)
}

func (tc TokenController) GenerateTokenB2B(expiredIn int, issuer string, privateKey string, clientId string) notificationTokenModels.NotificationTokenDTO {
	var xTimestamp = TokenServices.GenerateTimestamp()
	var token = TokenServices.GenerateToken(int64(expiredIn), issuer, privateKey, clientId)
	return TokenServices.GenerateNotificationTokenDTO(token, xTimestamp, clientId, expiredIn)
}

func (tc TokenController) GenerateInvalidSignatureResponse() notificationTokenModels.NotificationTokenDTO {
	var xTimestamp = TokenServices.GenerateTimestamp()
	return TokenServices.GenerateInvalidSignature(xTimestamp)
}

func (tc TokenController) DoGenerateRequestHeader(privateKey string, clientId string, tokenB2B string) createVaModels.RequestHeaderDTO {
	externalId := SnapUtils.GenerateExternalId()
	xTimestamp := TokenServices.GenerateTimestamp()
	signature, _ := TokenServices.CreateSignature(privateKey, clientId, xTimestamp)
	return snapUtils.GenerateRequestHeaderDto("", signature, xTimestamp, clientId, externalId, tokenB2B)
}
