package payment

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type PaymentRequestDTO struct {
	PartnerReferenceNo string                     `json:"partnerReferenceNo"`
	FeeType            string                     `json:"feeType,omitempty"` //ovo
	Amount             createVaModels.TotalAmount `json:"totalAmount"`
	PayOptionDetails   []PayOptionDetailsDTO      `json:"payOptionDetails,omitempty"` //allo, ovo
	AdditionalInfo     PaymentAdditionalInfoDTO   `json:"additionalInfo"`
}

type PayOptionDetailsDTO struct {
	PayMethod   string                     `json:"payMethod"`
	TransAmount createVaModels.TotalAmount `json:"transAmount"`
	FeeAmount   createVaModels.TotalAmount `json:"feeAmount"`
}

type PaymentAdditionalInfoDTO struct {
	Channel           string         `json:"channel"`
	Remarks           string         `json:"remarks,omitempty"` //allo, cimb
	SuccessPaymentUrl string         `json:"successPaymentUrl,omitempty"`
	FailedPaymentUrl  string         `json:"failedPaymentUrl,omitempty"`
	LineItems         []LineItemsDTO `json:"lineItems,omitempty"`   //allo
	PaymentType       string         `json:"paymentType,omitempty"` //bri, ovo
}

type LineItemsDTO struct {
	Name     string `json:"name,omitempty"`
	Price    string `json:"price,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

func (p *PaymentRequestDTO) ValidatePaymentRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(p.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}

	if p.AdditionalInfo.Channel == directDebitChannel.DirectDebitChannelNames[directDebitChannel.EMONEY_OVO_SNAP] {
		if p.FeeType != "" && !directDebitChannel.IsValidFeeType(p.FeeType) {
			return errors.New("value can only be OUR/BEN/SHA for EMONEY_OVO_SNAP")
		}

		if len(p.PayOptionDetails) == 0 {
			return errors.New("pay option details cannot be empty for EMONEY_OVO_SNAP")
		}

		if p.AdditionalInfo.PaymentType != "" && !directDebitChannel.IsValidPaymentType(p.AdditionalInfo.PaymentType) {
			return errors.New("additionalInfo.paymentType cannot be empty for EMONEY_OVO_SNAP")
		}
	}

	if p.AdditionalInfo.Channel == directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_ALLO_SNAP] {
		if len(p.AdditionalInfo.LineItems) == 0 {
			return errors.New("additionalInfo.lineItems cannot be empty for DIRECT_DEBIT_ALLO_SNAP")
		}

		if p.AdditionalInfo.Remarks == "" {
			return errors.New("additionalInfo.remarks cannot be empty for DIRECT_DEBIT_ALLO_SNAP")
		}
	}

	if p.AdditionalInfo.Channel == directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_CIMB_SNAP] {
		if p.AdditionalInfo.Remarks == "" {
			return errors.New("additionalInfo.remarks cannot be empty for DIRECT_DEBIT_CIMB_SNAP")
		}
	}

	if p.AdditionalInfo.Channel == directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_BRI_SNAP] {
		if p.AdditionalInfo.PaymentType != "" && !directDebitChannel.IsValidPaymentType(p.AdditionalInfo.PaymentType) {
			return errors.New("additionalInfo.paymentType cannot be empty for DIRECT_DEBIT_BRI_SNAP")
		}
	}

	return nil
}
