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

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "1117"); valid {
		inquiryResponseBodyDto.ResponseCode = "2003000"
		inquiryResponseBodyDto.ResponseMessage = "Success"
		return true, inquiryResponseBodyDto
	}

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "111"); valid {
		inquiryResponseBodyDto.ResponseCode = "4012701"
		inquiryResponseBodyDto.ResponseMessage = "Access Token Invalid (B2B)"
		return true, inquiryResponseBodyDto
	}

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "112"); valid {
		inquiryResponseBodyDto.ResponseCode = "4012700"
		inquiryResponseBodyDto.ResponseMessage = "Unauthorized . Signature Not Match"
		return true, inquiryResponseBodyDto
	}

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "113"); valid {
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

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "114"); valid {
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

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "115"); valid {
		inquiryResponseBodyDto.ResponseCode = "4092700"
		inquiryResponseBodyDto.ResponseMessage = "Conflict"
		return true, inquiryResponseBodyDto
	}

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "116"); valid {
		inquiryResponseBodyDto.ResponseCode = "2002400"
		inquiryResponseBodyDto.ResponseMessage = "Success"
		return true, inquiryResponseBodyDto
	}

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "117"); valid {
		inquiryResponseBodyDto.ResponseCode = "4042414"
		inquiryResponseBodyDto.ResponseMessage = "Bill has been paid"
		return true, inquiryResponseBodyDto
	}

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "118"); valid {
		inquiryResponseBodyDto.ResponseCode = "4042419"
		inquiryResponseBodyDto.ResponseMessage = "Bill expired"
		return true, inquiryResponseBodyDto
	}

	if _, valid := strings.CutPrefix(dto.AdditionalInfo.TrxId, "119"); valid {
		inquiryResponseBodyDto.ResponseCode = "4042412"
		inquiryResponseBodyDto.ResponseMessage = "Bill not found"
		return true, inquiryResponseBodyDto
	}

	return false, inquiryResponseBodyDto
}
