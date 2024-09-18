package models

import createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"

type UpdateVaResponseDTO struct {
	ResponseCode       string                             `json:"responseCode"`
	ResponseMessage    string                             `json:"responseMessage"`
	VirtualAccountData *createVaModels.VirtualAccountData `json:"virtualAccountData,omitempty"`
}
