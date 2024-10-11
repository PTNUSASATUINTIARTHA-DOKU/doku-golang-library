package doku

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/controllers"
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	accountUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountunbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	cardRegistrationModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistration"
	cardRegistrationUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistrationunbinding"
	checkStatusModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/checkstatus"
	jumpAppModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/jumpapp"
	notifDirectDebitModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/notification"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"
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
	tokenB2B2CExpiresIn          int
	tokenB2B2CGeneratedTimestamp string
}

func (snap *Snap) GetTokenB2B() tokenVaModels.TokenB2BResponseDTO {
	tokenB2BResponseDTO := TokenController.GetTokenB2B(snap.PrivateKey, snap.ClientId, snap.IsProduction)
	snap.SetTokenB2B(tokenB2BResponseDTO)
	return tokenB2BResponseDTO
}

func (snap *Snap) GetTokenB2B2C(authCode string) (tokenVaModels.TokenB2B2CResponseDTO, error) {
	tokenB2B2CResponseDTO, err := TokenController.GetTokenB2B2C(authCode, snap.PrivateKey, snap.ClientId, snap.IsProduction)
	if err != nil {
		return tokenVaModels.TokenB2B2CResponseDTO{
			ResponseCode:    "5007400",
			ResponseMessage: err.Error(),
		}, err
	}
	snap.SetTokenB2B2C(tokenB2B2CResponseDTO)
	return tokenB2B2CResponseDTO, nil
}

func (snap *Snap) SetTokenB2B(tokenB2BResponseDTO tokenVaModels.TokenB2BResponseDTO) {
	snap.tokenB2B = tokenB2BResponseDTO.AccessToken
	snap.tokenExpiresIn = tokenB2BResponseDTO.ExpiresIn - 10
	snap.tokenGeneratedTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
}

func (snap *Snap) SetTokenB2B2C(tokenB2B2CResponseDTO tokenVaModels.TokenB2B2CResponseDTO) {
	snap.tokenB2B2C = tokenB2B2CResponseDTO.AccessToken
	snap.tokenB2B2CExpiresIn = 890
	snap.tokenB2B2CGeneratedTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
}

func (snap *Snap) CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto) createVaModels.CreateVaResponseDto {

	if isSimulator, response := createVaRequestDto.ValidateSimulatorASPI(); isSimulator && !snap.IsProduction {
		resp, _ := json.Marshal(response)
		log.Println("RESPONSE: ", string(resp))
		return response
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
		snap.SecretKey,
		snap.ClientId,
		snap.tokenB2B,
		snap.IsProduction,
	)
	return createVaResponse
}

