package config

import (
	"strings"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Provide provides a new Viper instance configured using a configuration file
// located at /etc/agora, ~/agora, or the same directory the application
// executable is located. Will panic if there is an error finding or reading
// this file.
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

// ProviderSet provides a Viper instance.
var ProviderSet = wire.NewSet(Provide)
