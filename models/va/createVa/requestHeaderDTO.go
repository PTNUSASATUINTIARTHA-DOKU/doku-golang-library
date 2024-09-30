package models

import (
	"errors"
	"strings"

	directDebitUtils "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons/utils/directdebit"
)

type RequestHeaderDTO struct {
	XTimestamp            string `json:"xTimestamp"`
	XSignature            string `json:"xSignature"`
	XPartnerId            string `json:"xPartnerId"`
	XExternalId           string `json:"xExternalId"`
	XDeviceId             string `json:"xDeviceId,omitempty"`
	XIpAddress            string `json:"xIpAddress,omitempty"`
	ChannelId             string `json:"channelId"`
	Authorization         string `json:"authorization"`
	AuthorizationCustomer string `json:"authorizationCustomer"`
}

func (dto *RequestHeaderDTO) validateIpAddress(channel string) (bool, string) {
	if dto.XIpAddress == "" {
		return false, "Ip Address cannot be empty for this channel."
	} else if len(dto.XIpAddress) < 10 || len(dto.XIpAddress) > 15 {
		return false, "X-IP-ADDRESS must be in 10 to 15 characters."
	}

	return true, ""
}

func (dto *RequestHeaderDTO) validateDeviceId(channel string) (bool, string) {
	if dto.XDeviceId == "" {
		return false, "DeviceId cannot be empty for this channel."
	} else if len(dto.XDeviceId) > 64 {
		return false, "X-DEVICE-ID must be 64 characters or fewer. Ensure that X-DEVICE-ID is no longer than 64 characters."
	}

	return true, ""
}

func (dto *RequestHeaderDTO) ValidateAccountBinding(channel string) error {
	var validationErrors []string
	if channel == directDebitUtils.DirectDebitChannelNames[directDebitUtils.DIRECT_DEBIT_ALLO_SNAP] {
		if valid, message := dto.validateIpAddress(channel); !valid {
			validationErrors = append(validationErrors, message)
		}
		if valid, message := dto.validateDeviceId(channel); !valid {
			validationErrors = append(validationErrors, message)
		}
	}
	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "\n * "))
	}
	return nil
}

func (dto *RequestHeaderDTO) ValidateCheckBalance(channel string) error {
	if channel == directDebitUtils.DirectDebitChannelNames[directDebitUtils.DIRECT_DEBIT_ALLO_SNAP] {
		if valid, message := dto.validateIpAddress(channel); !valid {
			return errors.New(message)
		}
	}
	return nil
}

func (dto *RequestHeaderDTO) ValidatePaymentAndRefund(channel string) error {
	var validationErrors []string
	if channel == directDebitUtils.DirectDebitChannelNames[directDebitUtils.DIRECT_DEBIT_ALLO_SNAP] {
		if valid, message := dto.validateIpAddress(channel); !valid {
			validationErrors = append(validationErrors, message)
		}
	} else if channel == directDebitUtils.DirectDebitChannelNames[directDebitUtils.EMONEY_DANA_SNAP] || channel == directDebitUtils.DirectDebitChannelNames[directDebitUtils.EMONEY_SHOPEE_PAY_SNAP] {
		if valid, message := dto.validateIpAddress(channel); !valid {
			validationErrors = append(validationErrors, message)
		}
		if valid, message := dto.validateDeviceId(channel); !valid {
			validationErrors = append(validationErrors, message)
		}
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "\n * "))
	}
	return nil
}

func (dto *RequestHeaderDTO) ValidateAccountUnbinding(channel string) error {
	if channel == directDebitUtils.DirectDebitChannelNames[directDebitUtils.DIRECT_DEBIT_ALLO_SNAP] {
		if valid, message := dto.validateIpAddress(channel); !valid {
			return errors.New(message)
		}
	}
	return nil
}
