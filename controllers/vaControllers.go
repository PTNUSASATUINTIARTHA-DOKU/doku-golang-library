package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
	checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"
	inquiryVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/inquiry"
	updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

var vaServices services.VaServices
var tokenServices services.TokenServices
var snapUtils utils.SnapUtils

type VaControllerInterface interface {
	CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto, secretKey string, clientId string, tokenB2B string, isProduction bool) (createVaModels.CreateVaResponseDto, error)
	DoUpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO, clientId string, tokenB2B string, secretKey string, isProduction bool) (updateVaModels.UpdateVaResponseDTO, error)
	DoCheckStatusVa(checkStatusVARequestDto checkVaModels.CheckStatusVARequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) (checkVaModels.CheckStatusVaResponseDto, error)
	DoDeletePaymentCode(deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) (deleteVaModels.DeleteVaResponseDto, error)
	DirectInquiryRequestMapping(headerRequest *http.Request, inquiryRequestBodyDto inquiryVaModels.InquiryRequestBodyDTO) (string, error)
	DirectInquiryResponseMapping(xmlData string) (inquiryVaModels.InquiryResponseBodyDTO, error)
}

type VaController struct{}

func (vc VaController) CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto, secretKey string, clientId string, tokenB2B string, isProduction bool) (createVaModels.CreateVaResponseDto, error) {
	timestamp := tokenServices.GenerateTimestamp()
	externalId := snapUtils.GenerateExternalId()
	createVaRequestDto.AdditionalInfo.Origin = createVaModels.Origin{
		Product:       "SDK",
		Source:        "Golang",
		SourceVersion: commons.SDK_VERSION,
		System:        "doku-golang-library",
		ApiFormat:     "SNAP",
	}
	minifiedRequestBody, err := json.Marshal(createVaRequestDto)
	if err != nil {
		return createVaModels.CreateVaResponseDto{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	endPointUrl := commons.CREATE_VA
	httpMethod := "POST"
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	requestHeader := vaServices.GenerateRequestHeaderDto("H2H", signature, timestamp, clientId, externalId, tokenB2B)

	response, err := vaServices.CreateVa(
		requestHeader,
		createVaRequestDto,
		isProduction)

	return response, err
}

func (vc VaController) DoUpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO, clientId string, tokenB2B string, secretKey string, isProduction bool) (updateVaModels.UpdateVaResponseDTO, error) {
	timeStamp := tokenServices.GenerateTimestamp()
	endPointUrl := commons.UPDATE_VA
	httpMethod := "PUT"
	updateVaRequestDTO.AdditionalInfo.Origin = createVaModels.Origin{
		Product:       "SDK",
		Source:        "Golang",
		SourceVersion: commons.SDK_VERSION,
		System:        "doku-golang-library",
		ApiFormat:     "SNAP",
	}
	minifiedRequestBody, err := json.Marshal(updateVaRequestDTO)
	if err != nil {
		return updateVaModels.UpdateVaResponseDTO{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timeStamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	header := vaServices.GenerateRequestHeaderDto("SDK", signature, timeStamp, clientId, externalId, tokenB2B)
	return vaServices.DoUpdateVa(header, updateVaRequestDTO, isProduction)
}

func (vc VaController) DoCheckStatusVa(checkStatusVARequestDto checkVaModels.CheckStatusVARequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) (checkVaModels.CheckStatusVaResponseDto, error) {
	timeStamp := tokenServices.GenerateTimestamp()
	endPointUrl := commons.CHECK_VA
	httpMethod := "POST"
	minifiedRequestBody, err := json.Marshal(checkStatusVARequestDto)
	if err != nil {
		return checkVaModels.CheckStatusVaResponseDto{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timeStamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	header := vaServices.GenerateRequestHeaderDto("SDK", signature, timeStamp, clientId, externalId, tokenB2B)
	return vaServices.DoCheckStatusVa(header, checkStatusVARequestDto, isProduction)
}

func (vc VaController) DoDeletePaymentCode(deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) (deleteVaModels.DeleteVaResponseDto, error) {
	timeStamp := tokenServices.GenerateTimestamp()
	endPointUrl := commons.DELETE_VA
	httpMethod := "DELETE"
	minifiedRequestBody, err := json.Marshal(deleteVaRequestDto)
	if err != nil {
		return deleteVaModels.DeleteVaResponseDto{}, fmt.Errorf("error marshalling request body: %w", err)
	}
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timeStamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	header := vaServices.GenerateRequestHeaderDto("SDK", signature, timeStamp, clientId, externalId, tokenB2B)
	return vaServices.DoDeletePaymentCode(header, deleteVaRequestDto, isProduction)
}

func (vc VaController) DirectInquiryRequestMapping(headerRequest *http.Request, inquiryRequestBodyDto inquiryVaModels.InquiryRequestBodyDTO) (string, error) {
	return vaServices.DirectInquiryRequestMapping(headerRequest, inquiryRequestBodyDto)
}

func (vc VaController) DirectInquiryResponseMapping(xmlData string) (inquiryVaModels.InquiryResponseBodyDTO, error) {
	return vaServices.DirectInquiryResponseMapping(xmlData)
}
