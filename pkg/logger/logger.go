package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var instance = logrus.New()

func init() {
	instance.Formatter = new(logrus.TextFormatter)
	instance.Formatter.(*logrus.TextFormatter).DisableColors = false
	instance.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
	instance.Formatter.(*logrus.TextFormatter).FullTimestamp = true

	// TODO Automatically adjust log output level based on environment startup configuration
	instance.Level = logrus.TraceLevel
	instance.Out = os.Stdout
}

func GetLogger(level string) *logrus.Logger {
	switch strings.ToLower(level) {
	case "trace":
		instance.Level = logrus.TraceLevel
	case "debug":
		instance.Level = logrus.DebugLevel
	case "info":
		instance.Level = logrus.InfoLevel
	case "warn":
		instance.Level = logrus.WarnLevel
	case "error":
		instance.Level = logrus.ErrorLevel
	case "panic":
		instance.Level = logrus.PanicLevel
	}
	return instance
}
