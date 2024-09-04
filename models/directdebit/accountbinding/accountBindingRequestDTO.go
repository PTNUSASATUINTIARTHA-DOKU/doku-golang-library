package accountbinding

import (
	"errors"
	"strings"

	directDebitChannel "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils"
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
		return errors.New("additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'")
	}

	if dto.AdditionalInfo.Channel == directDebitChannel.DirectDebitChannelNames[directDebitChannel.DIRECT_DEBIT_ALLO_SNAP] {
		if dto.AdditionalInfo.DeviceModel == "" || dto.AdditionalInfo.OsType == "" || dto.AdditionalInfo.ChannelId == "" {
			return errors.New("value cannot be null for DIRECT_DEBIT_ALLO_SNAP")
		}

		if !strings.EqualFold(dto.AdditionalInfo.OsType, "ios") && !strings.EqualFold(dto.AdditionalInfo.OsType, "android") {
			return errors.New("osType value can only be ios/android")
		}

		if !strings.EqualFold(dto.AdditionalInfo.ChannelId, "app") && !strings.EqualFold(dto.AdditionalInfo.ChannelId, "web") {
			return errors.New("channelId value can only be app/web")
		}
	}

	return nil
}
