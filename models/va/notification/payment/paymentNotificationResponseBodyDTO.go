package payment

import notification "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification"

type PaymentNotificationResponseBodyDTO struct {
	ResponseCode       string                                      `json:"responseCode"`
	ResponseMessage    string                                      `json:"responseMessage"`
	VirtualAccountData notification.NotificationVirtualAccountData `json:"virtualAccountData"`
}
