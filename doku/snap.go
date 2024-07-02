package doku

import (
	"strconv"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/controllers"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models"
)

var TokenController controllers.TokenController
var VaController controllers.VaController

type Snap struct {
	// ----------------
	PrivateKey   string
	PublicKey    string
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

func (snap *Snap) setTokenB2B(tokenB2BResponseDTO models.TokenB2BResponseDTO) {
	snap.tokenB2B = tokenB2BResponseDTO.AccessToken
	snap.tokenExpiresIn = tokenB2BResponseDTO.ExpiresIn - 10
	snap.tokenGeneratedTimestamp = strconv.FormatInt(time.Now().Unix(), 10)
}

func (snap *Snap) CreateVa(createVaRequestDto models.CreateVaRequestDto) models.CreateVaResponseDto {
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
