package models

type DeleteVaResponseVirtualAccountData struct {
	PartnerServiceId string                         `json:"partnerServiceId"`
	CustomerNo       string                         `json:"customerNo"`
	VirtualAccountNo string                         `json:"virtualAccountNo"`
	TrxId            string                         `json:"trxId"`
	AdditionalInfo   DeleteVaResponseAdditionalInfo `json:"additionalInfo"`
}
