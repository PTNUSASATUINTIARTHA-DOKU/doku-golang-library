package utilsmock

import (
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	accountUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountunbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	cardRegistrationModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistration"
	cardRegistrationUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistrationunbinding"
	checkStatusModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/checkstatus"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"
	tokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"
	updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"
)

type MockGenerated struct{}

func (mr *MockGenerated) GetTokenB2BResponseDTO(responseCode string) tokenModels.TokenB2BResponseDTO {
	return tokenModels.TokenB2BResponseDTO{
		ResponseCode:    responseCode,
		ResponseMessage: "Successful",
		AccessToken:     "TOKEN_B2B",
		TokenType:       "Bearer ",
		ExpiresIn:       900,
	}
}

func (mr *MockGenerated) CreateVaRequestDTO() createVaModels.CreateVaRequestDto {
	return createVaModels.CreateVaRequestDto{
		PartnerServiceId:    "    1899",
		CustomerNo:          "20240704001",
		VirtualAccountNo:    "    189920240704001",
		VirtualAccountName:  "SDKMockTest",
		VirtualAccountEmail: "mock@testing.com",
		VirtualAccountPhone: "6281288932399",
		TrxId:               "INV_20240711001",
		TotalAmount: createVaModels.TotalAmount{
			Value:    "11000.00",
			Currency: "IDR",
		},
		AdditionalInfo: createVaModels.AdditionalInfo{
			Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
			VirtualAccountConfig: createVaModels.VirtualAccountConfig{
				ReusableStatus: false,
			},
		},
		VirtualAccountTrxType: "C",
		ExpiredDate:           "2024-11-24T10:55:00+07:00",
	}
}

func (mr *MockGenerated) CreateVaResponseDTO() createVaModels.CreateVaResponseDto {
	virtualAccountData := createVaModels.VirtualAccountData{
		PartnerServiceId:    "    1899",
		CustomerNo:          "20240704001",
		VirtualAccountNo:    "    189920240704001",
		VirtualAccountName:  "SDKMockTest",
		VirtualAccountEmail: "mock@testing.com",
		TrxId:               "INV_20240711001",
		TotalAmount: createVaModels.TotalAmount{
			Value:    "11000.00",
			Currency: "IDR",
		},
		AdditionalInfo: createVaModels.AdditionalInfoResponse{
			HowToPayPage: "howToPayPage",
			HowToPayApi:  "howToPayApi",
		},
	}
	return createVaModels.CreateVaResponseDto{
		ResponseCode:       "2002700",
		ResponseMessage:    "Successful",
		VirtualAccountData: &virtualAccountData,
	}
}

func (mr *MockGenerated) UpdateVaRequestDTO() updateVaModels.UpdateVaDTO {
	return updateVaModels.UpdateVaDTO{
		PartnerServiceId:    "    1899",
		CustomerNo:          "20240704001",
		VirtualAccountNo:    "    189920240704001",
		VirtualAccountName:  "SDKMockTest",
		VirtualAccountEmail: "mock@testing.com",
		VirtualAccountPhone: "6281288932399",
		TrxId:               "INV_20240711001",
		TotalAmount: createVaModels.TotalAmount{
			Value:    "11000.00",
			Currency: "IDR",
		},
		AdditionalInfo: updateVaModels.UpdateVaAdditionalInfoDTO{
			Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
			VirtualAccountConfig: updateVaModels.UpdateVaVirtualAccountConfigDTO{
				Status: "ACTIVE",
			},
		},
		VirtualAccountTrxType: "C",
		ExpiredDate:           "2024-11-24T10:55:00+07:00",
	}
}

