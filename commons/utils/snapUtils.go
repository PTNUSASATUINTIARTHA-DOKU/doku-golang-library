package utils

import (
	"math/rand"
	"strconv"
	"time"

	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type SnapUtils struct{}

func (su SnapUtils) GenerateExternalId() string {
	numbers := "0123456789"
	result := make([]byte, 10)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range result {
		result[i] = numbers[r.Intn(len(numbers))]
	}

	timestamp := time.Now().Unix()

	return string(result) + strconv.FormatInt(timestamp, 10)
}

func (su SnapUtils) GenerateRequestHeaderDto(
	channelId string,
	signature string,
	timestamp string,
	clientId string,
	externalId string,
	deviceId string,
	ipAddress string,
	tokenB2B string,
	tokenB2B2C string,
) createVaModels.RequestHeaderDTO {

	return createVaModels.RequestHeaderDTO{
		XTimestamp:            timestamp,
		XSignature:            signature,
		XPartnerId:            clientId,
		XExternalId:           externalId,
		XDeviceId:             deviceId,
		XIpAddress:            ipAddress,
		ChannelId:             channelId,
		Authorization:         tokenB2B,
		AuthorizationCustomer: tokenB2B2C,
	}
}
