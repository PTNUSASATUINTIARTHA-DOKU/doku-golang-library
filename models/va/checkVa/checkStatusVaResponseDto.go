package models

type CheckStatusVaResponseDto struct {
	ResponseCode       string                            `json:"responseCode"`
	ResponseMessage    string                            `json:"responseMessage"`
	VirtualAccountData *CheckStatusVirtualAccountData    `json:"virtualAccountData,omitempty"`
	AdditionalInfo     CheckStatusResponseAdditionalInfo `json:"additionalInfo"`
}