func (snap *Snap) UpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO) updateVaModels.UpdateVaResponseDTO {

	if isSimulator, response := updateVaRequestDTO.ValidateSimulatorASPI(); isSimulator && !snap.IsProduction {
		resp, _ := json.Marshal(response)
		log.Println("RESPONSE: ", string(resp))
		return response
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
	if isSimulator, response := checkStatusVaRequestDto.ValidateSimulatorASPI(); isSimulator && !snap.IsProduction {
		resp, _ := json.Marshal(response)
		log.Println("RESPONSE: ", string(resp))
		return response
	}
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

	if isSimulator, response := deleteVaRequestDto.ValidateSimulatorASPI(); isSimulator && !snap.IsProduction {
		resp, _ := json.Marshal(response)
		log.Println("RESPONSE: ", string(resp))
		return response
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

func (snap *Snap) ValidateTokenB2B(requestTokenB2B string) (bool, error) {
	return TokenController.ValidateTokenB2B(requestTokenB2B, snap.PublicKey)
}

func (snap *Snap) validateSignature(request *http.Request, publicKeyDOKU string) bool {
	return TokenController.ValidateSignature(request, snap.PrivateKey, snap.ClientId, publicKeyDOKU)
}

func (snap *Snap) ValidateSignatureAndGenerateToken(request *http.Request, publicKeyDOKU string) notificationTokenModels.NotificationTokenDTO {
	var isSignatureValid = snap.validateSignature(request, publicKeyDOKU)
	return snap.generateTokenB2B(isSignatureValid)
}

func (snap *Snap) GenerateNotificationResponse(isTokenValid bool, paymentNotificationRequestBodyDTO notificationPaymentModels.PaymentNotificationRequestBodyDTO) (notificationPaymentModels.PaymentNotificationResponseBodyDTO, error) {
	if isTokenValid {
		return NotificationController.GenerateNotificationResponse(paymentNotificationRequestBodyDTO), nil
	} else {
		return NotificationController.GenerateInvalidTokenResponse(paymentNotificationRequestBodyDTO), fmt.Errorf("invalid token")
	}
}

func (snap *Snap) ValidateTokenAndGenerateNotificationResponse(requestTokenB2B string, paymentNotificationRequestBodyDTO notificationPaymentModels.PaymentNotificationRequestBodyDTO) (notificationPaymentModels.PaymentNotificationResponseBodyDTO, error) {
	isTokenValid, err := snap.ValidateTokenB2B(requestTokenB2B)
	if err != nil {
		return notificationPaymentModels.PaymentNotificationResponseBodyDTO{}, fmt.Errorf("token validation failed: %w", err)
	}

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

func (snap *Snap) DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, deviceId string, ipAddress string) (accountBindingModels.AccountBindingResponseDTO, error) {
	err := accountBindingRequest.ValidateAccountBindingRequest()
	if err != nil {
		return accountBindingModels.AccountBindingResponseDTO{
			ResponseCode:    "500700",
			ResponseMessage: err.Error(),
		}, err
	}
	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenInvalid {
		snap.GetTokenB2B()
	}

	responseAccountBinding, err := DirectDebitController.DoAccountBinding(accountBindingRequest, snap.SecretKey, snap.ClientId, deviceId, ipAddress, snap.tokenB2B, snap.IsProduction)
	if err != nil {
		return accountBindingModels.AccountBindingResponseDTO{
			ResponseCode:    "5000700",
			ResponseMessage: err.Error(),
		}, err
	}
	return responseAccountBinding, nil
}

func (snap *Snap) DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, deviceId string, ipAddress string, authCode string) (balanceInquiryModels.BalanceInquiryResponseDto, error) {
	err := balanceInquiryRequestDto.ValidateBalanceInquiryRequest()
	if err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{
			ResponseCode:    "5001100",
			ResponseMessage: err.Error(),
		}, err
	}
	isTokenB2BInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenB2BInvalid {
		snap.GetTokenB2B()
	}

	isTokenB2B2CInvalid := TokenController.IsTokenInvalid(snap.tokenB2B2C, snap.tokenB2B2CExpiresIn, snap.tokenB2B2CGeneratedTimestamp)

	if isTokenB2B2CInvalid {
		snap.GetTokenB2B2C(authCode)
	}

	responseBalanceInquiry, err := DirectDebitController.DoBalanceInquiry(balanceInquiryRequestDto, snap.SecretKey, snap.ClientId, ipAddress, snap.tokenB2B, snap.tokenB2B2C, snap.IsProduction)
	if err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{
			ResponseCode:    "5001100",
			ResponseMessage: err.Error(),
		}, err
	}
	return responseBalanceInquiry, nil
}

func (snap *Snap) DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, ipAddress string, authCode string) (paymentModels.PaymentResponseDTO, error) {
	if err := paymentRequestDTO.ValidatePaymentRequest(); err != nil {
		return paymentModels.PaymentResponseDTO{
			ResponseCode:    "5005400",
			ResponseMessage: err.Error(),
		}, err
	}

	isTokenB2BInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenB2BInvalid {
		snap.GetTokenB2B()
	}

	isTokenB2B2CInvalid := TokenController.IsTokenInvalid(snap.tokenB2B2C, snap.tokenB2B2CExpiresIn, snap.tokenB2B2CGeneratedTimestamp)

	if isTokenB2B2CInvalid {
		responseB2B2C, err := snap.GetTokenB2B2C(authCode)
		if err != nil || responseB2B2C.ResponseCode[:3] == "401" {
			return paymentModels.PaymentResponseDTO{
				ResponseCode:    responseB2B2C.ResponseCode,
				ResponseMessage: responseB2B2C.ResponseMessage,
			}, err
		}
	}

	responsePayment, err := DirectDebitController.DoPayment(paymentRequestDTO, snap.SecretKey, snap.ClientId, ipAddress, snap.tokenB2B2C, snap.tokenB2B, snap.IsProduction)
	if err != nil {
		return paymentModels.PaymentResponseDTO{
			ResponseCode:    "5005400",
			ResponseMessage: err.Error(),
		}, err
	}
	return responsePayment, nil
}

