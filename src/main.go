package main

import (
	"fmt"
	config "github.com/miniwaveme/api/src/config"
	logger "github.com/miniwaveme/api/src/logger"
)

func main() {
	config.LoadConfig()
	conf := config.GetConfig()
	fmt.Print(conf.GetBool("miniwaveme.log.enabled"))

	logger.LoadLogger()
	log := logger.GetLogger()
	log.Info("application started")
}
