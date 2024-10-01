package notification

type NotificationPaymentDirectDebitHeaderRequestDTO struct {
	OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
}

type RequestHeaderDTO struct {
	XTimestamp            string `json:"xTimestamp"`
	XSignature            string `json:"xSignature"`
	XPartnerId            string `json:"xPartnerId"`
	XExternalId           string `json:"xExternalId"`
	XDeviceId             string `json:"xDeviceId"`
	XIpAddress            string `json:"xIpAddress,omitempty"`
	AuthorizationCustomer string `json:"authorizationCustomer"`
	Authorization         string `json:"authorization"`
}
