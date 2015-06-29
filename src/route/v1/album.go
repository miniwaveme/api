package route

import (
	"github.com/julienschmidt/httprouter"
	"api/src/controller"
)

func RegisterAlbumRoutes(r *httprouter.Router) {

	ac := controller.NewAlbumController()

	r.GET("/v1/album/list/:page", ac.GetAlbumList)
	r.GET("/v1/album/get/:id", ac.GetAlbum)
	r.PUT("/v1/album/:id", ac.UpdateAlbum)
	r.POST("/v1/album", ac.CreateAlbum)
	r.DELETE("/v1/album/:id", ac.RemoveAlbum)
}
