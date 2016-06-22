package document

import (
	"gopkg.in/mgo.v2/bson"
)

type Artist struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Image Image         `json:"image" bson:"image"`
}
