package spcuLog

import (
	"github.com/sirupsen/logrus"
)

// Time format for our log package
const iso8601DateFormat = "2006-01-02T15:04:05"

// SpcuLoggerConfig
// It specifies the details about a new SpcuLogger
type SpcuLoggerConfig struct {
	IsProductionMode bool `default:"true"`
}

// SpcuLogger
// It is a wrapper for logrus.Logger.
type SpcuLogger struct {
	*logrus.Logger
	config SpcuLoggerConfig
}

// NewLogger
// It is to create new SpcuLogger by a SpcuLoggerConfig.
func NewLogger(config SpcuLoggerConfig) *SpcuLogger {
	// Create new logrus logger
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.InfoLevel)
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: iso8601DateFormat,
	})
	// Adds file and line number
	logrusLogger.SetReportCaller(true)

	logrusLogger.Infof("Log package has been provided for:\n%+v", config)

	return &SpcuLogger{
		logrusLogger,
		config,
	}
}
