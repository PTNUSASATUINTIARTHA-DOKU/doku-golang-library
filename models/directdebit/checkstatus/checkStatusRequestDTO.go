package checkstatus

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type CheckStatusRequestDTO struct {
	OriginalPartnerReferenceNo string                              `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        string                              `json:"originalReferenceNo,omitempty"`
	OriginalExternalId         string                              `json:"originalExternalId,omitempty"`
	ServiceCode                string                              `json:"serviceCode"`
	TransactionDate            string                              `json:"transactionDate,omitempty"`
	Amount                     createVaModels.TotalAmount          `json:"amount,omitempty"`
	MerchantId                 string                              `json:"merchantId,omitempty"`
	SubMerchantId              string                              `json:"subMerchantId,omitempty"`
	ExternalStoreId            string                              `json:"externalStoreId,omitempty"`
	AdditionalInfo             CheckStatusAdditionalInfoRequestDTO `json:"additionalInfo,omitempty"`
}

type CheckStatusAdditionalInfoRequestDTO struct {
	DeviceId string `json:"deviceId,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

func (dto *CheckStatusRequestDTO) ValidateCheckStatusRequest() error {
	if !directDebitChannel.ValidateDirectDebitChannel(dto.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	} else if dto.ServiceCode != "55" {
		return errors.New("serviceCode must be 55")
	}

	return nil
}
