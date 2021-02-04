package jwt

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"time"
)

func Cfg(v *viper.Viper) (*Config, error) {
	dur, err := time.ParseDuration(v.GetString("jwt.duration"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Issuer:     v.GetString("jwt.issuer"),
		SigningKey: v.GetString("jwt.signingkey"),
		Duration:   dur,
	}

	return cfg, nil
}

func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

// Provider returns a new JWTService with the specified config.
func Provider(config *Config) *Service {
	return &Service{*config}
}

var (
	ProviderProductionSet = wire.NewSet(Provider, Cfg)
	ProviderTestSet		  = wire.NewSet(Provider, CfgTest)
)
