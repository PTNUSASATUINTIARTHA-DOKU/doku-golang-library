
# DOKU Golang Library
Welcome to the DOKU Golang library! This powerful tool simplifies access to the DOKU API for your server-side Go applications.

## Documentation
For detailed information, visit the full [DOKU API Docs](https://developers.doku.com/accept-payment/direct-api/snap).

## Requirements
- Go 1.22.2 or higher.

## Installation
Get started by installing the library:


```xml
go get github.com/lib/pq
go get -u github.com/golang-jwt/jwt/v4
go get github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library@latest
```


## Usage
This section will guide you through setting up the DOKU Golang library, creating payment requests, and handling notifications. Let’s get started!

### 1. Configuration
To configure the library, you'll need your account's Client ID, Secret Key, and Private Key. Here’s how:

1. **Client ID and Secret Key:** Retrieve these from the Integration menu in your [DOKU Dashboard](https://dashboard.doku.com/bo/login).
2. **Private Key:** Generate your Private Key following DOKU’s guide and insert the corresponding Public Key into the same menu.

> Your private key will not be transmitted or shared with DOKU. It remains on your server and is only used to sign the requests you send to DOKU.

```go
import "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/doku"

var dokuSnap *doku.Snap
func InitializeSnap() {
    doku.TokenController = controllers.TokenController{}
    doku.VaController = controllers.VaController{}
    dokuSnap = &doku.Snap{
	PrivateKey:   privateKey, 
	ClientId:     clientId, 
	IsProduction: false,
	SecretKey:    secretKey, 
	PublicKey:    publicKey,
    }
    dokuSnap.GetTokenB2B()
}
```

### 2. Payment Flow
This section guides you through the steps to process payments using the DOKU Golang library. You'll learn how to create a payment request and call the payment function.
#### a. Virtual Account
DOKU offers three ways to use a virtual account: DOKU-Generated Payment Code (DGPC), Merchant-Generated Payment Code (MGPC), and Direct Inquiry Payment Code (DIPC). You can find the full details [here](https://developers.doku.com/accept-payment/direct-api/snap/integration-guide/virtual-account).

> [!Important!]
>Each transaction can use only one feature at a time, but you can use multiple features across different transactions.

##### Create VA DGPC and MGPC
###### CreateVaRequestDTO Model
Create the request object to generate a VA number. Specify the acquirer in the request object. This function is applicable for DGPC and MGPC.

```go
import createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"

createVaRequestDTO = createVaModels.CreateVaRequestDto{
        PartnerServiceId: "    1899",
	CustomerNo:       "20240704001",
	VirtualAccountNo: "    189920240704001",
	VirtualAccountName: "SDK TEST",
	VirtualAccountEmail: "sdk@email.com",
	VirtualAccountPhone: "6281288932399",
	TrxId: "INV_20240711001",
	TotalAmount: createVaModels.TotalAmount{
			Value:    "10000.00",
			Currency: "IDR",
	},
	AdditionalInfo: createVaModels.AdditionalInfo{
		Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
		VirtualAccountConfig: createVaModels.VirtualAccountConfig{
			ReusableStatus: false,
		},
	},
	VirtualAccountTrxType: "C",
	ExpiredDate: "2024-07-29T09:54:04+07:00",
}
```

###### createVa Function
Call the `createVa` function to request the paycode from DOKU. You’ll receive the paycode and payment instructions to display to your customers. This function is applicable for DGPC and MGPC.

```go
createVaResponse := dokuSnap.CreateVa(createVaRequestDTO)
```

##### DIPC
###### inquiryResponse Function
If you use the DIPC feature, you can generate your own paycode and allow your customers to pay without direct communication with DOKU. After customers initiate the payment via the acquirer's channel, DOKU sends an inquiry request to you for validation. This function is applicable for DIPC.

> [!Important!]
>Before sending the inquiry, DOKU sends a token request. Use the `generateToken` function found in the Handling Payment Notification section.

```go
func main() {
     config.InitializeDB()
     defer config.CloseDB()
     config.InitializeSnap()
     http.HandleFunc("/v1.1/transfer-va/inquiry", handlers.DirectInquiryHandler)
     if err := http.ListenAndServe(":8091", nil); err != nil {
	fmt.Println("Server failed:", err)
     }
}

func DirectInquiryHandler(w http.ResponseWriter, r *http.Request) {
	directInquiryResponse, statusCode := services.ProcessDirectInquiry(w, r)
	responseInquiry, err := json.Marshal(directInquiryResponse)
	if err != nil {
		http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseInquiry) 
}

func ProcessDirectInquiry(w http.ResponseWriter, r *http.Request) (inquiry.InquiryResponseBodyDTO, int) {
	authHeader := r.Header.Get("Authorization")
	isTokenValid := config.Snap.ValidateTokenB2B(authHeader)

	var requestBodyInquiry inquiry.InquiryRequestBodyDTO

	err := json.NewDecoder(r.Body).Decode(&requestBodyInquiry)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest) // error message is optional 
	}

	defer r.Body.Close()

	if isTokenValid {
		inquiryData, err := GetDataInquiry(w, requestBodyInquiry.InquiryRequestId)

		if err == sql.ErrNoRows {
			return GenerateResponseDirectInquiryVaNotFound(), http.StatusNotFound
		}

		UpdateDirectInquiryRequest(w, r, requestBodyInquiry.TrxDateInit, requestBodyInquiry.VirtualAccountNo)

		return GenerateResponseDirectInquiry(inquiryData), http.StatusOK

	} else {
		return GenerateResponseDirectInquiryUnauthorized(), http.StatusUnauthorized
	}
}

func UpdateDirectInquiryRequest(w http.ResponseWriter, r *http.Request, trxDate string, vaNumber string) {

	query := `
		UPDATE direct_inquiry_va
		SET status_va = 'inquiry', settlement_time = $1
		WHERE va_number = $2;
	`

	_, err := config.DB.Exec(query, trxDate, vaNumber)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error Insert Header Request: %v", err), http.StatusInternalServerError)
		log.Println("Error inserting data:", err)
		return
	}
}

func GetDataInquiry(w http.ResponseWriter, inquiryRequestID string) (inquiry.InquiryRequestVirtualAccountDataDTO, error) {
	var queryGetInquiryJson = `
			SELECT inquiry_json_object
			FROM direct_inquiry
			WHERE inquiry_request_id = $1;
		`

	var inquiryJsonObject string

	err := config.DB.QueryRow(queryGetInquiryJson, inquiryRequestID).Scan(&inquiryJsonObject)

	// data not found
	if err == sql.ErrNoRows {
		http.Error(w, "Data not found", http.StatusNotFound) // error message is optional
		return inquiry.InquiryRequestVirtualAccountDataDTO{}, err
	}

	// others error
	if err != nil {
		http.Error(w, "Error retrieving inquiry data", http.StatusInternalServerError) -> // error message is optional
		return inquiry.InquiryRequestVirtualAccountDataDTO{}, err
	}

	var inquiryData inquiry.InquiryRequestVirtualAccountDataDTO
	err = json.Unmarshal([]byte(inquiryJsonObject), &inquiryData)
	// failed to parse data
	if err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusInternalServerError) // error message is optional
		return inquiry.InquiryRequestVirtualAccountDataDTO{}, err
	}

	return inquiryData, nil
}

func GenerateResponseDirectInquiry(inquiryData inquiry.InquiryRequestVirtualAccountDataDTO) inquiry.InquiryResponseBodyDTO {
	inquiryResponse := inquiry.InquiryResponseBodyDTO{
		ResponseCode:       "2002400",
		ResponseMessage:    "Successful",
		VirtualAccountData: &inquiryData,
	}
	return inquiryResponse
}

func GenerateResponseDirectInquiryUnauthorized() inquiry.InquiryResponseBodyDTO {
	inquiryResponse := inquiry.InquiryResponseBodyDTO{
		ResponseCode:    "4010000",
		ResponseMessage: "Unauthorized",
	}
	return inquiryResponse
}

func GenerateResponseDirectInquiryVaNotFound() inquiry.InquiryResponseBodyDTO {
	inquiryResponse := inquiry.InquiryResponseBodyDTO{
		ResponseCode:    "4012400",
		ResponseMessage: "Virtual Account Not Found",
	}
	return inquiryResponse
}

```

##### Update VA
###### UpdateVaRequestDto Model
Create the request object to update VA. Specify the acquirer in the request object.

```go
import updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"

updateVaRequestDTO := updateVaModels.UpdateVaDTO{
	PartnerServiceId:    "    1899",
	CustomerNo:          "000000000650",
	VirtualAccountNo:    "    1899000000000650",
	VirtualAccountName:  "SDK TEST",
	VirtualAccountEmail: "sdk@email.com",
	VirtualAccountPhone: "6281288932399",
	TrxId:               "INV_20240710001",
	TotalAmount: createVaModels.TotalAmount{
		Value: "10000.00",
		Currency: "IDR",
	},
	AdditionalInfo: updateVaModels.UpdateVaAdditionalInfoDTO{
		Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
		VirtualAccountConfig: updateVaModels.UpdateVaVirtualAccountConfigDTO{
			Status: "ACTIVE",
	        },
	},
	VirtualAccountTrxType: "C",
	ExpiredDate:           "2024-11-24T10:55:00+07:00",
}

```

###### updateVa Function
Call the `updateVa` function to update VA. It will return the updated VA.

```go
updateVaResponse := dokuSnap.updateVa(updateVaRequestDTO);
```

##### Delete VA
###### DeleteVaRequestDto Model
Create the request object to delete VA. Specify the acquirer in the request object.

```go
import 	deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"

requestDeleteVa := deleteVaModels.DeleteVaRequestDto{
   PartnerServiceId: "    1899",
   CustomerNo:       "000000000971",
   VirtualAccountNo: "    1899000000000971"
   TrxId:            "757",
   AdditionalInfo: deleteVaModels.DeleteVaRequestAdditionalInfo{
     Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
   },
}
```

###### deletePaymentCode Function
Call the `deletePaymentCode` function to delete VA.

```go
snap.DeletePaymentCode(requestDeleteVa)
```

##### Check Status VA
###### CheckStatusVaRequestDto Model
Create the request object to check status of your VA. Specify the acquirer in the request object.

```go
import checkVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"

inquiryRequestId := ""
paymentRequestId := ""
checkStatusVaRequestDto := checkVaModels.CheckStatusVARequestDto{
   PartnerServiceId: updateVaResponseDto.VirtualAccountData.PartnerServiceId,
   CustomerNo:       updateVaResponseDto.VirtualAccountData.CustomerNo,
   VirtualAccountNo: updateVaResponseDto.VirtualAccountData.VirtualAccountNo,
   InquiryRequestId: &inquiryRequestId,
   PaymentRequestId: &paymentRequestId,
}
```

###### checkStatusVa Function
Call the `checkStatusVa` function to check the status of your VA.

```go
snap.CheckStatusVa(checkStatusVaRequestDto)
```


### Handling Payment Notification
After your customers make a payment, you’ll receive a notification from DOKU to update the payment status on your end. DOKU first sends a token request (as with DIPC), then uses that token to send the payment notification.
##### validateAsymmetricSignatureAndGenerateToken function
Generate the response to DOKU, including the required token, by calling this function.

```go
/**
 * request -> *http.Request
 */
response := dokuSnap.ValidateSignatureAndGenerateToken(request)

```

##### validateTokenAndGenerateNotificationReponse function
Deserialize the raw notification data into a structured object using a Data Transfer Object (DTO). This allows you to update the order status, notify customers, or perform other necessary actions based on the notification details.

```go
response := dokuSnap.ValidateTokenAndGenerateNotificationResponse(r.Header.Get("Authorization"), requestBody)
```

##### generateNotificationResponse function
DOKU requires a response to the notification. Use this function to serialize the response data to match DOKU’s format.
You will need to validate the token first and provide the PaymentNotificationRequestBodyDto (you can use the model included in the SDK).

```go
/**
 * isTokenValid -> boolean
 * paymentNotificationRequestBodyDTO -> object
 */
dokuSnap.GenerateNotificationResponse(isTokenValid, paymentNotificationRequestBodyDTO)
```

### 4. Additional Features
Need to use our functions independently? No problem! Here’s how:
#### - v1 to SNAP converter
If you're one of our earliest users, you might still use our v1 APIs. In order to simplify your re-integration process to DOKU's SNAP API specification, DOKU provides you with a helper tools to directly convert v1 APIs to SNAP APIs specification.

##### a. convertRequestV1
Convert DOKU's inquiry and notification from SNAP format (JSON) to v1 format (XML). Feed the inquiry and notification directly to your app without manually mapping parameters or converting file formats.
This function expects an XML string request and return a SNAP format of the request.

```go
/**
 * header -> HttpServletRequest
 * InquiryRequestBodyDto -> object
 */
dokuSnap.directInquiryRequestMapping(header, InquiryRequestBodyDto);
```

##### b. convertResponseV1
Convert your inquiry response to DOKU from v1 format (XML) to SNAP format (Form data). Our library handles response code mapping, allowing you to directly use the converted response and send it to DOKU.
This function will return the response in form data format.

```go
/**
 * xmlString -> String
 */
dokuSnap.directInquiryResponseMapping(xmlString);
```
