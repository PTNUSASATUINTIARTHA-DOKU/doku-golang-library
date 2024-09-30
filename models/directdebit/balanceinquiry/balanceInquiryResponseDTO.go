package balanceinquiry

import (
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type BalanceInquiryResponseDto struct {
	ResponseCode    string            `json:"responseCode,omitempty"`
	ResponseMessage string            `json:"responseMessage,omitempty"`
	AccountInfos    []AccountInfosDto `json:"accountInfos,omitempty"`
}

type AccountInfosDto struct {
	BalanceType string                     `json:"balanceType"`
	Amount      createVaModels.TotalAmount `json:"amount"`
	FlatAmount  createVaModels.TotalAmount `json:"flatAmount"`
	HoldAmount  createVaModels.TotalAmount `json:"holdAmount"`
}
