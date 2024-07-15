package models

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type UpdateVaDTO struct {
	PartnerServiceId      string                     `json:"partnerServiceId"`
	CustomerNo            *string                    `json:"customerNo,omitempty"`
	VirtualAccountNo      *string                    `json:"virtualAccountNo,omitempty"`
	VirtualAccountName    string                     `json:"virtualAccountName"`
	VirtualAccountEmail   string                     `json:"virtualAccountEmail"`
	VirtualAccountPhone   string                     `json:"virtualAccountPhone"`
	TrxId                 string                     `json:"trxId"`
	TotalAmount           createVaModels.TotalAmount `json:"totalAmount"`
	AdditionalInfo        UpdateVaAdditionalInfoDTO  `json:"additionalInfo"`
	VirtualAccountTrxType string                     `json:"virtualAccounTrxType"`
	ExpiredDate           string                     `json:"expiredDate"`
}

func (dto *UpdateVaDTO) ValidateUpdateVaRequestDTO() {

	var validationErrors []string

	if valid, message := dto.validatePartnerServiceId(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if dto.CustomerNo != nil {
		if valid, message := dto.validateCustomerNo(); !valid {
			validationErrors = append(validationErrors, message)
		}
		if valid, message := dto.validateVirtualAccountNo(); !valid {
			validationErrors = append(validationErrors, message)
		}
	}

	if valid, message := dto.validateVirtualAccountName(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateVirtualAccountEmail(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateVirtualAccountPhone(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateTrxId(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateValue(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateCurrency(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateChannel(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateStatus(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateVirtualAccountTrxType(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if valid, message := dto.validateExpiredDate(); !valid {
		validationErrors = append(validationErrors, message)
	}

	if len(validationErrors) > 0 {
		panic(errors.New("Validation Failed: \n * " + strings.Join(validationErrors, "\n * ")))
	}
}

func (dto *UpdateVaDTO) validatePartnerServiceId() (bool, string) {
	if len(dto.PartnerServiceId) != 8 {
		return false, "PartnerServiceId must be exactly 8 characters long and equiped with left-padded spaces. Example: '  888994'."
	}
	if !regexp.MustCompile(`^\s{0,7}\d{1,8}$`).MatchString(dto.PartnerServiceId) {
		return false, "PartnerServiceId must consist of up to 8 digits of character. Remaining space in case of partner serivce id is less than 8 must be filled with spaces. Example: ' 888994' (2 spaces and 6 digits)."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateCustomerNo() (bool, string) {
	if len(*dto.CustomerNo) > 20 {
		return false, "CustomerNo must be 20 characters or fewer. Ensure that customerNo is no longer than 20 characters. Example: '00000000000000000001'."
	}
	if !regexp.MustCompile(`^\d+$`).MatchString(*dto.CustomerNo) {
		return false, "CustomerNo must consist of only digits. Ensure that customerNo contains only numbers. Example: '00000000000000000001'."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateVirtualAccountNo() (bool, string) {
	customerNo := *dto.CustomerNo
	virtualAccountNo := *dto.VirtualAccountNo
	if dto.VirtualAccountNo == nil {
		return false, "VirtualAccountNo cannot be null. Please provide a virtualAccountNo. Example: '  88899400000000000000000001'."
	}
	if virtualAccountNo != dto.PartnerServiceId+customerNo {
		return false, "VirtualAccountNo must be the concatenation of partnerServiceId and customerNo. Example: ' 88899400000000000000000001' (where partnerServiceId is ' 888994' and customerNo is '00000000000000000001')."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateVirtualAccountName() (bool, string) {
	if len(dto.VirtualAccountName) < 1 {
		return false, "VirtualAccountName must be at least 1 character long. Ensure that virtualAccountName is not empty. Example: 'Toru Yamashita'."
	}
	if len(dto.VirtualAccountName) > 255 {
		return false, "VirtualAccountName must be 255 characters or fewer. Ensure that virtualAccountName is no longer than 255 characters. Example: 'Toru Yamashita'."
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9.\\\-\/+,=_:'@% ]*$`).MatchString(dto.VirtualAccountName) {
		return false, "VirtualAccountName can only contain letters, numbers, spaces, and the following characters: .\\-/+,=_:'@%. Ensure that virtualAccountName does not contain invalid characters. Example: 'Toru.Yamashita-123'."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateVirtualAccountEmail() (bool, string) {
	if len(dto.VirtualAccountEmail) < 1 {
		return false, "VirtualAccountEmail must be at least 1 character long. Ensure that VirtualAccountEmail is not empty. Example: 'toru@example.com'."
	}
	if len(dto.VirtualAccountEmail) > 255 {
		return false, "VirtualAccountEmail must be 255 characters or fewer. Ensure that VirtualAccountEmail is no longer than 255 characters. Example: 'toru@example.com'."
	}
	if !regexp.MustCompile(`^[\w-]+@([\w-]+\.)+[\w-]{2,4}$`).MatchString(dto.VirtualAccountEmail) {
		return false, "VirtualAccountEmail is not in a valid email format. Ensure it contains an '@' symbol followed by a domain name. Example: 'toru@example.com'."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateVirtualAccountPhone() (bool, string) {
	if len(dto.VirtualAccountPhone) < 9 {
		return false, "VirtualAccountPhone must be at least 9 characters long. Ensure that VirtualAccountPhone is at least 9 characters long. Example: '628123456789'."
	}
	if len(dto.VirtualAccountPhone) > 30 {
		return false, "virtualAccountPhone must be 30 characters or fewer. Ensure that virtualAccountPhone is no longer than 30 characters. Example: '628123456789012345678901234567'."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateTrxId() (bool, string) {
	if len(dto.TrxId) < 1 {
		return false, "TrxId must be at least 1 character long. Ensure that TrxId is not empty. Example: '23219829713'."
	}
	if len(dto.TrxId) > 64 {
		return false, "TrxId must be 64 characters or fewer. Ensure that TrxId is no longer than 64 characters. Example: '23219829713'."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateValue() (bool, string) {
	if len(dto.TotalAmount.Value) < 4 {
		return false, "TotalAmount.Value must be at least 4 characters long and formatted as 0.00. Ensure that TotalAmount.Value is at least 4 characters long and in the correct format. Example: '100.00'."
	}
	if len(dto.TotalAmount.Value) > 19 {
		return false, "TotalAmount.Value must be 19 characters or fewer and formatted as 9999999999999999.99. Ensure that TotalAmount.Value is no longer than 19 characters and in the correct format. Example: '9999999999999999.99'."
	}
	if !regexp.MustCompile(`^(0|[1-9]\d{0,15})(\.\d{2})?$`).MatchString(dto.TotalAmount.Value) {
		return false, "TotalAmount.Value is invalid format."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateCurrency() (bool, string) {
	if dto.TotalAmount.Currency == "" {
		return false, "TotalAmount.Currency must be a string. Ensure that TotalAmount.Currency is enclosed in quotes. Example: 'IDR'."
	}
	if len(dto.TotalAmount.Currency) != 3 {
		return false, "TotalAmount.Currency must be exactly 3 characters long. Ensure that TotalAmount.Currency is exactly 3 characters. Example: 'IDR'."
	}
	if dto.TotalAmount.Currency != "IDR" {
		return false, "TotalAmount.currency must be 'IDR'. Ensure that TotalAmount.Currency is 'IDR'. Example: 'IDR'."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateChannel() (bool, string) {
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

func (dto *UpdateVaDTO) validateStatus() (bool, string) {
	if len(dto.AdditionalInfo.VirtualAccountConfig.Status) <= 1 {
		return false, "status must be at least 1 character long. Ensure that status is not empty. Example: ‘INACTIVE’."
	}
	if len(dto.AdditionalInfo.VirtualAccountConfig.Status) >= 20 {
		return false, "status must be 20 characters or fewer. Ensure that status is no longer than 20 characters. Example: ‘INACTIVE’."
	}
	if !(dto.AdditionalInfo.VirtualAccountConfig.Status == "ACTIVE" || dto.AdditionalInfo.VirtualAccountConfig.Status == "INACTIVE") {
		return false, "status must be either 'ACTIVE' or 'INACTIVE'. Ensure that status is one of these values. Example: ‘INACTIVE’."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateVirtualAccountTrxType() (bool, string) {
	if len(dto.VirtualAccountTrxType) != 1 {
		return false, "VirtualAccountTrxType must be exactly 1 character long. Ensure that VirtualAccountTrxType is either '1' or '2'. Example: '1'."
	}
	if !(dto.VirtualAccountTrxType == "1" || dto.VirtualAccountTrxType == "2") {
		return false, "VirtualAccountTrxType must be either '1' or '2'. Ensure that VirtualAccountTrxType is one of these values. Example: '1'."
	}
	return true, ""
}

func (dto *UpdateVaDTO) validateExpiredDate() (bool, string) {
	_, err := time.Parse("2006-01-02T15:04:05+07:00", dto.ExpiredDate)
	if err != nil {
		return false, "ExpiredDate must be in ISO-8601 format. Ensure that ExpiredDate follows the correct format. Example: '2023-01-01T10:55:00+07:00'."
	}
	return true, ""
}
