package log

import (
	"time"

	"github.com/justoxh/go-server/config"
	"github.com/justoxh/gokit/logger"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLog(cfg *config.LoggerConfig) *logrus.Logger {
	options := &logger.Options{
		Level:          cfg.Level,
		WithCallerHook: true,
		Formatter:      cfg.Formatter,
		DisableConsole: cfg.DisableConsole,
		Write:          cfg.Write,
		Path:           cfg.Path,
		FileName:       cfg.FileName,
		MaxAge:         cfg.MaxAge * time.Hour,
		RotationTime:   cfg.RotationTime * time.Hour,
		Debug:          cfg.Debug,
	}
	Log = logger.GetLoggerWithOptions("default", options)
	return Log
}
