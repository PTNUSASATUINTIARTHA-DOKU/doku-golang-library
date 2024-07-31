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
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	tokenModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/token"
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

	var tokenB2BResponse tokenModels.TokenB2BResponseDTO
	err = json.NewDecoder(res.Body).Decode(&tokenB2BResponse)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	return tokenB2BResponse
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
