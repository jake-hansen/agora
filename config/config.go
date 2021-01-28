package config

import (
	"github.com/spf13/viper"
	"strings"
)

var config *viper.Viper

// Init is used to initialize the configuration for Agora in the specified environment
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

func GetConfig() *viper.Viper {
	return config
}
