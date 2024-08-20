package inquiry

import (
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type InquiryRequestVirtualAccountDataDTO struct {
	PartnerServiceId      string                           `json:"partnerServiceId"`
	CustomerNo            string                           `json:"customerNo"`
	VirtualAccountNo      string                           `json:"virtualAccountNo"`
	VirtualAccountName    string                           `json:"virtualAccountName"`
	VirtualAccountEmail   string                           `json:"virtualAccountEmail"`
	VirtualAccountPhone   string                           `json:"virtualAccountPhone"`
	TotalAmount           createVaModels.TotalAmount       `json:"totalAmount"`
	VirtualAccountTrxType string                           `json:"virtualAccountTrxType"`
	ExpiredDate           string                           `json:"expiredDate"`
	AdditionalInfo        InquiryResponseAdditionalInfoDTO `json:"additionalInfo"`
	InquiryStatus         string                           `json:"inquiryStatus"`
	InquiryReason         InquiryReasonDTO                 `json:"inquiryReason"`
	InquiryRequestId      string                           `json:"inquiryRequestId"`
}
