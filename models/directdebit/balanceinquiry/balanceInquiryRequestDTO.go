package balanceinquiry

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
)

type BalanceInquiryRequestDto struct {
	AdditionalInfo BalanceInquiryAdditionalInfoRequestDto `json:"additionalInfo"`
}

type BalanceInquiryAdditionalInfoRequestDto struct {
	Channel string `json:"channel"`
}

func (dto *BalanceInquiryRequestDto) ValidateBalanceInquiryRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(dto.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}
	return nil
}
