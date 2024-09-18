package models

import (
	"errors"
	"regexp"
	"strings"
)

type CheckStatusVARequestDto struct {
	PartnerServiceId   string  `json:"partnerServiceId"`
	CustomerNo         string  `json:"customerNo"`
	VirtualAccountNo   string  `json:"virtualAccountNo"`
	VirtualAccountName *string `json:"virtualAccountName,omitempty"`
	InquiryRequestId   *string `json:"inquiryRequestId,omitempty"`
	PaymentRequestId   *string `json:"paymentRequestId,omitempty"`
	AdditionalInfo     *string `json:"additionalInfo,omitempty"`
}

func (dto *CheckStatusVARequestDto) ValidateCheckStatusVaRequestDto() error {

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

	if dto.InquiryRequestId != nil {
		if valid, message := dto.validateInquiryRequestId(); !valid {
			validationErrors = append(validationErrors, message)
		}
	}

	if dto.PaymentRequestId != nil {
		if valid, message := dto.validatePaymentRequestId(); !valid {
			validationErrors = append(validationErrors, message)
		}
	}

	if len(validationErrors) > 0 {
		return errors.New("Validation Failed: \n * " + strings.Join(validationErrors, "\n * "))
	}
	return nil
}

func (dto *CheckStatusVARequestDto) validatePartnerServiceId() (bool, string) {
	if len(dto.PartnerServiceId) != 8 {
		return false, "PartnerServiceId must be exactly 8 characters long and equiped with left-padded spaces. Example: '  888994'."
	}
	if !regexp.MustCompile(`^\s{0,7}\d{1,8}$`).MatchString(dto.PartnerServiceId) {
		return false, "PartnerServiceId must consist of up to 8 digits of character. Remaining space in case of partner serivce id is less than 8 must be filled with spaces. Example: ' 888994' (2 spaces and 6 digits)."
	}
	return true, ""
}

func (dto *CheckStatusVARequestDto) validateCustomerNo() (bool, string) {
	if len(dto.CustomerNo) > 20 {
		return false, "CustomerNo must be 20 characters or fewer. Ensure that customerNo is no longer than 20 characters. Example: '00000000000000000001'."
	}
	if !regexp.MustCompile(`^\d+$`).MatchString(dto.CustomerNo) {
		return false, "CustomerNo must consist of only digits. Ensure that customerNo contains only numbers. Example: '00000000000000000001'."
	}
	return true, ""
}

func (dto *CheckStatusVARequestDto) validateVirtualAccountNo() (bool, string) {
	if dto.VirtualAccountNo != dto.PartnerServiceId+dto.CustomerNo {
		return false, "VirtualAccountNo must be the concatenation of partnerServiceId and customerNo. Example: ' 88899400000000000000000001' (where partnerServiceId is ' 888994' and customerNo is '00000000000000000001')."
	}
	return true, ""
}

func (dto *CheckStatusVARequestDto) validateInquiryRequestId() (bool, string) {
	if len(*dto.InquiryRequestId) > 128 {
		return false, "InquiryRequestId must be 128 characters or fewer. Ensure that InquiryRequestId is no longer than 128 characters. Example: 'abcdef-123456-abcdef'."
	}
	return true, ""
}

func (dto *CheckStatusVARequestDto) validatePaymentRequestId() (bool, string) {
	if len(*dto.PaymentRequestId) > 128 {
		return false, "PaymentRequestId must be 128 characters or fewer. Ensure that PaymentRequestId is no longer than 128 characters. Example: 'abcdef-123456-abcdef'."
	}
	return true, ""
}

func (dto *CheckStatusVARequestDto) ValidateSimulatorASPI() (bool, CheckStatusVaResponseDto) {
	var checkStatusResponseDto CheckStatusVaResponseDto

	if _, valid := strings.CutPrefix(dto.VirtualAccountNo, "1113"); valid {
		checkStatusResponseDto.ResponseCode = "2002600"
		checkStatusResponseDto.ResponseMessage = "Success"
		return true, checkStatusResponseDto
	}

	if _, valid := strings.CutPrefix(dto.VirtualAccountNo, "111"); valid {
		checkStatusResponseDto.ResponseCode = "4012701"
		checkStatusResponseDto.ResponseMessage = "Access Token Invalid (B2B)"
		return true, checkStatusResponseDto
	}

	if _, valid := strings.CutPrefix(dto.VirtualAccountNo, "112"); valid {
		checkStatusResponseDto.ResponseCode = "4012700"
		checkStatusResponseDto.ResponseMessage = "Unauthorized . Signature Not Match"
		return true, checkStatusResponseDto
	}

	if _, valid := strings.CutPrefix(dto.VirtualAccountNo, "113"); valid {
		var vaData CheckStatusVirtualAccountData
		vaData.PartnerServiceId = "90341537"
		vaData.CustomerNo = "00000000000000000000"
		vaData.VirtualAccountNo = "0000000000000000000000000000"
		vaData.TrxId = "PGPWF123"

		checkStatusResponseDto.ResponseCode = "4002702"
		checkStatusResponseDto.ResponseMessage = "Invalid Mandatory Field partnerServiceId"
		checkStatusResponseDto.VirtualAccountData = &vaData
		return true, checkStatusResponseDto
	}

	if _, valid := strings.CutPrefix(dto.VirtualAccountNo, "114"); valid {
		var vaData CheckStatusVirtualAccountData
		vaData.PartnerServiceId = "90341537"
		vaData.CustomerNo = "00000000000000000000"
		vaData.VirtualAccountNo = "0000000000000000000000000000"
		vaData.TrxId = "PGPWF123"

		checkStatusResponseDto.ResponseCode = "4002701"
		checkStatusResponseDto.ResponseMessage = "Invalid Field Format totalAmount.currency"
		checkStatusResponseDto.VirtualAccountData = &vaData
		return true, checkStatusResponseDto
	}

	if _, valid := strings.CutPrefix(dto.VirtualAccountNo, "115"); valid {
		checkStatusResponseDto.ResponseCode = "4092700"
		checkStatusResponseDto.ResponseMessage = "Conflict"
		return true, checkStatusResponseDto
	}

	return false, checkStatusResponseDto
}
