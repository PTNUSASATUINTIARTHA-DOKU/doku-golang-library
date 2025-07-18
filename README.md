# DOKU GOLANG SDK Documentation

## Introduction
Welcome to the DOKU Golang SDK! This SDK simplifies access to the DOKU API for your GO applications, enabling seamless integration with payment and virtual account services.

If your looking for another language  [Node.js](https://github.com/PTNUSASATUINTIARTHA-DOKU/doku-nodejs-library), [PHP](https://github.com/PTNUSASATUINTIARTHA-DOKU/doku-php-library), [Python](https://github.com/PTNUSASATUINTIARTHA-DOKU/doku-python-library), [Java](https://github.com/PTNUSASATUINTIARTHA-DOKU/doku-java-library)

## Table of Contents
- [DOKU Golang SDK Documentation](#doku-php-sdk-documentation)
  - [1. Getting Started](#1-getting-started)
  - [2. Usage](#2-usage)
    - [Virtual Account](#virtual-account)
      - [I. Virtual Account (DGPC \& MGPC)](#i-virtual-account-dgpc--mgpc)
      - [II. Virtual Account (DIPC)](#ii-virtual-account-dipc)
      - [III. Check Virtual Account Status](#iii-check-virtual-account-status)
    - [B. Binding / Registration Operations](#b-binding--registration-operations)
      - [I. Account Binding](#i-account-binding)
      - [II. Card Registration](#ii-card-registration)
    - [C. Direct Debit and E-Wallet](#c-direct-debit-and-e-wallet)
      - [I. Request Payment](#i-request-payment)
      - [II. Request Payment Jump App](#ii-request-payment-jump-app)
  - [3. Other Operation](#3-other-operation)
    - [Check Transaction Status](#a-check-transaction-status)
    - [Refund](#b-refund)
    - [Balance Inquiry](#c-balance-inquiry)
  - [4. Error Handling and Troubleshooting](#4-error-handling-and-troubleshooting)




## 1. Getting Started

### Requirements
- Go 1.22.2 or higher

### Installation
To install the DOKU Snap SDK, use GO:
```bash
go get github.com/lib/pq
go get -u github.com/golang-jwt/jwt/v4
go get github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library@latest
```

### Configuration
Before using the Doku Snap SDK, you need to initialize it with your credentials:
1. **Client ID**, **Secret Key** and **DOKU Public Key**: Retrieve these from the Integration menu in your Doku Dashboard
2. **Private Key** and **Public Key** : Generate your Private Key and Public Key
   
How to generate Merchant privateKey and publicKey:
1. generate private key RSA : openssl genrsa -out private.key 2048
2. set passphrase your private key RSA : openssl pkcs8 -topk8 -inform PEM -outform PEM -in private.key -out pkcs8.key -v1 PBE-SHA1-3DES
3. generate public key RSA : openssl rsa -in private.key -outform PEM -pubout -out public.pem

The encryption model applied to messages involves both asymmetric and symmetric encryption, utilizing a combination of Private Key and Public Key, adhering to the following standards:

  1. Standard Asymmetric Encryption Signature: SHA256withRSA dengan Private Key ( Kpriv ) dan Public Key ( Kpub ) (256 bits)
  2. Standard Symmetric Encryption Signature HMAC_SHA512 (512 bits)
  3. Standard Symmetric Encryption AES-256 dengan client secret sebagai encryption key.

| **Parameter**       | **Description**                                    | **Required** |
|-----------------|----------------------------------------------------|--------------|
| `privateKey`    | The private key for the partner service.           | ✅          |
| `publicKey`     | The public key for the partner service.            | ✅           |
| `dokuPublicKey` | Key that merchants use to verify DOKU request      | ✅           |
| `clientId`      | The client ID associated with the service.         | ✅           |
| `secretKey`     | The secret key for the partner service.            | ✅           |
| `isProduction`  | Set to true for production environment             | ✅           |
| `issuer`        | Optional issuer for advanced configurations.       | ❌           |
| `authCode`      | Optional authorization code for advanced use.      | ❌           |


```go
import "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/doku"

privateKey := "YOUR_PRIVATE_KEY"
publicKey := "YOUR_PUBLIC_KEY"
clientId := "YOUR_CLIENT_ID"
secretKey := "YOUR_SECRET_KEY"
isProduction := false
issuer := "YOUR_ISSUER"

doku.TokenController = controllers.TokenController{}
doku.VaController = controllers.VaController{}
doku.NotificationController = controllers.NotificationController{}
doku.DirectDebitController = &controllers.DirectDebitController{}

var snap doku.Snap
snap = doku.Snap{
    PrivateKey:   privateKey,
    ClientId:     clientId,
    IsProduction: isProduction,
    SecretKey:    secretKey,
    Issuer: 	  issuer,
    PublicKey:    publicKey,
}
```

## 2. Usage

**Initialization**
Always start by initializing the Snap object.

```go
snap = doku.Snap{
    PrivateKey:   privateKey,
    ClientId:     clientId,
    IsProduction: isProduction,
    SecretKey:    secretKey,
    Issuer: 	  issuer,
    PublicKey:    publicKey,
}
```
### Virtual Account
#### I. Virtual Account (DGPC & MGPC)
##### DGPC
- **Description:** A pre-generated virtual account provided by DOKU.
- **Use Case:** Recommended for one-time transactions.
##### MGPC
- **Description:** Merchant generated virtual account.
- **Use Case:** Recommended for top up business model.

Parameters for **createVA** and **updateVA**
<table>
  <thead>
    <tr>
      <th><strong>Parameter</strong></th>
      <th colspan="2"><strong>Description</strong></th>
      <th><strong>Data Type</strong></th>
      <th><strong>Required</strong></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>partnerServiceId</code></td>
      <td colspan="2">The unique identifier for the partner service.</td>
      <td>String(20)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>customerNo</code></td>
      <td colspan="2">The customer's identification number.</td>
      <td>String(20)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>virtualAccountNo</code></td>
      <td colspan="2">The virtual account number associated with the customer.</td>
      <td>String(20)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>virtualAccountName</code></td>
      <td colspan="2">The name of the virtual account associated with the customer.</td>
      <td>String(255)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>virtualAccountEmail</code></td>
      <td colspan="2">The email address associated with the virtual account.</td>
      <td>String(255)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td><code>virtualAccountPhone</code></td>
      <td colspan="2">The phone number associated with the virtual account.</td>
      <td>String(9-30)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td><code>trxId</code></td>
      <td colspan="2">Invoice number in Merchants system.</td>
      <td>String(64)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="2"><code>totalAmount</code></td>
      <td colspan="2"><code>value</code>: Transaction Amount (ISO 4217) <br> <small>Example: "11500.00"</small></td>
      <td>String(16.2)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>Currency</code>: Currency <br> <small>Example: "IDR"</small></td>
      <td>String(3)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="4"><code>additionalInfo</code></td>
      <td colspan="2"><code>channel</code>: Channel that will be applied for this VA <br> <small>Example: VIRTUAL_ACCOUNT_BANK_CIMB</small></td>
      <td>String(20)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="3"><code>virtualAccountConfig</code></td>
      <td><code>reusableStatus</code>: Reusable Status For Virtual Account Transaction <br><small>value TRUE or FALSE</small></td>
      <td>Boolean</td>
      <td>❌</td>
    </tr>
    <tr>
      <td><code>minAmount</code>: Minimum Amount can be used only if <code>virtualAccountTrxType</code> is Open Amount (O). <br><small>Example: "10000.00"</small></td>
      <td>String(16.2)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td><code>maxAmount</code>: Maximum Amount can be used only if <code>virtualAccountTrxType</code> is Open Amount (O). <br><small>Example: "5000000.00"</small></td>
      <td>String(16.2)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td><code>virtualAccountTrxType</code></td>
      <td colspan="2">Transaction type for this transaction. C (Closed Amount), O (Open Amount)</td>
      <td>String(1)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>expiredDate</code></td>
      <td colspan="2">Expiration date for Virtual Account. ISO-8601 <br><small>Example: "2023-01-01T10:55:00+07:00"</small></td>
      <td>String</td>
      <td>❌</td>
    </tr>
    <tr>
      <td rowspan="2"><code>freeTexts</code></td>
      <td colspan="2"><code>English</code>: Free text for additional description. <br> <small>Example: "Free texts"</small></td>
      <td>String(64)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td colspan="2"><code>Indonesia</code>: Free text for additional description. <br> <small>Example: "Tulisan Bebas"</small></td>
      <td>String(64)</td>
      <td>❌</td>
    </tr>
  </tbody>
</table>


1. **Create Virtual Account**
    - **Function:** `createVa`
    ```go
      import createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"

      createVaRequestDTO := createVaModels.CreateVaRequestDto{
          PartnerServiceId: "8129014",
          CustomerNo:       "17223992157",
          VirtualAccountNo: "812901417223992157",
          VirtualAccountName: "Test Example",
          VirtualAccountEmail: "test.example@gmail.com",
          VirtualAccountPhone: "621722399214895",
          TrxId: "INV_CIMB_TEST_1",
          TotalAmount: createVaModels.TotalAmount{
              Value:    "12500.00",
              Currency: "IDR",
          },
          AdditionalInfo: createVaModels.AdditionalInfo{
              Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
              VirtualAccountConfig: createVaModels.VirtualAccountConfig{
                  ReusableStatus: true,
              },
          },
          VirtualAccountTrxType: "C",
          ExpiredDate: "2025-08-31T09:54:04+07:00",
          FreeTexts: []createVaModels.FreeTexts{
            {
              English: "Free Texts",
              Indonesia:   "Tulisan Bebas",
            },
          },
      }

      createVaResponse := snap.CreateVa(createVaRequestDTO)
    ```

2. **Update Virtual Account**
    - **Function:** `updateVa`

    ```go
      import updateVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/updateVa"

      updateVaRequestDTO := updateVaModels.UpdateVaDTO{
          PartnerServiceId:    "8129014",
          CustomerNo:          "17223992157",
          VirtualAccountNo:    "812901417223992157",
          VirtualAccountName:  "Test Example",
          VirtualAccountEmail: "test.example@gmail.com",
          VirtualAccountPhone: "621722399214895",
          TrxId:               "INV_CIMB_TEST_1",
          TotalAmount: createVaModels.TotalAmount{
              Value:    "12500.00",
              Currency: "IDR",
          },
          AdditionalInfo: updateVaModels.UpdateVaAdditionalInfoDTO{
              Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
              VirtualAccountConfig: updateVaModels.UpdateVaVirtualAccountConfigDTO{
                  Status: "ACTIVE",
              },
          },
          VirtualAccountTrxType: "C",
          ExpiredDate:           "2025-08-31T09:54:04+07:00",
      }

      updateVaResponse, err := snap.updateVa(updateVaRequestDTO);
    ```

3. **Delete Virtual Account**

    | **Parameter**        | **Description**                                                             | **Data Type**       | **Required** |
    |-----------------------|----------------------------------------------------------------------------|---------------------|--------------|
    | `partnerServiceId`    | The unique identifier for the partner service.                             | String(8)        | ✅           |
    | `customerNo`          | The customer's identification number.                                      | String(20)       | ✅           |
    | `virtualAccountNo`    | The virtual account number associated with the customer.                   | String(20)       | ✅           |
    | `trxId`               | Invoice number in Merchant's system.                                       | String(64)       | ✅           |
    | `additionalInfo`      | `channel`: Channel applied for this VA.<br><small>Example: VIRTUAL_ACCOUNT_BANK_CIMB</small> | String(30)       | ✅    |

    
  - **Function:** `deletePaymentCode`

    ```go
    import deleteVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/deleteVa"

    deleteVaRequest := deleteVaModels.DeleteVaRequestDto{
        PartnerServiceId: "8129014",
        CustomerNo:       "17223992157",
        VirtualAccountNo: "812901417223992157",
        TrxId:            "INV_CIMB_TEST_1",
        AdditionalInfo: deleteVaModels.DeleteVaRequestAdditionalInfo{
            Channel: "VIRTUAL_ACCOUNT_BANK_CIMB",
        },
    }

    deleteVaResponse, err := snap.DeletePaymentCode(deleteVaRequest)
    ```


#### II. Virtual Account (DIPC)
- **Description:** The VA number is registered on merchant side and DOKU will forward Acquirer inquiry request to merchant side when the customer make payment at the acquirer channel

- **Function:** `directInquiryVa`

    ```go
    func directInquiryVa(w http.ResponseWriter, r *http.Request) {
      authHeader := r.Header.Get("Authorization")
      isTokenValid := config.Snap.ValidateTokenB2B(authHeader)

      var requestBodyInquiry inquiry.InquiryRequestBodyDTO

      err := json.NewDecoder(r.Body).Decode(&requestBodyInquiry)
      if err != nil {
          http.Error(w, "Invalid Request Body", http.StatusBadRequest) message is optional
          return 
      }
      defer r.Body.Close()

      if !isTokenValid {
          inquiryResponse := inquiry.InquiryResponseBodyDTO{
              ResponseCode:    "4010000",
              ResponseMessage: "Unauthorized",
          }
          w.WriteHeader(http.StatusUnauthorized)
          json.NewEncoder(w).Encode(inquiryResponse)
          return
      }

      inquiryData, err := GetDataInquiry(w, requestBodyInquiry.InquiryRequestId)
      if err == sql.ErrNoRows {
          inquiryResponse := inquiry.InquiryResponseBodyDTO{
              ResponseCode:    "4012400",
              ResponseMessage: "Virtual Account Not Found",
          }
          w.WriteHeader(http.StatusNotFound)
          json.NewEncoder(w).Encode(inquiryResponse)
          return
      }

      UpdateDirectInquiryRequest(w, r, requestBodyInquiry.TrxDateInit, requestBodyInquiry.VirtualAccountNo)

      inquiryResponse := inquiry.InquiryResponseBodyDTO{
          ResponseCode:       "2002400",
          ResponseMessage:    "Successful",
          VirtualAccountData: &inquiryData,
      }
      w.WriteHeader(http.StatusOK)
      json.NewEncoder(w).Encode(inquiryResponse)
    }
    ```

#### III. Check Virtual Account Status
 | **Parameter**        | **Description**                                                             | **Data Type**       | **Required** |
|-----------------------|----------------------------------------------------------------------------|---------------------|--------------|
| `partnerServiceId`    | The unique identifier for the partner service.                             | String(8)        | ✅           |
| `customerNo`          | The customer's identification number.                                      | String(20)       | ✅           |
| `virtualAccountNo`    | The virtual account number associated with the customer.                   | String(20)       | ✅           |
| `inquiryRequestId`    | The customer's identification number.                                      | String(128)       | ❌           |
| `paymentRequestId`    | The virtual account number associated with the customer.                   | String(128)       | ❌           |
| `additionalInfo`      | The virtual account number associated with the customer.                   | String      | ❌           |

  - **Function:** `checkStatusVa`
    ```go
    import checkStatusVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/checkVa"

    inquiryRequestId := ""
    paymentRequestId := ""
    checkStatusVaRequestDto := checkStatusVaModels.CheckStatusVARequestDto{
        PartnerServiceId: "8129014",
        CustomerNo:       "17223992157",
        VirtualAccountNo: "812901417223992157",
        InquiryRequestId: &inquiryRequestId,
        PaymentRequestId: &paymentRequestId,
    }

    snap.CheckStatusVa(checkStatusVaRequestDto)
    ```

### B. Binding / Registration Operations
The card registration/account binding process must be completed before payment can be processed. The merchant will send the card registration request from the customer to DOKU.

Each card/account can only registered/bind to one customer on one merchant. Customer needs to verify OTP and input PIN.

| **Services**     | **Binding Type**      | **Details**                        |
|-------------------|-----------------------|-----------------------------------|
| Direct Debit      | Account Binding       | Supports **Allo Bank** and **CIMB** |
| Direct Debit      | Card Registration     | Supports **BRI**                    |
| E-Wallet          | Account Binding       | Supports **OVO**                    |

#### I. Account Binding 
1. **Binding**

<table>
  <thead>
    <tr>
      <th><strong>Parameter</strong></th>
      <th colspan="2"><strong>Description</strong></th>
      <th><strong>Data Type</strong></th>
      <th><strong>Required</strong></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>phoneNo</code></td>
      <td colspan="2">Phone Number Customer. <br> <small>Format: 628238748728423</small> </td>
      <td>String(9-16)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="13"><code>additionalInfo</code></td>
      <td colspan="2"><code>channel</code>: Payment Channel<br></td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>custIdMerchant</code>: Customer id from merchant</td>
      <td>String(64)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>customerName</code>: Customer name from merchant</td>
      <td>String(70)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td colspan="2"><code>email</code>: Customer email from merchant </td>
      <td>String(64)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td colspan="2"><code>idCard</code>: Customer id card from merchant</td>
      <td>String(20)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td colspan="2"><code>country</code>: Customer country </td>
      <td>String</td>
      <td>❌</td>
    </tr>
    <tr>
      <td colspan="2"><code>address</code>: Customer Address</td>
      <td>String(255)</td>
      <td>❌</td>
    </tr>
        <tr>
      <td colspan="2"><code>dateOfBirth</code> </td>
      <td>String(YYYYMMDD)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td colspan="2"><code>successRegistrationUrl</code>: Redirect URL when binding is success </td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>failedRegistrationUrl</code>: Redirect URL when binding is success fail</td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>deviceModel</code>: Device Model customer </td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>osType</code>: Format: ios/android </td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>channelId</code>: Format: app/web </td>
      <td>String</td>
      <td>✅</td>
    </tr>
    </tbody>
  </table> 

  - **Function:** `doAccountBinding`

    ```go
    import accountBindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountbinding"

    deviceId := "YOUR_DEVICE_ID"
    ipAddress := "YOUR_IP_ADDRESS"

    accountBindingRequest := accountBindingModels.AccountBindingRequestDTO{
        PhoneNo: "6288912121237",
        AdditionalInfo: accountBindingModels.AccountBindingAdditionalInfoRequestDto{
            Channel:                "DIRECT_DEBIT_CIMB_SNAP",
            CustIdMerchant:         "CUST123",
            CustomerName:           "John Doe",
            Email:                  "john.doe@example.com",
            IdCard:                 "99999",
            Country:                "Indonesia",
            Address:                "Jakarta",
            DateOfBirth:            "19990101",
            SuccessRegistrationUrl: "https://success.example.com",
            FailedRegistrationUrl:  "https://fail.example.com",
            DeviceModel:            "iPhone 12",
            OsType:                 "ios",
            ChannelId:              "CH001",
        },
    }

    accountBindingResponse, err := snap.DoAccountBinding(accountBindingRequest, deviceId, ipAddress)
    ```

1. **Unbinding**

    - **Function:** `getTokenB2B2C`
    ```go
    authCode := "YOUR_AUTH_CODE_FROM_ACCOUNT_BINDING"
    responseGetTokenB2B2C, err := snap.GetTokenB2B2C(authCode)
   ```
    - **Function:** `doAccountUnbinding`
    ```go
    import accountUnbindingModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/accountunbinding"

    accountUnbindingRequest := accountUnbindingModels.AccountUnbindingRequestDTO{
      TokenId: responseGetTokenB2B2C.AccessToken,
      AdditionalInfo: accountUnbindingModels.AccountUnbindingAdditionalInfoRequestDTO{
        Channel: "DIRECT_DEBIT_CIMB_SNAP",
      },
    }

    ipAddress := "YOUR_IP_ADDRESS"
    accountUnbindingResponse, err := snap.DoAccountUnbinding(accountUnbindingRequest, ipAddress)
    ```

#### II. Card Registration
1. **Registration**
    - **Function:** `doCardRegistration`

    ```go
    import cardRegistrationModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/cardregistration"

    cardRegistRequest := cardRegistrationModels.CardRegistrationRequestDTO{
      CardData: cardRegistrationModels.BankCardDataDTO{
        BankCardNo:   "7801",
        BankCardType: "D",
        Email:        "email@email.com",
        ExpiryDate:   "0525",
      },
      CustIdMerchant: "Jhon Doe",
      PhoneNo:        "6282124918109",
      AdditionalInfo: cardRegistrationModels.CardRegistrationAdditionalInfoRequestDTO{
        Channel:                "DIRECT_DEBIT_BRI_SNAP",
        DateOfBirth:            "19990101",
        SuccessRegistrationUrl: "https://success.example.com",
        FailedRegistrationUrl:  "https://fail.example.com",
      },
    }

    cardRegistrationResponse, err := snap.DoCardRegistration(cardRegistRequest, "DH")
    ```

2. **UnRegistration**
    - **Function:** `getTokenB2B2C`
    ```go
      authCode := "YOUR_AUTH_CODE_FROM_ACCOUNT_BINDING"
      responseGetTokenB2B2C, err := snap.GetTokenB2B2C(authCode)
    ```
    - **Function:** `doCardUnbinding`

    ```go
    cardUnbindingRequest := cardUnbindingModels.CardRegistrationUnbindingRequestDTO{
      TokenId: responseGetTokenB2B2C.AccessToken,
      AdditionalInfo: cardUnbindingModels.CardRegistrationUnbindingAdditionalInfoRequestDTO{
          Channel: "DIRECT_DEBIT_BRI_SNAP",
      },
    }

    ipAddress := "YOUR_IP_ADDRESS"
    cardUnbindingResponse, err := snap.DoCardRegistrationUnbinding(cardUnbindingRequest, ipAddress)
    ```

### C. Direct Debit and E-Wallet 

#### I. Request Payment
  Once a customer’s account or card is successfully register/bind, the merchant can send a payment request. This section describes how to send a unified request that works for both Direct Debit and E-Wallet channels.

| **Acquirer**       | **Channel Name**         | 
|-------------------|--------------------------|
| Allo Bank         | DIRECT_DEBIT_ALLO_SNAP   | 
| BRI               | DIRECT_DEBIT_BRI_SNAP    | 
| CIMB              | DIRECT_DEBIT_CIMB_SNAP   |
| OVO               | EMONEY_OVO_SNAP   | 

##### Common parameter
<table>
  <thead>
    <tr>
      <th><strong>Parameter</strong></th>
      <th colspan="2"><strong>Description</strong></th>
      <th><strong>Data Type</strong></th>
      <th><strong>Required</strong></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>partnerReferenceNo</code></td>
      <td colspan="2"> Reference No From Partner <br> <small>Format: 628238748728423</small> </td>
      <td>String(9-16)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="2"><code>amount</code></td>
      <td colspan="2"><code>value</code>: Transaction Amount (ISO 4217) <br> <small>Example: "11500.00"</small></td>
      <td>String(16.2)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>Currency</code>: Currency <br> <small>Example: "IDR"</small></td>
      <td>String(3)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="4"><code>additionalInfo</code> </td>
      <td colspan = "2" ><code>channel</code>: payment channel</td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>remarks</code>:Remarks from Partner</td>
      <td>String(40)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>successPaymentUrl</code>: Redirect Url if payment success</td>
      <td>String</td>
      <td>✅</td>
    </tr>
        <tr>
      <td colspan="2"><code>failedPaymentUrl</code>: Redirect Url if payment fail
      </td>
      <td>String</td>
      <td>✅</td>
    </tr>
    </tbody>
  </table> 

 ##### Allo Bank Specific Parameters

| **Parameter**                        | **Description**                                               | **Required** |
|--------------------------------------|---------------------------------------------------------------|--------------|
| `additionalInfo.remarks`             | Remarks from the partner                                      | ✅           |
| `additionalInfo.lineItems.name`      | Item name (String)                                            | ✅           |
| `additionalInfo.lineItems.price`     | Item price (ISO 4217)                                         | ✅           |
| `additionalInfo.lineItems.quantity`  | Item quantity (Integer)                                      | ✅           |
| `payOptionDetails.payMethod`         | Balance type (options: BALANCE/POINT/PAYLATER)                | ✅           |
| `payOptionDetails.transAmount.value` | Transaction amount                                            | ✅           |
| `payOptionDetails.transAmount.currency` | Currency (ISO 4217, e.g., "IDR")                             | ✅           |


#####  CIMB Specific Parameters

| **Parameter**                        | **Description**                                               | **Required** |
|--------------------------------------|---------------------------------------------------------------|--------------|
| `additionalInfo.remarks`             | Remarks from the partner                                      | ✅           |


#####  DANA Specific Parameters

| **Parameter**                           | **Description**                                                | **Required** |
| ------------------------------------|---------------------------------------------------------------|--------------|
| `additionalInfo.orderTitle`              | Order title from merchant (optional)                          | ❌           |
| `additionalInfo.supportDeepLinkCheckoutUrl` | Value ('true' for Jumpapp behavior, 'false' for webview, default: 'false') | ❌           |


#####  OVO Specific Parameters

| **Parameter**                           | **Description**                                                | **Required** |
|------------------------------------------|---------------------------------------------------------------|--------------|
| `feeType`                                | Fee type from partner (values: OUR, BEN, SHA)                  | ❌           |
| `payOptionDetails.payMethod`             | Payment method format: CASH, POINTS                            | ✅           |
| `payOptionDetails.transAmount.value`    | Transaction amount (ISO 4217)                                  | ✅           |
| `payOptionDetails.transAmount.currency` | Currency (ISO 4217, e.g., "IDR")                               | ✅           |
| `payOptionDetails.feeAmount.value`      | Fee amount (if applicable)                                     | ✅           |
| `payOptionDetails.feeAmount.currency`   | Currency for the fee                                          | ✅           |
| `additionalInfo.paymentType`            | Transaction type (values: SALE, RECURRING)                     | ✅           |

  
Here’s how you can use the `doPayment` function for both payment types:
  - **Function:** `doPayment`
    
    ```go
    paymentModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/payment"
    createVaModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/va/createVa"

    paymentRequest := paymentModels.PaymentRequestDTO{
      PartnerReferenceNo: "INV-101",
      FeeType:            "OUR",  // Only for OVO (Optional) - OUR/BEN/SHA
      Amount: createVaModels.TotalAmount{
          Value:    "10000.00", 
          Currency: "IDR",     
      },
      PayOptionDetails: []paymentModels.PayOptionDetailsDTO{ // Only for OVO or ALLO_BANK (Required)
          {
              PayMethod: "CASH", // CASH / POINTS (OVO) or BALANCE/POINT/PAYLATER (ALLO_BANK)
              TransAmount: createVaModels.TotalAmount{
                  Value:    "10000.00", 
                  Currency: "IDR",         
              },
              FeeAmount: createVaModels.TotalAmount{
                  Value:    "1100.00", 
                  Currency: "IDR",
              },
          },
      },
      AdditionalInfo: paymentModels.PaymentAdditionalInfoDTO{
          Channel:           "EMONEY_OVO_SNAP", //  "DIRECT_DEBIT_CIMB / DIRECT_DEBIT_BRI_SNAP / DIRECT_DEBIT_ALLO_SNAP / EMONEY_OVO_SNAP"
          Remarks:           "Payment Order",  
          SuccessPaymentUrl: "https://success.example.com",
          FailedPaymentUrl:  "https://fail.example.com",
          LineItems: []paymentModels.LineItemsDTO{ // Only for ALLO_BANK, 
              {
                  Name:     "Bag",     
                  Price:    "10000.00",     
                  Quantity: "1",  
              },
          },
          PaymentType: "SALE", // Only For OVO and BRI: SALE, RECURRING (Optional)
      },
      ChargeToken: "",
    }
    authCode := "YOUR_AUTH_CODE_FROM_BINDING"
    ipAddress := "YOUR_IP_ADDRESS"
    paymentResponse, err := snap.DoPayment(dataRequest, ipAddress, authCode)
    ```
#### II. Request Payment Jump App
| **Acquirer**      | **Channel Name**        | 
|-------------------|-------------------------|
| DANA              | EMONEY_DANA_SNAP        | 
| ShopeePay         | EMONEY_SHOPEE_PAY_SNAP  |

The following fields are common across **DANA and ShopeePay** requests:
<table>
  <thead>
    <tr>
      <th><strong>Parameter</strong></th>
      <th colspan="2"><strong>Description</strong></th>
      <th><strong>Data Type</strong></th>
      <th><strong>Required</strong></th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>partnerReferenceNo</code></td>
      <td colspan="2"> Reference No From Partner <br> <small>Examplae : INV-0001</small> </td>
      <td>String(9-16)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>validUpto</code></td>
      <td colspan = "2" >Expired time payment url </td>
      <td>String</td>
      <td>❌</td>
    </tr>
    <tr>
      <td><code>pointOfInitiation</code></td>
      <td colspan = "2" >Point of initiation from partner,<br> value: app/pc/mweb </td>
      <td>String</td>
      <td>❌</td>
    </tr>
    <tr>
      <td rowspan = "3" > <code>urlParam</code></td>
      <td colspan = "2"><code>url</code>: URL after payment sucess </td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>type</code>: Pay Return<br> <small>always PAY_RETURN </small></td>
      <td>String</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>isDeepLink</code>: Is Merchant use deep link or not<br> <small>Example: "Y/N"</small></td>
      <td>String(1)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td rowspan="2"><code>amount</code></td>
      <td colspan="2"><code>value</code>: Transaction Amount (ISO 4217) <br> <small>Example: "11500.00"</small></td>
      <td>String(16.2)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td colspan="2"><code>Currency</code>: Currency <br> <small>Example: "IDR"</small></td>
      <td>String(3)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td><code>additionalInfo</code> </td>
      <td colspan = "2" ><code>channel</code>: payment channel</td>
      <td>String</td>
      <td>✅</td>
    </tr>
    </tbody>
  </table> 

##### DANA

DANA spesific parameters
<table>
    <thead>
    <tr>
      <th><strong>Parameter</strong></th>
      <th colspan="2"><strong>Description</strong></th>
      <th><strong>Data Type</strong></th>
      <th><strong>Required</strong></th>
    </tr>
    </thead>
    <tbody>
    <tr>
      <td rowspan = "2" ><code>additionalInfo</code></td>
      <td colspan = "2" ><code>orderTitle</code>: Order title from merchant</td>
      <td>String</td>
      <td>❌</td>
    </tr>
    <tr>
      <td colspan = "2" ><code>supportDeepLinkCheckoutUrl</code> : Value 'true' for Jumpapp behaviour, 'false' for webview, false by default</td>
      <td>String</td>
      <td>❌</td>
    </tr>
    </tbody>
  </table> 
For Shopeepay and Dana you can use the `doPaymentJumpApp` function for for Jumpapp behaviour

- **Function:** `doPaymentJumpApp`

```go
jumpAppModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/jumpapp"

paymentJumpAppRequest := jumpAppModels.PaymentJumpAppRequestDTO{
  PartnerReferenceNo: "INV-101",
  ValidUpTo:          "2025-12-31T23:59:59Z",
  PointOfInitiation:  "app", // app/pc/mweb
  UrlParam: []jumpAppModels.UrlParamDTO{{
    Url: "https://your.url/endpoint",
    Type: "PAY_RETURN",
    IsDeepLink: "Y",
  },},
  Amount: createVaModels.TotalAmount{
    Value: "10000.00",
    Currency: "IDR",
  },
  AdditionalInfo: jumpAppModels.PaymentJumpAppAdditionalInfoRequestDTO{
    Channel: "EMONEY_DANA_SNAP" // EMONEY_DANA_SNAP or EMONEY_SHOPEE_PAY_SNAP
    OrderTitle: "Payment Order",
    Metadata: "Your Metadata", // Only for ShopeePay (Optional)
    SupportDeepLinkCheckoutUrl: true, // Only for DANA (Optional)
  },
}

deviceId := "YOUR_DEVICE_ID"
ipAddress := "YOUR_IP_ADDRESS"

snap.DoPaymentJumpApp(jumpappRequest, deviceId, ipAddress)
```

## 3. Other Operation

### A. Check Transaction Status

  ```go	
  checkStatusModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/checkstatus"

  func DoCheckStatus(checkStatusRequestDTO checkStatusModels.CheckStatusRequestDTO) (checkStatusModels.CheckStatusResponseDTO, error) {
	return snap.DoCheckStatus(checkStatusRequestDTO)
}
  ```

### B. Refund

  ```go
  refundModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/refund"

  func DoRefund(refundRequestDTO refundModels.RefundRequestDTO, ipAddress string, authCode string, deviceId string) (refundModels.RefundResponseDTO, error){
	return snap.DoRefund(refundRequestDTO, ipAddress, authCode, deviceId)
  }
  ```

### C. Balance Inquiry

  ```go
  balanceInquiryModels "github.com/PTNUSASATUINTIARTHA-DOKU/doku-golang-library/models/directdebit/balanceinquiry"

  func DoBalanceInquiry(balanceInquiryRequestDto balanceInquiryModels.BalanceInquiryRequestDto, deviceId string, ipAddress string, authCode string) (balanceInquiryModels.BalanceInquiryResponseDto, error) {
	return snap.DoBalanceInquiry(balanceInquiryRequestDto, deviceId, ipAddress, authCode)
  }
  ```

## 4. Error Handling and Troubleshooting

The SDK returns errors for various conditions. Always check for errors when making API calls:
 ```go
  import ( "log" )

  createVaResponse, err := snap.CreateVa(createVaRequestDto)
  if err != nil {
    // Handle the error appropriately
    log.Printf("Error: %s", err.Error())
    // You can also return or take other actions based on the error
    return
  }

  // Process successful result
  log.Printf("VA created successfully: %+v", createVaResponse)
 ```

This section provides common errors and solutions:

| Error Code | Description                           | Solution                                     |
|------------|---------------------------------------|----------------------------------------------|
| `4010000`  | Unauthorized                          | Check if Client ID and Secret Key are valid. |
| `4012400`  | Virtual Account Not Found             | Verify the virtual account number provided.  |
| `2002400`  | Successful                            | Transaction completed successfully.          |



