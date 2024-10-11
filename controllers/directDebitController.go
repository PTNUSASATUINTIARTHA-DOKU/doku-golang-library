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
	checkStatusModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/checkstatus"
	jumpAppModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/jumpapp"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

type DirectDebitInterface interface {
	DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) (accountBindingModels.AccountBindingResponseDTO, error)
	DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) (balanceInquiryModels.BalanceInquiryResponseDto, error)
	DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B2C string, tokenB2B string, isProduction bool) (paymentModels.PaymentResponseDTO, error)
	DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) (accountUnbindingModels.AccountUnbindingResponseDTO, error)
	DoPaymentJumpApp(paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) (jumpAppModels.PaymentJumpAppResponseDTO, error)
	DoCardRegistration(cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, secretKey string, clientId string, channelId string, tokenB2B string, isProduction bool) (cardRegistrationModels.CardRegistrationResponseDTO, error)
	DoRefund(refundRequestDTO refundModels.RefundRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, deviceId string, isProduction bool) (refundModels.RefundResponseDTO, error)
	DoCheckStatus(checkStatusRequestDTO checkStatusModels.CheckStatusRequestDTO, secretKey string, clientId string, tokenB2B string, isProduction bool) (checkStatusModels.CheckStatusResponseDTO, error)
	DoCardRegistrationUnbinding(cardRegistrationUnbindingRequestDTO registrationCardUnbindingModels.CardRegistrationUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) (registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO, error)
}

var directDebitService services.DirectDebitService

type DirectDebitController struct{}

func (dd *DirectDebitController) DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) (accountBindingModels.AccountBindingResponseDTO, error) {
	endPointUrl := commons.DIRECT_DEBIT_ACCOUNT_BINDING
	minifiedRequestBody, err := json.Marshal(accountBindingRequest)
	if err != nil {
		return accountBindingModels.AccountBindingResponseDTO{}, fmt.Errorf("error marshalling response JSON: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, deviceId, ipAddress, tokenB2B, "")
	return directDebitService.DoAccountBindingProcess(requestHeader, accountBindingRequest, isProduction)
}

func (dd *DirectDebitController) DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) (balanceInquiryModels.BalanceInquiryResponseDto, error) {
	url := commons.DIRECT_DEBIT_BALANCE_INQUIRY_URL
	minifiedRequestBody, err := json.Marshal(balanceInquiryRequestDto)
	if err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{}, fmt.Errorf("error marshalling response JSON: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoBalanceInquiryProcess(requestHeader, balanceInquiryRequestDto, isProduction)
}

func (dd *DirectDebitController) DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B2C string, tokenB2B string, isProduction bool) (paymentModels.PaymentResponseDTO, error) {
	url := commons.DIRECT_DEBIT_PAYMENT
	minifiedRequestBody, err := json.Marshal(paymentRequestDTO)
	if err != nil {
		return paymentModels.PaymentResponseDTO{}, fmt.Errorf("error marshalling response JSON: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("DH", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoPaymentProcess(requestHeader, paymentRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) (accountUnbindingModels.AccountUnbindingResponseDTO, error) {
	url := commons.DIRECT_DEBIT_ACCOUNT_UNBINDING
	minifiedRequestBody, err := json.Marshal(accountUnbindingRequestDTO)
	if err != nil {
		return accountUnbindingModels.AccountUnbindingResponseDTO{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, "")
	return directDebitService.DoAccountUnbindingProcess(requestHeader, accountUnbindingRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoPaymentJumpApp(paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) (jumpAppModels.PaymentJumpAppResponseDTO, error) {
	url := commons.DIRECT_DEBIT_PAYMENT
	minifiedRequestBody, err := json.Marshal(paymentJumpAppRequestDTO)
	if err != nil {
		return jumpAppModels.PaymentJumpAppResponseDTO{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("DH", signature, timestamp, clientId, externalId, deviceId, ipAddress, tokenB2B, "")
	return directDebitService.DoPaymentJumpAppProcess(requestHeader, paymentJumpAppRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoCardRegistration(cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, secretKey string, clientId string, channelId string, tokenB2B string, isProduction bool) (cardRegistrationModels.CardRegistrationResponseDTO, error) {
	url := commons.DIRECT_DEBIT_CARD_REGISTRATION
	minifiedRequestBody, err := json.Marshal(cardRegistrationRequestDTO)
	if err != nil {
		return cardRegistrationModels.CardRegistrationResponseDTO{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	externalId := snapUtils.GenerateExternalId()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	requestHeader := snapUtils.GenerateRequestHeaderDto(channelId, signature, timestamp, clientId, externalId, "", "", tokenB2B, "")
	return directDebitService.DoCardRegistrationProcess(requestHeader, cardRegistrationRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoRefund(refundRequestDTO refundModels.RefundRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, deviceId string, isProduction bool) (refundModels.RefundResponseDTO, error) {
	url := commons.DIRECT_DEBIT_REFUND
	minifiedRequestBody, err := json.Marshal(refundRequestDTO)
	if err != nil {
		return refundModels.RefundResponseDTO{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	externalId := snapUtils.GenerateExternalId()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, deviceId, ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoRefundProcess(requestHeader, refundRequestDTO, isProduction)

}

func (dd *DirectDebitController) DoCheckStatus(checkStatusRequestDTO checkStatusModels.CheckStatusRequestDTO, secretKey string, clientId string, tokenB2B string, isProduction bool) (checkStatusModels.CheckStatusResponseDTO, error) {

	url := commons.DIRECT_DEBIT_CHECK_STATUS

	minifiedRequestBody, err := json.Marshal(checkStatusRequestDTO)
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{}, fmt.Errorf("error marshalling request body: %w", err)
	}

	timestamp := tokenServices.GenerateTimestamp()
	externalId := snapUtils.GenerateExternalId()
	signature := tokenServices.GenerateSymetricSignature("POST", url, tokenB2B, minifiedRequestBody, timestamp, secretKey)

	requestHeader := snapUtils.GenerateRequestHeaderDto(
		"", signature, timestamp, clientId, externalId, "", "", tokenB2B, "",
	)

	checkStatusResponse, err := directDebitService.DoCheckStatusProcess(requestHeader, checkStatusRequestDTO, isProduction)
	if err != nil {
		return checkStatusResponse, fmt.Errorf("error in DoCheckStatusProcess: %v", err)
	}

	return checkStatusResponse, nil
}

func (dd *DirectDebitController) DoCardRegistrationUnbinding(cardRegistrationUnbindingRequestDTO registrationCardUnbindingModels.CardRegistrationUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) (registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO, error) {
	url := commons.DIRECT_DEBIT_CARD_UNBINDING
	minifiedRequestBody, err := json.Marshal(cardRegistrationUnbindingRequestDTO)
	if err != nil {
		return registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	httpMethod := "POST"
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature(httpMethod, url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, "")
	return directDebitService.DoCardRegistrationUnbindingProcess(requestHeader, cardRegistrationUnbindingRequestDTO, isProduction)
}
