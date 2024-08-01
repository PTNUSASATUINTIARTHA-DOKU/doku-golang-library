package converter

import "encoding/xml"

type InquiryResponse struct {
	XMLName          xml.Name `xml:"INQUIRY_RESPONSE"`
	PaymentCode      string   `xml:"PAYMENTCODE"`
	Amount           string   `xml:"AMOUNT"`
	PurchaseAmount   string   `xml:"PURCHASEAMOUNT"`
	MinAmount        string   `xml:"MINAMOUNT"`
	MaxAmount        string   `xml:"MAXAMOUNT"`
	TransIdMerchant  string   `xml:"TRANSIDMERCHANT"`
	Words            string   `xml:"WORDS"`
	RequestDateTime  string   `xml:"REQUESTDATETIME"`
	Currency         string   `xml:"CURRENCY"`
	PurchaseCurrency string   `xml:"PURCHASECURRENCY"`
	SessionId        string   `xml:"SESSIONID"`
	Name             string   `xml:"NAME"`
	Email            string   `xml:"EMAIL"`
	Basket           string   `xml:"BASKET"`
	AdditionalData   string   `xml:"ADDITIONALDATA"`
	ResponseCode     string   `xml:"RESPONSECODE"`
}
