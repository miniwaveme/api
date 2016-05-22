package main

import (
	"github.com/miniwaveme/api/src/conf"
	"github.com/miniwaveme/api/src/logger"
	"github.com/miniwaveme/api/src/route"
	"github.com/miniwaveme/api/src/router"
)

func main() {

	c := conf.C()

	log := logger.GetLogger()
	log.Info("application started")

	r := router.New()
	route.RegisterRoutesV1(r)
	r.Use(logger.LogRequestMiddleware)

	log.Fatal(r.Run(":" + c.GetString("api_port")))
}
