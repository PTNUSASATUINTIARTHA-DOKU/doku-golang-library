package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"
	accountUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountunbinding"
	balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"
	cardRegistrationModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistration"
	registrationCardUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistrationunbinding"
	checkStatusModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/checkstatus"
	jumpAppModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/jumpapp"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type DirectDebitService struct{}

func (dd *DirectDebitService) DoAccountBindingProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, accountBindingRequestDTO accountBindingModels.AccountBindingRequestDTO, isProduction bool) (accountBindingModels.AccountBindingResponseDTO, error) {
	if err := requestHeaderDTO.ValidateAccountBinding(accountBindingRequestDTO.AdditionalInfo.Channel); err != nil {
		return accountBindingModels.AccountBindingResponseDTO{}, err
	}
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_ACCOUNT_BINDING
	header := map[string]string{
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"X-DEVICE-ID":   requestHeaderDTO.XDeviceId,
		"X-IP-ADDRESS":  requestHeaderDTO.XIpAddress,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(accountBindingRequestDTO)
	if err != nil {
		return accountBindingModels.AccountBindingResponseDTO{}, fmt.Errorf("error body response: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return accountBindingModels.AccountBindingResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return accountBindingModels.AccountBindingResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var accountBindingResponseDTO accountBindingModels.AccountBindingResponseDTO
	if err := json.Unmarshal(respBody, &accountBindingResponseDTO); err != nil {
		return accountBindingModels.AccountBindingResponseDTO{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	return accountBindingResponseDTO, nil
}

func (dd *DirectDebitService) DoBalanceInquiryProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, isProduction bool) (balanceInquiryModels.BalanceInquiryResponseDto, error) {
	if err := requestHeaderDTO.ValidateCheckBalance(balanceInquiryRequestDto.AdditionalInfo.Channel); err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{}, err
	}
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_BALANCE_INQUIRY_URL
	header := map[string]string{
		"X-TIMESTAMP":            requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":            requestHeaderDTO.XSignature,
		"X-PARTNER-ID":           requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID":          requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":           requestHeaderDTO.XIpAddress,
		"Authorization-customer": "Bearer " + requestHeaderDTO.AuthorizationCustomer,
		"Authorization":          "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":           "application/json",
	}

	bodyRequest, err := json.Marshal(balanceInquiryRequestDto)
	if err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{}, fmt.Errorf("error marshal body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var balanceInquiryResponseDto balanceInquiryModels.BalanceInquiryResponseDto
	if err := json.Unmarshal(respBody, &balanceInquiryResponseDto); err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}
	return balanceInquiryResponseDto, nil
}

func (dd *DirectDebitService) DoPaymentProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, paymentRequestDTO paymentModels.PaymentRequestDTO, isProduction bool) (paymentModels.PaymentResponseDTO, error) {
	if err := requestHeaderDTO.ValidatePaymentAndRefund(paymentRequestDTO.AdditionalInfo.Channel); err != nil {
		return paymentModels.PaymentResponseDTO{}, err
	}
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_PAYMENT
	header := map[string]string{
		"X-TIMESTAMP":            requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":            requestHeaderDTO.XSignature,
		"X-PARTNER-ID":           requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID":          requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":           requestHeaderDTO.XIpAddress,
		"CHANNEL-ID":             requestHeaderDTO.ChannelId,
		"Authorization-Customer": "Bearer " + requestHeaderDTO.AuthorizationCustomer,
		"Authorization":          "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":           "application/json",
	}

	bodyRequest, err := json.Marshal(paymentRequestDTO)
	if err != nil {
		return paymentModels.PaymentResponseDTO{}, fmt.Errorf("error marshal body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return paymentModels.PaymentResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return paymentModels.PaymentResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var paymentResponse paymentModels.PaymentResponseDTO
	if err := json.Unmarshal(respBody, &paymentResponse); err != nil {
		return paymentModels.PaymentResponseDTO{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}
	return paymentResponse, nil
}

func (dd *DirectDebitService) DoAccountUnbindingProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, isProduction bool) (accountUnbindingModels.AccountUnbindingResponseDTO, error) {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_ACCOUNT_UNBINDING
	header := map[string]string{
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":  requestHeaderDTO.XIpAddress,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(accountUnbindingRequestDTO)
	if err != nil {
		return accountUnbindingModels.AccountUnbindingResponseDTO{}, fmt.Errorf("error body response: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return accountUnbindingModels.AccountUnbindingResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return accountUnbindingModels.AccountUnbindingResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var accountUnbindingResponse accountUnbindingModels.AccountUnbindingResponseDTO
	if err := json.Unmarshal(respBody, &accountUnbindingResponse); err != nil {
		return accountUnbindingModels.AccountUnbindingResponseDTO{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}
	return accountUnbindingResponse, nil
}

func (dd *DirectDebitService) DoPaymentJumpAppProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, isProduction bool) (jumpAppModels.PaymentJumpAppResponseDTO, error) {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_PAYMENT
	header := map[string]string{
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":  requestHeaderDTO.XIpAddress,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"CHANNEL-ID":    requestHeaderDTO.ChannelId,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(paymentJumpAppRequestDTO)
	if err != nil {
		return jumpAppModels.PaymentJumpAppResponseDTO{}, fmt.Errorf("error marshal body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return jumpAppModels.PaymentJumpAppResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return jumpAppModels.PaymentJumpAppResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var paymentJumpAppResponse jumpAppModels.PaymentJumpAppResponseDTO
	if err := json.Unmarshal(respBody, &paymentJumpAppResponse); err != nil {
		return jumpAppModels.PaymentJumpAppResponseDTO{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}
	return paymentJumpAppResponse, nil
}

func (dd *DirectDebitService) DoCardRegistrationProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, isProduction bool) (cardRegistrationModels.CardRegistrationResponseDTO, error) {
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_CARD_REGISTRATION
	header := map[string]string{
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"CHANNEL-ID":    requestHeaderDTO.ChannelId,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(cardRegistrationRequestDTO)
	if err != nil {
		return cardRegistrationModels.CardRegistrationResponseDTO{}, fmt.Errorf("error marshal body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return cardRegistrationModels.CardRegistrationResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return cardRegistrationModels.CardRegistrationResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var cardRegistrationResponseDTO cardRegistrationModels.CardRegistrationResponseDTO
	if err := json.Unmarshal(respBody, &cardRegistrationResponseDTO); err != nil {
		return cardRegistrationModels.CardRegistrationResponseDTO{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}
	return cardRegistrationResponseDTO, nil
}

func (dd *DirectDebitService) DoRefundProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, refundRequestDTO refundModels.RefundRequestDTO, isProduction bool) (refundModels.RefundResponseDTO, error) {
	if err := requestHeaderDTO.ValidatePaymentAndRefund(refundRequestDTO.AdditionalInfo.Channel); err != nil {
		return refundModels.RefundResponseDTO{}, err
	}
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_REFUND
	header := map[string]string{
		"X-TIMESTAMP":            requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":            requestHeaderDTO.XSignature,
		"X-PARTNER-ID":           requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID":          requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":           requestHeaderDTO.XIpAddress,
		"Authorization-Customer": "Bearer " + requestHeaderDTO.AuthorizationCustomer,
		"Authorization":          "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":           "application/json",
	}

	bodyRequest, err := json.Marshal(refundRequestDTO)
	if err != nil {
		return refundModels.RefundResponseDTO{}, fmt.Errorf("error marshal body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return refundModels.RefundResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return refundModels.RefundResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return refundModels.RefundResponseDTO{}, fmt.Errorf("error reading response body: %w", err)
	}
	fmt.Println("RESPONSE: ", string(respBody))

	var refundResponseDTO refundModels.RefundResponseDTO
	if err := json.Unmarshal(respBody, &refundResponseDTO); err != nil {
		return refundModels.RefundResponseDTO{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return refundResponseDTO, nil
}

func (dd *DirectDebitService) DoCheckStatusProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, checkStatusRequestDTO checkStatusModels.CheckStatusRequestDTO, isProduction bool) (checkStatusModels.CheckStatusResponseDTO, error) {

	var checkStatusResponse checkStatusModels.CheckStatusResponseDTO

	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_CHECK_STATUS

	header := map[string]string{
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(checkStatusRequestDTO)
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{}, fmt.Errorf("error marshal body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{}, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(respBody, &checkStatusResponse); err != nil {
		return checkStatusModels.CheckStatusResponseDTO{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return checkStatusResponse, nil
}

func (dd *DirectDebitService) DoCardRegistrationUnbindingProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, registrationCardUnbindingRequestDTO registrationCardUnbindingModels.CardRegistrationUnbindingRequestDTO, isProduction bool) (registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO, error) {
	if err := requestHeaderDTO.ValidateAccountUnbinding(registrationCardUnbindingRequestDTO.AdditionalInfo.Channel); err != nil {
		return registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO{}, err
	}
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_CARD_UNBINDING
	header := map[string]string{
		"X-TIMESTAMP":   requestHeaderDTO.XTimestamp,
		"X-SIGNATURE":   requestHeaderDTO.XSignature,
		"X-PARTNER-ID":  requestHeaderDTO.XPartnerId,
		"X-EXTERNAL-ID": requestHeaderDTO.XExternalId,
		"X-IP-ADDRESS":  requestHeaderDTO.XIpAddress,
		"Authorization": "Bearer " + requestHeaderDTO.Authorization,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(registrationCardUnbindingRequestDTO)
	if err != nil {
		return registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO{}, fmt.Errorf("error marshal body request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO{}, fmt.Errorf("error creating HTTP request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO{}, fmt.Errorf("error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var cardRegistrationUnbindingResponse registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO
	if err := json.Unmarshal(respBody, &cardRegistrationUnbindingResponse); err != nil {
		return registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}
	return cardRegistrationUnbindingResponse, nil
}
