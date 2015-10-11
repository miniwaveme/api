package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/miniwaveme/api/src/conf"
	"github.com/miniwaveme/api/src/logger"
	"github.com/miniwaveme/api/src/route"
	"net/http"
)

func main() {

	c := conf.C()

	log := logger.GetLogger()
	log.Info("applicationclear started")

	r := httprouter.New()
	route.RegisterRoutesV1(r)

	log.Fatal(http.ListenAndServe(":"+c.GetString("api_port"), r))
}
