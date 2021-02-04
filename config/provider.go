package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"strings"
)

func Provide() *viper.Viper {
	cfg := viper.New()
	cfg.SetConfigType("yaml")
	cfg.SetConfigName("config")
	cfg.AddConfigPath("/etc/agora/")
	cfg.AddConfigPath("$HOME/agora")
	cfg.AddConfigPath(".")

	cfg.SetEnvPrefix("agora")
	cfg.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	cfg.SetEnvKeyReplacer(replacer)

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}
	return cfg
}

var ProviderSet = wire.NewSet(Provide)
