package services

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	tokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
	notificationTokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/notification/token"
	"github.com/golang-jwt/jwt/v4"
)

var config commons.Config

type TokenServices struct{}

func (ts TokenServices) GenerateTimestamp() string {
	now := time.Now()
	_, offset := now.Zone()
	offsetHours := offset / 3600
	offsetMinutes := (offset % 3600) / 60
	timestamp := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%+03d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), offsetHours, offsetMinutes)
	return timestamp
}

func (ts TokenServices) CreateSignature(privateKeyPem string, clientID string, xTimestamp string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		return "", errors.New("failed to decode PEM block containing private key")
	}

	var rsaPrivateKey *rsa.PrivateKey
	var err error

	if block.Type == "PRIVATE KEY" {
		privateKey, parseErr := x509.ParsePKCS8PrivateKey(block.Bytes)
		if parseErr != nil {
			return "", parseErr
		}

		var ok bool
		rsaPrivateKey, ok = privateKey.(*rsa.PrivateKey)
		if !ok {
			return "", errors.New("not an RSA private key")
		}
	} else if block.Type == "RSA PRIVATE KEY" {
		rsaPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("unsupported private key type")
	}

	stringToSign := clientID + "|" + xTimestamp
	hashed := sha256.Sum256([]byte(stringToSign))
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func (ts TokenServices) CompareSignatures(clientId, timestamp, signature, publicKeyDOKU string) (bool, error) {
	strToSign := fmt.Sprintf("%s|%s", clientId, timestamp)

	block, _ := pem.Decode([]byte(publicKeyDOKU))
	if block == nil {
		fmt.Println("failed to parse PEM block containing the public key")
		return false, errors.New("failed to parse PEM block containing the public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse public key")
		return false, fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("public key is not of type RSA")
		return false, errors.New("public key is not of type RSA")
	}

	hashed := sha256.Sum256([]byte(strToSign))

	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("failed to decode base64 signature")
		return false, fmt.Errorf("failed to decode base64 signature: %v", err)
	}

	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed[:], sigBytes)
	if err != nil {
		fmt.Printf("signature verification failed: %v", err)
		return false, fmt.Errorf("signature verification failed: %v", err)
	}

	return true, nil
}

func (ts TokenServices) GenerateSymetricSignature(httpMethod string, endPointUrl string, tokenB2B string, minifiedRequestBody []byte, timestamp, clientSecret string) string {
	minifiedJson := string(minifiedRequestBody)
	hash := sha256.New()
	hash.Write([]byte(minifiedJson))
	lowercaseHexHash := strings.ToLower(hex.EncodeToString(hash.Sum(nil)))
	strToSign := httpMethod + ":" + endPointUrl + ":" + tokenB2B + ":" + lowercaseHexHash + ":" + timestamp
	hmac := hmac.New(sha512.New, []byte(clientSecret))
	hmac.Write([]byte(strToSign))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))

	return signature
}

func (ts TokenServices) CreateTokenB2BRequestDTO(signature string, timestamp string, clientId string) tokenModels.TokenB2BRequestDTO {
	var tokenB2BRequestDTO = tokenModels.TokenB2BRequestDTO{
		Signature: signature,
		Timestamp: timestamp,
		ClientID:  clientId,
		GrantType: "client_credentials",
	}
	return tokenB2BRequestDTO
}

func (ts TokenServices) CreateTokenB2B(tokenB2BRequestDTO tokenModels.TokenB2BRequestDTO, isProduction bool) tokenModels.TokenB2BResponseDTO {

	baseUrl := config.GetBaseUrl(isProduction) + commons.ACCESS_TOKEN

	var requestBody = map[string]string{
		"grantType": tokenB2BRequestDTO.GrantType,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
	}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
	}

	req.Header = http.Header{
		"X-SIGNATURE":  {tokenB2BRequestDTO.Signature},
		"X-TIMESTAMP":  {tokenB2BRequestDTO.Timestamp},
		"X-CLIENT-KEY": {tokenB2BRequestDTO.ClientID},
		"Content-Type": {"application/json"},
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
	}
	defer res.Body.Close()

	respBody, _ := io.ReadAll(res.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var tokenB2BResponse tokenModels.TokenB2BResponseDTO
	if err := json.Unmarshal(respBody, &tokenB2BResponse); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return tokenB2BResponse
}

