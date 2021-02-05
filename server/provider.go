package server

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/router"
	"github.com/spf13/viper"
)

func Cfg(v *viper.Viper) (*Config, error) {
	cfg := &Config{
		Address: config.Build().GetString("server.address"),
	}

	return cfg, nil
}

func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

func Provide(cfg *Config, router *router.Router) *Server {
	return New(*cfg, router)
}

var (
	ProviderProductionSet = wire.NewSet(Provide, Cfg)
)