package cookieservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/spf13/viper"
	"net/http"
)

func Cfg(v *viper.Viper) *Config {
	domain:= v.GetString("api.domain")
	env := v.GetString("environment")

	prodOrStaging := env == "prod" || env == "staging"

	secureCookies := prodOrStaging

	cfg := &Config{
		Domain: domain,
		SecureCookies: secureCookies,
		SameSite: http.SameSiteLaxMode,
	}

	if prodOrStaging {
		cfg.SameSite = http.SameSiteNoneMode
	}

	return cfg
}

func Provide(cfg *Config) *CookieService {
	return NewCookieService(*cfg)
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.CookieService), new(*CookieService)),Cfg)
)
