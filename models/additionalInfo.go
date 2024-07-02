package models

type AdditionalInfo struct {
	Channel              string               `json:"channel"`
	VirtualAccountConfig VirtualAccountConfig `json:"virtualAccountConfig"`
}
