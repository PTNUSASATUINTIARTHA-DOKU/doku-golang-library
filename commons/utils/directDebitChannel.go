package utils

type DirectDebitChannelEnum string

const (
	DIRECT_DEBIT_ALLO_SNAP    DirectDebitChannelEnum = "DIRECT_DEBIT_ALLO_SNAP"
	DIRECT_DEBIT_CIMB_SNAP    DirectDebitChannelEnum = "DIRECT_DEBIT_CIMB_SNAP"
	DIRECT_DEBIT_MANDIRI_SNAP DirectDebitChannelEnum = "DIRECT_DEBIT_MANDIRI_SNAP"
	DIRECT_DEBIT_BRI_SNAP     DirectDebitChannelEnum = "DIRECT_DEBIT_BRI_SNAP"
	EMONEY_OVO_SNAP           DirectDebitChannelEnum = "EMONEY_OVO_SNAP"
	EMONEY_SHOPEE_PAY_SNAP    DirectDebitChannelEnum = "EMONEY_SHOPEE_PAY_SNAP"
	EMONEY_DANA_SNAP          DirectDebitChannelEnum = "EMONEY_DANA_SNAP"
)
