package payment

type PaymentNotificationResponseDTO struct {
	Header PaymentNotificationResponseHeaderDTO `json:"header"`
	Body   PaymentNotificationResponseBodyDTO   `json:"body"`
}
