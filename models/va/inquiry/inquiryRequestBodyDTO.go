package inquiry

import "strings"

type InquiryRequestBodyDTO struct {
	PartnerServiceId string                          `json:"partnerServiceId"`
	CustomerNo       string                          `json:"customerNo"`
	VirtualAccountNo string                          `json:"virtualAccountNo"`
	ChannelCode      string                          `json:"channelCode"`
	TrxDateInit      string                          `json:"trxDateInit"`
	Language         string                          `json:"language"`
	InquiryRequestId string                          `json:"inquiryRequestId"`
	AdditionalInfo   InquiryRequestAdditionalInfoDTO `json:"additionalInfo"`
}

func (dto *InquiryRequestBodyDTO) ValidateSimulatorASPI() (bool, InquiryResponseBodyDTO) {
	var inquiryResponseBodyDto InquiryResponseBodyDTO

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "1117"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "1117")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "2003000"
		inquiryResponseBodyDto.ResponseMessage = "Success"
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "111"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "111")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "4012701"
		inquiryResponseBodyDto.ResponseMessage = "Access Token Invalid (B2B)"
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "112"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "112")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "4012700"
		inquiryResponseBodyDto.ResponseMessage = "Unauthorized . Signature Not Match"
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "113"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "113")
		return validVaNo
	}() {
		var vaData InquiryRequestVirtualAccountDataDTO
		vaData.PartnerServiceId = "90341537"
		vaData.CustomerNo = "00000000000000000000"
		vaData.VirtualAccountNo = "0000000000000000000000000000"
		vaData.AdditionalInfo.TrxId = "PGPWF123"

		inquiryResponseBodyDto.ResponseCode = "4002702"
		inquiryResponseBodyDto.ResponseMessage = "Invalid Mandatory Field partnerServiceId"
		inquiryResponseBodyDto.VirtualAccountData = &vaData
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "114"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "114")
		return validVaNo
	}() {
		var vaData InquiryRequestVirtualAccountDataDTO
		vaData.PartnerServiceId = "90341537"
		vaData.CustomerNo = "00000000000000000000"
		vaData.VirtualAccountNo = "0000000000000000000000000000"
		vaData.AdditionalInfo.TrxId = "PGPWF123"

		inquiryResponseBodyDto.ResponseCode = "4002701"
		inquiryResponseBodyDto.ResponseMessage = "Invalid Field Format totalAmount.currency"
		inquiryResponseBodyDto.VirtualAccountData = &vaData
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "115"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "115")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "4092700"
		inquiryResponseBodyDto.ResponseMessage = "Conflict"
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "116"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "116")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "2002400"
		inquiryResponseBodyDto.ResponseMessage = "Success"
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "117"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "117")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "4042414"
		inquiryResponseBodyDto.ResponseMessage = "Bill has been paid"
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "118"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "118")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "4042419"
		inquiryResponseBodyDto.ResponseMessage = "Bill expired"
		return true, inquiryResponseBodyDto
	}

	if _, validTrxId := strings.CutPrefix(dto.AdditionalInfo.TrxId, "119"); validTrxId || func() bool {
		_, validVaNo := strings.CutPrefix(dto.VirtualAccountNo, "119")
		return validVaNo
	}() {
		inquiryResponseBodyDto.ResponseCode = "4042412"
		inquiryResponseBodyDto.ResponseMessage = "Bill not found"
		return true, inquiryResponseBodyDto
	}

	return false, inquiryResponseBodyDto
}
