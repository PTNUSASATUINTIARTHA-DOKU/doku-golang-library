package utilsmock

import (
	"errors"
	"net/http"

	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	accountUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountunbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	cardRegistrationModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistration"
	registrationCardUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistrationunbinding"
	checkStatusModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/checkstatus"
	jumpAppModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/jumpapp"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"
	models "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
	deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"
	inquiryVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/inquiry"
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

func (m *MockController) ValidateTokenB2B(requestTokenB2B string, publicKey string) (bool, error) {
	return false, nil
}

func (m *MockController) ValidateSignature(request *http.Request, privateKey string, clientId string, publicKeyDOKU string) bool {
	return false
}

func (m *MockController) GenerateTokenB2B(expiredIn int, issuer string, privateKey string, clientId string) notificationTokenModels.NotificationTokenDTO {
	return notificationTokenModels.NotificationTokenDTO{}
}

func (m *MockController) GetTokenB2B2C(authCode string, privateKey string, clientId string, isProduction bool) (models.TokenB2B2CResponseDTO, error) {
	return models.TokenB2B2CResponseDTO{}, nil
}

func (m *MockController) GenerateInvalidSignatureResponse() notificationTokenModels.NotificationTokenDTO {
	return notificationTokenModels.NotificationTokenDTO{}
}

func (m *MockController) DoGenerateRequestHeader(privateKey string, clientId string, tokenB2B string) createVaModels.RequestHeaderDTO {
	return createVaModels.RequestHeaderDTO{}
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
	args := m.Called(checkStatusVARequestDto, privateKey, clientId, tokenB2B, secretKey, isProduction)
	return args.Get(0).(checkVaModels.CheckStatusVaResponseDto)
}

func (m *MockController) DoDeletePaymentCode(deleteVaRequestDto deleteVaModels.DeleteVaRequestDto, privateKey string, clientId string, tokenB2B string, secretKey string, isProduction bool) deleteVaModels.DeleteVaResponseDto {
	args := m.Called(deleteVaRequestDto, privateKey, clientId, tokenB2B, secretKey, isProduction)
	return args.Get(0).(deleteVaModels.DeleteVaResponseDto)
}

func (m *MockController) DirectInquiryRequestMapping(headerRequest *http.Request, inquiryRequestBodyDto inquiryVaModels.InquiryRequestBodyDTO) (string, error) {
	return "", errors.New("mock function unit testing")
}

func (m *MockController) DirectInquiryResponseMapping(xmlData string) (inquiryVaModels.InquiryResponseBodyDTO, error) {
	return inquiryVaModels.InquiryResponseBodyDTO{}, errors.New("mock function unit testing")
}

// End VaController

// DirectDebitController
func (m *MockController) DoAccountBinding(accountBindingRequest accountBindingModels.AccountBindingRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) (accountBindingModels.AccountBindingResponseDTO, error) {
	args := m.Called(accountBindingRequest, secretKey, clientId, deviceId, ipAddress, tokenB2B, isProduction)
	return args.Get(0).(accountBindingModels.AccountBindingResponseDTO), nil
}

func (m *MockController) DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, isProduction bool) (balanceInquiryModels.BalanceInquiryResponseDto, error) {
	args := m.Called(balanceInquiryRequestDto, secretKey, clientId, ipAddress, tokenB2B, tokenB2B2C, isProduction)
	return args.Get(0).(balanceInquiryModels.BalanceInquiryResponseDto), nil
}

func (m *MockController) DoPayment(paymentRequestDTO paymentModels.PaymentRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B2C string, tokenB2B string, isProduction bool) (paymentModels.PaymentResponseDTO, error) {
	args := m.Called(paymentRequestDTO, secretKey, clientId, ipAddress, tokenB2B2C, tokenB2B, isProduction)
	return args.Get(0).(paymentModels.PaymentResponseDTO), nil
}

func (m *MockController) DoAccountUnbinding(accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) (accountUnbindingModels.AccountUnbindingResponseDTO, error) {
	args := m.Called(accountUnbindingRequestDTO, secretKey, clientId, ipAddress, tokenB2B, isProduction)
	return args.Get(0).(accountUnbindingModels.AccountUnbindingResponseDTO), nil
}

func (m *MockController) DoPaymentJumpApp(paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, secretKey string, clientId string, deviceId string, ipAddress string, tokenB2B string, isProduction bool) (jumpAppModels.PaymentJumpAppResponseDTO, error) {
	args := m.Called(paymentJumpAppRequestDTO, secretKey, clientId, deviceId, ipAddress, tokenB2B, isProduction)
	return args.Get(0).(jumpAppModels.PaymentJumpAppResponseDTO), nil
}

func (m *MockController) DoCardRegistration(cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, secretKey string, clientId string, channelId string, tokenB2B string, isProduction bool) (cardRegistrationModels.CardRegistrationResponseDTO, error) {
	args := m.Called(cardRegistrationRequestDTO, secretKey, clientId, channelId, tokenB2B, isProduction)
	return args.Get(0).(cardRegistrationModels.CardRegistrationResponseDTO), nil
}

func (m *MockController) DoRefund(refundRequestDTO refundModels.RefundRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, tokenB2B2C string, deviceId string, isProduction bool) (refundModels.RefundResponseDTO, error) {
	args := m.Called(refundRequestDTO, secretKey, clientId, ipAddress, tokenB2B, tokenB2B2C, deviceId, isProduction)
	return args.Get(0).(refundModels.RefundResponseDTO), nil
}

func (m *MockController) DoCheckStatus(checkStatusRequestDTO checkStatusModels.CheckStatusRequestDTO, secretKey string, clientId string, tokenB2B string, isProduction bool) (checkStatusModels.CheckStatusResponseDTO, error) {
	args := m.Called(checkStatusRequestDTO, secretKey, clientId, tokenB2B, isProduction)
	return args.Get(0).(checkStatusModels.CheckStatusResponseDTO), nil
}

func (m *MockController) DoCardRegistrationUnbinding(cardRegistrationUnbindingRequestDTO registrationCardUnbindingModels.CardRegistrationUnbindingRequestDTO, secretKey string, clientId string, ipAddress string, tokenB2B string, isProduction bool) (registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO, error) {
	args := m.Called(cardRegistrationUnbindingRequestDTO, secretKey, clientId, ipAddress, tokenB2B, isProduction)
	return args.Get(0).(registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO), nil
}
