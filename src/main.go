package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/miniwaveme/api/src/config"
	"github.com/miniwaveme/api/src/logger"
	"github.com/miniwaveme/api/src/route"
	"net/http"
	"os"
)

func main() {
	config.LoadConfig()

	logger.LoadLogger()
	log := logger.GetLogger()
	log.Info("application started")

	r := httprouter.New()
	route.RegisterRoutesV1(r)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), r))
}
