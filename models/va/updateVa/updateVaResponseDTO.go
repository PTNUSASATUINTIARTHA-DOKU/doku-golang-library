package models

type UpdateVaResponseDTO struct {
	ResponseCode       string      `json:"responseCode"`
	ResponseMessage    string      `json:"responseMessage"`
	VirtualAccountData UpdateVaDTO `json:"virtualAccountData,omitempty"`
}
