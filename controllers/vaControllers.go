package controllers

import (
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

var vaServices services.VaServices
var tokenServices services.TokenServices

type VaController struct{}

func (vc VaController) CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto, privateKey string, clientId string, tokenB2B string, isProduction bool) createVaModels.CreateVaResponseDto {
	timeStamp := tokenServices.GenerateTimestamp()
	externalId := vaServices.GenerateExternalId()
	signature, _ := tokenServices.CreateSignature(privateKey, clientId, timeStamp)
	requestHeader := vaServices.GenerateRequestHeaderDto("SDK", signature, timeStamp, clientId, externalId, tokenB2B)

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
	signature := tokenServices.GenerateSymetricSignature(httpMethod, endPointUrl, tokenB2B, updateVaRequestDTO, timeStamp, secretKey)
	externalId := vaServices.GenerateExternalId()
	header := vaServices.GenerateRequestHeaderDto("SDK", signature, timeStamp, clientId, externalId, tokenB2B)
	return vaServices.DoUpdateVa(header, updateVaRequestDTO, isProduction)
}
