package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	ArtistController struct{}
)

func NewArtistController() *ArtistController {
	return &ArtistController{}
}

func (ac ArtistController) GetArtist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (ac ArtistController) CreateArtist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (ac ArtistController) UpdateArtist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (ac ArtistController) RemoveArtist(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
