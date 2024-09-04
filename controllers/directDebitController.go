package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

type DirectDebitInterface interface {
	DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) accountBindingModels.AccountBindingResponseDto
	DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) balanceInquiryModels.BalanceInquiryResponseDto
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
