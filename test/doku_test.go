package test

import (
	"strings"
	"testing"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/doku"
	utilsmock "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/test/utilsMock"
	"github.com/stretchr/testify/assert"
)

var mockGenerated utilsmock.MockGenerated
var mockController = new(utilsmock.MockController)

func TestGetTokenB2BSuccess(t *testing.T) {

	doku.TokenController = mockController
	mockController.On("GetTokenB2B", "privateKeyPem", "BRN-0221-1693209567392", false).Return(mockGenerated.GetTokenB2BResponseDTO("2007300"))
	snap := doku.Snap{
		PrivateKey:   "privateKeyPem",
		ClientId:     "BRN-0221-1693209567392",
		IsProduction: false,
	}
	actualResponse := snap.GetTokenB2B()
	assert.Equal(t, "2007300", actualResponse.ResponseCode)

}
func TestGetTokenB2BInvalidClientId(t *testing.T) {

	doku.TokenController = mockController
	mockController.On("GetTokenB2B", "privateKeyPem", "BRN", false).Return(mockGenerated.GetTokenB2BResponseDTO("5007300"))
	snap := doku.Snap{
		PrivateKey:   "privateKeyPem",
		ClientId:     "BRN",
		IsProduction: false,
	}
	actualResponse := snap.GetTokenB2B()
	assert.Equal(t, "5007300", actualResponse.ResponseCode)

}

func TestCreateVaSuccess(t *testing.T) {

	doku.VaController = mockController
	mockController.On("CreateVa", mockGenerated.CreateVaRequestDTO(), "privateKeyPem", "clientId", "TOKEN_B2B", false).Return(mockGenerated.CreateVaResponseDTO())
	snap := doku.Snap{
		PrivateKey:   "privateKeyPem",
		ClientId:     "clientId",
		IsProduction: false,
	}
	snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
	actualResponse := snap.CreateVa(mockGenerated.CreateVaRequestDTO())
	assert.Equal(t, "2002700", actualResponse.ResponseCode)

}

