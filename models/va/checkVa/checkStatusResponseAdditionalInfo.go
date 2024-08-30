package models

type CheckStatusResponseAdditionalInfo struct {
	Acquirer AcquirerDetails `json:"acquirer"`
}

type AcquirerDetails struct {
	Id string `json:"id"`
}
