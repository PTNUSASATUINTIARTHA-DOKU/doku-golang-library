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
) createVaModels.CreateVaResponseDto {

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

	createVaRequestDto.Origin = createVaModels.Origin{
		Product:       "SDK",
		Source:        "Golang",
		SourceVersion: "1.0.0",
		System:        "doku-golang-library",
		ApiFormat:     "SNAP",
	}

	bodyRequest, err := json.Marshal(createVaRequestDto)
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
	var createVaResponseDTO createVaModels.CreateVaResponseDto
	if err := json.Unmarshal(respBody, &createVaResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return createVaResponseDTO

}

func (vs VaServices) DoUpdateVa(requestHeaderDTO createVaModels.RequestHeaderDTO, updateVaRequestDTO updateVaModels.UpdateVaDTO, isProduction bool) updateVaModels.UpdateVaResponseDTO {
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
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyRequest))
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
	var updateVaResponseDTO updateVaModels.UpdateVaResponseDTO
	if err := json.Unmarshal(respBody, &updateVaResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return updateVaResponseDTO
}

func (vs VaServices) DoCheckStatusVa(requestHeaderDTO createVaModels.RequestHeaderDTO, checkStatusVARequestDto checkVaModels.CheckStatusVARequestDto, isProduction bool) checkVaModels.CheckStatusVaResponseDto {
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
	var checkStatusVaResponseDTO checkVaModels.CheckStatusVaResponseDto
	if err := json.Unmarshal(respBody, &checkStatusVaResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return checkStatusVaResponseDTO
}

func (vs VaServices) DoDeletePaymentCode(requestHeaderDTO createVaModels.RequestHeaderDTO, deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, isProduction bool) deleteVaModels.DeleteVaResponseDto {
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
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(bodyRequest))
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
	var deleteVaResponseDto deleteVaModels.DeleteVaResponseDto
	if err := json.Unmarshal(respBody, &deleteVaResponseDto); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return deleteVaResponseDto
}

func (vs VaServices) V1SnapConverter(xmlData []byte) (map[string]interface{}, error) {
	var xmlResponse inquiryModels.InquiryResponse
	var response map[string]interface{}

	err := xml.Unmarshal(xmlData, &xmlResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling XML: %v", err)
	}

	switch xmlResponse.ResponseCode {
	case "0000":
		xmlResponse.ResponseCode = "2002400"
	case "3000", "3001", "3006":
		xmlResponse.ResponseCode = "4042412"
	case "3002":
		xmlResponse.ResponseCode = "4042414"
	case "3004":
		xmlResponse.ResponseCode = "4032400"
	case "9999":
		xmlResponse.ResponseCode = "5002401"
	}

	if xmlResponse.Currency >= "360" {
		xmlResponse.Currency = "IDR"
	}

	response = map[string]interface{}{
		"virtualAccountData": map[string]interface{}{
			"additionalInfo": map[string]interface{}{
				"trxId": xmlResponse.TransIdMerchant,
				"virtualAccountConfig": map[string]interface{}{
					"minAmount": xmlResponse.MinAmount,
					"maxAmount": xmlResponse.MaxAmount,
				},
			},
			"totalAmount": map[string]interface{}{
				"value":    xmlResponse.Amount,
				"currency": xmlResponse.Currency,
			},
			"virtualAccountName":  xmlResponse.Name,
			"virtualAccountEmail": xmlResponse.Email,
			"virtualAccountNo":    xmlResponse.PaymentCode,
			"customerNo":          xmlResponse.PaymentCode,
		},
		"responseCode": xmlResponse.ResponseCode,
	}

	return response, nil
}

func (vs VaServices) SnapV1Converter(jsonData []byte) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	vaData, ok := data["virtualAccountData"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("virtualAccountData not found")
	}

	additionalInfo, ok := vaData["additionalInfo"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("additionalInfo not found")
	}

	form := url.Values{}
	if partnerServiceId, ok := vaData["partnerServiceId"].(string); ok {
		form.Set("MALLID", partnerServiceId)
	} else {
		return "", fmt.Errorf("partnerServiceId not found or is not a string")
	}

	channel, ok := additionalInfo["channel"].(string)
	if !ok {
		return "", fmt.Errorf("channel not found or is not a string")
	}

	v1ChannelId := commons.GetVAChannelIdV1(channel)
	form.Set("PAYMENTCHANNEL", v1ChannelId)

	if paymentCode, ok := vaData["virtualAccountNo"].(string); ok {
		form.Set("PAYMENTCODE", paymentCode)
	} else {
		return "", fmt.Errorf("virtualAccountNo not found or is not a string")
	}

	form.Set("STATUSTYPE", "/")

	if inquiryRequestId, ok := vaData["inquiryRequestId"].(string); ok {
		form.Set("OCOID", inquiryRequestId)
	} else {
		return "", fmt.Errorf("inquiryRequestId not found or is not a string")
	}

	return form.Encode(), nil
}
