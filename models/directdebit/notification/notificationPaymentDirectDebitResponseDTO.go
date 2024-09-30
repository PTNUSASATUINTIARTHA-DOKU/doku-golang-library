package notification

type NotificationPaymentDirectDebitResponseDTO struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	ApprovalCode    string `json:"approvalCode,omitempty"`
}
