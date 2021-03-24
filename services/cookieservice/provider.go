package cookieservice

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func Cfg(v *viper.Viper) *Config {
	domain:= v.GetString("api.domain")
	env := v.GetString("environment")

	secureCookies := env == "prod"

	cfg := &Config{
		Domain: domain,
		SecureCookies: secureCookies,
	}

	return cfg
}

func Provide(cfg *Config) *CookieService {
	return NewCookieService(*cfg)
}

var (
	ProviderSet = wire.NewSet(Provide, Cfg)
)