func (ts TokenServices) ValidateTokenB2B(requestTokenB2B string, publicKey string) (bool, error) {
	requestTokenB2B = strings.TrimPrefix(requestTokenB2B, "Bearer ")

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return false, fmt.Errorf("invalid public key format")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("failed to parse public key: %w", err)
	}

	_, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("invalid public key type")
	}

	_, _, err = new(jwt.Parser).ParseUnverified(requestTokenB2B, jwt.MapClaims{})
	if err != nil {
		return false, fmt.Errorf("failed to parse token: %w", err)
	}

	if err != nil {
		return false, fmt.Errorf("invalid token: %w", err)
	}

	return true, nil
}

func (ts TokenServices) IsTokenExpired(tokenExpiresIn int, tokenGeneratedTimestamp string) bool {

	now := int(time.Now().Unix())

	timeInt, _ := strconv.Atoi(tokenGeneratedTimestamp)
	var expirationTime = timeInt + tokenExpiresIn

	if expirationTime < now {
		return true
	} else {
		return false
	}
}

func (ts TokenServices) IsTokenEmpty(tokenB2B string) bool {
	return tokenB2B == ""
}

func (ts TokenServices) GenerateToken(expiredIn int64, issuer string, privateKey string, clientId string) string {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		fmt.Println("Failed to parse PEM block containing the private key")
		return ""
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed to parse private key:", err)
		return ""
	}

	rsaPrivateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Failed to cast parsed key to *rsa.PrivateKey")
		return ""
	}

	expiration := time.Now().Unix() + expiredIn
	payload := jwt.MapClaims{
		"exp":      expiration,
		"issuer":   issuer,
		"clientId": clientId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, payload)

	tokenString, err := token.SignedString(rsaPrivateKey)
	if err != nil {
		fmt.Println("Error when converting token to string:", err)
		return ""
	}

	return tokenString
}

func (ts TokenServices) GenerateNotificationTokenDTO(token string, timestamp string, clientId string, expiresIn int) notificationTokenModels.NotificationTokenDTO {
	var tokenHeader = notificationTokenModels.NotificationTokenHeaderDTO{
		XTimeStamp: timestamp,
		XClientKey: clientId,
	}

	var tokenBody = notificationTokenModels.NotificationTokenBodyDTO{
		ResponseCode:    "2007300",
		ResponseMessage: "Successful",
		AccessToken:     token,
		TokenType:       "Bearer",
		ExpiresIn:       expiresIn,
		AdditionalInfo:  "",
	}

	var response = notificationTokenModels.NotificationTokenDTO{
		Header: tokenHeader,
		Body:   tokenBody,
	}

	return response
}

func (ts TokenServices) GenerateInvalidSignature(timestamp string) notificationTokenModels.NotificationTokenDTO {
	var tokenHeader = notificationTokenModels.NotificationTokenHeaderDTO{
		XClientKey: "",
		XTimeStamp: timestamp,
	}

	var tokenBody = notificationTokenModels.NotificationTokenBodyDTO{
		ResponseCode:    "4017300",
		ResponseMessage: "Unauthorized.Invalid Signature",
		AccessToken:     "",
		TokenType:       "",
		ExpiresIn:       0,
		AdditionalInfo:  "",
	}

	var response = notificationTokenModels.NotificationTokenDTO{
		Header: tokenHeader,
		Body:   tokenBody,
	}

	return response
}

func (ts TokenServices) CreateTokenB2B2CRequestDTO(authCode string) tokenModels.TokenB2B2CRequestDTO {
	return tokenModels.TokenB2B2CRequestDTO{
		GrantType: "authorization_code",
		AuthCode:  authCode,
	}
}

func (ts TokenServices) HitTokenB2B2CApi(tokenB2B2CRequestDTO tokenModels.TokenB2B2CRequestDTO, timestamp string, signature string, clientId string, isProduction bool) (tokenModels.TokenB2B2CResponseDTO, error) {
	url := config.GetBaseUrl(isProduction) + commons.ACCESS_TOKEN_B2B2C
	header := map[string]string{
		"X-TIMESTAMP":  timestamp,
		"X-SIGNATURE":  signature,
		"X-CLIENT-KEY": clientId,
		"Content-Type": "application/json",
	}

	bodyRequest, err := json.Marshal(tokenB2B2CRequestDTO)
	if err != nil {
		return tokenModels.TokenB2B2CResponseDTO{}, fmt.Errorf("error body response: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return tokenModels.TokenB2B2CResponseDTO{}, fmt.Errorf("error body request: %w", err)
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return tokenModels.TokenB2B2CResponseDTO{}, fmt.Errorf("error response: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("RESPONSE: ", string(respBody))

	var tokenB2B2CResponseDTO tokenModels.TokenB2B2CResponseDTO
	if err := json.Unmarshal(respBody, &tokenB2B2CResponseDTO); err != nil {
		return tokenModels.TokenB2B2CResponseDTO{}, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	return tokenB2B2CResponseDTO, nil
}
