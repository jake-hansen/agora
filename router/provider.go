package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/router/handlers"
	"github.com/spf13/viper"
)

// Cfg provides a new Config using values from a Viper and contins the given middleware and handlers.
func Cfg(v *viper.Viper, middleware []gin.HandlerFunc, handlerManager *handlers.HandlerManager) (*Config, error) {
	cfg := &Config{
		Environment:    v.GetString("environment"),
		Middleware:     middleware,
		HandlerManager: handlerManager,
		RootEndpoint:   v.GetString("api.endpoints.root"),
	}

	return cfg, nil
}

// CfgTest provides the passed Config.
func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

// Provide provides a new Router using the given Config.
func Provide(cfg *Config) *Router {
	return New(*cfg)
}

var (
	// ProviderProductionSet provides a new Router for use in production.
	ProviderProductionSet = wire.NewSet(Provide, Cfg, handlers.ProvideHandlerManager)

	// ProviderTestSet provides a new Router for use in testing.
	ProviderTestSet = wire.NewSet(Provide, CfgTest)
)
