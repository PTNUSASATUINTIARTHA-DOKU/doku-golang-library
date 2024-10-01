package accountunbinding

type AccountUnbindingResponseDTO struct {
	ResponseCode    string             `json:"responseCode"`
	ResponseMessage string             `json:"responseMessage"`
	AdditionalInfo  *AdditionalInfoDTO `json:"additionalInfoDto,omitempty"`
}

type AdditionalInfoDTO struct {
	RedirectUrl string `json:"redirectUrl,omitempty"`
}
