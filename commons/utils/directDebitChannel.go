package utils

import "strings"

type DirectDebitChannel int

const (
	DIRECT_DEBIT_ALLO_SNAP DirectDebitChannel = iota + 1
	DIRECT_DEBIT_CIMB_SNAP
	DIRECT_DEBIT_MANDIRI_SNAP
	DIRECT_DEBIT_BRI_SNAP
	EMONEY_OVO_SNAP
	EMONEY_SHOPEE_PAY_SNAP
	EMONEY_DANA_SNAP
)

var DirectDebitChannelNames = map[DirectDebitChannel]string{
	DIRECT_DEBIT_ALLO_SNAP:    "DIRECT_DEBIT_ALLO_SNAP",
	DIRECT_DEBIT_CIMB_SNAP:    "DIRECT_DEBIT_CIMB_SNAP",
	DIRECT_DEBIT_MANDIRI_SNAP: "DIRECT_DEBIT_MANDIRI_SNAP",
	DIRECT_DEBIT_BRI_SNAP:     "DIRECT_DEBIT_BRI_SNAP",
	EMONEY_OVO_SNAP:           "EMONEY_OVO_SNAP",
	EMONEY_SHOPEE_PAY_SNAP:    "EMONEY_SHOPEE_PAY_SNAP",
	EMONEY_DANA_SNAP:          "EMONEY_DANA_SNAP",
}

func ValidateDirectDebitChannel(channel string) bool {
	for _, validChannel := range DirectDebitChannelNames {
		if channel == validChannel {
			return true
		}
	}
	return false
}

func IsValidFeeType(feeType string) bool {
	switch strings.ToUpper(feeType) {
	case "OUR", "BEN", "SHA":
		return true
	default:
		return false
	}
}

func IsValidPaymentType(paymentType string) bool {
	switch strings.ToUpper(paymentType) {
	case "SALE", "RECURRING":
		return true
	default:
		return false
	}
}
