package app

import (
	"fmt"

	"github.com/yuseferi/just-ad-ch/log"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(level string) (logger *zap.Logger, err error) {
	cfg := log.NewConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	atom := zap.NewAtomicLevel()
	err = atom.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}

	cfg.Level = atom

	logger, err = cfg.Build()
	if err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)

	return logger, err
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Println(v ...interface{}) {
	l.logger.Warn(fmt.Sprint(v...))
}
