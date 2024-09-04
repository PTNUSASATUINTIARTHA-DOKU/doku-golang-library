package accountbinding

type AccountBindingResponseDto struct {
	ResponseCode    string                                   `json:"responseCode,omitempty"`
	ResponseMessage string                                   `json:"responseMessage,omitempty"`
	ReferenceNo     string                                   `json:"referenceNo,omitempty"`
	RedirectUrl     string                                   `json:"redirectUrl,omitempty"`
	AdditionalInfo  *AccountBindingAdditionalInfoResponseDto `json:"additionalInfo,omitempty"`
}

type AccountBindingAdditionalInfoResponseDto struct {
	CustIdMerchant string `json:"custIdMerchant,omitempty"`
	Status         string `json:"status,omitempty"`
	AuthCode       string `json:"authCode,omitempty"`
}
