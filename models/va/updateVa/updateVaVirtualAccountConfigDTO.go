package models

type UpdateVaVirtualAccountConfigDTO struct {
	Status    string  `json:"status"`
	MinAmount *string `json:"minAmount"`
	MaxAmount *string `json:"maxAmount"`
}
