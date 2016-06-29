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
	AlbumController struct{}
)

func NewAlbumController() *AlbumController {
	return &AlbumController{}
}

func (ac AlbumController) GetAlbumList(w http.ResponseWriter, r *http.Request, p router.Params) {
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
	albums, err := manager.FindAlbums(skip, pageSize)
	if err != nil {
		router.WriteError(w, 500, "database error")
		return
	}

	router.WriteJSON(w, albums)
}

func (ac AlbumController) GetAlbum(w http.ResponseWriter, r *http.Request, p router.Params) {
	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	album, err := manager.FindAlbumById(albumId)
	if err != nil {
		router.WriteError(w, 404, "Album not found")
		return
	}

	router.WriteJSON(w, album)
}

func (ac AlbumController) CreateAlbum(w http.ResponseWriter, r *http.Request, p router.Params) {
	r.ParseMultipartForm(0)

	artistId := r.FormValue("artistId")
	if artistId == "" {
		router.WriteError(w, 404, "artistId parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(artistId) == false {
		router.WriteError(w, 400, "Invalid artistId")
		return
	}

	artist, err := manager.FindArtistById(artistId)
	if err != nil {
		router.WriteError(w, 404, "Artist not found")
		return
	}

	albumName := r.FormValue("name")
	if albumName == "" {
		router.WriteError(w, 400, "name parameter must be specified")
		return
	}

	albumYearFormValue := r.FormValue("year")
	if albumYearFormValue == "" {
		router.WriteError(w, 400, "year parameter must be specified")
		return
	}

	albumYear, err := strconv.Atoi(albumYearFormValue)
	if err != nil {
		router.WriteError(w, 404, "invalid year value")
		return
	}

	var img image.Image = nil
	_, header, err := r.FormFile("cover")
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

	album, err := manager.CreateAlbum(albumName, albumYear, artist, img)
	if err != nil {
		panic(err)
	}

	router.WriteJSON(w, album)
}

func (ac AlbumController) UpdateAlbum(w http.ResponseWriter, r *http.Request, p router.Params) {

	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(albumId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	r.ParseMultipartForm(0)
	artistId := r.FormValue("artistId")
	if artistId == "" {
		router.WriteError(w, 404, "artistId parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(artistId) == false {
		router.WriteError(w, 400, "Invalid artistId")
		return
	}

	artist, err := manager.FindArtistById(artistId)
	if err != nil {
		router.WriteError(w, 404, "Artist not found")
		return
	}

	albumName := r.FormValue("name")
	if albumName == "" {
		router.WriteError(w, 400, "name parameter must be specified")
		return
	}

	albumYearFormValue := r.FormValue("year")
	if albumYearFormValue == "" {
		router.WriteError(w, 400, "year parameter must be specified")
		return
	}

	albumYear, err := strconv.Atoi(albumYearFormValue)
	if err != nil {
		router.WriteError(w, 404, "invalid year value")
		return
	}

	var img image.Image = nil
	_, header, err := r.FormFile("cover")
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

	album, err := manager.UpdateAlbum(albumId, albumName, albumYear, artist, img)
	if err != nil {
		panic(err)
	}

	router.WriteJSON(w, album)
}

func (ac AlbumController) RemoveAlbum(w http.ResponseWriter, r *http.Request, p router.Params) {
	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(albumId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	err := manager.RemoveAlbum(albumId)
	if err != nil {
		router.WriteError(w, 404, "album not found")
	}
}

func (ac AlbumController) AddAlbumTrack(w http.ResponseWriter, r *http.Request, p router.Params) {
	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(albumId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	nbTrackParam, exists := p["nb"]
	if exists == false {
		router.WriteError(w, 400, "track number parameter must be specified")
		return
	}

	nbTrack, err := strconv.Atoi(nbTrackParam)
	if err != nil {
		router.WriteError(w, 400, "track number parameter value")
		return
	}

	trackName := r.FormValue("name")
	if trackName == "" {
		router.WriteError(w, 400, "track name parameter must be specified")
		return
	}

	album, err := manager.FindAlbumById(albumId)
	if err != nil {
		router.WriteError(w, 404, "Album not found")
		return
	}

	if album.TrackExist(nbTrack) == true {
		router.WriteError(w, 400, "Track already exists")
		return
	}

	album, err = manager.AddAlbumTrack(albumId, nbTrack, trackName)
	if err != nil {
		panic(err)
	}

	router.WriteJSON(w, album)
}

func (ac AlbumController) UpdateAlbumTrack(w http.ResponseWriter, r *http.Request, p router.Params) {
	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(albumId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	nbTrackParam, exists := p["nb"]
	if exists == false {
		router.WriteError(w, 404, "track number parameter must be specified")
		return
	}

	nbTrack, err := strconv.Atoi(nbTrackParam)
	if err != nil {
		router.WriteError(w, 404, "track number parameter value")
		return
	}

	trackName := r.FormValue("name")
	if trackName == "" {
		router.WriteError(w, 400, "track name parameter must be specified")
		return
	}

	album, err := manager.FindAlbumById(albumId)
	if err != nil {
		router.WriteError(w, 404, "Album not found")
		return
	}

	if album.TrackExist(nbTrack) == false {
		router.WriteError(w, 400, "Track doesn't exists")
		return
	}

	album, err = manager.UpdateAlbumTrack(albumId, nbTrack, trackName)
	if err != nil {
		panic(err)
	}

	router.WriteJSON(w, album)
}

func (ac AlbumController) RemoveAlbumTrack(w http.ResponseWriter, r *http.Request, p router.Params) {
	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(albumId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	nbTrackParam, exists := p["nb"]
	if exists == false {
		router.WriteError(w, 404, "track number parameter must be specified")
		return
	}

	nbTrack, err := strconv.Atoi(nbTrackParam)
	if err != nil {
		router.WriteError(w, 404, "track number parameter value")
		return
	}

	album, err := manager.FindAlbumById(albumId)
	if err != nil {
		router.WriteError(w, 404, "Album not found")
		panic(err)
	}

	if album.TrackExist(nbTrack) == false {
		router.WriteError(w, 404, "Track doesn't exists")
		return
	}

	album, err = manager.RemoveAlbumTrack(albumId, nbTrack)
	if err != nil {
		panic(err)
	}

	router.WriteJSON(w, album)
}

func (ac AlbumController) AddTrackArtist(w http.ResponseWriter, r *http.Request, p router.Params) {
	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(albumId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	artistId, exists := p["artistId"]
	if exists == false {
		router.WriteError(w, 404, "artistId parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(artistId) == false {
		router.WriteError(w, 400, "Invalid artistId")
		return
	}

	artist, err := manager.FindArtistById(artistId)
	if err != nil {
		router.WriteError(w, 404, "Artist not found")
		return
	}

	nbTrackParam, exists := p["nb"]
	if exists == false {
		router.WriteError(w, 404, "track number parameter must be specified")
		return
	}

	nbTrack, err := strconv.Atoi(nbTrackParam)
	if err != nil {
		router.WriteError(w, 404, "track number parameter value")
		return
	}

	album, err := manager.FindAlbumById(albumId)
	if err != nil {
		router.WriteError(w, 404, "Album not found")
		panic(err)
	}

	if album.TrackExist(nbTrack) == false {
		router.WriteError(w, 404, "Track doesn't exists")
		return
	}

	album, err = manager.AddTrackArtist(albumId, nbTrack, artist)

	router.WriteJSON(w, album)
}

func (ac AlbumController) RemoveTrackArtist(w http.ResponseWriter, r *http.Request, p router.Params) {
	albumId, exists := p["id"]
	if exists == false {
		router.WriteError(w, 404, "id parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(albumId) == false {
		router.WriteError(w, 400, "Invalid id")
		return
	}

	artistId, exists := p["artistId"]
	if exists == false {
		router.WriteError(w, 404, "artistId parameter must be specified")
		return
	}

	if bson.IsObjectIdHex(artistId) == false {
		router.WriteError(w, 400, "Invalid artistId")
		return
	}

	artist, err := manager.FindArtistById(artistId)
	if err != nil {
		router.WriteError(w, 404, "Artist not found")
		return
	}

	nbTrackParam, exists := p["nb"]
	if exists == false {
		router.WriteError(w, 404, "track number parameter must be specified")
		return
	}

	nbTrack, err := strconv.Atoi(nbTrackParam)
	if err != nil {
		router.WriteError(w, 404, "track number parameter value")
		return
	}

	album, err := manager.FindAlbumById(albumId)
	if err != nil {
		router.WriteError(w, 404, "Album not found")
		panic(err)
	}

	if album.TrackExist(nbTrack) == false {
		router.WriteError(w, 404, "Track doesn't exists")
		return
	}

	album, err = manager.RemoveTrackArtist(albumId, nbTrack, artist)

	router.WriteJSON(w, album)
}
