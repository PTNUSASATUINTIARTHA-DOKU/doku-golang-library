package models

type Origin struct {
	Product       string `json:"product,omitempty"`
	Source        string `json:"source,omitempty"`
	SourceVersion string `json:"sourceVersion,omitempty"`
	System        string `json:"system,omitempty"`
	ApiFormat     string `json:"apiFormat,omitempty"`
}

type AdditionalInfo struct {
	Channel              string               `json:"channel"`
	VirtualAccountConfig VirtualAccountConfig `json:"virtualAccountConfig"`
	Origin               Origin               `json:"origin"`
}
