package refund

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type RefundRequestDTO struct {
	OriginalPartnerReferenceNo string                         `json:"originalPartnerReferenceNo"`
	OriginalExternalId         string                         `json:"originalExternalId,omitempty"`
	RefundAmount               createVaModels.TotalAmount     `json:"refundAmount"`
	Reason                     string                         `json:"reason,omitempty"`
	PartnerRefundNo            string                         `json:"partnerRefundNo"`
	AdditionalInfo             RefundAdditionalInfoRequestDTO `json:"additionalInfo"`
}

type RefundAdditionalInfoRequestDTO struct {
	Channel string `json:"channel"`
}

func (dto *RefundRequestDTO) ValidateRefundRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(dto.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}

	return nil
}
