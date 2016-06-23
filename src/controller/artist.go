package controller

import (
	"github.com/miniwaveme/api/src/manager"
	"github.com/miniwaveme/api/src/router"
	"gopkg.in/mgo.v2/bson"
	"image"
	"image/jpeg"
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
	r.ParseMultipartForm(0)

	artistName := r.FormValue("name")
	if artistName == "" {
		router.WriteError(w, 400, "name parameter must be specified")
		return
	}

	var img image.Image = nil
	_, header, err := r.FormFile("picture")
	if err == nil && header != nil {
		file, err := header.Open()
		if err != nil {
			router.WriteError(w, 400, "cant open file")
			return
		}

		img, err = jpeg.Decode(file)
		if err != nil {
			router.WriteError(w, 400, "cant decode JPEG")
			return
		}
		file.Close()
	}
	artist, err := manager.CreateArtist(artistName, img)
	if err != nil {
		panic(err)
	}

	router.WriteJSON(w, artist)
}

func (ac ArtistController) UpdateArtist(w http.ResponseWriter, r *http.Request, p router.Params) {
	r.ParseMultipartForm(0)

	artistName := r.FormValue("name")
	if artistName == "" {
		router.WriteError(w, 400, "name parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(p["id"]) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	var img image.Image = nil
	_, header, err := r.FormFile("picture")
	if err == nil && header != nil {

		file, err := header.Open()
		if err != nil {
			router.WriteError(w, 400, "cant open file")
			return
		}

		img, err = jpeg.Decode(file)
		if err != nil {
			router.WriteError(w, 400, "cant decode JPEG")
			return
		}
		file.Close()
	}
	artist, err := manager.UpdateArtist(p["id"], artistName, img)
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
