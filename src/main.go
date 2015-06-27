package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	config "github.com/miniwaveme/api/src/config"
	"os"
)

func main() {
	f, _ := os.OpenFile("../log/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)
	log.SetFormatter(&log.TextFormatter{DisableColors: true})
	log.Debug("debug")
	log.Info("info")
	log.Warn("warning")
	log.Error("error")
	//log.Fatal("fatal") On fatal app is shutdown
	//log.Panic("panic") On panic app is shutdown
	fmt.Println("Hello World")

	config.LoadConfig()
	conf := config.GetConfig()
	fmt.Print(conf.GetBool("miniwaveme.log.enabled"))
}