func (snap *Snap) DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, ipAddress string) (accountUnbindingModels.AccountUnbindingResponseDTO, error) {
	if err := accountUnbindingRequestDTO.ValidateAccountUnbindingRequest(); err != nil {
		return accountUnbindingModels.AccountUnbindingResponseDTO{
			ResponseCode:    "5000500",
			ResponseMessage: err.Error(),
		}, err
	}

	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenInvalid {
		snap.GetTokenB2B()
	}

	responseAccountUnbinding, err := DirectDebitController.DoAccountUnbinding(accountUnbindingRequestDTO, snap.SecretKey, snap.ClientId, ipAddress, snap.tokenB2B, snap.IsProduction)
	if err != nil {
		return accountUnbindingModels.AccountUnbindingResponseDTO{
			ResponseCode:    "5000500",
			ResponseMessage: err.Error(),
		}, err
	}

	return responseAccountUnbinding, nil
}

func (snap *Snap) DoPaymentJumpApp(paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, deviceId string, ipAddress string) (jumpAppModels.PaymentJumpAppResponseDTO, error) {

	if err := paymentJumpAppRequestDTO.ValidatePaymentJumpAppRequest(); err != nil {
		return jumpAppModels.PaymentJumpAppResponseDTO{
			ResponseCode:    "5005400",
			ResponseMessage: err.Error(),
		}, err
	}

	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenInvalid {
		snap.GetTokenB2B()
	}

	responsePaymentJumpApp, err := DirectDebitController.DoPaymentJumpApp(paymentJumpAppRequestDTO, snap.SecretKey, snap.ClientId, deviceId, ipAddress, snap.tokenB2B, snap.IsProduction)
	if err != nil {
		return jumpAppModels.PaymentJumpAppResponseDTO{
			ResponseCode:    "5005400",
			ResponseMessage: err.Error(),
		}, err
	}

	return responsePaymentJumpApp, nil
}

func (snap *Snap) DoCardRegistration(cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, channelId string) (cardRegistrationModels.CardRegistrationResponseDTO, error) {
	if err := cardRegistrationRequestDTO.ValidateCardRegistrationRequest(); err != nil {
		return cardRegistrationModels.CardRegistrationResponseDTO{
			ResponseCode:    "5000700",
			ResponseMessage: err.Error(),
		}, err
	}

	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenInvalid {
		snap.GetTokenB2B()
	}

	responseCardRegistration, err := DirectDebitController.DoCardRegistration(cardRegistrationRequestDTO, snap.SecretKey, snap.ClientId, channelId, snap.tokenB2B, snap.IsProduction)
	if err != nil {
		return cardRegistrationModels.CardRegistrationResponseDTO{
			ResponseCode:    "5000700",
			ResponseMessage: err.Error(),
		}, err
	}

	return responseCardRegistration, nil
}

