package cardregistration

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
)

type CardRegistrationRequestDTO struct {
	CardData       string                                   `json:"cardData"`
	CustIdMerchant string                                   `json:"custIdMerchant"`
	PhoneNo        string                                   `json:"phoneNo"`
	AdditionalInfo CardRegistrationAdditionalInfoRequestDTO `json:"additionalInfo"`
}

type CardRegistrationAdditionalInfoRequestDTO struct {
	Channel                string `json:"channel"`
	CustomerName           string `json:"customerName"`
	Email                  string `json:"email"`
	IdCard                 string `json:"idCard"`
	Country                string `json:"country"`
	Address                string `json:"address"`
	DateOfBirth            string `json:"dateOfBirth"`
	SuccessRegistrationUrl string `json:"successRegistrationUrl"`
	FailedRegistrationUrl  string `json:"failedRegistrationUrl"`
}

func (cr *CardRegistrationRequestDTO) ValidateCardRegistrationRequest() error {

	if !directDebitChannel.ValidateDirectDebitChannel(cr.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}
	return nil
}
