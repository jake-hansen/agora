package log

import (
	"go.uber.org/zap"
)

type Log struct {
	*zap.Logger
}

func NewLog(cfg *zap.Config) (*Log, error) {
	logger, err := cfg.Build()
	if err != nil {
		return nil,  err
	}

	return &Log{logger}, nil
}
