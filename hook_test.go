package gologging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel} // set level
}
func (s SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("sample hook", entry.Level, entry.Message) // do something here if hooked
	return nil
}
func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("this is info")
	logger.Warn("this is warn") // will hook. call Fire func
}
