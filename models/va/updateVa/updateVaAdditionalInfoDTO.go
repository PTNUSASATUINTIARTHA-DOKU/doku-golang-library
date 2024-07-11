package models

type UpdateVaAdditionalInfoDTO struct {
	Channel              string                          `json:"channel"`
	VirtualAccountConfig UpdateVaVirtualAccountConfigDTO `json:"virtualAccountConfig"`
}
