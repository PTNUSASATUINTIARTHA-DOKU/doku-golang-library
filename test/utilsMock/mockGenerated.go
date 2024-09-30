package utilsmock

import (
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
		AdditionalInfo: checkVaModels.CheckStatusResponseAdditionalInfo{
			Acquirer: checkVaModels.AcquirerDetails{
				Id: "BANK_CIMB",
			},
		},
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
