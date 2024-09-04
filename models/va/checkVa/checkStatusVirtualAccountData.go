package models

import (
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type CheckStatusVirtualAccountData struct {
	PaymentFlagReason    CheckStatusResponsePaymentFlagReason `json:"paymentFlagReason"`
	PartnerServiceId     string                               `json:"partnerServiceId"`
	CustomerNo           string                               `json:"customerNo"`
	VirtualAccountNo     string                               `json:"virtualAccountNo"`
	TrxId                string                               `json:"trxId"`
	InquiryRequestId     *string                              `json:"inquiryRequestId"`
	PaymentRequestId     *string                              `json:"paymentRequestId"`
	VirtualAccountNumber *string                              `json:"virtualAccountNumber"`
	PaidAmount           createVaModels.TotalAmount           `json:"paidAmount"`
	BillDetails          []CheckStatusBillDetail              `json:"billDetails"`
	AdditionalInfo       CheckStatusResponseAdditionalInfo    `json:"additionalInfo"`
}

type CheckStatusBillDetail struct {
	BillAmount createVaModels.TotalAmount `json:"billAmount"`
}
