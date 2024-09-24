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
	GetTokenB2B2C(authCode string, privateKey string, clientId string, isProduction bool) tokenModels.TokenB2B2CResponseDTO
	IsTokenInvalid(tokenB2B string, tokenExpiresIn int, tokenGeneratedTimestamp string) bool
	ValidateTokenB2B(requestTokenB2B string, publicKey string) (bool, error)
	ValidateSignature(request *http.Request, privateKey string, clientId string, publicKeyDOKU string) bool
	GenerateTokenB2B(expiredIn int, issuer string, privateKey string, clientId string) notificationTokenModels.NotificationTokenBodyDTO
	GenerateInvalidSignatureResponse() notificationTokenModels.NotificationTokenBodyDTO
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

func (tc TokenController) GetTokenB2B2C(authCode string, privateKey string, clientId string, isProduction bool) tokenModels.TokenB2B2CResponseDTO {
	timestamp := tokenServices.GenerateTimestamp()
	signature, _ := tokenServices.CreateSignature(privateKey, clientId, timestamp)
	tokenB2B2CRequestDTO := tokenServices.CreateTokenB2B2CRequestDTO(authCode)
	return tokenServices.HitTokenB2B2CApi(tokenB2B2CRequestDTO, timestamp, signature, clientId, isProduction)
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

func (tc TokenController) ValidateTokenB2B(requestTokenB2B string, publicKey string) (bool, error) {
	return TokenServices.ValidateTokenB2B(requestTokenB2B, publicKey)
}

func (tc TokenController) ValidateSignature(request *http.Request, privateKey string, clientId string, publicKeyDOKU string) bool {
	timestamp := request.Header.Get("x-timestamp")
	requestSignature := request.Header.Get("x-signature")
	compareSignature, _ := tokenServices.CompareSignatures(clientId, timestamp, requestSignature, publicKeyDOKU)
	return compareSignature
}

func (tc TokenController) GenerateTokenB2B(expiredIn int, issuer string, privateKey string, clientId string) notificationTokenModels.NotificationTokenBodyDTO {
	var xTimestamp = TokenServices.GenerateTimestamp()
	var token = TokenServices.GenerateToken(int64(expiredIn), issuer, privateKey, clientId)
	return TokenServices.GenerateNotificationTokenDTO(token, xTimestamp, clientId, expiredIn)
}

func (tc TokenController) GenerateInvalidSignatureResponse() notificationTokenModels.NotificationTokenBodyDTO {
	var xTimestamp = TokenServices.GenerateTimestamp()
	return TokenServices.GenerateInvalidSignature(xTimestamp)
}

func (tc TokenController) DoGenerateRequestHeader(privateKey string, clientId string, tokenB2B string) createVaModels.RequestHeaderDTO {
	externalId := SnapUtils.GenerateExternalId()
	xTimestamp := TokenServices.GenerateTimestamp()
	signature, _ := TokenServices.CreateSignature(privateKey, clientId, xTimestamp)
	return snapUtils.GenerateRequestHeaderDto("", signature, xTimestamp, clientId, externalId, "", "", tokenB2B, "")
}
