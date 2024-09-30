package notification

import (
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type NotificationPaymentDirectDebitRequestDTO struct {
	OriginalPartnerReferenceNo string                                       `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo        string                                       `json:"originalReferenceNo"`
	OriginalExternalId         string                                       `json:"originalExternalId"`
	LatestTransactionStatus    string                                       `json:"latestTransactionStatus"`
	TransactionStatusDesc      string                                       `json:"transactionStatusDesc"`
	Amount                     createVaModels.TotalAmount                   `json:"amount"`
	AdditionalInfo             NotificationPaymentDirectDebitAdditionalInfo `json:"additionalInfo"`
}

type NotificationPaymentDirectDebitAdditionalInfo struct {
	ChannelId      string                       `json:"channelId"`
	AcquirerId     string                       `json:"acquirerId"`
	CustIdMerchant string                       `json:"custIdMerchant"`
	AccountType    string                       `json:"accountType"`
	LineItems      []paymentModels.LineItemsDTO `json:"LineItems"`
	Origin         createVaModels.Origin        `json:"origin"`
}
