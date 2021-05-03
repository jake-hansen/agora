package corsmiddleware

import (
	"github.com/gin-contrib/cors"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Cfg provides a Config using values obtained from a Viper instance.
func Cfg(v *viper.Viper) (*cors.Config, error) {
	maxAge := v.GetDuration("api.cors.maxage")
	allowedOrigins := v.GetStringSlice("api.cors.origins.allowed")

	c := cors.Config{
		AllowAllOrigins:        false,
		AllowOrigins:           allowedOrigins,
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowCredentials:       true,
		MaxAge:                 maxAge,
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowWildcard:          false,
		AllowBrowserExtensions: true,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}

	return &c, nil
}

// Provide returns the CORSMiddleware configured with the provided Config.
func Provide(cfg *cors.Config) *CORSMiddleware {
	return &CORSMiddleware{Config: *cfg}
}

var (
	// ProviderProductionSet provides CORSMiddleware for use in production.
	ProviderProductionSet = wire.NewSet(Provide, Cfg)
)
