package controllers

import (
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/services"
)

var VaService services.VaServices

type VaController struct{}

func (vc VaController) CreateVa(createVaRequestDto models.CreateVaRequestDto, privateKey string, clientId string, tokenB2B string, isProduction bool) models.CreateVaResponseDto {
	timeStamp := TokenServices.GenerateTimestamp()
	requestHeader := VaService.CreateVaRequestHeaderDto(createVaRequestDto, privateKey, timeStamp, clientId, tokenB2B)

	response := VaService.CreateVa(
		requestHeader,
		createVaRequestDto,
		isProduction)

	return response
}
