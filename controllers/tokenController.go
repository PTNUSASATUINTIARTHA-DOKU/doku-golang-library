package controllers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
	tokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	notificationTokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/token"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

type TokenControllerInterface interface {
	VerifyClientKey(privateKey, clientID string) error
	GetTokenB2B(privateKey string, clientId string, isProduction bool) tokenModels.TokenB2BResponseDTO
	GetTokenB2B2C(authCode string, privateKey string, clientId string, isProduction bool) (tokenModels.TokenB2B2CResponseDTO, error)
	IsTokenInvalid(tokenB2B string, tokenExpiresIn int, tokenGeneratedTimestamp string) bool
	ValidateTokenB2B(requestTokenB2B string, publicKey string) (bool, error)
	ValidateSignature(request *http.Request, privateKey string, clientId string, publicKeyDOKU string) bool
	GenerateTokenB2B(expiredIn int, issuer string, privateKey string, clientId string) notificationTokenModels.NotificationTokenDTO
	GenerateInvalidSignatureResponse() notificationTokenModels.NotificationTokenDTO
	DoGenerateRequestHeader(privateKey string, clientId string, tokenB2B string) createVaModels.RequestHeaderDTO
}

var TokenServices services.TokenServices
var SnapUtils utils.SnapUtils

type TokenController struct{}

func (tc TokenController) VerifyClientKey(privateKey, clientID string) error {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return errors.New("failed to decode PEM block containing private key")
	}

	var (
		rsaPrivateKey *rsa.PrivateKey
		err           error
	)

	switch block.Type {
	case "PRIVATE KEY":
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return err
		}

		var ok bool
		if rsaPrivateKey, ok = privateKey.(*rsa.PrivateKey); !ok {
			return errors.New("not an RSA private key")
		}
	case "RSA PRIVATE KEY":
		if rsaPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return err
		}
	default:
		return errors.New("unsupported private key type")
	}

	hashed := sha256.Sum256(fmt.Appendf(nil, "%s|%s", clientID, TokenServices.GenerateTimestamp()))
	if _, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:]); err != nil {
		return err
	}

	return nil
}

func (tc TokenController) GetTokenB2B(privateKey string, clientId string, isProduction bool) tokenModels.TokenB2BResponseDTO {
	var xtimestamp = TokenServices.GenerateTimestamp()
	var signature, _ = TokenServices.CreateSignature(privateKey, clientId, xtimestamp)
	var createTokenB2BRequestDTO tokenModels.TokenB2BRequestDTO = TokenServices.CreateTokenB2BRequestDTO(signature, xtimestamp, clientId)
	return TokenServices.CreateTokenB2B(createTokenB2BRequestDTO, isProduction)
}

func (tc TokenController) GetTokenB2B2C(authCode string, privateKey string, clientId string, isProduction bool) (tokenModels.TokenB2B2CResponseDTO, error) {
	timestamp := tokenServices.GenerateTimestamp()
	signature, err := tokenServices.CreateSignature(privateKey, clientId, timestamp)
	if err != nil {
		return tokenModels.TokenB2B2CResponseDTO{}, err
	}
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
	return snapUtils.GenerateRequestHeaderDto("", signature, xTimestamp, clientId, externalId, "", "", tokenB2B, "")
}
