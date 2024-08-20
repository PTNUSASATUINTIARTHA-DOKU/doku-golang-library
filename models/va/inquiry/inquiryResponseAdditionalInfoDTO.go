package inquiry

import (
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type InquiryResponseAdditionalInfoDTO struct {
	Channel              string                              `json:"channel"`
	TrxId                string                              `json:"trxId"`
	VirtualAccountConfig createVaModels.VirtualAccountConfig `json:"virtualAccountConfig"`
}
