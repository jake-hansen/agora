package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	handlers2 "github.com/jake-hansen/agora/router/handlers"
	"github.com/spf13/viper"
)

func Cfg(v *viper.Viper, middleware []gin.HandlerFunc, handlers []handlers2.Handler) (*Config, error) {
	cfg := &Config{
		Environment:  v.GetString("environment"),
		Middleware:   middleware,
		Handlers:     handlers,
		RootEndpoint: v.GetString("api.endpoints.root"),
	}

	return cfg, nil
}

func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

func Provide(cfg *Config) *Router {
	return New(*cfg)
}

var (
	ProviderProductionSet = wire.NewSet(Provide, Cfg)
	ProviderTestSet		  = wire.NewSet(Provide, CfgTest)
) 