func (mr *MockGenerated) UpdateVaResponseDTO() updateVaModels.UpdateVaResponseDTO {
	virtualAccountData := updateVaModels.UpdateVaDTO{
		PartnerServiceId:    "    1899",
		CustomerNo:          "20240704001",
		VirtualAccountNo:    "    189920240704001",
		VirtualAccountName:  "SDKMockTest",
		VirtualAccountEmail: "mock@testing.com",
		TrxId:               "INV_20240711001",
		TotalAmount: createVaModels.TotalAmount{
			Value:    "11000.00",
			Currency: "IDR",
		},
		AdditionalInfo: updateVaModels.UpdateVaAdditionalInfoDTO{
			Channel: "",
			VirtualAccountConfig: updateVaModels.UpdateVaVirtualAccountConfigDTO{
				ReusableStatus: true,
				Status:         "",
			},
		},
	}
	return updateVaModels.UpdateVaResponseDTO{
		ResponseCode:       "2002700",
		ResponseMessage:    "Successful",
		VirtualAccountData: virtualAccountData,
	}
}

func (mr *MockGenerated) CheckStatusVaRequest() checkVaModels.CheckStatusVARequestDto {
	return checkVaModels.CheckStatusVARequestDto{
		PartnerServiceId: "    1899",
		CustomerNo:       "000000000968",
		VirtualAccountNo: "    1899000000000968",
	}
}

func (mr *MockGenerated) CheckStatusVa() checkVaModels.CheckStatusVaResponseDto {
	virtualAccountData := checkVaModels.CheckStatusVirtualAccountData{
		PaymentFlagReason: checkVaModels.CheckStatusResponsePaymentFlagReason{
			English:   "Pending",
			Indonesia: "Belum Terbayar",
		},
		PartnerServiceId: "    1899",
		CustomerNo:       "000000000966",
		VirtualAccountNo: "    1899000000000966",
		PaidAmount: createVaModels.TotalAmount{
			Value:    "11000.00",
			Currency: "IDR",
		},
		BillDetails: []checkVaModels.CheckStatusBillDetail{{
			BillAmount: createVaModels.TotalAmount{
				Value:    "11000.00",
				Currency: "IDR",
			},
		}},
		TrxId: "7041",
	}
	return checkVaModels.CheckStatusVaResponseDto{
		ResponseCode:       "2002600",
		ResponseMessage:    "Successful",
		VirtualAccountData: &virtualAccountData,
	}
}

func (mr *MockGenerated) DeletePaymentCodeRequest() deleteVaModels.DeleteVaRequestDto {
	return deleteVaModels.DeleteVaRequestDto{
		PartnerServiceId: "    1899",
		CustomerNo:       "000000000971",
		VirtualAccountNo: "    1899000000000971",
		TrxId:            "757",
		AdditionalInfo: deleteVaModels.DeleteVaRequestAdditionalInfo{
			Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
		},
	}
}

func (mr *MockGenerated) DeletePaymentCode() deleteVaModels.DeleteVaResponseDto {
	virtualAccountData := deleteVaModels.DeleteVaResponseVirtualAccountData{
		PartnerServiceId: "    1899",
		CustomerNo:       "000000000971",
		VirtualAccountNo: "    1899000000000971",
		TrxId:            "757",
		AdditionalInfo: deleteVaModels.DeleteVaResponseAdditionalInfo{
			Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
		},
	}
	return deleteVaModels.DeleteVaResponseDto{
		ResponseCode:       "2003100",
		ResponseMessage:    "Successful",
		VirtualAccountData: &virtualAccountData,
	}
}

func (mr *MockGenerated) GetTokenB2B2CResponseDTO(responseCode string) tokenModels.TokenB2B2CResponseDTO {
	return tokenModels.TokenB2B2CResponseDTO{
		ResponseCode:          responseCode,
		ResponseMessage:       "Successful",
		AccessToken:           "TOKEN_B2B2C",
		TokenType:             "Bearer ",
		AccessTokenExpiryTime: "900",
	}
}

func (mr *MockGenerated) AccountBindingRequest() accountBindingModels.AccountBindingRequestDTO {
	return accountBindingModels.AccountBindingRequestDTO{
		PhoneNo: "6289912121237",
		AdditionalInfo: accountBindingModels.AccountBindingAdditionalInfoRequestDto{
			Channel:                "DIRECT_DEBIT_CIMB_SNAP",
			CustIdMerchant:         "SDK7",
			CustomerName:           "SDK GOLANG",
			Email:                  "sdkgo@gmail.com",
			IdCard:                 "99999",
			Country:                "Indonesia",
			Address:                "Jakarta",
			DateOfBirth:            "19990101",
			SuccessRegistrationUrl: "https://sandbox.doku.com/bo/login/",
			FailedRegistrationUrl:  "https://www.doku.com/id-ID",
			DeviceModel:            "ios18",
			OsType:                 "ios",
			ChannelId:              "app",
		},
	}
}

