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
	checkStatusModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/checkstatus"
	registrationCardUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistrationunbinding"
	jumpAppModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/jumpapp"
	paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
	refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"
	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type DirectDebitService struct{}

func (dd *DirectDebitService) DoAccountBindingProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, accountBindingRequestDTO accountBindingModels.AccountBindingRequestDTO, isProduction bool) accountBindingModels.AccountBindingResponseDto {
	if err := requestHeaderDTO.ValidateAccountBinding(accountBindingRequestDTO.AdditionalInfo.Channel); err != nil {
		return accountBindingModels.AccountBindingResponseDto{
			ResponseCode:    "500",
			ResponseMessage: err.Error(),
		}
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
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var accountBindingResponseDTO accountBindingModels.AccountBindingResponseDto
	if err := json.Unmarshal(respBody, &accountBindingResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return accountBindingResponseDTO
}

func (dd *DirectDebitService) DoBalanceInquiryProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, isProduction bool) balanceInquiryModels.BalanceInquiryResponseDto {
	if err := requestHeaderDTO.ValidateCheckBalance(balanceInquiryRequestDto.AdditionalInfo.Channel); err != nil {
		return balanceInquiryModels.BalanceInquiryResponseDto{
			ResponseCode:    "500",
			ResponseMessage: err.Error(),
		}
	}
	url := config.GetBaseUrl(isProduction) + commons.DIRECT_DEBIT_BALANCE_INQUIRY_URL
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

	bodyRequest, err := json.Marshal(balanceInquiryRequestDto)
	if err != nil {
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var balanceInquiryResponseDto balanceInquiryModels.BalanceInquiryResponseDto
	if err := json.Unmarshal(respBody, &balanceInquiryResponseDto); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}
	return balanceInquiryResponseDto
}

func (dd *DirectDebitService) DoPaymentProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, paymentRequestDTO paymentModels.PaymentRequestDTO, isProduction bool) paymentModels.PaymentResponseDTO {
	if err := requestHeaderDTO.ValidatePaymentAndRefund(paymentRequestDTO.AdditionalInfo.Channel); err != nil {
		return paymentModels.PaymentResponseDTO{
			ResponseCode:    "500",
			ResponseMessage: err.Error(),
		}
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
		fmt.Println("Error parse body request :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var paymentResponse paymentModels.PaymentResponseDTO
	if err := json.Unmarshal(respBody, &paymentResponse); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}
	return paymentResponse
}

func (dd *DirectDebitService) DoAccountUnbindingProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, accountUnbindingRequestDTO accountUnbindingModels.AccountUnbindingRequestDTO, isProduction bool) accountUnbindingModels.AccountUnbindingResponseDTO {
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
		fmt.Println("Error parse body request :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var accountUnbindingResponse accountUnbindingModels.AccountUnbindingResponseDTO
	if err := json.Unmarshal(respBody, &accountUnbindingResponse); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}
	return accountUnbindingResponse
}

func (dd *DirectDebitService) DoPaymentJumpAppProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, paymentJumpAppRequestDTO jumpAppModels.PaymentJumpAppRequestDTO, isProduction bool) jumpAppModels.PaymentJumpAppResponseDTO {
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
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var paymentJumpAppResponse jumpAppModels.PaymentJumpAppResponseDTO
	if err := json.Unmarshal(respBody, &paymentJumpAppResponse); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}
	return paymentJumpAppResponse
}

func (dd *DirectDebitService) DoCardRegistrationProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, cardRegistrationRequestDTO cardRegistrationModels.CardRegistrationRequestDTO, isProduction bool) cardRegistrationModels.CardRegistrationResponseDTO {
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
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var cardRegistrationResponseDTO cardRegistrationModels.CardRegistrationResponseDTO
	if err := json.Unmarshal(respBody, &cardRegistrationResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}
	return cardRegistrationResponseDTO
}

func (dd *DirectDebitService) DoRefundProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, refundRequestDTO refundModels.RefundRequestDTO, isProduction bool) refundModels.RefundResponseDTO {
	if err := requestHeaderDTO.ValidatePaymentAndRefund(refundRequestDTO.AdditionalInfo.Channel); err != nil {
		return refundModels.RefundResponseDTO{
			ResponseCode:    "500",
			ResponseMessage: err.Error(),
		}
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
		fmt.Println("Error body response :", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		fmt.Println("Error body request :", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error response :", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var refundResponseDTO refundModels.RefundResponseDTO
	if err := json.Unmarshal(respBody, &refundResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}
	return refundResponseDTO
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
		return checkStatusModels.CheckStatusResponseDTO{
			ResponseCode:    "500",
			ResponseMessage: fmt.Sprintf("error marshalling request body: %v", err),
		}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{
			ResponseCode:    "500",
			ResponseMessage: fmt.Sprintf("error creating HTTP request: %v", err),
		}, err
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{
			ResponseCode:    "500",
			ResponseMessage: fmt.Sprintf("error sending HTTP request: %v", err),
		}, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return checkStatusModels.CheckStatusResponseDTO{
			ResponseCode:    "500",
			ResponseMessage: fmt.Sprintf("error reading response body: %v", err),
		}, err
	}

	if err := json.Unmarshal(respBody, &checkStatusResponse); err != nil {
		return checkStatusModels.CheckStatusResponseDTO{
			ResponseCode:    "500",
			ResponseMessage: fmt.Sprintf("error unmarshalling response body: %v", err),
		}, err
	}

	return checkStatusResponse, nil
}

func (dd *DirectDebitService) DoCardRegistrationUnbindingProcess(requestHeaderDTO createVaModels.RequestHeaderDTO, registrationCardUnbindingRequestDTO registrationCardUnbindingModels.CardRegistrationUnbindingRequestDTO, isProduction bool) registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO {
    if err := requestHeaderDTO.ValidateAccountUnbinding(registrationCardUnbindingRequestDTO.AdditionalInfo.Channel); err != nil {
        return registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO{
            ResponseCode:    "500",
            ResponseMessage: err.Error(),
        }
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
        fmt.Println("Error parse body request :", err)
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
    if err != nil {
        fmt.Println("Error body request :", err)
    }

    for key, value := range header {
        req.Header.Set(key, value)
    }

    client := &http.Client{
        Timeout: time.Second * 30,
    }
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error response :", err)
    }
    defer resp.Body.Close()

    respBody, _ := io.ReadAll(resp.Body)
    fmt.Println("RESPONSE: ", string(respBody))

    var cardRegistrationUnbindingResponse registrationCardUnbindingModels.CardRegistrationUnbindingResponseDTO
    if err := json.Unmarshal(respBody, &cardRegistrationUnbindingResponse); err != nil {
        fmt.Println("error unmarshaling response JSON: ", err)
    }
    return cardRegistrationUnbindingResponse
}