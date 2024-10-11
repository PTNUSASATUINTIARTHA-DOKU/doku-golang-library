package payment

import (
	"errors"
	"regexp"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type PaymentRequestDTO struct {
	PartnerReferenceNo string                     `json:"partnerReferenceNo"`
	FeeType            string                     `json:"feeType,omitempty"` //ovo
	Amount             createVaModels.TotalAmount `json:"amount"`
	PayOptionDetails   []PayOptionDetailsDTO      `json:"payOptionDetails,omitempty"` //allo, ovo
	AdditionalInfo     PaymentAdditionalInfoDTO   `json:"additionalInfo"`
	ChargeToken        string                     `json:"chargeToken"`
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
	if err := p.validateChannel(); err != nil {
		return err
	}

	if err := p.validateAmount(); err != nil {
		return err
	}

	if err := p.validateCurrency(); err != nil {
		return err
	}

	switch p.AdditionalInfo.Channel {
	case directDebitChannel.DirectDebitChannelNames[directDebitChannel.EMONEY_OVO_SNAP]:
		return p.validateEMoneyOVOSnap()

	case directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_ALLO_SNAP]:
		return p.validateDirectDebitAlloSnap()

	case directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_CIMB_SNAP]:
		return p.validateDirectDebitCIMBSnap()

	case directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_BRI_SNAP]:
		return p.validateDirectDebitBRISnap()
	}

	return nil
}

func (p *PaymentRequestDTO) validateChannel() error {
	if !directDebitChannel.ValidateDirectDebitChannel(p.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}
	return nil
}

func (p *PaymentRequestDTO) validateAmount() error {
	valueLength := len(p.Amount.Value)
	if valueLength < 4 {
		return errors.New("Amount.Value must be at least 4 characters long and formatted as 0.00. Example: '100.00'")
	}
	if valueLength > 19 {
		return errors.New("Amount.Value must be 19 characters or fewer and formatted as 9999999999999999.99. Example: '9999999999999999.99'")
	}
	if !regexp.MustCompile(`^(0|[1-9]\d{0,15}\.\d{2})?$`).MatchString(p.Amount.Value) {
		return errors.New("Amount.Value is in invalid format")
	}
	return nil
}

func (p *PaymentRequestDTO) validateEMoneyOVOSnap() error {
	if p.FeeType != "" && !directDebitChannel.IsValidFeeType(p.FeeType) {
		return errors.New("value can only be OUR/BEN/SHA for EMONEY_OVO_SNAP")
	}
	if len(p.PayOptionDetails) == 0 {
		return errors.New("pay option details cannot be empty for EMONEY_OVO_SNAP")
	}
	if err := p.validateAdditionalInfo(); err != nil {
		return err
	}
	return nil
}

func (p *PaymentRequestDTO) validateDirectDebitAlloSnap() error {
	if len(p.AdditionalInfo.LineItems) == 0 {
		return errors.New("additionalInfo.lineItems cannot be empty for DIRECT_DEBIT_ALLO_SNAP")
	}
	if p.AdditionalInfo.Remarks == "" {
		return errors.New("additionalInfo.remarks cannot be empty for DIRECT_DEBIT_ALLO_SNAP")
	}
	if len(p.AdditionalInfo.Remarks) > 40 {
		return errors.New("additionalInfo.remarks must be 40 characters or fewer. Example: 'remarks'")
	}
	return p.validateAdditionalInfo()
}

func (p *PaymentRequestDTO) validateDirectDebitCIMBSnap() error {
	if p.AdditionalInfo.Remarks == "" {
		return errors.New("additionalInfo.remarks cannot be empty for DIRECT_DEBIT_CIMB_SNAP")
	}
	if len(p.AdditionalInfo.Remarks) > 40 {
		return errors.New("additionalInfo.remarks must be 40 characters or fewer. Example: 'remarks'")
	}
	return p.validateAdditionalInfo()
}

func (p *PaymentRequestDTO) validateDirectDebitBRISnap() error {
	if p.AdditionalInfo.PaymentType != "" && !directDebitChannel.IsValidPaymentType(p.AdditionalInfo.PaymentType) {
		return errors.New("additionalInfo.paymentType is invalid for DIRECT_DEBIT_BRI_SNAP")
	}
	return p.validateAdditionalInfo()
}

func (p *PaymentRequestDTO) validateAdditionalInfo() error {
	if p.AdditionalInfo.SuccessPaymentUrl == "" {
		return errors.New("additionalInfo.SuccessPaymentUrl cannot be null. Example: 'https://www.doku.com'")
	}
	if p.AdditionalInfo.FailedPaymentUrl == "" {
		return errors.New("additionalInfo.FailedPaymentUrl cannot be null. Example: 'https://www.doku.com'")
	}
	return nil
}

func (p *PaymentRequestDTO) validateCurrency() error {
	value := p.Amount.Currency

	if value == "" {
		return errors.New("must be a string; ensure that refundAmount.Currency is enclosed in quotes, e.g., 'IDR'")
	}

	if len(value) != 3 {
		return errors.New("refundAmount.currency must be exactly 3 characters long, e.g., 'IDR'")
	}

	if value != "IDR" {
		return errors.New("refundAmount.currency must be 'IDR', e.g., 'IDR'")
	}

	return nil
}
