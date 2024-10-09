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
	mockController.On("CreateVa", mockGenerated.CreateVaRequestDTO(), "", "clientId", "TOKEN_B2B", false).Return(mockGenerated.CreateVaResponseDTO())
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

// Direct Debit Unit Test
func TestDirectDebit(t *testing.T) {

	t.Run("TestAccountBindingSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoAccountBinding", mockGenerated.AccountBindingRequest(), "secretKey", "clientId", "deviceId", "ipAddress", "TOKEN_B2B", false).Return(mockGenerated.AccountBindingResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		actualResponse, _ := snap.DoAccountBinding(mockGenerated.AccountBindingRequest(), "deviceId", "ipAddress")
		assert.Equal(t, "2000700", actualResponse.ResponseCode)
	})

	t.Run("TestAccountBindingPhoneNumberIsNull", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.PhoneNo = ""
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "phoneNo cannot be null. Please provide a phoneNo. Example: '62813941306101'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountBindingPhoneNumberLessThan9", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.PhoneNo = "12345678"
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "phoneNo must be at least 9 digits. Example: '62813941306101'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountBindingPhoneNumberMoreThan18", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.PhoneNo = "123456789123456788"
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "phoneNo must be 16 characters or fewer. Example: '62813941306101'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountBindingChannelInvalid", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.AdditionalInfo.Channel = ""
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.channel is not valid. Ensure it is one of the valid channels like 'DIRECT_DEBIT_ALLO_SNAP'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountBindingCustIdMerchantisNull", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.AdditionalInfo.CustIdMerchant = ""
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.custIdMerchant cannot be null. Example: 'cust-001'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountBindingCustIdMerchantMoreThan64", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.AdditionalInfo.CustIdMerchant = "INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100INV100"
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.custIdMerchant must be 64 characters or fewer. Example: 'cust-001'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountBindingUrlSuccessIsNull", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.AdditionalInfo.SuccessRegistrationUrl = ""
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.SuccessRegistrationUrl cannot be null. Example: 'https://www.doku.com'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountBindingUrlFailedIsNull", func(t *testing.T) {
		request := mockGenerated.AccountBindingRequest()
		request.AdditionalInfo.FailedRegistrationUrl = ""
		err := request.ValidateAccountBindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.FailedRegistrationUrl cannot be null. Example: 'https://www.doku.com'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountUnbindingChannelInvalid", func(t *testing.T) {
		request := mockGenerated.AccountUnbindingRequest()
		request.AdditionalInfo.Channel = ""
		err := request.ValidateAccountUnbindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestAccountUnbindingSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoAccountUnbinding", mockGenerated.AccountUnbindingRequest(), "secretKey", "clientId", "ipAddress", "TOKEN_B2B", false).Return(mockGenerated.AccountUnbindingResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		actualResponse, _ := snap.DoAccountUnbinding(mockGenerated.AccountUnbindingRequest(), "ipAddress")
		assert.Equal(t, "2000900", actualResponse.ResponseCode)
	})

	t.Run("TestAccountUnbindingTokenInvalid", func(t *testing.T) {
		request := mockGenerated.AccountUnbindingRequest()
		request.TokenId = "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8w"
		err := request.ValidateAccountUnbindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "tokenId must be 2048 characters or fewer. Ensure that tokenId is no longer than 2048 characters. Example: 'eyJhbGciOiJSUzI1NiJ...'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestBalanceInquirySuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoBalanceInquiry", mockGenerated.BalanceInquiryRequest(), "secretKey", "clientId", "ipAddress", "TOKEN_B2B", "TOKEN_B2B2C", false).Return(mockGenerated.BalanceInquiryResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		snap.SetTokenB2B2C(mockGenerated.GetTokenB2B2CResponseDTO("2007400"))
		actualResponse, _ := snap.DoBalanceInquiry(mockGenerated.BalanceInquiryRequest(), "deviceId", "ipAddress", "authCode")
		assert.Equal(t, "2001100", actualResponse.ResponseCode)
	})

	t.Run("TestBalanceInquiryChannelInvalid", func(t *testing.T) {
		request := mockGenerated.BalanceInquiryRequest()
		request.AdditionalInfo.Channel = ""
		err := request.ValidateBalanceInquiryRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCheckStatusDirectDebitSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoCheckStatus", mockGenerated.CheckStatusRequest(), "secretKey", "clientId", "TOKEN_B2B", false).Return(mockGenerated.CheckStatusResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		actualResponse, _ := snap.DoCheckStatus(mockGenerated.CheckStatusRequest())
		assert.Equal(t, "2005500", actualResponse.ResponseCode)
	})

	t.Run("TestCheckStatusDirectDebitAndEwalletWithInvalidServiceCode", func(t *testing.T) {
		request := mockGenerated.CheckStatusRequest()
		request.ServiceCode = ""
		err := request.ValidateCheckStatusRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "serviceCode must be 55"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCheckStatusEwallettSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CheckStatusRequest()
		request.AdditionalInfo.Channel = "EMONEY_OVO_SNAP"
		mockController.On("DoCheckStatus", mockGenerated.CheckStatusRequest(), "secretKey", "clientId", "TOKEN_B2B", false).Return(mockGenerated.CheckStatusResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		actualResponse, _ := snap.DoCheckStatus(mockGenerated.CheckStatusRequest())
		assert.Equal(t, "2005500", actualResponse.ResponseCode)
	})

	t.Run("TestCardRegistrationSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoCardRegistration", mockGenerated.CardRegistrationRequest(), "secretKey", "clientId", "channelId", "TOKEN_B2B", false).Return(mockGenerated.CardRegistrationResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		actualResponse, _ := snap.DoCardRegistration(mockGenerated.CardRegistrationRequest(), "channelId")
		assert.Equal(t, "2000100", actualResponse.ResponseCode)
	})

	t.Run("TestCardRegistrationCardDataInvalid", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationRequest()
		request.CardData = ""
		err := request.ValidateCardRegistrationRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "cardData cannot be null. Please provide cardData. Example: '5cg2G2719+jxU1RfcGmeCyQrLagUaAWJWWhLpmmb'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCardRegistrationCustIdMerchantNull", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationRequest()
		request.CustIdMerchant = ""
		err := request.ValidateCardRegistrationRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "custIdMerchant cannot be null"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCardRegistrationCustIdMerchantMoreThan64", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationRequest()
		request.CustIdMerchant = "12345678901234567890123456789012345678901234567890123456789011234567890123456789012345678901234567890123456789012345678901"
		err := request.ValidateCardRegistrationRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "custIdMerchant must be 64 characters or fewer"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCardRegistrationChannelInvalid", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationRequest()
		request.AdditionalInfo.Channel = ""
		err := request.ValidateCardRegistrationRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.channel is not valid"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCardRegistrationUrlSuccessNull", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationRequest()
		request.AdditionalInfo.SuccessRegistrationUrl = ""
		err := request.ValidateCardRegistrationRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.SuccessRegistrationUrl cannot be null. Example: 'https://www.doku.com'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCardRegistrationUrlFailedNull", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationRequest()
		request.AdditionalInfo.FailedRegistrationUrl = ""
		err := request.ValidateCardRegistrationRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.FailedRegistrationUrl cannot be null. Example: 'https://www.doku.com'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCardRegistrationUnbindingSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoCardRegistrationUnbinding", mockGenerated.CardRegistrationUnbindingRequest(), "secretKey", "clientId", "ipAddress", "TOKEN_B2B", false).Return(mockGenerated.CardRegistrationUnbindingResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		actualResponse, _ := snap.DoCardRegistrationUnbinding(mockGenerated.CardRegistrationUnbindingRequest(), "ipAddress")
		assert.Equal(t, "2000500", actualResponse.ResponseCode)
	})

	t.Run("TestCardRegistrationUnbindingTokenIdInvalid", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationUnbindingRequest()
		request.TokenId = "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8weyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE3MjkwNTA5ODcsImlzcyI6IkRPS1UiLCJjbGllbnRJZCI6IkJSTi0wMjA1LTE3MjcxNzQ2Mzk4MDEiLCJhY2NvdW50SWQiOiIyZjAzYjY1ODdlZDE5ZmJkYjE5MTJjOTEzNzljMTEwZiJ9.S0kMrOOR8_kur2iU3YbzlMkVtWexM_jziYFY1uaJI3bYdNZ7TnPD1ZYOI-_v4tzQn6on0Rozp00_WdQFMmdoXu9lIBkprEz9e2rN2_tg1tUSXPG6SW5umgf9IV0n1Ro2M5Xfvh4zRFboAU4SvqlSbVM57Vk0LBMTWn8ah0NaBIL40p-UB1UfZ8q5-jyshFszS7S59c21fMA8FXFH_Zz6hjk7HWaAjYPPRmuAkEs3liWYaAoGS_eHL0p_t_IpBlOBsMe6dJhpeDllwole7sptJ1Hckux6mSB4zIKUIrQHZW8F3hTCV1Mx1Hkome7e_6f0VJDsclXbe48xWVtidd2C8w"
		err := request.ValidateCardRegistrationUnbindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "tokenId must be 2048 characters or fewer"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestCardRegistrationUnbindingChannelInvalid", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.CardRegistrationUnbindingRequest()
		request.AdditionalInfo.Channel = ""
		err := request.ValidateCardRegistrationUnbindingRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.channel is not valid"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoRefund", mockGenerated.RefundRequest(), "secretKey", "clientId", "ipAddress", "TOKEN_B2B", "TOKEN_B2B2C", "deviceId", false).Return(mockGenerated.RefundResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		snap.SetTokenB2B2C(mockGenerated.GetTokenB2B2CResponseDTO("2007400"))
		actualResponse, _ := snap.DoRefund(mockGenerated.RefundRequest(), "ipAddress", "auth", "deviceId")
		assert.Equal(t, "2000700", actualResponse.ResponseCode)
	})

	t.Run("TestRefundChannelInvalid", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.RefundRequest()
		request.AdditionalInfo.Channel = ""
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.channel is not valid"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundAmountLessThan1", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.RefundRequest()
		request.RefundAmount.Value = ""
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "refundAmount.Value must be at least 4 characters long and formatted as 0.00."
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundAmountMoreThan19", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.RefundRequest()
		request.RefundAmount.Value = "99999999999999999.00"
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "refundAmount.Value must be 19 characters or fewer and formatted as 9999999999999999.99. Example: '9999999999999999.99'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundAmountInvalidFormat", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.RefundRequest()
		request.RefundAmount.Value = "999900"
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "refundAmount.Value is in invalid format"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundCurrencyLessThan1", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.RefundRequest()
		request.RefundAmount.Currency = ""
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "must be a string; ensure that refundAmount.Currency is enclosed in quotes, e.g., 'IDR'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundCurrencyLessThan3", func(t *testing.T) {
		request := mockGenerated.RefundRequest()
		request.RefundAmount.Currency = "ID"
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "refundAmount.currency must be exactly 3 characters long, e.g., 'IDR'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundCurrencyInvalidFormat", func(t *testing.T) {
		request := mockGenerated.RefundRequest()
		request.RefundAmount.Currency = "JPY"
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "refundAmount.currency must be 'IDR', e.g., 'IDR'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundOriginalPartnerReferenceNoIsNull", func(t *testing.T) {
		request := mockGenerated.RefundRequest()
		request.OriginalPartnerReferenceNo = ""
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "originalPartnerReferenceNo cannot be null"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundOriginalPartnerReferenceNoMoreThan12", func(t *testing.T) {
		request := mockGenerated.RefundRequest()
		request.OriginalPartnerReferenceNo = "1234567890123"
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "originalPartnerReferenceNo must be 12 characters or fewer."
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundPartnerRefundNoIsNull", func(t *testing.T) {
		request := mockGenerated.RefundRequest()
		request.PartnerRefundNo = ""
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "partnerRefundNo cannot be null."
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestRefundPartnerRefundNoMoreThan12", func(t *testing.T) {
		request := mockGenerated.RefundRequest()
		request.PartnerRefundNo = "1234567890123"
		err := request.ValidateRefundRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "partnerRefundNo must be 12 characters or fewer."
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentDirectDebitSuccess", func(t *testing.T) {
		doku.DirectDebitController = mockController
		mockController.On("DoPayment", mockGenerated.PaymentDirectDebitRequest(), "secretKey", "clientId", "ipAddress", "TOKEN_B2B2C", "TOKEN_B2B", false).Return(mockGenerated.PaymentDirectDebitResponse())
		snap := doku.Snap{
			PrivateKey:   "privateKeyPem",
			ClientId:     "clientId",
			SecretKey:    "secretKey",
			IsProduction: false,
		}
		snap.SetTokenB2B(mockGenerated.GetTokenB2BResponseDTO("2007300"))
		snap.SetTokenB2B2C(mockGenerated.GetTokenB2B2CResponseDTO("2007400"))
		actualResponse, _ := snap.DoPayment(mockGenerated.PaymentDirectDebitRequest(), "ipAddress", "")
		assert.Equal(t, "2005400", actualResponse.ResponseCode)
	})

	t.Run("TestPaymentAmountLessThan1", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.PaymentDirectDebitRequest()
		request.Amount.Value = ""
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "Amount.Value must be at least 4 characters long and formatted as 0.00."
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentAmountMoreThan19", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.PaymentDirectDebitRequest()
		request.Amount.Value = "99999999999999999.00"
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "Amount.Value must be 19 characters or fewer and formatted as 9999999999999999.99. Example: '9999999999999999.99'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentChannelInvalid", func(t *testing.T) {
		request := mockGenerated.PaymentDirectDebitRequest()
		request.AdditionalInfo.Channel = ""
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.channel is not valid. Ensure that additionalInfo.channel is one of the valid channels. Example: 'DIRECT_DEBIT_ALLO_SNAP"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentCurrencyLessThan1", func(t *testing.T) {
		doku.DirectDebitController = mockController
		request := mockGenerated.PaymentDirectDebitRequest()
		request.Amount.Currency = ""
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "must be a string; ensure that refundAmount.Currency is enclosed in quotes, e.g., 'IDR'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentCurrencyLessThan3", func(t *testing.T) {
		request := mockGenerated.PaymentDirectDebitRequest()
		request.Amount.Currency = "ID"
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "refundAmount.currency must be exactly 3 characters long, e.g., 'IDR'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentCurrencyInvalidFormat", func(t *testing.T) {
		request := mockGenerated.PaymentDirectDebitRequest()
		request.Amount.Currency = "JPY"
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "refundAmount.currency must be 'IDR', e.g., 'IDR'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentUrlSuccessIsNull", func(t *testing.T) {
		request := mockGenerated.PaymentDirectDebitRequest()
		request.AdditionalInfo.SuccessPaymentUrl = ""
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.SuccessPaymentUrl cannot be null. Example: 'https://www.doku.com'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

	t.Run("TestPaymentUrlFailedIsNull", func(t *testing.T) {
		request := mockGenerated.PaymentDirectDebitRequest()
		request.AdditionalInfo.FailedPaymentUrl = ""
		err := request.ValidatePaymentRequest()
		if err == nil {
			t.Fatalf("expected an error, got nil")
		}
		expectedError := "additionalInfo.FailedPaymentUrl cannot be null. Example: 'https://www.doku.com'"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("expected error message to contain '%s', got '%s'", expectedError, err.Error())
		}
	})

}
