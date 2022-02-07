package gologging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("this is info. this is application.log file test")
	logger.Warn("this is warn. this is application.log file test")
	logger.Error("this is error. this is application.log file test")

}
