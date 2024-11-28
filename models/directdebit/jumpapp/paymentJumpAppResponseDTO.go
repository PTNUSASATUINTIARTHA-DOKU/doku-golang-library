package jumpapp

type PaymentJumpAppResponseDTO struct {
	ResponseCode       string                                   `json:"responseCode"`
	ResponseMessage    string                                   `json:"responseMessage"`
	WebRedirectUrl     string                                   `json:"webRedirectUrl,omitempty"`
	PartnerReferenceNo string                                   `json:"partnerReferenceNo,omitempty"`
	AdditionalInfo     *PaymentJumpAppAdditionalInfoResponseDTO `json:"additionalInfo,omitempty"`
}

type PaymentJumpAppAdditionalInfoResponseDTO struct {
	WebRedirectUrl string `json:"webRedirectUrl,omitempty"`
}
