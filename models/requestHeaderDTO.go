package models

type RequestHeaderDTO struct {
	XTimestamp    string `json:"xTimestamp"`
	XSignature    string `json:"xSignature"`
	XPartnerId    string `json:"xPartnerId"`
	XExternalId   string `json:"xExternalId"`
	ChannelId     string `json:"channelId"`
	Authorization string `json:"authorization"`
}
