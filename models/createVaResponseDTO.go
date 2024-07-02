package models

type CreateVaResponseDto struct {
	ResponseCode       string             `json:"responseCode"`
	ResponseMessage    string             `json:"responseMessage"`
	VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
}
