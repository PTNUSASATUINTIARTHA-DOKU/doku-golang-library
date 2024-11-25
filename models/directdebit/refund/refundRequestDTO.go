package refund

import (
	"errors"
	"fmt"
	"regexp"

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
	Channel string                `json:"channel"`
	Origin  createVaModels.Origin `json:"origin"`
}

func (dto *RefundRequestDTO) ValidateRefundRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(dto.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}

	if err := dto.validateOriginalPartnerReferenceNo(); err != nil {
		return err
	}

	if err := dto.validateAmount(); err != nil {
		return err
	}

	if err := dto.validateCurrency(); err != nil {
		return err
	}

	if err := dto.validatePartnerRefundNo(); err != nil {
		return err
	}

	return nil
}

func (dto *RefundRequestDTO) validateOriginalPartnerReferenceNo() error {
	value := dto.OriginalPartnerReferenceNo
	channel := dto.AdditionalInfo.Channel

	if value == "" {
		return errors.New("originalPartnerReferenceNo cannot be null. Please provide an originalPartnerReferenceNo. Example: 'INV-0001'")
	}

	switch channel {
	case "EMONEY_OVO_SNAP":
		if len(value) > 32 {
			return fmt.Errorf("originalPartnerReferenceNo must be 32 characters or fewer. Current length is %d. Example: 'INV-001'", len(value))
		}
	case "EMONEY_DANA_SNAP", "EMONEY_SHOPEE_PAY_SNAP", "DIRECT_DEBIT_ALLO_SNAP":
		if len(value) > 64 {
			return fmt.Errorf("originalPartnerReferenceNo must be 64 characters or fewer. Current length is %d. Example: 'INV-001'", len(value))
		}
	case "DIRECT_DEBIT_CIMB_SNAP", "DIRECT_DEBIT_BRI_SNAP":
		if len(value) > 12 {
			return fmt.Errorf("originalPartnerReferenceNo must be 12 characters or fewer. Current length is %d. Example: 'INV-001'", len(value))
		}
	}

	return nil
}

func (dto *RefundRequestDTO) validateAmount() error {
	valueLength := len(dto.RefundAmount.Value)
	if valueLength < 4 {
		return errors.New("refundAmount.Value must be at least 4 characters long and formatted as 0.00. Example: '100.00'")
	}
	if valueLength > 19 {
		return errors.New("refundAmount.Value must be 19 characters or fewer and formatted as 9999999999999999.99. Example: '9999999999999999.99'")
	}
	if !regexp.MustCompile(`^(0|[1-9]\d{0,15}\.\d{2})?$`).MatchString(dto.RefundAmount.Value) {
		return errors.New("refundAmount.Value is in invalid format")
	}
	return nil
}

func (dto *RefundRequestDTO) validateCurrency() error {
	value := dto.RefundAmount.Currency

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

func (dto *RefundRequestDTO) validatePartnerRefundNo() error {
	value := dto.PartnerRefundNo
	channel := dto.AdditionalInfo.Channel

	if value == "" {
		return errors.New("partnerRefundNo cannot be null. Please provide a partnerRefundNo. Example: 'INV-0001'")
	}

	if channel == "EMONEY_DANA_SNAP" || channel == "EMONEY_SHOPEE_PAY_SNAP" || channel == "EMONEY_OVO_SNAP" {
		if len(value) > 64 {
			return errors.New("partnerRefundNo must be 64 characters or fewer. Ensure that partnerRefundNo is no longer than 64 characters. Example: 'INV-REF-001'")
		}
	} else if channel == "DIRECT_DEBIT_CIMB_SNAP" || channel == "DIRECT_DEBIT_BRI_SNAP" {
		if len(value) > 12 {
			return errors.New("partnerRefundNo must be 12 characters or fewer. Ensure that partnerRefundNo is no longer than 12 characters. Example: 'INV-REF-001'")
		}
	} else if channel == "DIRECT_DEBIT_ALLO_SNAP" {
		if len(value) < 32 || len(value) > 64 {
			return errors.New("partnerRefundNo must be 64 characters and at least 32 characters. Ensure that partnerRefundNo is no longer than 64 characters and at least 32 characters. Example: 'INV-REF-001'")
		}
	}

	return nil
}
