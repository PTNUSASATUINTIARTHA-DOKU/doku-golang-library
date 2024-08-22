package notification

import createVa "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"

type NotificationVirtualAccountData struct {
	PartnerServiceId   string `json:"partnerServiceId"`
	CustomerNo         string `json:"customerNo"`
	VirtualAccountNo   string `json:"virtualAccountNo"`
	VirtualAccountName string `json:"virtualAccountName"`
	PaymentRequestId   string `json:"paymentRequestId"`
	AdditionalInfo     createVa.AdditionalInfo
}
