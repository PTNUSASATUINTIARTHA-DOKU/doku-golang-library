package cardregistration

type CardRegistrationResponseDTO struct {
	ResponseCode    string                                    `json:"responseCode"`
	ResponseMessage string                                    `json:"responseMessage"`
	ReferenceNo     string                                    `json:"referenceNo"`
	RedirectUrl     string                                    `json:"redirectUrl"`
	AdditionalInfo  CardRegistrationAdditionalInfoResponseDTO `json:"additionalInfo"`
}

type CardRegistrationAdditionalInfoResponseDTO struct {
	CustIdMerchant string `json:"custIdMerchant"`
	Status         string `json:"status"`
	AuthCode       string `json:"authCode"`
}
