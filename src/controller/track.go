package controller

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/miniwaveme/api/src/db"
	"github.com/miniwaveme/api/src/document"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type (
	TrackController struct{}
)

func NewTrackController() *TrackController {
	return &TrackController{}
}

func (tc TrackController) GetTrackList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var results []document.Track
	err := db.DS().DefaultDB().C("track").Find(nil).All(&results)

	if err != nil {
		logrus.Error(err)
	}

	json.NewEncoder(w).Encode(results)
}

func (tc TrackController) GetTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	result := document.Track{}

	err := db.DS().DefaultDB().C("track").FindId(bson.ObjectIdHex(id)).One(&result)

	if err != nil {
		logrus.Error(err)
	}

	json.NewEncoder(w).Encode(result)
}

func (tc TrackController) CreateTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var track document.Track
	err := json.NewDecoder(r.Body).Decode(&track)

	if err != nil {
		logrus.Error(err)
	}

	track.Id = bson.NewObjectId()
	err = db.DS().DefaultDB().C("track").Insert(track)

	if err != nil {
		panic(err)
		logrus.Error(err)
	}
}

func (tc TrackController) UpdateTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (tc TrackController) RemoveTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	err := db.DS().DefaultDB().C("track").RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		logrus.Error(err)
	}
}
