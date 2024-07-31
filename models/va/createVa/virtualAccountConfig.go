package models

type VirtualAccountConfig struct {
	ReusableStatus bool    `json:"reusableStatus"`
	MinAmount      *string `json:"minAmount"`
	MaxAmount      *string `json:"maxAmount"`
}
