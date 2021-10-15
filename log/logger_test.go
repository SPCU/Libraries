package spcuLog

import "testing"

func TestSpcuLogger(t *testing.T) {
	logger, err := NewLogger(SpcuLoggerConfig{})
	if err != nil {
		t.Fatalf(err.Error())
	}
	logger.Traceln("This is a trace log.")
	logger.Infoln("This is an information log.")
	logger.Warningln("This is a warning log.")
	logger.Warnln("This is a warn log.")
	logger.Debugln("This is a debug log.")
	logger.Errorln("This is an error log.")
}

func TestSpcuLoggerWithFileOutput(t *testing.T) {
	logger, err := NewLogger(SpcuLoggerConfig{
		IsReportCaller: true,
		OutputFileConfig: &OutputFileConfig{
			Path:             ".",
			MaxNumberOfFiles: 10,
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	logger.Traceln("This is a trace log.")
	logger.Infoln("This is an information log.")
	logger.Warningln("This is a warning log.")
	logger.Warnln("This is a warn log.")
	logger.Debugln("This is a debug log.")
	logger.Errorln("This is an error log.")
}

func TestSpcuLoggerFileThreshold(t *testing.T) {
	logger, err := NewLogger(SpcuLoggerConfig{
		IsReportCaller: true,
		OutputFileConfig: &OutputFileConfig{
			Path:             ".",
			Name:             "counter",
			MaxNumberOfFiles: 20,
			MaxSizeInBytes:   256000,
		},
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	for i := 0; i < 50000; i++ {
		logger.Infoln(i, "This is an information log.")
		logger.Warnln(i, "This is a warn log.")
		logger.Errorln(i, "This is an error log.")
	}
}
