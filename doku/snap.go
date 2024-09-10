package doku

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/controllers"
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	tokenVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"
	inquiryVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/inquiry"
	notificationPaymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/payment"
	notificationTokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/token"
	updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"
)

var TokenController controllers.TokenControllerInterface
var VaController controllers.VaControllerInterface
var DirectDebitController controllers.DirectDebitInterface
var NotificationController controllers.NotificationInterface

type Snap struct {
	// ----------------
	PrivateKey   string
	PublicKey    string
	SecretKey    string
	Issuer       string
	ClientId     string
	IsProduction bool
	// ----------------
	tokenB2B                     string
	tokenExpiresIn               int
	tokenGeneratedTimestamp      string
	tokenB2B2C                   string
	tokenB2B2CExpiresIn          string
	tokenB2B2CGeneratedTimestamp string
}

func (snap *Snap) GetTokenB2B() tokenVaModels.TokenB2BResponseDTO {
	tokenB2BResponseDTO := TokenController.GetTokenB2B(snap.PrivateKey, snap.ClientId, snap.IsProduction)
	snap.SetTokenB2B(tokenB2BResponseDTO)
	return tokenB2BResponseDTO
}

func (snap *Snap) GetTokenB2B2C(authCode string) tokenVaModels.TokenB2B2CResponseDTO {
	tokenB2B2CResponseDTO := TokenController.GetTokenB2B2C(authCode, snap.PrivateKey, snap.ClientId, snap.IsProduction)
	snap.SetTokenB2B2C(tokenB2B2CResponseDTO)
	return tokenB2B2CResponseDTO
}

func (snap *Snap) SetTokenB2B(tokenB2BResponseDTO tokenVaModels.TokenB2BResponseDTO) {
	snap.tokenB2B = tokenB2BResponseDTO.AccessToken
	snap.tokenExpiresIn = tokenB2BResponseDTO.ExpiresIn - 10
	snap.tokenGeneratedTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
}

func (snap *Snap) SetTokenB2B2C(tokenB2B2CResponseDTO tokenVaModels.TokenB2B2CResponseDTO) {
	snap.tokenB2B2C = tokenB2B2CResponseDTO.AccessToken
	snap.tokenB2B2CExpiresIn = tokenB2B2CResponseDTO.AccessTokenExpiryTime
	snap.tokenB2B2CGeneratedTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
}

func (snap *Snap) CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto) createVaModels.CreateVaResponseDto {

	if isSimulator, errorResponse := createVaRequestDto.ValidateSimulatorASPI(); isSimulator && !snap.IsProduction {
		resp, _ := json.Marshal(errorResponse)
		log.Println("RESPONSE: ", string(resp))
		return errorResponse
	}

	if err := createVaRequestDto.ValidateVaRequestDto(); err != nil {
		log.Println(err)
	}

	isTokenInvalid := TokenController.IsTokenInvalid(
		snap.tokenB2B,
		snap.tokenExpiresIn,
		snap.tokenGeneratedTimestamp,
	)
	if isTokenInvalid {
		snap.GetTokenB2B()
	}
	createVaResponse := VaController.CreateVa(
		createVaRequestDto,
		snap.PrivateKey,
		snap.ClientId,
		snap.tokenB2B,
		snap.IsProduction,
	)
	return createVaResponse
}

func (snap *Snap) UpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO) updateVaModels.UpdateVaResponseDTO {

	if isSimulator, errorResponse := updateVaRequestDTO.ValidateSimulatorASPI(); isSimulator && !snap.IsProduction {
		resp, _ := json.Marshal(errorResponse)
		log.Println("RESPONSE: ", string(resp))
		return errorResponse
	}

	if err := updateVaRequestDTO.ValidateUpdateVaRequestDTO(); err != nil {
		log.Println(err)
	}

	isTokenInvalid := TokenController.IsTokenInvalid(
		snap.tokenB2B,
		snap.tokenExpiresIn,
		snap.tokenGeneratedTimestamp,
	)
	if isTokenInvalid {
		snap.GetTokenB2B()
	}
	updateVaResponse := VaController.DoUpdateVa(updateVaRequestDTO, snap.ClientId, snap.tokenB2B, snap.SecretKey, snap.IsProduction)

	return updateVaResponse
}

func (snap *Snap) CheckStatusVa(checkStatusVaRequestDto checkVaModels.CheckStatusVARequestDto) checkVaModels.CheckStatusVaResponseDto {

	checkStatusVaRequestDto.ValidateCheckStatusVaRequestDto()
	isTokenInvalid := TokenController.IsTokenInvalid(
		snap.tokenB2B,
		snap.tokenExpiresIn,
		snap.tokenGeneratedTimestamp,
	)
	if isTokenInvalid {
		snap.GetTokenB2B()
	}
	checkStatusVaResponseDTO := VaController.DoCheckStatusVa(checkStatusVaRequestDto, snap.PrivateKey, snap.ClientId, snap.tokenB2B, snap.SecretKey, snap.IsProduction)

	return checkStatusVaResponseDTO
}

