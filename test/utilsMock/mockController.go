package utilsmock

import (
	"net/http"

	models "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"
	notificationTokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/token"
	updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"
	"github.com/stretchr/testify/mock"
)

type MockController struct {
	mock.Mock
}

// TokenController

func (m *MockController) GetTokenB2B(privateKey string, clientId string, isProduction bool) models.TokenB2BResponseDTO {
	args := m.Called(privateKey, clientId, isProduction)
	return args.Get(0).(models.TokenB2BResponseDTO)
}

func (m *MockController) IsTokenInvalid(tokenB2B string, tokenExpiresIn int, tokenGeneratedTimestamp string) bool {
	return false
}

func (m *MockController) ValidateTokenB2B(requestTokenB2B string, publicKey string) bool {
	return false
}

func (m *MockController) ValidateSignature(request *http.Request, privateKey string, clientId string) bool {
	return false
}

func (m *MockController) GenerateTokenB2B(expiredIn int, issuer string, privateKey string, clientId string) notificationTokenModels.NotificationTokenDTO {
	return notificationTokenModels.NotificationTokenDTO{}
}

func (m *MockController) GenerateInvalidSignatureResponse() notificationTokenModels.NotificationTokenDTO {
	return notificationTokenModels.NotificationTokenDTO{}
}

// End TokenController

// VaController

func (m *MockController) CreateVa(createVaRequestDto createVaModels.CreateVaRequestDto, privateKey string, clientId string, tokenB2B string, isProduction bool) createVaModels.CreateVaResponseDto {
	args := m.Called(createVaRequestDto, privateKey, clientId, tokenB2B, isProduction)
	return args.Get(0).(createVaModels.CreateVaResponseDto)
}

func (m *MockController) DoUpdateVa(updateVaRequestDTO updateVaModels.UpdateVaDTO, clientId string, tokenB2B string, secretKey string, isProduction bool) updateVaModels.UpdateVaResponseDTO {
	args := m.Called(updateVaRequestDTO, clientId, tokenB2B, secretKey, isProduction)
	return args.Get(0).(updateVaModels.UpdateVaResponseDTO)
}

func (m *MockController) DoCheckStatusVa(checkStatusVARequestDto checkVaModels.CheckStatusVARequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) checkVaModels.CheckStatusVaResponseDto {
	return checkVaModels.CheckStatusVaResponseDto{}
}

func (m *MockController) DoDeletePaymentCode(deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) deleteVaModels.DeleteVaResponseDto {
	return deleteVaModels.DeleteVaResponseDto{}
}

// End VaController
