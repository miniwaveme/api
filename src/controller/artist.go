package controller

import (
	"github.com/miniwaveme/api/src/manager"
	"github.com/miniwaveme/api/src/router"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type (
	ArtistController struct{}
)

func NewArtistController() *ArtistController {
	return &ArtistController{}
}

func (ac ArtistController) GetArtistList(w http.ResponseWriter, r *http.Request, p router.Params) {
	r.ParseForm()

	pageParam := r.FormValue("page")
	if pageParam == "" {
		pageParam = "1"
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		router.WriteError(w, 404, "invalid page value")
		return
	}

	pageSizeParam := r.FormValue("pageSize")
	if pageSizeParam == "" {
		pageSizeParam = "50"
	}

	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		router.WriteError(w, 404, "invalid pageSize value")
		return
	}

	skip := (page - 1) * pageSize
	artists, err := manager.FindArtists(skip, pageSize)
	if err != nil {
		router.WriteError(w, 500, "database error")
		return
	}

	router.WriteJSON(w, artists)
}

func (ac ArtistController) GetArtist(w http.ResponseWriter, r *http.Request, p router.Params) {

	artistId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	artist, err := manager.FindArtistById(artistId)
	if err != nil {
		router.WriteError(w, 404, "Artist not found")
		return
	}

	router.WriteJSON(w, artist)
}

func (ac ArtistController) CreateArtist(w http.ResponseWriter, r *http.Request, p router.Params) {
	r.ParseForm()
	artistName := r.FormValue("name")

	if artistName == "" {
		router.WriteError(w, 400, "name parameter must be specified")
		return
	}

	artist, err := manager.CreateArtist(artistName)
	if err != nil {
		panic(err)
	}

	router.WriteJSON(w, artist)
}

func (ac ArtistController) UpdateArtist(w http.ResponseWriter, r *http.Request, p router.Params) {
	r.ParseForm()
	artistName := r.FormValue("name")

	if artistName == "" {
		router.WriteError(w, 400, "name parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(p["id"]) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	artist, err := manager.UpdateArtist(p["id"], artistName)
	if err != nil {
		router.WriteError(w, 404, "Artist not found")
		return
	}

	router.WriteJSON(w, artist)
}

func (ac ArtistController) RemoveArtist(w http.ResponseWriter, r *http.Request, p router.Params) {
	artistId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(artistId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	err := manager.RemoveArtist(artistId)
	if err != nil {
		router.WriteError(w, 404, "Artist not found")
	}
}
