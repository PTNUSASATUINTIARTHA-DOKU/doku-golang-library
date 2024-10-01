package cardregistration

type CardRegistrationResponseDTO struct {
	ResponseCode    string                                     `json:"responseCode"`
	ResponseMessage string                                     `json:"responseMessage"`
	ReferenceNo     string                                     `json:"referenceNo,omitempty"`
	RedirectUrl     string                                     `json:"redirectUrl,omitempty"`
	AdditionalInfo  *CardRegistrationAdditionalInfoResponseDTO `json:"additionalInfo,omitempty"`
}

type CardRegistrationAdditionalInfoResponseDTO struct {
	CustIdMerchant string `json:"custIdMerchant,omitempty"`
	Status         string `json:"status,omitempty"`
	AuthCode       string `json:"authCode,omitempty"`
}
