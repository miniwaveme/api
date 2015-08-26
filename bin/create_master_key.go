package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/miniwaveme/api/src/db"
	"github.com/miniwaveme/api/src/document"
	"github.com/miniwaveme/api/src/logger"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Master Key"
	app.Usage = "Create the Master Key"
	app.Action = func(c *cli.Context) {

		log := logger.GetLogger()

		col := db.DS().DefaultDB().C("key")

		oid := bson.NewObjectId()
		err := col.Insert(&document.Key{Id: oid, AppKey: uuid.NewV4().String(), AppSecret: uuid.NewV4().String(), Roles: []string{"ROLE_MASTER"}})

		if err != nil {
			log.Fatal(err)
		}

		result := document.Key{}
		err = col.Find(bson.M{"_id": oid}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s | %s %s", color.YellowString("app_key:"), color.GreenString(result.AppKey), color.YellowString("app_secret:"), color.GreenString(result.AppSecret))
		fmt.Println("")
	}

	app.Run(os.Args)
}
