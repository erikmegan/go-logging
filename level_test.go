package gologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("this is trace") // will not shown in console
	logger.Debug("this is debug") // will not shown in console
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
	// logger.Fatal("this is fatal") // will exit
	// logger.Panic("this is panic") // will exit
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // defult is InfoLevel

	logger.Trace("this is trace") // will shown in console
	logger.Debug("this is debug") // will shown in console
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
