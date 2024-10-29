package services

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	inquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/converter"
	checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"
	inquiryVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/inquiry"
	updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"
)

type VaServices struct{}

func (vs VaServices) GenerateRequestHeaderDto(
	channelId string,
	signature string,
	timestamp string,
	clientId string,
	externalId string,
	tokenB2B string) createVaModels.RequestHeaderDTO {

	return createVaModels.RequestHeaderDTO{
		XTimestamp:    timestamp,
		XSignature:    signature,
		XPartnerId:    clientId,
		XExternalId:   externalId,
		ChannelId:     channelId,
		Authorization: tokenB2B,
	}
}

func (vs VaServices) CreateVa(
	requestHeaderDto createVaModels.RequestHeaderDTO,
	createVaRequestDto createVaModels.CreateVaRequestDto,
	isProduction bool,
) (createVaModels.CreateVaResponseDto, error) {

	url := config.GetBaseUrl(isProduction) + commons.CREATE_VA

	header := map[string]string{
		"X-PARTNER-ID":  requestHeaderDto.XPartnerId,
		"X-TIMESTAMP":   requestHeaderDto.XTimestamp,
		"X-SIGNATURE":   requestHeaderDto.XSignature,
		"Authorization": "Bearer " + requestHeaderDto.Authorization,
		"X-EXTERNAL-ID": requestHeaderDto.XExternalId,
		"CHANNEL-ID":    requestHeaderDto.ChannelId,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(createVaRequestDto)
	if err != nil {
		return createVaModels.CreateVaResponseDto{}, fmt.Errorf("error body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return createVaModels.CreateVaResponseDto{}, fmt.Errorf("error body response: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return createVaModels.CreateVaResponseDto{}, fmt.Errorf("error body response: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))
	var createVaResponseDTO createVaModels.CreateVaResponseDto
	if err := json.Unmarshal(respBody, &createVaResponseDTO); err != nil {
		return createVaModels.CreateVaResponseDto{}, fmt.Errorf("error body response: %w", err)
	}

	return createVaResponseDTO, nil

}

func (vs VaServices) DoUpdateVa(requestHeaderDTO createVaModels.RequestHeaderDTO, updateVaRequestDTO updateVaModels.UpdateVaDTO, isProduction bool) (updateVaModels.UpdateVaResponseDTO, error) {
	url := config.GetBaseUrl(isProduction) + commons.UPDATE_VA

	header := map[string]string{
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"CHANNEL-ID":    requestHeaderDTO.ChannelId,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(updateVaRequestDTO)
	if err != nil {
		return updateVaModels.UpdateVaResponseDTO{}, fmt.Errorf("error body request : %w", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return updateVaModels.UpdateVaResponseDTO{}, fmt.Errorf("error body response : %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return updateVaModels.UpdateVaResponseDTO{}, fmt.Errorf("error response : %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))
	var updateVaResponseDTO updateVaModels.UpdateVaResponseDTO
	if err := json.Unmarshal(respBody, &updateVaResponseDTO); err != nil {
		return updateVaModels.UpdateVaResponseDTO{}, fmt.Errorf("error body response : %w", err)
	}

	return updateVaResponseDTO, nil
}

func (vs VaServices) DoCheckStatusVa(requestHeaderDTO createVaModels.RequestHeaderDTO, checkStatusVARequestDto checkVaModels.CheckStatusVARequestDto, isProduction bool) (checkVaModels.CheckStatusVaResponseDto, error) {
	url := config.GetBaseUrl(isProduction) + commons.CHECK_VA

	header := map[string]string{
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"CHANNEL-ID":    requestHeaderDTO.ChannelId,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(checkStatusVARequestDto)
	if err != nil {
		return checkVaModels.CheckStatusVaResponseDto{}, fmt.Errorf("error body request : %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return checkVaModels.CheckStatusVaResponseDto{}, fmt.Errorf("error body request : %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return checkVaModels.CheckStatusVaResponseDto{}, fmt.Errorf("error response : %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))
	var checkStatusVaResponseDTO checkVaModels.CheckStatusVaResponseDto
	if err := json.Unmarshal(respBody, &checkStatusVaResponseDTO); err != nil {
		return checkVaModels.CheckStatusVaResponseDto{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	return checkStatusVaResponseDTO, nil
}

func (vs VaServices) DoDeletePaymentCode(requestHeaderDTO createVaModels.RequestHeaderDTO, deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, isProduction bool) (deleteVaModels.DeleteVaResponseDto, error) {
	url := config.GetBaseUrl(isProduction) + commons.DELETE_VA

	header := map[string]string{
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"CHANNEL-ID":    requestHeaderDTO.ChannelId,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(deleteVaRequestDto)
	if err != nil {
		return deleteVaModels.DeleteVaResponseDto{}, fmt.Errorf("error body response : %w", err)
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return deleteVaModels.DeleteVaResponseDto{}, fmt.Errorf("error body request : %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return deleteVaModels.DeleteVaResponseDto{}, fmt.Errorf("error response : %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))
	var deleteVaResponseDto deleteVaModels.DeleteVaResponseDto
	if err := json.Unmarshal(respBody, &deleteVaResponseDto); err != nil {
		return deleteVaModels.DeleteVaResponseDto{}, fmt.Errorf("error unmarshaling response JSON: ", err)
	}

	return deleteVaResponseDto, nil
}

func (vs VaServices) DirectInquiryResponseMapping(xmlData string) (inquiryVaModels.InquiryResponseBodyDTO, error) {
	var xmlResponse inquiryModels.InquiryResponse
	var response inquiryVaModels.InquiryResponseBodyDTO
	var responseMessage = ""

	xmlBytes := []byte(xmlData)

	err := xml.Unmarshal(xmlBytes, &xmlResponse)
	if err != nil {
		return response, fmt.Errorf("error unmarshaling XML: %v", err)
	}

	switch xmlResponse.ResponseCode {
	case "0000":
		xmlResponse.ResponseCode = "2002400"
		responseMessage = "Inquiry Success"
	case "3000", "3001":
		xmlResponse.ResponseCode = "4042412"
		responseMessage = "Invalid Virtual Account Number"
	case "3006":
		xmlResponse.ResponseCode = "4042412"
		responseMessage = "Billing Not Found"
	case "3002":
		xmlResponse.ResponseCode = "4042414"
		responseMessage = "Inquiry Decline by merchant"
	case "3004":
		xmlResponse.ResponseCode = "4032400"
		responseMessage = "Billing Was Expired"
	case "9999":
		xmlResponse.ResponseCode = "5002401"
		responseMessage = "Unexpected Failure"
	}

	if xmlResponse.Currency >= "360" {
		xmlResponse.Currency = "IDR"
	}

	response = inquiryVaModels.InquiryResponseBodyDTO{
		ResponseCode:    xmlResponse.ResponseCode,
		ResponseMessage: responseMessage,
		VirtualAccountData: &inquiryVaModels.InquiryRequestVirtualAccountDataDTO{
			CustomerNo:          xmlResponse.PaymentCode,
			VirtualAccountNo:    xmlResponse.PaymentCode,
			VirtualAccountName:  xmlResponse.Name,
			VirtualAccountEmail: xmlResponse.Email,
			TotalAmount: createVaModels.TotalAmount{
				Value:    xmlResponse.Amount,
				Currency: xmlResponse.Currency,
			},
			AdditionalInfo: inquiryVaModels.InquiryResponseAdditionalInfoDTO{
				TrxId: xmlResponse.TransIdMerchant,
				VirtualAccountConfig: createVaModels.VirtualAccountConfig{
					MinAmount: &xmlResponse.MinAmount,
					MaxAmount: &xmlResponse.MaxAmount,
				},
			},
		},
	}

	return response, nil
}

func (vs VaServices) DirectInquiryRequestMapping(headerRequest *http.Request, jsonData inquiryVaModels.InquiryRequestBodyDTO) (string, error) {
	form := url.Values{}

	partnerServiceId := headerRequest.Header.Get("X-PARTNER-ID")
	if partnerServiceId == "" {
		return "", fmt.Errorf("X-PARTNER-ID header not found or is empty")
	}
	form.Set("MALLID", partnerServiceId)

	channel := jsonData.AdditionalInfo.Channel
	if channel == "" {
		return "", fmt.Errorf("channel not found or is empty")
	}

	v1ChannelId := commons.GetVAChannelIdV1(channel)
	form.Set("PAYMENTCHANNEL", v1ChannelId)

	if jsonData.VirtualAccountNo != "" {
		form.Set("PAYMENTCODE", jsonData.VirtualAccountNo)
	} else {
		return "", fmt.Errorf("virtualAccountNo not found or is empty")
	}

	form.Set("STATUSTYPE", "/")

	if jsonData.InquiryRequestId != "" {
		form.Set("OCOID", jsonData.InquiryRequestId)
	} else {
		return "", fmt.Errorf("inquiryRequestId not found or is empty")
	}

	return form.Encode(), nil
}
