package logger

import (
	"github.com/Sirupsen/logrus"
	"os"
)

var (
	gLogger = loadLogger()
)

func GetLogger() *logrus.Logger {
	return gLogger
}

func loadLogger() *logrus.Logger {

	logger := logrus.New()
	logger.Out = os.Stderr

	return logger
}
