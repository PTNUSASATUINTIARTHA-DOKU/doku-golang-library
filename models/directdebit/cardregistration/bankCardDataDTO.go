package cardregistration

type BankCardDataDTO struct {
	BankCardNo   string `json:"bankCardNo"`
	BankCardType string `json:"bankCardType"`
	ExpiryDate   string `json:"expiryDate"`
}
