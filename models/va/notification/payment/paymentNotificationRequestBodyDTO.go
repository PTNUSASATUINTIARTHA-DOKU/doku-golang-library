package payment

import va "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"

type PaymentNotificationRequestBodyDTO struct {
	PartnerServiceId    string                                      `json:"partnerServiceId"`
	CustomerNo          string                                      `json:"customerNo"`
	VirtualAccountNo    string                                      `json:"virtualAccountNo"`
	VirtualAccountName  string                                      `json:"virtualAccountName"`
	TrxId               string                                      `json:"trxId"`
	PaymentRequestId    string                                      `json:"paymentRequestId"`
	PaidAmount          va.TotalAmount                              `json:"paidAmount"`
	VirtualAccountEmail string                                      `json:"virtualAccountEmail"`
	VirtualAccountPhone string                                      `json:"virtualAccountPhone"`
	AdditionalInfo      PaymentNotificationRequestAdditionalInfoDTO `json:"additionalInfo"`
}

type PaymentNotificationRequestAdditionalInfoDTO struct {
	Channel              string                  `json:"channel"`
	VirtualAccountConfig va.VirtualAccountConfig `json:"virtualAccountConfig"`
	SenderName           string                  `json:"senderName"`
	SourceAccountNo      string                  `json:"sourceAccountNo"`
	SourceBankCode       string                  `json:"sourceBankCode"`
	SourceBankName       string                  `json:"sourceBankName"`
}
