package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	accountUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountunbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

type DirectDebitInterface interface {
	DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) accountBindingModels.AccountBindingResponseDto
	DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) balanceInquiryModels.BalanceInquiryResponseDto
	DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B2C string, tokenB2B string, isProduction bool) paymentModels.PaymentResponseDTO
	DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) accountUnbindingModels.AccountUnbindingResponseDTO
}

var config commons.Config
var directDebitService services.DirectDebitService

type DirectDebitController struct{}

func (dd *DirectDebitController) DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) accountBindingModels.AccountBindingResponseDto {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_ACCOUNT_BINDING
	minifiedRequestBody, err := json.Marshal(accountBindingRequest)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature("POST", url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, deviceId, ipAddress, tokenB2B, "")
	return directDebitService.DoAccountBindingProcess(requestHeader, accountBindingRequest, isProduction)
}

func (dd *DirectDebitController) DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) balanceInquiryModels.BalanceInquiryResponseDto {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_BALANCE_INQUIRY_URL
	minifiedRequestBody, err := json.Marshal(balanceInquiryRequestDto)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature("POST", url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoBalanceInquiryProcess(requestHeader, balanceInquiryRequestDto, isProduction)
}

func (dd *DirectDebitController) DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B2C string, tokenB2B string, isProduction bool) paymentModels.PaymentResponseDTO {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_PAYMENT
	minifiedRequestBody, err := json.Marshal(paymentRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature("POST", url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("SDK", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, tokenB2B2C)
	return directDebitService.DoPaymentProcess(requestHeader, paymentRequestDTO, isProduction)
}

func (dd *DirectDebitController) DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) accountUnbindingModels.AccountUnbindingResponseDTO {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_ACCOUNT_UNBINDING
	minifiedRequestBody, err := json.Marshal(accountUnbindingRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	timestamp := tokenServices.GenerateTimestamp()
	signature := tokenServices.GenerateSymetricSignature("POST", url, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	requestHeader := snapUtils.GenerateRequestHeaderDto("SDK", signature, timestamp, clientId, externalId, "", ipAddress, tokenB2B, "")
	return directDebitService.DoAccountUnbindingProcess(requestHeader, accountUnbindingRequestDTO, isProduction)
}
