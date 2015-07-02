package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	AlbumController struct{}
)

func NewAlbumController() *AlbumController {
	return &AlbumController{}
}

func (ac AlbumController) GetAlbumList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (ac AlbumController) GetAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (ac AlbumController) CreateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (ac AlbumController) UpdateAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (ac AlbumController) RemoveAlbum(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
