package checkstatus

import (
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type CheckStatusResponseDTO struct {
	ResponseCode               string                                `json:"responseCode"`
	ResponseMessage            string                                `json:"responseMessage"`
	OriginalPartnerReferenceNo string                                `json:"originalPartnerReferenceNo,omitempty"`
	OriginalReferenceNo        string                                `json:"originalReferenceNo,omitempty"`
	ApprovalCode               string                                `json:"approvalCode,omitempty"`
	OriginalExternalId         string                                `json:"originalExternalId,omitempty"`
	ServiceCode                string                                `json:"serviceCode,omitempty"`
	LatestTransactionStatus    string                                `json:"latestTransactionStatus,omitempty"`
	TransactionStatusDesc      string                                `json:"transactionStatusDesc,omitempty"`
	OriginalResponseCode       string                                `json:"originalResponseCode,omitempty"`
	OriginalResponseMessage    string                                `json:"originalResponseMessage,omitempty"`
	SessionId                  string                                `json:"sessionId,omitempty"`
	RequestID                  string                                `json:"requestID,omitempty"`
	RefundHistory              []*RefundHistoryDTO                   `json:"refundHistory,omitempty"`
	TransAmount                *createVaModels.TotalAmount           `json:"transAmount,omitempty"`
	FeeAmount                  *createVaModels.TotalAmount           `json:"feeAmount,omitempty"`
	PaidTime                   string                                `json:"paidTime,omitempty"`
	AdditionalInfo             *CheckStatusAdditionalInfoResponseDTO `json:"additionalInfo,omitempty"`
}

type CheckStatusAdditionalInfoResponseDTO struct {
	DeviceId string      `json:"deviceId,omitempty"`
	Channel  string      `json:"channel,omitempty"`
	Acquirer interface{} `json:"acquirer,omitempty"`
}

type RefundHistoryDTO struct {
	RefundNo           string                      `json:"refundNo,omitempty"`
	PartnerReferenceNo string                      `json:"partnerReferenceNo,omitempty"`
	RefundAmount       *createVaModels.TotalAmount `json:"refundAmount,omitempty"`
	RefundStatus       string                      `json:"refundStatus,omitempty"`
	RefundDate         string                      `json:"refundDate,omitempty"`
	Reason             string                      `json:"reason,omitempty"`
}