func (snap *Snap) DoRefund(refundRequestDTO refundModels.RefundRequestDTO, ipAddress string, authCode string, deviceId string) (refundModels.RefundResponseDTO, error) {
	if err := refundRequestDTO.ValidateRefundRequest(); err != nil {
		return refundModels.RefundResponseDTO{
			ResponseCode:    "5000700",
			ResponseMessage: err.Error(),
		}, err
	}
	isTokenB2BInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenB2BInvalid {
		snap.GetTokenB2B()
	}

	isTokenB2B2CInvalid := TokenController.IsTokenInvalid(snap.tokenB2B2C, snap.tokenB2B2CExpiresIn, snap.tokenB2B2CGeneratedTimestamp)

	if isTokenB2B2CInvalid {
		snap.GetTokenB2B2C(authCode)
	}

	responseRefund, err := DirectDebitController.DoRefund(refundRequestDTO, snap.SecretKey, snap.ClientId, ipAddress, snap.tokenB2B, snap.tokenB2B2C, deviceId, snap.IsProduction)
	if err != nil {
		return refundModels.RefundResponseDTO{
			ResponseCode:    "5000700",
			ResponseMessage: err.Error(),
		}, err
	}
	return responseRefund, nil
}

func (snap *Snap) DoCheckStatus(checkStatusRequestDTO checkStatusModels.CheckStatusRequestDTO) (checkStatusModels.CheckStatusResponseDTO, error) {
	err := checkStatusRequestDTO.ValidateCheckStatusRequest()
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{
			ResponseCode:    "5002600",
			ResponseMessage: err.Error(),
		}, err
	}

	isTokenB2BInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenB2BInvalid {
		snap.GetTokenB2B()
	}

	responseCheckStatus, err := DirectDebitController.DoCheckStatus(checkStatusRequestDTO, snap.SecretKey, snap.ClientId, snap.tokenB2B, snap.IsProduction)
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{
			ResponseCode:    "5002600",
			ResponseMessage: err.Error(),
		}, err
	}

	return responseCheckStatus, nil
}

func (snap *Snap) DoCardRegistrationUnbinding(cardRegistrationUnbindingRequestDTO cardRegistrationUnbindingModels.CardRegistrationUnbindingRequestDTO, ipAddress string) (cardRegistrationUnbindingModels.CardRegistrationUnbindingResponseDTO, error) {
	if err := cardRegistrationUnbindingRequestDTO.ValidateCardRegistrationUnbindingRequest(); err != nil {
		return cardRegistrationUnbindingModels.CardRegistrationUnbindingResponseDTO{
			ResponseCode:    "5000500",
			ResponseMessage: err.Error(),
		}, err
	}

	isTokenInvalid := TokenController.IsTokenInvalid(snap.tokenB2B, snap.tokenExpiresIn, snap.tokenGeneratedTimestamp)

	if isTokenInvalid {
		snap.GetTokenB2B()
	}

	responseCardRegisrtationUnbinding, err := DirectDebitController.DoCardRegistrationUnbinding(cardRegistrationUnbindingRequestDTO, snap.SecretKey, snap.ClientId, ipAddress, snap.tokenB2B, snap.IsProduction)
	if err != nil {
		return cardRegistrationUnbindingModels.CardRegistrationUnbindingResponseDTO{
			ResponseCode:    "5000500",
			ResponseMessage: err.Error(),
		}, err
	}

	return responseCardRegisrtationUnbinding, err
}

func (snap *Snap) DirectDebitPaymentNotification(requestTokenB2B2C string) (notifDirectDebitModels.NotificationPaymentDirectDebitResponseDTO, error) {
	requestTokenB2B2C = strings.TrimPrefix(requestTokenB2B2C, "Bearer ")
	isTokenB2B2CValid, errB2BC := snap.ValidateTokenB2B(requestTokenB2B2C)
	if errB2BC != nil {
		return notifDirectDebitModels.NotificationPaymentDirectDebitResponseDTO{
			ResponseCode:    "5007400",
			ResponseMessage: errB2BC.Error(),
		}, errB2BC
	}
	return snap.GenerateDirectDebitNotificationResponse(isTokenB2B2CValid), nil
}

func (snap *Snap) GenerateDirectDebitNotificationResponse(isTokenB2B2CValid bool) notifDirectDebitModels.NotificationPaymentDirectDebitResponseDTO {
	if isTokenB2B2CValid {
		return NotificationController.GenerateDirectDebitNotificationResponse()
	} else {
		return NotificationController.GenerateDirectDebitInvalidTokenResponse()
	}
}