func (snap *Snap) DeletePaymentCode(deleteVaRequestDto deleteVaModels.DeleteVaRequestDto) deleteVaModels.DeleteVaResponseDto {

	if isSimulator, errorResponse := deleteVaRequestDto.ValidateSimulatorASPI(); isSimulator && !snap.IsProduction {
		resp, _ := json.Marshal(errorResponse)
		log.Println("RESPONSE: ", string(resp))
		return errorResponse
	}
	deleteVaRequestDto.ValidateDeleteVaRequest()
	isTokenInvalid := TokenController.IsTokenInvalid(
		snap.tokenB2B,
		snap.tokenExpiresIn,
		snap.tokenGeneratedTimestamp,
	)
	if isTokenInvalid {
		snap.GetTokenB2B()
	}
	deleteVaResponseDto := VaController.DoDeletePaymentCode(deleteVaRequestDto, snap.PrivateKey, snap.ClientId, snap.tokenB2B, snap.SecretKey, snap.IsProduction)

	return deleteVaResponseDto
}

func (snap *Snap) generateTokenB2B(isSignatureValid bool) notificationTokenModels.NotificationTokenDTO {
	if isSignatureValid {
		return TokenController.GenerateTokenB2B(snap.tokenExpiresIn, snap.Issuer, snap.PrivateKey, snap.ClientId)
	} else {
		return TokenController.GenerateInvalidSignatureResponse()
	}
}

func (snap *Snap) ValidateTokenB2B(requestTokenB2B string) bool {
	return TokenController.ValidateTokenB2B(requestTokenB2B, snap.PublicKey)
}

func (snap *Snap) validateSignature(request *http.Request) bool {
	return TokenController.ValidateSignature(request, snap.PrivateKey, snap.ClientId)
}

func (snap *Snap) ValidateSignatureAndGenerateToken(request *http.Request) notificationTokenModels.NotificationTokenDTO {
	var isSignatureValid = snap.validateSignature(request)
	return snap.generateTokenB2B(isSignatureValid)
}

func (snap *Snap) GenerateNotificationResponse(isTokenValid bool, paymentNotificationRequestBodyDTO notificationPaymentModels.PaymentNotificationRequestBodyDTO) notificationPaymentModels.PaymentNotificationResponseBodyDTO {
	if isTokenValid {
		return NotificationController.GenerateNotificationResponse(paymentNotificationRequestBodyDTO)
	} else {
		return NotificationController.GenerateInvalidTokenResponse(paymentNotificationRequestBodyDTO)
	}
}

func (snap *Snap) ValidateTokenAndGenerateNotificationResponse(requestTokenB2B string, paymentNotificationRequestBodyDTO notificationPaymentModels.PaymentNotificationRequestBodyDTO) notificationPaymentModels.PaymentNotificationResponseBodyDTO {
	isTokenValid := snap.ValidateTokenB2B(requestTokenB2B)
	return snap.GenerateNotificationResponse(isTokenValid, paymentNotificationRequestBodyDTO)
}

func (snap *Snap) GenerateRequestHeader() createVaModels.RequestHeaderDTO {
	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)
	if isTokenInvalid {
		snap.GetTokenB2B()
	}
	return TokenController.DoGenerateRequestHeader(snap.PrivateKey, snap.ClientId, snap.tokenB2B)
}

func (snap *Snap) DirectInquiryRequestMapping(headerRequest *http.Request, inquiryRequestBodyDto inquiryVaModels.InquiryRequestBodyDTO) (string, error) {
	return VaController.DirectInquiryRequestMapping(headerRequest, inquiryRequestBodyDto)
}

func (snap *Snap) DirectInquiryResponseMapping(xmlData string) (inquiryVaModels.InquiryResponseBodyDTO, error) {
	return VaController.DirectInquiryResponseMapping(xmlData)
}

func (snap *Snap) DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, deviceId string, ipAddress string) accountBindingModels.AccountBindingResponseDto {
	accountBindingRequest.ValidateAccountBindingRequest()
	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenInvalid {
		snap.GetTokenB2B()
	}
	return DirectDebitController.DoAccountBinding(accountBindingRequest, snap.SecretKey, snap.ClientId, deviceId, ipAddress, snap.tokenB2B, snap.IsProduction)
}

func (snap *Snap) DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, deviceId string, ipAddress string) balanceInquiryModels.BalanceInquiryResponseDto {
	balanceInquiryRequestDto.ValidateBalanceInquiryRequest()
	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenInvalid {
		snap.GetTokenB2B()
	}
	return DirectDebitController.DoBalanceInquiry(balanceInquiryRequestDto, snap.SecretKey, snap.ClientId, ipAddress, snap.tokenB2B, snap.tokenB2B2C, snap.IsProduction)
}
