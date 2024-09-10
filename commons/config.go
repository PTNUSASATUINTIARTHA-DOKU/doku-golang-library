package commons

const (
	SANDBOX_BASE_URL                 = "https://api-uat.doku.com"
	PRODUCTION_BASE_URL              = "https://dashboard.doku.com"
	ACCESS_TOKEN                     = "/authorization/v1/access-token/b2b"
	CREATE_VA                        = "/virtual-accounts/bi-snap-va/v1.1/transfer-va/create-va"
	UPDATE_VA                        = "/virtual-accounts/bi-snap-va/v1.1/transfer-va/update-va"
	CHECK_VA                         = "/orders/v1.0/transfer-va/status"
	DELETE_VA                        = "/virtual-accounts/bi-snap-va/v1.1/transfer-va/delete-va"
	DIRECT_DEBIT_ACCOUNT_BINDING     = "/direct-debit/core/v1/registration-account-binding"
	ACCESS_TOKEN_B2B2C               = "/authorization/v1/access-token/b2b2c"
	DIRECT_DEBIT_BALANCE_INQUIRY_URL = "/direct-debit/core/v1/balance-inquiry"
	DIRECT_DEBIT_PAYMENT             = "/direct-debit/core/v1/debit/payment-host-to-host"
	DIRECT_DEBIT_ACCOUNT_UNBINDING   = "/direct-debit/core/v1/registration-account-unbinding"
)

type Config struct{}

func (c Config) GetBaseUrl(isProduction bool) string {
	if isProduction {
		return PRODUCTION_BASE_URL
	}
	return SANDBOX_BASE_URL
}
