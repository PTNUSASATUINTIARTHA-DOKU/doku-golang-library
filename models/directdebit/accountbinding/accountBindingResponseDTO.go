package accountbinding

type AccountBindingResponseDTO struct {
	ResponseCode    string                                   `json:"responseCode,omitempty"`
	ResponseMessage string                                   `json:"responseMessage,omitempty"`
	ReferenceNo     string                                   `json:"referenceNo,omitempty"`
	RedirectUrl     string                                   `json:"redirectUrl,omitempty"`
	AdditionalInfo  *AccountBindingAdditionalInfoResponseDTO `json:"additionalInfo,omitempty"`
}

type AccountBindingAdditionalInfoResponseDTO struct {
	CustIdMerchant string `json:"custIdMerchant,omitempty"`
	Status         string `json:"status,omitempty"`
	AuthCode       string `json:"authCode,omitempty"`
}
