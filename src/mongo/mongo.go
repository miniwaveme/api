package mongo

import (
	"fmt"
	"github.com/miniwaveme/api/src/config"
	"gopkg.in/mgo.v2"
)

// ------------------------------------------------------------------------------------------------

var Ds = initDataStore()

func initDataStore() *DataStore {
	return &DataStore{newSession(), getDbName()}
}

type DataStore struct {
	session *mgo.Session
	dbName  string
}

func (ds DataStore) GetSession() *mgo.Session {
	return ds.session.Copy()
}

func (ds DataStore) GetDbName() string {
	return ds.dbName
}

// ------------------------------------------------------------------------------------------------

func newSession() *mgo.Session {
	session, err := mgo.Dial(getUrl())
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session
}

// ------------------------------------------------------------------------------------------------

func getUrl() string {
	return fmt.Sprintf("%s:%s/%s", getDbUrl(), getDbPort(), getDbName())
}

func getDbUrl() string {
	return config.GetConfig().GetString("db_url")
}

func getDbName() string {
	return config.GetConfig().GetString("db_name")
}

func getDbPort() string {
	return config.GetConfig().GetString("db_port")
}
