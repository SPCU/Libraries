package spcuLog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	. "os"
)

// Time format for our log package
const iso8601DateFormat = "2006-01-02T15:04:05"

// SpcuLoggerConfig
// It specifies the details about a new SpcuLogger
type SpcuLoggerConfig struct {
	IsProductionMode bool `default:"true"`
	IsReportCaller   bool `default:"false"`
	OutputPath       *string
}

// SpcuLogger
// It is a wrapper for logrus.Logger.
type SpcuLogger struct {
	*logrus.Logger
	config SpcuLoggerConfig
}

// NewLogger
// It is to create new SpcuLogger by a SpcuLoggerConfig.
func NewLogger(config SpcuLoggerConfig) (*SpcuLogger, error) {
	// Create new logrus logger
	ll := logrus.New()
	ll.SetLevel(logrus.InfoLevel)
	ll.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: iso8601DateFormat,
	})
	// Adds file and line number
	ll.SetReportCaller(config.IsReportCaller)
	// Check if it is going to be saved on file
	if config.OutputPath != nil {
		f, err := OpenFile(*config.OutputPath, O_RDWR|O_CREATE|O_APPEND, 0666)
		if err != nil {
			fmt.Println("Could Not Open Log File : " + err.Error())
		}
		ll.SetOutput(f)
	}

	ll.Infof("A SPCU logger has been provided: %+v\n", config)

	return &SpcuLogger{
		ll,
		config,
	}, nil
}
