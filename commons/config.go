package commons

const (
	SANDBOX_BASE_URL             = "https://api-uat.doku.com"
	PRODUCTION_BASE_URL          = "https://dashboard.doku.com"
	ACCESS_TOKEN                 = "/authorization/v1/access-token/b2b"
	CREATE_VA                    = "/virtual-accounts/bi-snap-va/v1.1/transfer-va/create-va"
	UPDATE_VA                    = "/virtual-accounts/bi-snap-va/v1.1/transfer-va/update-va"
	CHECK_VA                     = "/orders/v1.0/transfer-va/status"
	DELETE_VA                    = "/virtual-accounts/bi-snap-va/v1.1/transfer-va/delete-va"
	DIRECT_DEBIT_ACCOUNT_BINDING = "/direct-debit/core/v1/registration-account-binding"
)

type Config struct{}

func (c Config) GetBaseUrl(isProduction bool) string {
	if isProduction {
		return PRODUCTION_BASE_URL
	}
	return SANDBOX_BASE_URL
}
