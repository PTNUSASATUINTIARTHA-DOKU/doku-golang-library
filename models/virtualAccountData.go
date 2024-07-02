package models

type VirtualAccountData struct {
	PartnerServiceId    string                 `json:"partnerServiceId"`
	CustomerNo          string                 `json:"customerNo"`
	VirtualAccountNo    string                 `json:"virtualAccountNo"`
	VirtualAccountName  string                 `json:"virtualAccountName"`
	VirtualAccountEmail string                 `json:"virtualAccountEmail"`
	TrxId               string                 `json:"trxId"`
	TotalAmount         TotalAmount            `json:"totalAmount"`
	AdditionalInfo      AdditionalInfoResponse `json:"additionalInfo"`
}
