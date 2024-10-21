package cardregistration

type BankCardDataDTO struct {
	BankCardNo         string `json:"bankCardNo"`
	BankCardType       string `json:"bankCardType"`
	IdentificationNo   string `json:"identificationNo,omitempty"`
	IdentificationType string `json:"identificationType,omitempty"`
	Email              string `json:"email,omitempty"`
	ExpiryDate         string `json:"expiryDate"`
}
