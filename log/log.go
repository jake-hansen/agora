package log

import (
	"go.uber.org/zap"
)

// Log is a log implementation for the application that uses a Zap logger.
type Log struct {
	ZapLogger *zap.Logger
}

// NewLog returns a new Log configured with the provided Config.
func NewLog(cfg *zap.Config) (*Log, error) {
	logger, err := cfg.Build()
	if err != nil {
		return nil,  err
	}

	return &Log{logger}, nil
}
