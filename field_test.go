package gologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "erik").Info("this is info with field")
	logger.WithField("username", "erik").WithField("name", "my name").Info("this is info with multiple field")
}

func TestFieldUsingJson(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "erik").Info("this is info with field")
	logger.WithField("username", "erik").WithField("name", "my name").Info("this is info with multiple field")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.WithFields(logrus.Fields{"username": "erik", "name": "my name"}).Info("this is info with fields")
}
