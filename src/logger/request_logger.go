package logger

import (
	"github.com/Sirupsen/logrus"
	"net/http"
)

type loggingHandler struct {
	logger  *logrus.Logger
	handler http.Handler
}

func (lh loggingHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lh.logger.Info("Request receive")
	lh.handler.ServeHTTP(w, req)
}

func LogRequestMiddleware(h http.Handler) http.Handler {
	return loggingHandler{GetLogger(), h}
}
