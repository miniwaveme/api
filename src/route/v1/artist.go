package route

import (
	"github.com/julienschmidt/httprouter"
	"api/src/controller"
)

func RegisterArtistRoutes(r *httprouter.Router) {

	ac := controller.NewArtistController()

	r.GET("/v1/artist/list/:page", ac.GetArtistList)
	r.GET("/v1/artist/get/:id", ac.GetArtist)
	r.PUT("/v1/artist/:id", ac.UpdateArtist)
	r.POST("/v1/artist", ac.CreateArtist)
	r.DELETE("/v1/artist/:id", ac.RemoveArtist)
}
