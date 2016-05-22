package route

import (
	"github.com/miniwaveme/api/src/controller"
	"github.com/miniwaveme/api/src/router"
)

func RegisterRoutesV1(r *router.Router) {
	// Register app keys routes
	kc := controller.NewKeyController()
	r.AddRoute("GET", "/v1/key/list", kc.GetKeyList)
	r.AddRoute("GET", "/v1/key/get/:id", kc.GetKey)
	r.AddRoute("PUT", "/v1/key/:id", kc.UpdateKey)
	r.AddRoute("POST", "/v1/key", kc.CreateKey)
	r.AddRoute("DELETE", "/v1/key/:id", kc.RemoveKey)

	// Register artists routes
	arc := controller.NewArtistController()
	r.AddRoute("GET", "/v1/artist/list", arc.GetArtistList)
	r.AddRoute("GET", "/v1/artist/get/:id", arc.GetArtist)
	r.AddRoute("PUT", "/v1/artist/:id", arc.UpdateArtist)
	r.AddRoute("POST", "/v1/artist", arc.CreateArtist)
	r.AddRoute("DELETE", "/v1/artist/:id", arc.RemoveArtist)

	// Register albums routes
	alc := controller.NewAlbumController()
	r.AddRoute("GET", "/v1/album/list", alc.GetAlbumList)
	r.AddRoute("GET", "/v1/album/get/:id", alc.GetAlbum)
	r.AddRoute("PUT", "/v1/album/:id", alc.UpdateAlbum)
	r.AddRoute("POST", "/v1/album", alc.CreateAlbum)
	r.AddRoute("DELETE", "/v1/album/:id", alc.RemoveAlbum)
	r.AddRoute("POST", "/v1/album/:id/track/:nb", alc.AddAlbumTrack)
	r.AddRoute("PUT", "/v1/album/:id/track/:nb", alc.UpdateAlbumTrack)
	r.AddRoute("DELETE", "/v1/album/:id/track/:nb", alc.RemoveAlbumTrack)
}
