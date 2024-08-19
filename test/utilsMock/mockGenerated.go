package utilsmock

import (
	tokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
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
	return createVaModels.CreateVaResponseDto{
		ResponseCode:    "2002700",
		ResponseMessage: "Successful",
		VirtualAccountData: createVaModels.VirtualAccountData{
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
		},
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
	return updateVaModels.UpdateVaResponseDTO{
		ResponseCode:    "2002700",
		ResponseMessage: "Successful",
		VirtualAccountData: createVaModels.VirtualAccountData{
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
		},
	}
}
