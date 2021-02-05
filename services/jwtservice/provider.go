package jwtservice

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"time"
)

func Cfg(v *viper.Viper) (*Config, error) {
	dur, err := time.ParseDuration(v.GetString("jwtservice.duration"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Issuer:     v.GetString("jwtservice.issuer"),
		SigningKey: v.GetString("jwtservice.signingkey"),
		Duration:   dur,
	}

	return cfg, nil
}

func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

// Provide returns a new JWTService with the specified config.
func Provide(config *Config) *Service {
	return &Service{*config}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(JWTService), new(*Service)), Cfg)
	ProviderTestSet		  = wire.NewSet(Provide, wire.Bind(new(JWTService), new(*Service)), CfgTest)
)
