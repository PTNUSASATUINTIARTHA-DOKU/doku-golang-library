package models

type DeleteVaResponseDto struct {
	ResponseCode       string                              `json:"responseCode"`
	ResponseMessage    string                              `json:"responseMessage"`
	VirtualAccountData *DeleteVaResponseVirtualAccountData `json:"virtualAccountData"`
}
