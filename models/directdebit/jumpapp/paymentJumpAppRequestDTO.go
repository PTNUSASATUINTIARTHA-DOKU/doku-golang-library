package jumpapp

import (
	"errors"
	"strings"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type PaymentJumpAppRequestDTO struct {
	PartnerReferenceNo string                                 `json:"partnerReferenceNo"`
	ValidUpTo          string                                 `json:"validUpTo,omitempty"`
	PointOfInitiation  string                                 `json:"pointOfInitiation,omitempty"`
	UrlParam           UrlParamDTO                            `json:"urlParam"`
	Amount             createVaModels.TotalAmount             `json:"amount"`
	AdditionalInfo     PaymentJumpAppAdditionalInfoRequestDTO `json:"additionalInfo"`
}

type UrlParamDTO struct {
	Url        string `json:"url"`
	Type       string `json:"type"`
	IsDeepLink string `json:"isDeepLink"`
}

type PaymentJumpAppAdditionalInfoRequestDTO struct {
	Channel    string `json:"channel"`
	OrderTitle string `json:"orderTitle"`
	Metadata   string `json:"metadata"`
}

func (dto *PaymentJumpAppRequestDTO) ValidatePaymentJumpAppRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(dto.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}

	if dto.PointOfInitiation != "" {
		if !strings.EqualFold(dto.PointOfInitiation, "app") && !strings.EqualFold(dto.PointOfInitiation, "pc") && !strings.EqualFold(dto.PointOfInitiation, "mweb") {
			return errors.New("pointOfInitiation value can only be app/pc/mweb")
		}
	}

	if !strings.EqualFold(dto.UrlParam.Type, "PAY_RETURN") {
		return errors.New("urlParam.type must always be PAY_RETURN")
	}

	return nil
}
