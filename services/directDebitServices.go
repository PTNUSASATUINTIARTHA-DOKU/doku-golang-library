package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type DirectDebitService struct{}

func (dd *DirectDebitService) DoAccountBindingProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, accountBindingRequestDTO accountBindingModels.AccountBindingRequestDTO, isProduction bool) accountBindingModels.AccountBindingResponseDto {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_ACCOUNT_BINDING
	header := map[string]string{
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":  requestHeaderDTO.XIpAddress,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(accountBindingRequestDTO)
	if err != nil {
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var accountBindingResponseDTO accountBindingModels.AccountBindingResponseDto
	if err := json.Unmarshal(respBody, &accountBindingResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return accountBindingResponseDTO
}

func (dd *DirectDebitService) DoBalanceInquiryProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, isProduction bool) balanceInquiryModels.BalanceInquiryResponseDto {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_BALANCE_INQUIRY_URL
	header := map[string]string{
		"X-TIMESTAMP":            requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":            requestHeaderDTO.XSignature,
		"X-PARTNER-ID":           requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID":          requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":           requestHeaderDTO.XIpAddress,
		"Authorization-Customer": "Bearer " + requestHeaderDTO.AuthorizationCustomer,
		"Authorization":          "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":           "application/json",
	}

	bodyRequest, err := json.Marshal(balanceInquiryRequestDto)
	if err != nil {
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var balanceInquiryResponseDto balanceInquiryModels.BalanceInquiryResponseDto
	if err := json.Unmarshal(respBody, &balanceInquiryResponseDto); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}
	return balanceInquiryResponseDto
}