func (mr *MockGenerated) AccountBindingResponse() accountBindingModels.AccountBindingResponseDTO {
	return accountBindingModels.AccountBindingResponseDTO{
		ResponseCode:    "2000700",
		ResponseMessage: "Successful",
		ReferenceNo:     "xxx-xxx-xxx",
		RedirectUrl:     "https://sandbox.doku.com/direct-debit/ui/binding/core/2238241007113952193107119118405011355991",
		AdditionalInfo: &accountBindingModels.AccountBindingAdditionalInfoResponseDTO{
			CustIdMerchant: "SDK20",
			AuthCode:       "2238241007113952193107119118405011355991",
		},
	}
}

func (mr *MockGenerated) AccountUnbindingRequest() accountUnbindingModels.AccountUnbindingRequestDTO {
	return accountUnbindingModels.AccountUnbindingRequestDTO{
		TokenId: "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTAxMDEsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA4LTE3MjcyMzU4NDU1MDciLCJhY2NvdW50SWQiOiI5OGVkMjEwZWNlMzAwNTk0MDM2ZjM2ODRmMTkxMGRhNSJ9.OLaDwsbM3Mgz9C-12Gl_8UaGbScvbHCKHtDgz4N36oaP1XGM11XcGVqLS8bKMNc28ww8Z9rLsmYPZignXyzSnoYmdRlY8p_q-6g3sDYU-eps7Wet5dpXm-kgQfi4YiwSh_USMss6ZWiWZxVBHHqLte6M7yCYv2eTNtWC6hSm9Yihgrry99_W3_OVQEd4Vlfg8bpS3E53KTy-pt9pifnenjEQMC3tnndSJYNjeXF6OPb8Y906wvO1-yCQOKn7Xkm7FrQYIQ6McYvTD00V729olj_bLMnMMURbXJC9NM_UR4mS042b_PgNV3N-cCriG_GuSQ0uW2c-8S5N6p-l6Dfryg",
		AdditionalInfo: accountUnbindingModels.AccountUnbindingAdditionalInfoRequestDTO{
			Channel: "DIRECT_DEBIT_CIMB_SNAP",
		},
	}
}

func (mr *MockGenerated) AccountUnbindingResponse() accountUnbindingModels.AccountUnbindingResponseDTO {
	return accountUnbindingModels.AccountUnbindingResponseDTO{
		ResponseCode:    "2000900",
		ResponseMessage: "Successful",
	}
}

func (mr *MockGenerated) BalanceInquiryRequest() balanceInquiryModels.BalanceInquiryRequestDto {
	return balanceInquiryModels.BalanceInquiryRequestDto{
		AdditionalInfo: balanceInquiryModels.BalanceInquiryAdditionalInfoRequestDto{
			Channel: "DIRECT_DEBIT_CIMB_SNAP",
		},
	}
}

func (mr *MockGenerated) BalanceInquiryResponse() balanceInquiryModels.BalanceInquiryResponseDto {
	return balanceInquiryModels.BalanceInquiryResponseDto{
		ResponseCode:    "2001100",
		ResponseMessage: "Successful",
	}
}

func (mr *MockGenerated) CheckStatusRequest() checkStatusModels.CheckStatusRequestDTO {
	return checkStatusModels.CheckStatusRequestDTO{
		OriginalPartnerReferenceNo: "2020102900000000000001",
		OriginalReferenceNo:        "2020102977770000000009",
		OriginalExternalId:         "30443786930722726463280097920912",
		ServiceCode:                "55",
		TransactionDate:            "2020-12-21T14:56:11+07:00",
		Amount: createVaModels.TotalAmount{
			Value:    "12345678.00",
			Currency: "IDR",
		},
		MerchantId:      "23489182303312",
		SubMerchantId:   "23489182303312",
		ExternalStoreId: "183908924912387",
		AdditionalInfo: checkStatusModels.CheckStatusAdditionalInfoRequestDTO{
			DeviceId: "12345679237",
			Channel:  "DIRECT_DEBIT_ALLO_SNAP",
		},
	}

}