func TestCreateVaPartnerIdNot8Digits(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.PartnerServiceId = "1234"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must be exactly 8 characters long and equiped with left-padded spaces. Example: '  888994"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaPartnerIdInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.PartnerServiceId = "    890E"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must consist of up to 8 digits of character. Remaining space in case of partner serivce id is less than 8 must be filled with spaces. Example: ' 888994' (2 spaces and 6 digits)"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaCustomerNoInvalidLength(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.CustomerNo = "1234567890123456789011"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must be 20 characters or fewer. Ensure that customerNo is no longer than 20 characters. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaCustomerNoInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.CustomerNo = "1234567E"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must consist of only digits. Ensure that customerNo contains only numbers. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccNameIsLessThan1(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountName = ""
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountName must be at least 1 character long. Ensure that virtualAccountName is not empty. Example: 'Toru Yamashita"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccNameIsMoreThan255(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountName = "ImCyDjTlTqJu9Rrq1uSuKxNNqcNdcD8EuXigmUMZsge3fvkSOyZ8FwMfyDGeOXxaDENzXzHrnXTfHIqXaKLz5Uq7zaGkjNL0DiTRn7vnBEigFFkJlhftfqiT2ml82pYI1ZUmuuR3N1zaAQNYZvg3asANmoDVGmJYnMdGTyWtD3PPb2t8Nwm57Qd1BfSZIiC7A4cGFSyzYZNp2ObxP4zUeMoa0TPV2WbnLKJ761qP594vMXt9Om4pzdcwK3aAWHQd"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountName must be 255 characters or fewer. Ensure that virtualAccountName is no longer than 255 characters. Example: 'Toru Yamashita'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccNameIsInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountName = "!!!AAA1"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountName can only contain letters, numbers, spaces, and the following characters"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccIsLessThan1(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountEmail = ""
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountEmail must be at least 1 character long"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccEmailIsMoreThan255(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountEmail = "ImCyDjTlTqJu9Rrq1uSuKxNNqcNdcD8EuXigmUMZsge3fvkSOyZ8FwMfyDGeOXxaDENzXzHrnXTfHIqXaKLz5Uq7zaGkjNL0DiTRn7vnBEigFFkJlhftfqiT2ml82pYI1ZUmuuR3N1zaAQNYZvg3asANmoDVGmJYnMdGTyWtD3PPb2t8Nwm57Qd1BfSZIiC7A4cGFSyzYZNp2ObxP4zUeMoa0TPV2WbnLKJ761qP594vMXt9Om4pzdcwK3aAWHQd"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountEmail must be 255 characters or fewer."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccEmailIsInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountEmail = "doku-golang-library@mailcom"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountEmail is not in a valid email format"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccPhoneLessThan9(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountPhone = "12345678"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountPhone must be at least 9 characters long."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaVirtualAccPhoneMoreThan30(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountPhone = "1234567890123456789012345678901"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "virtualAccountPhone must be 30 characters or fewer. Ensure that virtualAccountPhone is no longer than 30 characters."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTrxIdIsLessThan1(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.TrxId = ""
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TrxId must be at least 1 character long. Ensure that TrxId is not empty. Example: '23219829713'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTrxIdIsMoreThan64(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.TrxId = "FcGcsrqYNNQotmv7b2dSFdVbUmiexl0s1wE7H23gpXsFzcXUHXXnRLBUuREMuWxVx"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TrxId must be 64 characters or fewer. Ensure that TrxId is no longer than 64 characters. Example: '23219829713'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTotalAmountIsLessThan4(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.TotalAmount.Value = "0"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Value must be at least 4 characters long and formatted as 0.00."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTotalAmountIsMoreThan19(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.TotalAmount.Value = "12345678901234567890"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Value must be 19 characters or fewer and formatted as 9999999999999999.99. Ensure that TotalAmount.Value is no longer than 19 characters and in the correct format"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTotalAmountIsInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.TotalAmount.Value = "1000E"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Value is invalid format."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTotalAmountCurrencyNot3Char(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.TotalAmount.Currency = "ID"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Currency must be exactly 3 characters long."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTotalAmountCurrencyNotIDR(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.TotalAmount.Currency = "INR"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.currency must be 'IDR'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaChannelIsLessThan1(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.AdditionalInfo.Channel = ""
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.Channel must be at least 1 character long"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaChannelIsMoreThan30(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.AdditionalInfo.Channel = "VIRTUAL_ACCOUNT_BANK_MANDIRI_TEST"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.Channel must be 30 characters or fewer."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaChannelIsInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.AdditionalInfo.Channel = "5Vl3mjMJpA6NuUNHWrucSymfjlWPCb"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.channel is not valid. Ensure that AdditionalInfo.channel is one of the valid channels."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTrxTypeIsNot1Digit(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountTrxType = "CO"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountTrxType must be exactly 1 character long. Ensure that VirtualAccountTrxType is either 'C', 'O' or 'V'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaTrxTypeIsInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.VirtualAccountTrxType = "A"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountTrxType must be either 'C', 'O' or 'V'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCreateVaExpiredIsInvalidFormat(t *testing.T) {

	request := mockGenerated.CreateVaRequestDTO()
	request.ExpiredDate = "2024-07-11"
	err := request.ValidateVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "ExpiredDate must be in ISO-8601 format."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaSuccess(t *testing.T) {

	doku.VaController = mockController
	mockController.On("DoUpdateVa", mockGenerated.UpdateVaRequestDTO(), "clientId", "TOKEN_B2B", "SECRET_KEY", false).Return(mockGenerated.UpdateVaResponseDTO())
	snap := doku.Snap{
		PrivateKey:   "privateKeyPem",
		ClientId:     "clientId",
		IsProduction: false,
		SecretKey:    "SECRET_KEY",
	}
	snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
	actualResponse := snap.UpdateVa(mockGenerated.UpdateVaRequestDTO())
	assert.Equal(t, "2002700", actualResponse.ResponseCode)

}

func TestUpdateVaPartnerIdNot8Digits(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.PartnerServiceId = "1234"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must be exactly 8 characters long and equiped with left-padded spaces. Example: '  888994"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaPartnerIdInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.PartnerServiceId = "    890E"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must consist of up to 8 digits of character. Remaining space in case of partner serivce id is less than 8 must be filled with spaces. Example: ' 888994' (2 spaces and 6 digits)"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaCustomerNoInvalidLength(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.CustomerNo = "1234567890123456789011"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must be 20 characters or fewer. Ensure that customerNo is no longer than 20 characters. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaCustomerNoInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.CustomerNo = "1234567E"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must consist of only digits. Ensure that customerNo contains only numbers. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaVirtualAccNameIsLessThan1(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountName = ""
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountName must be at least 1 character long. Ensure that virtualAccountName is not empty. Example: 'Toru Yamashita"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaVirtualAccNameIsMoreThan255(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountName = "ImCyDjTlTqJu9Rrq1uSuKxNNqcNdcD8EuXigmUMZsge3fvkSOyZ8FwMfyDGeOXxaDENzXzHrnXTfHIqXaKLz5Uq7zaGkjNL0DiTRn7vnBEigFFkJlhftfqiT2ml82pYI1ZUmuuR3N1zaAQNYZvg3asANmoDVGmJYnMdGTyWtD3PPb2t8Nwm57Qd1BfSZIiC7A4cGFSyzYZNp2ObxP4zUeMoa0TPV2WbnLKJ761qP594vMXt9Om4pzdcwK3aAWHQd"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountName must be 255 characters or fewer. Ensure that virtualAccountName is no longer than 255 characters. Example: 'Toru Yamashita'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaVirtualAccNameIsInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountName = "!!!AAA1"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountName can only contain letters, numbers, spaces, and the following characters"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVirtualAccIsLessThan1(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountEmail = ""
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountEmail must be at least 1 character long"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaVirtualAccEmailIsMoreThan255(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountEmail = "ImCyDjTlTqJu9Rrq1uSuKxNNqcNdcD8EuXigmUMZsge3fvkSOyZ8FwMfyDGeOXxaDENzXzHrnXTfHIqXaKLz5Uq7zaGkjNL0DiTRn7vnBEigFFkJlhftfqiT2ml82pYI1ZUmuuR3N1zaAQNYZvg3asANmoDVGmJYnMdGTyWtD3PPb2t8Nwm57Qd1BfSZIiC7A4cGFSyzYZNp2ObxP4zUeMoa0TPV2WbnLKJ761qP594vMXt9Om4pzdcwK3aAWHQd"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountEmail must be 255 characters or fewer."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaVirtualAccEmailIsInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountEmail = "doku-golang-library@mailcom"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountEmail is not in a valid email format"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaVirtualAccPhoneLessThan9(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountPhone = "12345678"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountPhone must be at least 9 characters long."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaVirtualAccPhoneMoreThan30(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountPhone = "1234567890123456789012345678901"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "virtualAccountPhone must be 30 characters or fewer. Ensure that virtualAccountPhone is no longer than 30 characters."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTrxIdIsLessThan1(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.TrxId = ""
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TrxId must be at least 1 character long. Ensure that TrxId is not empty. Example: '23219829713'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTrxIdIsMoreThan64(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.TrxId = "FcGcsrqYNNQotmv7b2dSFdVbUmiexl0s1wE7H23gpXsFzcXUHXXnRLBUuREMuWxVx"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TrxId must be 64 characters or fewer. Ensure that TrxId is no longer than 64 characters. Example: '23219829713'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTotalAmountIsLessThan4(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.TotalAmount.Value = "0"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Value must be at least 4 characters long and formatted as 0.00."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTotalAmountIsMoreThan19(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.TotalAmount.Value = "12345678901234567890"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Value must be 19 characters or fewer and formatted as 9999999999999999.99. Ensure that TotalAmount.Value is no longer than 19 characters and in the correct format"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTotalAmountIsInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.TotalAmount.Value = "1000E"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Value is invalid format."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTotalAmountCurrencyNot3Char(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.TotalAmount.Currency = "ID"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.Currency must be exactly 3 characters long."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTotalAmountCurrencyNotIDR(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.TotalAmount.Currency = "INR"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TotalAmount.currency must be 'IDR'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaChannelIsLessThan1(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.AdditionalInfo.Channel = ""
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.Channel must be at least 1 character long"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaChannelIsMoreThan30(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.AdditionalInfo.Channel = "VIRTUAL_ACCOUNT_BANK_MANDIRI_TEST"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.Channel must be 30 characters or fewer."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaChannelIsInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.AdditionalInfo.Channel = "5Vl3mjMJpA6NuUNHWrucSymfjlWPCb"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.channel is not valid. Ensure that AdditionalInfo.channel is one of the valid channels."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTrxTypeIsNot1Digit(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountTrxType = "CO"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountTrxType must be exactly 1 character long. Ensure that VirtualAccountTrxType is either 'C', 'O' or 'V'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaTrxTypeIsInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.VirtualAccountTrxType = "A"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountTrxType must be either 'C', 'O' or 'V'."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestUpdateVaExpiredIsInvalidFormat(t *testing.T) {

	request := mockGenerated.UpdateVaRequestDTO()
	request.ExpiredDate = "2024-07-11"
	err := request.ValidateUpdateVaRequestDTO()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "ExpiredDate must be in ISO-8601 format."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCheckStatusVaSuccess(t *testing.T) {
	doku.VaController = mockController
	mockController.On("DoCheckStatusVa", mockGenerated.CheckStatusVaRequest(), "privateKeyPem", "clientId", "", "", false).Return(mockGenerated.CheckStatusVa())
	snap := doku.Snap{
		PrivateKey:   "privateKeyPem",
		ClientId:     "clientId",
		IsProduction: false,
	}
	actualResponse := snap.CheckStatusVa(mockGenerated.CheckStatusVaRequest())
	assert.Equal(t, "2002600", actualResponse.ResponseCode)
}

func TestCheckStatusVaPartnerIdNot8Digits(t *testing.T) {

	request := mockGenerated.CheckStatusVaRequest()
	request.PartnerServiceId = "1234"
	err := request.ValidateCheckStatusVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must be exactly 8 characters long and equiped with left-padded spaces. Example: '  888994"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCheckStatusVaPartnerIdInvalidFormat(t *testing.T) {

	request := mockGenerated.CheckStatusVaRequest()
	request.PartnerServiceId = "    890E"
	err := request.ValidateCheckStatusVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must consist of up to 8 digits of character. Remaining space in case of partner serivce id is less than 8 must be filled with spaces. Example: ' 888994' (2 spaces and 6 digits)"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCheckStatusVaCustomerNoInvalidLength(t *testing.T) {

	request := mockGenerated.CheckStatusVaRequest()
	request.CustomerNo = "1234567890123456789011"
	err := request.ValidateCheckStatusVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must be 20 characters or fewer. Ensure that customerNo is no longer than 20 characters. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCheckStatusVaCustomerNoInvalidFormat(t *testing.T) {

	request := mockGenerated.CheckStatusVaRequest()
	request.CustomerNo = "1234567E"
	err := request.ValidateCheckStatusVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must consist of only digits. Ensure that customerNo contains only numbers. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCheckStatusVaInquiryRequestIdInvalidLength(t *testing.T) {

	request := mockGenerated.CheckStatusVaRequest()
	inquiryRequestId := "CIwxu2v0XgURbX2RYclSfsw4N6fd29YIgvgv1LJpkmSPItG7jrC8ARlKyRhfkgiVnSJvKWRBAu8u0wPyGg0N8mWA8vcSCEvcYsVWut7NNctBkNLT6Le2rBRiEMchWfv4z"
	request.InquiryRequestId = &inquiryRequestId
	err := request.ValidateCheckStatusVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "InquiryRequestId must be 128 characters or fewer. Ensure that InquiryRequestId is no longer than 128 characters"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestCheckStatusVaPaymentRequestIdInvalidLength(t *testing.T) {

	request := mockGenerated.CheckStatusVaRequest()
	paymentRequestId := "CIwxu2v0XgURbX2RYclSfsw4N6fd29YIgvgv1LJpkmSPItG7jrC8ARlKyRhfkgiVnSJvKWRBAu8u0wPyGg0N8mWA8vcSCEvcYsVWut7NNctBkNLT6Le2rBRiEMchWfv4z"
	request.PaymentRequestId = &paymentRequestId
	err := request.ValidateCheckStatusVaRequestDto()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PaymentRequestId must be 128 characters or fewer. Ensure that PaymentRequestId is no longer than 128 characters."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeSuccess(t *testing.T) {
	doku.VaController = mockController
	mockController.On("DoDeletePaymentCode", mockGenerated.DeletePaymentCodeRequest(), "privateKeyPem", "clientId", "", "", false).Return(mockGenerated.DeletePaymentCode())
	snap := doku.Snap{
		PrivateKey:   "privateKeyPem",
		ClientId:     "clientId",
		IsProduction: false,
	}
	actualResponse := snap.DeletePaymentCode(mockGenerated.DeletePaymentCodeRequest())
	assert.Equal(t, "2003100", actualResponse.ResponseCode)
}

func TestDeletePaymentCodeNot8Digits(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.PartnerServiceId = "1234"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must be exactly 8 characters long and equiped with left-padded spaces. Example: '  888994"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}
}

func TestDeletePaymentCodeInvalidFormat(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.PartnerServiceId = "    890E"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "PartnerServiceId must consist of up to 8 digits of character. Remaining space in case of partner serivce id is less than 8 must be filled with spaces. Example: ' 888994' (2 spaces and 6 digits)"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeCustomerNoIsNull(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.CustomerNo = ""
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo cannot be null. Please provide a CustomerNo. "
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeCustomerNoInvalidLength(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.CustomerNo = "1234567890123456789011"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must be 20 characters or fewer. Ensure that customerNo is no longer than 20 characters. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeCustomerNoInvalidFormat(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.CustomerNo = "1234567E"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "CustomerNo must consist of only digits. Ensure that customerNo contains only numbers. Example: '00000000000000000001'"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeVirtualAccountNoIsNull(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.VirtualAccountNo = ""
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountNo cannot be null. Please provide a virtualAccountNo. "
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeVirtualAccountNoInvalidFormat(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.VirtualAccountNo = "    189920240704000"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "VirtualAccountNo must be the concatenation of partnerServiceId and customerNo."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeTrxIdLessThan1(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.TrxId = ""
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TrxId must be at least 1 character long. Ensure that TrxId is not empty. "
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeTrxIdMoreThan64(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.TrxId = "CIwxu2v0XgURbX2RYclSfsw4N6fd29YIgvgv1LJpkmSPItG7jrC8ARlKyRhfkgiVnSJvKWRBAu"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "TrxId must be 64 characters or fewer. Ensure that TrxId is no longer than 64 characters."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeChannelLessThan1(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.AdditionalInfo.Channel = ""
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.Channel must be at least 1 character long. "
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeChannelMoreThan30(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.AdditionalInfo.Channel = "CIwxu2v0XgURbX2RYclSfsw4N6fd29YIgvgv1LJpkmSPItG7jrC8ARlKyRhfkgiVnSJvKWRBAu"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.Channel must be 30 characters or fewer. Ensure that AdditionalInfo.Channel is no longer than 30 characters."
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}

func TestDeletePaymentCodeChannelInvalidFormat(t *testing.T) {

	request := mockGenerated.DeletePaymentCodeRequest()
	request.AdditionalInfo.Channel = "iVnSJvKWRBAu"
	err := request.ValidateDeleteVaRequest()
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	expectedError := "AdditionalInfo.channel is not valid. Ensure that AdditionalInfo.channel is one of the valid channels"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
	}

}
