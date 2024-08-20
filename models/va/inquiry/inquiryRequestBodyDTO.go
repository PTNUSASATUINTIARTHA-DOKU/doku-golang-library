package inquiry

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