func (mr *MockGenerated) CheckStatusResponse() checkStatusModels.CheckStatusResponseDTO {
	return checkStatusModels.CheckStatusResponseDTO{
		ResponseCode:               "2005500",
		ResponseMessage:            "Request has been processed successfully",
		OriginalPartnerReferenceNo: "2020102900000000000001",
		OriginalReferenceNo:        "2020102977770000000009",
		ApprovalCode:               "201039000200",
		OriginalExternalId:         "30443786930722726463280097920912",
		ServiceCode:                "55",
		LatestTransactionStatus:    "00",
		TransactionStatusDesc:      "success",
		OriginalResponseCode:       "2005500",
		OriginalResponseMessage:    "Request has been processed successfully",
		SessionId:                  "883737GHY8839",
		RequestID:                  "3763773",
		RefundHistory: []*checkStatusModels.RefundHistoryDTO{
			{
				RefundNo:           "96194816941239812",
				PartnerReferenceNo: "239850918204981205970",
				RefundAmount: &createVaModels.TotalAmount{
					Value:    "12345678.00",
					Currency: "IDR",
				},
				RefundStatus: "00",
				RefundDate:   "2020-12-23T07:44:16+07:00",
				Reason:       "Customer Complain",
			},
			{
				RefundNo:           "96194123981251341",
				PartnerReferenceNo: "2398509123131981205970",
				RefundAmount: &createVaModels.TotalAmount{
					Value:    "112345678.00",
					Currency: "IDR",
				},
				RefundStatus: "00",
				RefundDate:   "2020-12-23T07:54:16+07:00",
				Reason:       "Customer Complain",
			},
		},
		TransAmount: &createVaModels.TotalAmount{
			Value:    "112345678.00",
			Currency: "IDR",
		},
		FeeAmount: &createVaModels.TotalAmount{
			Value:    "112345678.00",
			Currency: "IDR",
		},
		PaidTime: "2020-12-21T14:56:11+07:00",
		AdditionalInfo: &checkStatusModels.CheckStatusAdditionalInfoResponseDTO{
			DeviceId: "12345679237",
			Channel:  "mobilephone",
		},
	}

}

func (mr *MockGenerated) CardRegistrationRequest() cardRegistrationModels.CardRegistrationRequestDTO {
	return cardRegistrationModels.CardRegistrationRequestDTO{
		CardData:       "EMMRGWbWdaBbJubgXCLOleE0DlsqY875SlR6wrbEHar/zVgpYbJO18WGSACWmOrzD1gnQA9hZFnVsuIpvrHnsEKwSd7HMIbFknqtX8ulk6dplYpoWObZAJdpbjTW3RFp|5Dbs7CCZbkCcQkiu8544Fg==",
		CustIdMerchant: "SDKBRIGO4",
		PhoneNo:        "6287888697821",
		AdditionalInfo: cardRegistrationModels.CardRegistrationAdditionalInfoRequestDTO{
			Channel:                "DIRECT_DEBIT_BRI_SNAP",
			DateOfBirth:            "19990101",
			SuccessRegistrationUrl: "https://sandbox.doku.com/bo/login/",
			FailedRegistrationUrl:  "https://doku.com",
		},
	}

}

func (mr *MockGenerated) CardRegistrationResponse() cardRegistrationModels.CardRegistrationResponseDTO {
	return cardRegistrationModels.CardRegistrationResponseDTO{
		ResponseCode:    "2000100",
		ResponseMessage: "Successful",
		ReferenceNo:     "644715189102",
		RedirectUrl:     "https://sandbox.doku.com/direct-debit/ui/otp/bri/binding/2238241007124257729107119118427011910531",
		AdditionalInfo: &cardRegistrationModels.CardRegistrationAdditionalInfoResponseDTO{
			CustIdMerchant: "SDKBRIGO4",
			Status:         "PENDING",
			AuthCode:       "2238241007124257729107119118427011910531",
		},
	}

}

