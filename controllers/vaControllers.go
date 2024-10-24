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
	CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto, secretKey string, clientId string, tokenB2B string, isProduction bool) createVaModels.CreateVaResponseDto
	DoUpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO, clientId string, tokenB2B string, secretKey string, isProduction bool) updateVaModels.UpdateVaResponseDTO
	DoCheckStatusVa(checkStatusVARequestDto checkVaModels.CheckStatusVARequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) checkVaModels.CheckStatusVaResponseDto
	DoDeletePaymentCode(deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) deleteVaModels.DeleteVaResponseDto
	DirectInquiryRequestMapping(headerRequest *http.Request, inquiryRequestBodyDto inquiryVaModels.InquiryRequestBodyDTO) (string, error)
	DirectInquiryResponseMapping(xmlData string) (inquiryVaModels.InquiryResponseBodyDTO, error)
}

type VaController struct{}

func (vc VaController) CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto, secretKey string, clientId string, tokenB2B string, isProduction bool) createVaModels.CreateVaResponseDto {
	timestamp := tokenServices.GenerateTimestamp()
	externalId := snapUtils.GenerateExternalId()
	createVaRequestDto.Origin = createVaModels.Origin{
		Product:       "SDK",
		Source:        "Golang",
		SourceVersion: "v1.0.9",
		System:        "doku-golang-library",
		ApiFormat:     "SNAP",
	}
	minifiedRequestBody, err := json.Marshal(createVaRequestDto)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	endPointUrl := commons.CREATE_VA
	httpMethod := "POST"
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timestamp, secretKey)
	requestHeader := vaServices.GenerateRequestHeaderDto("H2H", signature, timestamp, clientId, externalId, tokenB2B)

	response := vaServices.CreateVa(
		requestHeader,
		createVaRequestDto,
		isProduction)

	return response
}

func (vc VaController) DoUpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO, clientId string, tokenB2B string, secretKey string, isProduction bool) updateVaModels.UpdateVaResponseDTO {
	timeStamp := tokenServices.GenerateTimestamp()
	endPointUrl := commons.UPDATE_VA
	httpMethod := "PUT"
	minifiedRequestBody, err := json.Marshal(updateVaRequestDTO)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timeStamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	header := vaServices.GenerateRequestHeaderDto("SDK", signature, timeStamp, clientId, externalId, tokenB2B)
	return vaServices.DoUpdateVa(header, updateVaRequestDTO, isProduction)
}

func (vc VaController) DoCheckStatusVa(checkStatusVARequestDto checkVaModels.CheckStatusVARequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) checkVaModels.CheckStatusVaResponseDto {
	timeStamp := tokenServices.GenerateTimestamp()
	endPointUrl := commons.CHECK_VA
	httpMethod := "POST"
	minifiedRequestBody, err := json.Marshal(checkStatusVARequestDto)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, minifiedRequestBody, timeStamp, secretKey)
	externalId := snapUtils.GenerateExternalId()
	header := vaServices.GenerateRequestHeaderDto("SDK", signature, timeStamp, clientId, externalId, tokenB2B)
	return vaServices.DoCheckStatusVa(header, checkStatusVARequestDto, isProduction)
}

func (vc VaController) DoDeletePaymentCode(deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) deleteVaModels.DeleteVaResponseDto {
	timeStamp := tokenServices.GenerateTimestamp()
	endPointUrl := commons.DELETE_VA
	httpMethod := "DELETE"
	minifiedRequestBody, err := json.Marshal(deleteVaRequestDto)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
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
