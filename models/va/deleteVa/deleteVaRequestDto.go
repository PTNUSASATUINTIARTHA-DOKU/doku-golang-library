package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
)

type DeleteVaRequestDto struct {
	PartnerServiceId string                        `json:"partnerServiceId"`
	CustomerNo       string                        `json:"customerNo"`
	VirtualAccountNo string                        `json:"virtualAccountNo"`
	TrxId            string                        `json:"trxId"`
	AdditionalInfo   DeleteVaRequestAdditionalInfo `json:"additionalInfo"`
}

func (dto *DeleteVaRequestDto) ValidateDeleteVaRequest() error {

	var validationErrors []string

	if valid, message := dto.validatePartnerServiceId(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateCustomerNo(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateVirtualAccountNo(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateTrxId(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateChannel(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if len(validationErrors) > 0 {
		return errors.New("Validation Failed: \n * " + strings.Join(validationErrors, "\n * "))
	}
	return nil
}

func (dto *DeleteVaRequestDto) validatePartnerServiceId() (bool, string) {
	if len(dto.PartnerServiceId) != 8 {
		return false, "PartnerServiceId must be exactly 8 characters long and equiped with left-padded spaces. Example: '  888994'."
	}
	if !regexp.MustCompile(`^\s{0,7}\d{1,8}$`).MatchString(dto.PartnerServiceId) {
		return false, "PartnerServiceId must consist of up to 8 digits of character. Remaining space in case of partner serivce id is less than 8 must be filled with spaces. Example: ' 888994' (2 spaces and 6 digits)."
	}
	return true, ""
}

func (dto *DeleteVaRequestDto) validateCustomerNo() (bool, string) {
	if dto.CustomerNo == "" {
		return false, "CustomerNo cannot be null. Please provide a CustomerNo. Example: '00000000000000000001'."
	}
	if len(dto.CustomerNo) > 20 {
		return false, "CustomerNo must be 20 characters or fewer. Ensure that customerNo is no longer than 20 characters. Example: '00000000000000000001'."
	}
	if !regexp.MustCompile(`^\d+$`).MatchString(dto.CustomerNo) {
		return false, "CustomerNo must consist of only digits. Ensure that customerNo contains only numbers. Example: '00000000000000000001'."
	}
	return true, ""
}

func (dto *DeleteVaRequestDto) validateVirtualAccountNo() (bool, string) {
	if dto.VirtualAccountNo == "" {
		return false, "VirtualAccountNo cannot be null. Please provide a virtualAccountNo. Example: '  88899400000000000000000001'."
	}
	if dto.VirtualAccountNo != dto.PartnerServiceId+dto.CustomerNo {
		return false, "VirtualAccountNo must be the concatenation of partnerServiceId and customerNo. Example: ' 88899400000000000000000001' (where partnerServiceId is ' 888994' and customerNo is '00000000000000000001')."
	}
	return true, ""
}

func (dto *DeleteVaRequestDto) validateTrxId() (bool, string) {
	if len(dto.TrxId) < 1 {
		return false, "TrxId must be at least 1 character long. Ensure that TrxId is not empty. Example: '23219829713'."
	}
	if len(dto.TrxId) > 64 {
		return false, "TrxId must be 64 characters or fewer. Ensure that TrxId is no longer than 64 characters. Example: '23219829713'."
	}
	return true, ""
}

func (dto *DeleteVaRequestDto) validateChannel() (bool, string) {
	if len(dto.AdditionalInfo.Channel) < 1 {
		return false, "AdditionalInfo.Channel must be at least 1 character long. Ensure that AdditionalInfo.Channel is not empty. Example: 'VIRTUAL_ACCOUNT_MANDIRI'."
	}
	if len(dto.AdditionalInfo.Channel) > 30 {
		return false, "AdditionalInfo.Channel must be 30 characters or fewer. Ensure that AdditionalInfo.Channel is no longer than 30 characters. Example: 'VIRTUAL_ACCOUNT_MANDIRI'."
	}
	if !commons.ValidateVAChannel(dto.AdditionalInfo.Channel) {
		return false, "AdditionalInfo.channel is not valid. Ensure that AdditionalInfo.channel is one of the valid channels. Example: 'VIRTUAL_ACCOUNT_MANDIRI'."
	}
	return true, ""
}
