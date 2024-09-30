package cardregistrationunbinding

type CardRegistrationUnbindingResponseDTO struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ReferenceNo     string `json:"referenceNo,omitempty"`
	RedirectUrl     string `json:"redirectUrl,omitempty"`
}
