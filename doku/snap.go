package doku

import (
	"strconv"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/controllers"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models"
)

var TokenController controllers.TokenController

type Snap struct {
	// ----------------
	PrivateKey   string
	PublicKey    string
	Issuer       string
	ClientId     string
	IsProduction bool
	// ----------------
	TokenB2B                string
	TokenExpiresIn          int
	TokenGeneratedTimestamp string
}

func (snap *Snap) GetTokenB2B() {
	tokenB2BResponseDTO := TokenController.GetTokenB2B(snap.PrivateKey, snap.ClientId, snap.IsProduction)
	snap.setTokenB2B(tokenB2BResponseDTO)
}

func (snap *Snap) setTokenB2B(tokenB2BResponseDTO models.TokenB2BResponseDTO) {
	snap.TokenB2B = tokenB2BResponseDTO.AccessToken
	snap.TokenExpiresIn = tokenB2BResponseDTO.ExpiresIn - 10
	snap.TokenGeneratedTimestamp = strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}
