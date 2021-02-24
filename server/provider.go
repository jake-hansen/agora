package server

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/router"
	"github.com/spf13/viper"
)

// Cfg provides a new Config using values from a Viper.
func Cfg(v *viper.Viper) (*Config, error) {
	cfg := &Config{
		Address: v.GetString("server.address"),
	}

	return cfg, nil
}

// CfgTest provides the passed Config.
func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

// Provide provides a new Server containing the given Config and Router.
func Provide(cfg *Config, router *router.Router) *Server {
	return New(*cfg, router)
}

var (
	// ProviderProductionSet provides a new Server for use in production.
	ProviderProductionSet = wire.NewSet(Provide, Cfg)
)
