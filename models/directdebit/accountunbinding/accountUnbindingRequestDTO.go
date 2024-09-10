package accountunbinding

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
)

type AccountUnbindingRequestDTO struct {
	TokenId        string                                   `json:"tokenId"`
	AdditionalInfo AccountUnbindingAdditionalInfoRequestDTO `json:"additionalInfo"`
}

type AccountUnbindingAdditionalInfoRequestDTO struct {
	Channel string `json:"channel"`
}

func (au *AccountUnbindingRequestDTO) ValidateAccountUnbindingRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(au.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}
	return nil
}
