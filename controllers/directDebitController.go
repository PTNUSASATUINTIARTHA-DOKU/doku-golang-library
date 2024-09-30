package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	accountUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountunbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	cardRegistrationModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistration"
	registrationCardUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistrationunbinding"
	jumpAppModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/jumpapp"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

type DirectDebitInterface interface {
	DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) accountBindingModels.AccountBindingResponseDto
	DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) balanceInquiryModels.BalanceInquiryResponseDto
	DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B2C string, tokenB2B string, isProduction bool) paymentModels.PaymentResponseDTO
	DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) accountUnbindingModels.AccountUnbindingResponseDTO
	DoPaymentJumpApp(paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) jumpAppModels.PaymentJumpAppResponseDTO
	DoCardRegistration(cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, secretKey string, clientId string, channelId string, tokenB2B string, isProduction bool) cardRegistrationModels.CardRegistrationResponseDTO
	DoRefund(refundRequestDTO refundModels.RefundRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) refundModels.RefundResponseDTO
	DoCardRegistrationUnbinding(cardRegistrationUnbindingRequestDTO registrationCardUnbindingModels.CardRegistrationUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO
}

var directDebitService services.DirectDebitService

type DirectDebitController struct{}

func (dd *DirectDebitController) DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) accountBindingModels.AccountBindingResponseDto {
	endPointUrl := commons.DIRECT_DEBIT_ACCOUNT_BINDING
	minifiedRequestBody, err := json.Marshal(accountBindingRequest)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, deviceId, ipAddress, tokenB2B, "")
	return directDebitService.DoAccountBindingProcess(requestHeader, accountBindingRequest, isProduction)
}

func (dd *DirectDebitController) DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) balanceInquiryModels.BalanceInquiryResponseDto {
	url := commons.DIRECT_DEBIT_BALANCE_INQUIRY_URL
	minifiedRequestBody, err := json.Marshal(balanceInquiryRequestDto)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoBalanceInquiryProcess(requestHeader, balanceInquiryRequestDto, isProduction)
}

func (dd *DirectDebitController) DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B2C string, tokenB2B string, isProduction bool) paymentModels.PaymentResponseDTO {
	url := commons.DIRECT_DEBIT_PAYMENT
	minifiedRequestBody, err := json.Marshal(paymentRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("SDK", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoPaymentProcess(requestHeader, paymentRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) accountUnbindingModels.AccountUnbindingResponseDTO {
	url := commons.DIRECT_DEBIT_ACCOUNT_UNBINDING
	minifiedRequestBody, err := json.Marshal(accountUnbindingRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("SDK", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, "")
	return directDebitService.DoAccountUnbindingProcess(requestHeader, accountUnbindingRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoPaymentJumpApp(paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) jumpAppModels.PaymentJumpAppResponseDTO {
	url := commons.DIRECT_DEBIT_PAYMENT
	minifiedRequestBody, err := json.Marshal(paymentJumpAppRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("H2H", signature, timestamp, clientId, externalId, deviceId, ipAddress, tokenB2B, "")
	return directDebitService.DoPaymentJumpAppProcess(requestHeader, paymentJumpAppRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoCardRegistration(cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, secretKey string, clientId string, channelId string, tokenB2B string, isProduction bool) cardRegistrationModels.CardRegistrationResponseDTO {
	url := commons.DIRECT_DEBIT_CARD_REGISTRATION
	minifiedRequestBody, err := json.Marshal(cardRegistrationRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	externalId := snapUtils.GenerateExternalId()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	requestHeader := snapUtils.GenerateRequestHeaderDto(channelId, signature, timestamp, clientId, externalId, "", "", tokenB2B, "")
	return directDebitService.DoCardRegistrationProcess(requestHeader, cardRegistrationRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoRefund(refundRequestDTO refundModels.RefundRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) refundModels.RefundResponseDTO {
	url := commons.DIRECT_DEBIT_REFUND
	minifiedRequestBody, err := json.Marshal(refundRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	externalId := snapUtils.GenerateExternalId()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	requestHeader := snapUtils.GenerateRequestHeaderDto("SDK", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoRefundProcess(requestHeader, refundRequestDTO, isProduction)

}

func (dd *DirectDebitController) DoCardRegistrationUnbinding(cardRegistrationUnbindingRequestDTO registrationCardUnbindingModels.CardRegistrationUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO {
	url := commons.DIRECT_DEBIT_CARD_UNBINDING
	minifiedRequestBody, err := json.Marshal(cardRegistrationUnbindingRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, "")
	return directDebitService.DoCardRegistrationUnbindingProcess(requestHeader, cardRegistrationUnbindingRequestDTO, isProduction)
}
