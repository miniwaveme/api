package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miniwaveme/api/src/config"
	"github.com/miniwaveme/api/src/logger"
	"github.com/miniwaveme/api/src/mongo"
	"github.com/miniwaveme/api/src/route"
	"net/http"
)

func main() {

	conf := config.GetConfig()

	log := logger.GetLogger()
	log.Info("application started")

	r := httprouter.New()
	route.RegisterRoutesV1(r)

	ds := mongo.Ds
	mgoDatabase := ds.GetSession().DB(ds.GetDbName())
	fmt.Println(mgoDatabase.Name)

	log.Fatal(http.ListenAndServe(":"+conf.GetString("api_port"), r))
}
