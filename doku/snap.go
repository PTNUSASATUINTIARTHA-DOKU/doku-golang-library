package doku

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/controllers"
	tokenVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	notificationTokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/token"
	updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"
)

var TokenController controllers.TokenController
var VaController controllers.VaController

type Snap struct {
	// ----------------
	PrivateKey   string
	PublicKey    string
	SecretKey    string
	Issuer       string
	ClientId     string
	IsProduction bool
	// ----------------
	tokenB2B                string
	tokenExpiresIn          int
	tokenGeneratedTimestamp string
}

func (snap *Snap) GetTokenB2B() {
	tokenB2BResponseDTO := TokenController.GetTokenB2B(snap.PrivateKey, snap.ClientId, snap.IsProduction)
	snap.setTokenB2B(tokenB2BResponseDTO)
}

func (snap *Snap) setTokenB2B(tokenB2BResponseDTO tokenVaModels.TokenB2BResponseDTO) {
	snap.tokenB2B = tokenB2BResponseDTO.AccessToken
	snap.tokenExpiresIn = tokenB2BResponseDTO.ExpiresIn - 10
	snap.tokenGeneratedTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
}

func (snap *Snap) CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto) createVaModels.CreateVaResponseDto {
	createVaRequestDto.ValidateVaRequestDto()
	isTokenInvalid := TokenController.IsTokenInvalid(
		snap.tokenB2B,
		snap.tokenExpiresIn,
		snap.tokenGeneratedTimestamp,
	)
	if isTokenInvalid {
		TokenController.GetTokenB2B(
			snap.PrivateKey,
			snap.ClientId,
			snap.IsProduction,
		)
	}
	createVaResponse := VaController.CreateVa(
		createVaRequestDto,
		snap.PrivateKey,
		snap.ClientId,
		snap.tokenB2B,
		snap.IsProduction,
	)
	return createVaResponse
}

func (snap *Snap) UpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO) updateVaModels.UpdateVaResponseDTO {

	updateVaRequestDTO.ValidateUpdateVaRequestDTO()
	isTokenInvalid := TokenController.IsTokenInvalid(
		snap.tokenB2B,
		snap.tokenExpiresIn,
		snap.tokenGeneratedTimestamp,
	)
	if isTokenInvalid {
		TokenController.GetTokenB2B(
			snap.PrivateKey,
			snap.ClientId,
			snap.IsProduction,
		)
	}
	updateVaResponse := VaController.DoUpdateVa(updateVaRequestDTO, snap.ClientId, snap.tokenB2B, snap.SecretKey, snap.IsProduction)

	return updateVaResponse
}

func (snap *Snap) CheckStatusVa(checkStatusVaRequestDto checkVaModels.CheckStatusVARequestDto) checkVaModels.CheckStatusVaResponseDto {

	checkStatusVaRequestDto.ValidateCheckStatusVaRequestDto()
	isTokenInvalid := TokenController.IsTokenInvalid(
		snap.tokenB2B,
		snap.tokenExpiresIn,
		snap.tokenGeneratedTimestamp,
	)
	if isTokenInvalid {
		TokenController.GetTokenB2B(
			snap.PrivateKey,
			snap.ClientId,
			snap.IsProduction,
		)
	}
	checkStatusVaResponseDTO := VaController.DoCheckStatusVa(checkStatusVaRequestDto, snap.PrivateKey, snap.ClientId, snap.tokenB2B, snap.SecretKey, snap.IsProduction)

	return checkStatusVaResponseDTO
}

func (snap *Snap) generateTokenB2B(isSignatureValid bool) notificationTokenModels.NotificationTokenDTO {
	if isSignatureValid {
		return TokenController.GenerateTokenB2B(snap.tokenExpiresIn, snap.Issuer, snap.PrivateKey, snap.ClientId)
	} else {
		return TokenController.GenerateInvalidSignatureResponse()
	}
}

func (snap *Snap) ValidateTokenB2B(requestTokenB2B string) bool {
	return TokenController.ValidateTokenB2B(requestTokenB2B, snap.PublicKey)
}

func (snap *Snap) validateSignature(request *http.Request) bool {
	return TokenController.ValidateSignature(request, snap.PrivateKey, snap.ClientId)
}

func (snap *Snap) ValidateSignatureAndGenerateToken(request *http.Request) notificationTokenModels.NotificationTokenDTO {
	var isSignatureValid = snap.validateSignature(request)
	return snap.generateTokenB2B(isSignatureValid)
}
