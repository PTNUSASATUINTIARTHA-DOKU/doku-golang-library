package models

import (
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type UpdateVaAdditionalInfoDTO struct {
	Channel              string                          `json:"channel"`
	VirtualAccountConfig UpdateVaVirtualAccountConfigDTO `json:"virtualAccountConfig"`
	Origin               createVaModels.Origin           `json:"origin"`
}
