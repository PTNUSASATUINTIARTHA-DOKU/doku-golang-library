package cardregistrationunbinding

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
)

type CardRegistrationUnbindingRequestDTO struct {
	TokenId        string                                            `json:"tokenId"`
	AdditionalInfo CardRegistrationUnbindingAdditionalInfoRequestDTO `json:"additionalInfo"`
}

type CardRegistrationUnbindingAdditionalInfoRequestDTO struct {
	Channel string `json:"channel"`
}

func (au *CardRegistrationUnbindingRequestDTO) ValidateCardRegistrationUnbindingRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(au.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	} else if len(au.TokenId) > 2048 {
		return errors.New("tokenId must be 2048 characters or fewer. Ensure that tokenId is no longer than 2048 characters. Example: 'eyJhbGciOiJSUzI1NiJ...'")
	}
	return nil
}