func (mr *MockGenerated) CardRegistrationUnbindingRequest() cardRegistrationUnbindingModels.CardRegistrationUnbindingRequestDTO {
	return cardRegistrationUnbindingModels.CardRegistrationUnbindingRequestDTO{
		TokenId: "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkyMzQwNTIsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA4LTE3MjcyMzU4NDU1MDciLCJhY2NvdW50SWQiOiJiOTY2ODBjOGZhNjcwZTAwZjdkZmVkYzU1ZTYwODJmNCJ9.P-6Kf8ovIxoo",
		AdditionalInfo: cardRegistrationUnbindingModels.CardRegistrationUnbindingAdditionalInfoRequestDTO{
			Channel: "DIRECT_DEBIT_BRI_SNAP",
		},
	}
}

func (mr *MockGenerated) CardRegistrationUnbindingResponse() cardRegistrationUnbindingModels.CardRegistrationUnbindingResponseDTO {
	return cardRegistrationUnbindingModels.CardRegistrationUnbindingResponseDTO{
		ResponseCode:    "2000500",
		ResponseMessage: "Successful",
		ReferenceNo:     "UNB-0001",
		RedirectUrl:     "https://doku.com/direct-debit/ui/binding/2238230713001534401107183161486001168389",
	}
}

func (mr *MockGenerated) RefundRequest() refundModels.RefundRequestDTO {
	return refundModels.RefundRequestDTO{
		OriginalPartnerReferenceNo: "INV-0001",
		PartnerRefundNo:            "INV-REF-0001",
		RefundAmount: createVaModels.TotalAmount{
			Value:    "15000.00",
			Currency: "IDR",
		},
		AdditionalInfo: refundModels.RefundAdditionalInfoRequestDTO{
			Channel: "DIRECT_DEBIT_CIMB_SNAP",
		},
	}
}

func (mr *MockGenerated) RefundResponse() refundModels.RefundResponseDTO {
	return refundModels.RefundResponseDTO{
		ResponseCode:    "2000700",
		ResponseMessage: "Successful",
		RefundAmount: &createVaModels.TotalAmount{
			Value:    "10000.00",
			Currency: "IDR",
		},
		OriginalPartnerReferenceNo: "Ra7o1bLJAh2oV9eb33129stQc5xFm5s7",
		OriginalReferenceNo:        "Ra7o1bLJAh2oV9eb33129stQc5xFm5s7",
		RefundNo:                   "Ra7o1bLJAh2oV9eb33129stQc5xFm5s7",
		PartnerRefundNo:            "Ra7o1bLJAh2oV9eb33129stQc5xFm5s7",
		RefundTime:                 "2024-01-01T09:09:00.123",
	}

}

func (mr *MockGenerated) PaymentDirectDebitRequest() paymentModels.PaymentRequestDTO {
	return paymentModels.PaymentRequestDTO{
		PartnerReferenceNo: "INV-0001",
		Amount: createVaModels.TotalAmount{
			Value:    "10000.00",
			Currency: "IDR",
		},
		AdditionalInfo: paymentModels.PaymentAdditionalInfoDTO{
			Channel:           "DIRECT_DEBIT_CIMB_SNAP",
			SuccessPaymentUrl: "www.merchant.com/success",
			FailedPaymentUrl:  "www.merchant.com/failed",
			Remarks:           "Remarks",
		},
	}

}

func (mr *MockGenerated) PaymentDirectDebitResponse() paymentModels.PaymentResponseDTO {
	return paymentModels.PaymentResponseDTO{
		ResponseCode:       "2005400",
		ResponseMessage:    "Successful",
		WebRedirectUrl:     "https://app-uat.doku.com/link/283702597342040",
		PartnerReferenceNo: "INV-0001",
	}

}
