package refund

import (
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type RefundResponseDTO struct {
	ResponseCode               string                     `json:"responseCode"`
	ResponseMessage            string                     `json:"responseMessage"`
	RefundAmount               createVaModels.TotalAmount `json:"refundAmount"`
	OriginalPartnerReferenceNo string                     `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        string                     `json:"originalReferenceNo,omitempty"`
	RefundNo                   string                     `json:"refundNo,omitempty"`
	PartnerRefundNo            string                     `json:"partnerRefundNo,omitempty"`
	RefundTime                 string                     `json:"refundTime,omitempty"`
}
