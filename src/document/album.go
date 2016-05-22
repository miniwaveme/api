package document

import (
	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Tracks    []Track       `json:"tracks" bson:"tracks"`
	Artist    Artist        `json:"artist" bson:"-"`
	ArtistRef bson.ObjectId `json:"-" bson:"artist"`
	Cover     Image         `json:"cover" bson:"cover"`
	Year      int           `json:"year" bson:"year"`
}

func (a Album) TrackExist(nbTrack int) bool {
	for _, track := range a.Tracks {
		if nbTrack == track.Number {
			return true
		}
	}
	return false
}
