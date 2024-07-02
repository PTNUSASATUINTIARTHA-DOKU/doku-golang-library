package models

type AdditionalInfoResponse struct {
	Channel      string `json:"channel"`
	HowToPayPage string `json:"howToPayPage"`
	HowToPayApi  string `json:"howToPayApi"`
}
