package services

import (
	notifDirectDebitModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/notification"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	notification "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification"
	paymentNotifModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/payment"
)

type NotificationServices struct{}

func (ns NotificationServices) GenerateNotificationResponse(paymentNotificationRequestBodyDTO paymentNotifModels.PaymentNotificationRequestBodyDTO) paymentNotifModels.PaymentNotificationResponseBodyDTO {
	return paymentNotifModels.PaymentNotificationResponseBodyDTO{
		ResponseCode:    "2002500",
		ResponseMessage: "success",
		VirtualAccountData: notification.NotificationVirtualAccountData{
			PartnerServiceId:   paymentNotificationRequestBodyDTO.PartnerServiceId,
			CustomerNo:         paymentNotificationRequestBodyDTO.CustomerNo,
			VirtualAccountNo:   paymentNotificationRequestBodyDTO.VirtualAccountNo,
			VirtualAccountName: paymentNotificationRequestBodyDTO.VirtualAccountName,
			PaymentRequestId:   paymentNotificationRequestBodyDTO.PaymentRequestId,
			AdditionalInfo: createVaModels.AdditionalInfo{
				Channel: paymentNotificationRequestBodyDTO.AdditionalInfo.Channel,
				VirtualAccountConfig: createVaModels.VirtualAccountConfig{
					ReusableStatus: paymentNotificationRequestBodyDTO.AdditionalInfo.VirtualAccountConfig.ReusableStatus,
				},
			},
		},
	}
}

func (ns NotificationServices) GenerateInvalidTokenNotificationResponse(paymentNotificationRequestBodyDTO paymentNotifModels.PaymentNotificationRequestBodyDTO) paymentNotifModels.PaymentNotificationResponseBodyDTO {
	return paymentNotifModels.PaymentNotificationResponseBodyDTO{
		ResponseCode:    "4012701",
		ResponseMessage: "Invalid Token (B2B)",
		VirtualAccountData: notification.NotificationVirtualAccountData{
			PartnerServiceId:   paymentNotificationRequestBodyDTO.PartnerServiceId,
			CustomerNo:         paymentNotificationRequestBodyDTO.CustomerNo,
			VirtualAccountNo:   paymentNotificationRequestBodyDTO.VirtualAccountNo,
			VirtualAccountName: paymentNotificationRequestBodyDTO.VirtualAccountName,
			PaymentRequestId:   paymentNotificationRequestBodyDTO.PaymentRequestId,
		},
	}
}

func (ns NotificationServices) GenerateDirectDebitNotificationResponse() notifDirectDebitModels.NotificationPaymentDirectDebitResponseDTO {
	return notifDirectDebitModels.NotificationPaymentDirectDebitResponseDTO{
		ResponseCode:    "2005600",
		ResponseMessage: "Request has been processed successfully",
		ApprovalCode:    "201039000200",
	}
}

func (ns NotificationServices) GenerateDirectDebitInvalidTokenNotificationResponse() notifDirectDebitModels.NotificationPaymentDirectDebitResponseDTO {
	return notifDirectDebitModels.NotificationPaymentDirectDebitResponseDTO{
		ResponseCode:    "5005600",
		ResponseMessage: "Invalid Token",
	}
}
