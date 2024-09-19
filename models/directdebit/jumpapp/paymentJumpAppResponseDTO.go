package jumpapp

type PaymentJumpAppResponseDTO struct {
	ResponseCode       string `json:"responseCode"`
	ResponseMessage    string `json:"responseMessage"`
	WebRedirectUrl     string `json:"webRedirectUrl"`
	PartnerReferenceNo string `json:"partnerReferenceNo"`
}
