package commons

const (
	SANDBOX_BASE_URL    = "https://api-uat.doku.com"
	PRODUCTION_BASE_URL = "https://dashboard.doku.com"
	ACCESS_TOKEN        = "/authorization/v1/access-token/b2b"
	CREATE_VA           = "/virtual-accounts/bi-snap-va/v1/transfer-va/create-va"
)

type Config struct{}

func (c Config) GetBaseUrl(isProduction bool) string {
	if isProduction {
		return PRODUCTION_BASE_URL
	}
	return SANDBOX_BASE_URL
}
