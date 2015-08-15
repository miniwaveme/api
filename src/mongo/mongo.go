package mongo

import (
	"fmt"
	"github.com/miniwaveme/api/src/config"
	"gopkg.in/mgo.v2"
)

type DataStore struct {
	session *mgo.Session
}

var (
	dataStore = initDataStore()
)

func GetSession() *mgo.Session {
	return dataStore.session.Copy()
}

func newSession() *mgo.Session {

	appConfig := config.GetConfig()

	session, err := mgo.Dial(fmt.Sprintf("%s:%s/%s", appConfig.GetString("db_url"), appConfig.GetString("db_port"), config.GetConfig().GetString("db_name")))
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session
}

func initDataStore() *DataStore {
	return &DataStore{newSession()}
}
