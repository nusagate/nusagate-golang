package nusagate

const (
	PRODUCTION_URL = "https://api.nusagate.com"
	SANDBOX_URL    = "https://api.sandbox.nusagate.com"
)

type ConfigOption struct {
	IsProduction bool
	ApiKey       string
	SecretKey    string
	BaseUrl      string
}

func getBaseUrl(isProduction bool) string {
	if isProduction {
		return PRODUCTION_URL
	} else {
		return SANDBOX_URL
	}
}

func DefaultConfig(isProduction bool, apiKey string, secretKey string) *ConfigOption {
	return &ConfigOption{
		ApiKey:    apiKey,
		SecretKey: secretKey,
		BaseUrl:   getBaseUrl(isProduction),
	}
}
