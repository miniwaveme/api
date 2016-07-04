package document

import (
	"gopkg.in/mgo.v2/bson"
)

type Track struct {
	Number     int             `json:"number" bson:"number"`
	Name       string          `json:"name" bson:"name"`
	ArtistsRef []bson.ObjectId `json:"artists" bson:"artists"`
	Duration   int             `json:"duration" bson:"duration"`
	Bpm        int             `json:"bpm" bson:"bpm"`
	Path       string          `json:"path" bson:"path"`
	Url        string          `json:"url" bson:"url"`
}
