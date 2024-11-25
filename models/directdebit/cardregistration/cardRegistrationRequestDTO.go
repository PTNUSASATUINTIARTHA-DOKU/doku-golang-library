package cardregistration

import (
	"encoding/json"
	"errors"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type CardRegistrationRequestDTO struct {
	CardData       interface{}                              `json:"cardData"`
	CustIdMerchant string                                   `json:"custIdMerchant"`
	PhoneNo        string                                   `json:"phoneNo"`
	AdditionalInfo CardRegistrationAdditionalInfoRequestDTO `json:"additionalInfo"`
}

type CardRegistrationAdditionalInfoRequestDTO struct {
	Channel                string                `json:"channel"`
	CustomerName           string                `json:"customerName"`
	Email                  string                `json:"email"`
	IdCard                 string                `json:"idCard"`
	Country                string                `json:"country"`
	Address                string                `json:"address"`
	DateOfBirth            string                `json:"dateOfBirth"`
	SuccessRegistrationUrl string                `json:"successRegistrationUrl"`
	FailedRegistrationUrl  string                `json:"failedRegistrationUrl"`
	Origin                 createVaModels.Origin `json:"origin"`
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

	if err := cr.validateCardData(); err != nil {
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

func (cr *CardRegistrationRequestDTO) validateCardData() error {
	if cr.CardData == nil {
		return errors.New("cardData cannot be null or empty")
	}

	switch cardData := cr.CardData.(type) {
	case string:
		if cardData == "" {
			return errors.New("cardData is an empty string")
		}
	case BankCardDataDTO:
		if cardData.BankCardNo == "" || cardData.BankCardType == "" || cardData.ExpiryDate == "" {
			return errors.New("bank card data fields cannot be empty")
		}
	case map[string]interface{}:
		var bankCardData BankCardDataDTO
		cardDataJSON, err := json.Marshal(cardData)
		if err != nil {
			return errors.New("unable to marshal cardData")
		}

		err = json.Unmarshal(cardDataJSON, &bankCardData)
		if err != nil {
			return errors.New("cardData is of an unsupported type or has invalid fields")
		}

		if bankCardData.BankCardNo == "" || bankCardData.BankCardType == "" || bankCardData.ExpiryDate == "" {
			return errors.New("bank card data fields cannot be empty")
		}
	default:
		return errors.New("cardData is of an unsupported type")
	}

	return nil
}
