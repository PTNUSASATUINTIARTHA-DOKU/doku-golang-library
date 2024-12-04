package balanceinquiry

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type BalanceInquiryRequestDto struct {
	AdditionalInfo BalanceInquiryAdditionalInfoRequestDto `json:"additionalInfo"`
}

type BalanceInquiryAdditionalInfoRequestDto struct {
	Channel string                `json:"channel"`
	Origin  createVaModels.Origin `json:"origin"`
}

func (dto *BalanceInquiryRequestDto) ValidateBalanceInquiryRequest(authCode string) error {
	if !directDebitChannel.ValidateDirectDebitChannel(dto.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	} else if authCode == "" {
		return errors.New("authCode cannot be an empty string or null")
	}
	return nil
}
