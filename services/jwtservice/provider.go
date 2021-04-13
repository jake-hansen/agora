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

	refreshDur, err := time.ParseDuration(v.GetString("jwtservice.refreshduration"))
	if err != nil {
		return nil, fmt.Errorf("jwtservice: %w", err)
	}

	cfg := &Config{
		Issuer:     v.GetString("jwtservice.issuer"),
		SigningKey: v.GetString("jwtservice.signingkey"),
		Duration:   dur,
		RefreshDuration: refreshDur,
	}

	return cfg, nil
}

// Provide returns a new JWTService with the specified config.
func Provide(config *Config) *JWTServiceImpl {
	return &JWTServiceImpl{*config}
}

var (
	// ProviderProductionSet provides a new JWTServiceImpl for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(JWTService), new(*JWTServiceImpl)), Cfg)
)
