package services

import (
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
