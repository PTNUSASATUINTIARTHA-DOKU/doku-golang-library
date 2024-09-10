package payment

type PaymentResponseDTO struct {
	ResponseCode       string `json:"responseCode,omitempty"`
	ResponseMessage    string `json:"responseMessage,omitempty"`
	WebRedirectUrl     string `json:"webRedirectUrl,omitempty"`
	PartnerReferenceNo string `json:"partnerReferenceNo,omitempty"`
}
