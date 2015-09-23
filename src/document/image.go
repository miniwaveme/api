package document

import (
	"gopkg.in/mgo.v2/bson"
)

type Image struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Path string        `json:"path" bson:"path"`
	Url  string        `json:"url" bson:"url"`
}
