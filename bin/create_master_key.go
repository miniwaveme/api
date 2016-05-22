package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/miniwaveme/api/src/logger"
	"github.com/miniwaveme/api/src/manager"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Master Key"
	app.Usage = "Create the Master Key"
	app.Action = func(c *cli.Context) {

		log := logger.GetLogger()

		appKey, appSecret, err := manager.CreateMasterKey()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s | %s %s", color.YellowString("app_key:"), color.GreenString(appKey), color.YellowString("app_secret:"), color.GreenString(appSecret))
		fmt.Println("")
	}

	app.Run(os.Args)
}
