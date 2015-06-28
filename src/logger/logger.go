package logger

import (
	"github.com/Sirupsen/logrus"
	"os"
)

var (
	gLogger *logrus.Logger
)

func GetLogger() *logrus.Logger {
	return gLogger
}

func LoadLogger() {
	gLogger = logrus.New()
	gLogger.Out = os.Stderr
}
