package document

import (
	"gopkg.in/mgo.v2/bson"
)

type Artist struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Images []Image       `json:"images" bson:"images"`
}
