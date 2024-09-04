package models

type RequestHeaderDTO struct {
	XTimestamp            string `json:"xTimestamp"`
	XSignature            string `json:"xSignature"`
	XPartnerId            string `json:"xPartnerId"`
	XExternalId           string `json:"xExternalId"`
	XDeviceId             string `json:"xDeviceId"`
	XIpAddress            string `json:"xIpAddress"`
	ChannelId             string `json:"channelId"`
	Authorization         string `json:"authorization"`
	AuthorizationCustomer string `json:"authorizationCustomer"`
}
