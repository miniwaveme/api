package mongo

import (
	"fmt"
	"github.com/miniwaveme/api/src/config"
	"gopkg.in/mgo.v2"
)

var (
	databaseName  = config.GetConfig().GetString("db_name")
	gMongoSession = newSession()
)

func GetDatabase() *mgo.Database {

	return gMongoSession.DB(databaseName)
}

func newSession() *mgo.Session {

	appConfig := config.GetConfig()

	session, err := mgo.Dial(fmt.Sprintf("%s:%s", appConfig.GetString("db_url"), appConfig.GetString("db_port")))
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}
