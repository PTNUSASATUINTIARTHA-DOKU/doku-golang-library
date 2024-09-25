package models

type UpdateVaVirtualAccountConfigDTO struct {
	ReusableStatus bool    `json:"reusableStatus,omitempty"`
	Status         string  `json:"status"`
	MinAmount      *string `json:"minAmount,omitempty"`
	MaxAmount      *string `json:"maxAmount,omitempty"`
}
