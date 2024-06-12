package services

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models"
)

var config commons.Config

type TokenService struct{}

func (ts TokenService) GenerateTimestamp() string {
	now := time.Now()
	_, offset := now.Zone()
	offsetHours := offset / 3600
	offsetMinutes := (offset % 3600) / 60
	timestamp := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%+03d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), offsetHours, offsetMinutes)
	return timestamp
}

func (ts TokenService) CreateSignature(privateKeyPem string, clientID string, xTimestamp string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil || block.Type != "PRIVATE KEY" {
		return "", errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not an RSA private key")
	}

	stringToSign := clientID + "|" + xTimestamp
	hashed := sha256.Sum256([]byte(stringToSign))
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), err
}

func (ts TokenService) CreateTokenB2BRequestDTO(signature string, timestamp string, clientId string) models.TokenB2BRequestDTO {
	var tokenB2BRequestDTO = models.TokenB2BRequestDTO{
		Signature: signature,
		Timestamp: timestamp,
		ClientID:  clientId,
		GrantType: "client_credentials",
	}
	return tokenB2BRequestDTO
}

func (ts TokenService) CreateTokenB2B(tokenB2BRequestDTO models.TokenB2BRequestDTO, isProduction bool) models.TokenB2BResponseDTO {

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

	var tokenB2BResponse models.TokenB2BResponseDTO
	err = json.NewDecoder(res.Body).Decode(&tokenB2BResponse)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println("Response :::", tokenB2BResponse)
	return tokenB2BResponse
}
