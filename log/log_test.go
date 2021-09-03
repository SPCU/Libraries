package spcuLog

import "testing"

func TestSpcuLogger(t *testing.T) {
	logger, err := NewLogger(SpcuLoggerConfig{})
	if err != nil {
		t.Fatalf(err.Error())
	}
	logger.Traceln("This is a trace log.")
	logger.Println("This is a print mode.")
	logger.Infoln("This is an information log.")
	logger.Warningln("This is a warning log.")
	logger.Warnln("This is a warn log.")
	logger.Debugln("This is a debug log.")
	logger.Errorln("This is an error log.")
}

func TestSpcuLoggerWithFileOutput(t *testing.T) {
	logPath := "temp.log"
	logger, err := NewLogger(SpcuLoggerConfig{
		IsReportCaller: true,
		OutputPath:     &logPath,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}
	logger.Traceln("This is a trace log.")
	logger.Println("This is a print mode.")
	logger.Infoln("This is an information log.")
	logger.Warningln("This is a warning log.")
	logger.Warnln("This is a warn log.")
	logger.Debugln("This is a debug log.")
	logger.Errorln("This is an error log.")
}
