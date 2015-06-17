package main

import (
	"fmt"
	config "github.com/miniwaveme/api/src/config"
	logger "github.com/miniwaveme/api/src/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func main() {
	config.LoadConfig()
	conf := config.GetConfig()
	fmt.Print(conf.GetBool("miniwaveme.log.enabled"))

	logger.LoadLogger()
	log := logger.GetLogger()
	log.Info("application started")

    r := httprouter.New()
	route.RegisterAlbumRoutes(r)
	route.RegisterArtistRoutes(r)
	route.RegisterKeyRoutes(r)
	route.RegisterTrackRoutes(r)

	http.ListenAndServe("localhost"+":"+os.Getenv("API_PORT"), r)
}
