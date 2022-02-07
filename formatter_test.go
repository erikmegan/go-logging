package gologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hello logging")
	logger.Warn("hello logging")
	logger.Error("hello logging")
}
