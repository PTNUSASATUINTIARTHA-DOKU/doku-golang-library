package utils

import (
	"fmt"
	"math/rand"
	"time"

	createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"
)

type SnapUtils struct{}

func (su SnapUtils) GenerateExternalId() string {

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
