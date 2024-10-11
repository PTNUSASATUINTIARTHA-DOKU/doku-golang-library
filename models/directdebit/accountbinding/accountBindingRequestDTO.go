package accountbinding

import (
	"errors"
	"strings"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
)

type AccountBindingRequestDTO struct {
	PhoneNo        string                                 `json:"phoneNo"`
	AdditionalInfo AccountBindingAdditionalInfoRequestDto `json:"additionalInfo"`
}

type AccountBindingAdditionalInfoRequestDto struct {
	Channel                string `json:"channel"`
	CustIdMerchant         string `json:"custIdMerchant"`
	CustomerName           string `json:"customerName"`
	Email                  string `json:"email"`
	IdCard                 string `json:"idCard"`
	Country                string `json:"country"`
	Address                string `json:"address"`
	DateOfBirth            string `json:"dateOfBirth"`
	SuccessRegistrationUrl string `json:"successRegistrationUrl"`
	FailedRegistrationUrl  string `json:"failedRegistrationUrl"`
	DeviceModel            string `json:"deviceModel"`
	OsType                 string `json:"osType"`
	ChannelId              string `json:"channelId"`
}

func (dto *AccountBindingRequestDTO) ValidateAccountBindingRequest() error {

	if !directDebitChannel.ValidateDirectDebitChannel(dto.AdditionalInfo.Channel) {
		return errors.New("additionalInfo.channel is not valid. Ensure it is one of the valid channels like 'DIRECT_DEBIT_ALLO_SNAP'")
	}

	if dto.AdditionalInfo.Channel == directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_ALLO_SNAP] {
		if err := validateDirectDebitAlloSnap(dto.AdditionalInfo); err != nil {
			return err
		}
	}

	if err := validatePhoneNo(dto.PhoneNo); err != nil {
		return err
	}

	if err := validateCustIdMerchant(dto.AdditionalInfo.CustIdMerchant); err != nil {
		return err
	}

	if err := dto.validateAdditionalInfo(); err != nil {
		return err
	}

	return nil
}

func validatePhoneNo(phoneNo string) error {
	if phoneNo == "" {
		return errors.New("phoneNo cannot be null. Please provide a phoneNo. Example: '62813941306101'")
	}
	if len(phoneNo) < 9 {
		return errors.New("phoneNo must be at least 9 digits. Example: '62813941306101'")
	}
	if len(phoneNo) > 16 {
		return errors.New("phoneNo must be 16 characters or fewer. Example: '62813941306101'")
	}
	return nil
}

func validateCustIdMerchant(custIdMerchant string) error {
	if custIdMerchant == "" {
		return errors.New("additionalInfo.custIdMerchant cannot be null. Example: 'cust-001'")
	}
	if len(custIdMerchant) > 64 {
		return errors.New("additionalInfo.custIdMerchant must be 64 characters or fewer. Example: 'cust-001'")
	}
	return nil
}

func (dto *AccountBindingRequestDTO) validateAdditionalInfo() error {
	if dto.AdditionalInfo.SuccessRegistrationUrl == "" {
		return errors.New("additionalInfo.SuccessRegistrationUrl cannot be null. Example: 'https://www.doku.com'")
	}
	if dto.AdditionalInfo.FailedRegistrationUrl == "" {
		return errors.New("additionalInfo.FailedRegistrationUrl cannot be null. Example: 'https://www.doku.com'")
	}
	return nil
}

func validateDirectDebitAlloSnap(additionalInfo AccountBindingAdditionalInfoRequestDto) error {
	if additionalInfo.DeviceModel == "" || additionalInfo.OsType == "" || additionalInfo.ChannelId == "" {
		return errors.New("DeviceModel, OsType, and ChannelId cannot be null for DIRECT_DEBIT_ALLO_SNAP")
	}

	if !strings.EqualFold(additionalInfo.OsType, "ios") && !strings.EqualFold(additionalInfo.OsType, "android") {
		return errors.New("osType value can only be 'ios' or 'android'")
	}

	if !strings.EqualFold(additionalInfo.ChannelId, "app") && !strings.EqualFold(additionalInfo.ChannelId, "web") {
		return errors.New("channelId value can only be 'app' or 'web'")
	}

	return nil
}
