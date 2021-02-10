package jwtservice

import (
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Cfg provides a new Config using values from a Viper.
func Cfg(v *viper.Viper) (*Config, error) {
	dur, err := time.ParseDuration(v.GetString("jwtservice.duration"))
	if err != nil {
		return nil, fmt.Errorf("jwtservice: %w", err)
	}

	cfg := &Config{
		Issuer:     v.GetString("jwtservice.issuer"),
		SigningKey: v.GetString("jwtservice.signingkey"),
		Duration:   dur,
	}

	return cfg, nil
}

// CfgTest provides the passed Config.
func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

// Provide returns a new JWTService with the specified config.
func Provide(config *Config) *Service {
	return &Service{*config}
}

var (
	// ProviderProductionSet provides a new Service for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(JWTService), new(*Service)), Cfg)

	// ProviderTestSet provides a new Service for testing.
	ProviderTestSet = wire.NewSet(Provide, wire.Bind(new(JWTService), new(*Service)), CfgTest)
)
