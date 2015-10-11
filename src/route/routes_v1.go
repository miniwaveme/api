package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/miniwaveme/api/src/controller"
)

func RegisterRoutesV1(r *httprouter.Router) *httprouter.Router {

	// Register app keys routes
	kc := controller.NewKeyController()
	r.GET("/v1/key/list/:page", kc.GetKeyList)
	r.GET("/v1/key/get/:id", kc.GetKey)
	r.PUT("/v1/key/:id", kc.UpdateKey)
	r.POST("/v1/key", kc.CreateKey)
	r.DELETE("/v1/key/:id", kc.RemoveKey)

	// Register tracks routes
	tc := controller.NewTrackController()
	r.GET("/v1/track/list/:page", tc.GetTrackList)
	r.GET("/v1/track/get/:id", tc.GetTrack)
	r.PUT("/v1/track/:id", tc.UpdateTrack)
	r.POST("/v1/track", tc.CreateTrack)
	r.DELETE("/v1/track/:id", tc.RemoveTrack)

	// Register artists routes
	arc := controller.NewArtistController()
	r.GET("/v1/artist/list/:page", arc.GetArtistList)
	r.GET("/v1/artist/get/:id", arc.GetArtist)
	r.PUT("/v1/artist/:id", arc.UpdateArtist)
	r.POST("/v1/artist", arc.CreateArtist)
	r.DELETE("/v1/artist/:id", arc.RemoveArtist)

	// Register albums routes
	alc := controller.NewAlbumController()
	r.GET("/v1/album/list/:page", alc.GetAlbumList)
	r.GET("/v1/album/get/:id", alc.GetAlbum)
	r.PUT("/v1/album/:id", alc.UpdateAlbum)
	r.POST("/v1/album", alc.CreateAlbum)
	r.DELETE("/v1/album/:id", alc.RemoveAlbum)

	return r
}
