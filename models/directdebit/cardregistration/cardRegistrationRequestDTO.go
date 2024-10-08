package cardregistration

import (
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
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
		return errors.New("additionalInfo.channel is not valid. Ensure it is one of the valid channels like 'DIRECT_DEBIT_ALLO_SNAP'")
	}

	if err := validateCustIdMerchant(cr.CustIdMerchant); err != nil {
		return err
	}

	if err := cr.validateAdditionalInfo(); err != nil {
		return err
	}

	if err := validateCardData(cr.CardData); err != nil {
		return err
	}

	return nil
}

func validateCustIdMerchant(custIdMerchant string) error {
	if custIdMerchant == "" {
		return errors.New("custIdMerchant cannot be null. Please provide a custIdMerchant. Example: 'cust-001'")
	}
	if len(custIdMerchant) > 64 {
		return errors.New("custIdMerchant must be 64 characters or fewer")
	}
	return nil
}

func (cr *CardRegistrationRequestDTO) validateAdditionalInfo() error {
	if cr.AdditionalInfo.SuccessRegistrationUrl == "" {
		return errors.New("additionalInfo.SuccessRegistrationUrl cannot be null. Example: 'https://www.doku.com'")
	}
	if cr.AdditionalInfo.FailedRegistrationUrl == "" {
		return errors.New("additionalInfo.FailedRegistrationUrl cannot be null. Example: 'https://www.doku.com'")
	}
	return nil
}

func validateCardData(cardData string) error {
	if cardData == "" {
		return errors.New("cardData cannot be null. Please provide cardData. Example: '5cg2G2719+jxU1RfcGmeCyQrLagUaAWJWWhLpmmb'")
	}
	return nil
}
