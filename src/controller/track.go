package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	TrackController struct{}
)

func NewTrackController() *TrackController {
	return &TrackController{}
}

func (tc TrackController) GetTrackList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (tc TrackController) GetTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (tc TrackController) CreateTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (tc TrackController) UpdateTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (tc TrackController) RemoveTrack(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
