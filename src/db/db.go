package db

import (
	"fmt"
	"github.com/miniwaveme/api/src/conf"
	"gopkg.in/mgo.v2"
)

var ds = initDataStore()

func initDataStore() *DataStore {
	return &DataStore{newSession(), conf.C().GetString("db_name")}
}

type DataStore struct {
	session *mgo.Session
	dbName  string
}

func DS() *DataStore {
	return ds
}

func (ds DataStore) DefaultDB() *mgo.Database {
	return ds.Sess().DB(ds.dbName)
}

func (ds DataStore) Sess() *mgo.Session {
	return ds.session.Copy()
}

func (ds DataStore) GetDbName() string {
	return ds.dbName
}

func newSession() *mgo.Session {

	session, err := mgo.Dial(getUrl())
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session
}

func getUrl() string {

	return fmt.Sprintf("%s:%s", conf.C().GetString("db_url"), conf.C().GetString("db_port"))
}
