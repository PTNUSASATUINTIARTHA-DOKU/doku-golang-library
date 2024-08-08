package notification

type NotificationVirtualAccountData struct {
	PartnerServiceId   string `json:"partnerServiceId"`
	CustomerNo         string `json:"customerNo"`
	VirtualAccountNo   string `json:"virtualAccountNo"`
	VirtualAccountName string `json:"virtualAccountName"`
	PaymentRequestId   string `json:"paymentRequestId"`
}
