package logger

import (
	"io"
	"os"

	"github.com/YEgorLu/time-tracker/internal/config"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(v ...any)
	Error(v ...any)
	Info(v ...any)
}

var logger Logger

func init() {
	log := logrus.New()
	configure(log)
	logger = log
}

func Get() Logger {
	return logger
}

func configure(log *logrus.Logger) {
	log.SetFormatter(&logrus.JSONFormatter{})
	writers := []io.Writer{os.Stdout}
	if config.App.LogsPath != "" {
		logsFile, err := os.OpenFile(config.App.LogsPath, os.O_WRONLY|os.O_APPEND, 0444)
		if err != nil {
			log.Warn("error opening logs file ", config.App.LogsPath, " using only stdout")
		} else {
			writers = append(writers, logsFile)
		}
	}
	log.SetOutput(io.MultiWriter(writers...))
}
