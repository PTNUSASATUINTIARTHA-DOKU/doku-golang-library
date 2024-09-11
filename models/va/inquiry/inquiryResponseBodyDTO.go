package inquiry

type InquiryResponseBodyDTO struct {
	ResponseCode       string                               `json:"responseCode"`
	ResponseMessage    string                               `json:"responseMessage"`
	VirtualAccountData *InquiryRequestVirtualAccountDataDTO `json:"virtualAccountData,omitempty"`
}
