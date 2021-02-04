package config

import (
	"github.com/spf13/viper"
	"strings"
)

var config *viper.Viper

// Init is used to initialize the configuration for Agora in the specified environment. Configuration files are searched
// for in the following locations: /etc/agora, $HOME/agora, and the directory the executable is run from.
func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("/etc/agora/")
	config.AddConfigPath("$HOME/agora")
	config.AddConfigPath(".")

	config.SetEnvPrefix("agora")
	config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
}

func ProvideViper() *viper.Viper {
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

// GetConfig returns a pointer to the Viper instance for the application.
func GetConfig() *viper.Viper {
	return config
}
