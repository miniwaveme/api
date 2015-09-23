package repo

import (
	"github.com/miniwaveme/api/src/db"
	"github.com/miniwaveme/api/src/document"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func CreateMasterKey() (string, string, error) {
	return CreateAppKey(document.RoleMaster)
}

func CreateAppKey(role string) (string, string, error) {

	col := db.DS().DefaultDB().C("key")
	oid := bson.NewObjectId()
	err := col.Insert(&document.Key{Id: oid, AppKey: uuid.NewV4().String(), AppSecret: uuid.NewV4().String(), Roles: []string{role}})

	if err != nil {
		return "", "", err
	}

	result := document.Key{}
	err = col.Find(bson.M{"_id": oid}).One(&result)

	if err != nil {
		return "", "", err
	}

	return result.AppKey, result.AppSecret, nil
}
