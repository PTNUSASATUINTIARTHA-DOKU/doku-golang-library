package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/commons"
	"github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models"
)

var tokenServices TokenServices

type VaServices struct{}

func generateExternalId() string {

	uuid := fmt.Sprintf("%x-%x-4%x-y%x-%x",
		rand.Int63n(0x100000000),
		rand.Int63n(0x10000),
		rand.Int63n(0x1000),
		rand.Int63n(0x1000),
		rand.Int63n(0x1000000000000),
	)

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	externalId := fmt.Sprintf("%s-%d", uuid, timestamp)
	return externalId
}

func (vs VaServices) CreateVaRequestHeaderDto(
	createVaRequestDto models.CreateVaRequestDto,
	privateKey string,
	timestamp string,
	clientId string,
	tokenB2B string) models.RequestHeaderDTO {

	var createSignature, _ = tokenServices.CreateSignature(privateKey, clientId, timestamp)

	return models.RequestHeaderDTO{
		XTimestamp:    timestamp,
		XSignature:    createSignature,
		XPartnerId:    clientId,
		XExternalId:   generateExternalId(),
		ChannelId:     createVaRequestDto.AdditionalInfo.Channel,
		Authorization: tokenB2B,
	}
}

func (vs VaServices) CreateVa(
	requestHeaderDto models.RequestHeaderDTO,
	createVaRequestDto models.CreateVaRequestDto,
	isProduction bool,
) models.CreateVaResponseDto {

	url := config.GetBaseUrl(isProduction) + commons.CREATE_VA

	header := map[string]string{
		"X-PARTNER-ID":  requestHeaderDto.XPartnerId,
		"X-TIMESTAMP":   requestHeaderDto.XTimestamp,
		"X-SIGNATURE":   requestHeaderDto.XSignature,
		"Authorization": "Bearer " + requestHeaderDto.Authorization,
		"X-EXTERNAL-ID": requestHeaderDto.XExternalId,
		"CHANNEL-ID":    requestHeaderDto.ChannelId,
		"Content-Type":  "application/json",
	}

	bodyRequest, err := json.Marshal(createVaRequestDto)
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
	var createVaResponseDTO models.CreateVaResponseDto
	if err := json.Unmarshal(respBody, &createVaResponseDTO); err != nil {
		fmt.Println("error unmarshaling response JSON: ", err)
	}

	return createVaResponseDTO

}
