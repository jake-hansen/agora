package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Provide(cfg *zap.Config) (*Log, func(), error) {
	log, err := NewLog(cfg)
	cleanup := func() {
		log.ZapLogger.Sync()
	}
	return log, cleanup, err
}

func Cfg(v *viper.Viper) *zap.Config {
	switch environment := v.GetString("environment"); environment {
	case "dev":
		return devZapConfig()
	default:
		return productionZapConfig()
	}
}

func productionZapConfig() *zap.Config {
	return &zap.Config {
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zapcore.EncoderConfig {
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func devZapConfig() *zap.Config {
	cfg := zap.NewDevelopmentConfig()
	return &cfg
}

var (
	ProviderProductionSet = wire.NewSet(Provide, Cfg)
	ProviderTestSet = wire.NewSet()
)
