package accountunbinding

type AccountUnbindingResponseDTO struct {
	ResponseCode    string `json:"responseCode,omitempty"`
	ResponseMessage string `json:"responseMessage,omitempty"`
	ReferenceNo     string `json:"referenceNo,omitempty"`
}
