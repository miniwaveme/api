package route

import (
	"github.com/julienschmidt/httprouter"
	"api/src/controller"
)

func RegisterKeyRoutes(r *httprouter.Router) {

	kc := controller.NewKeyController()

	r.GET("/v1/key/list/:page", kc.GetKeyList)
	r.GET("/v1/key/get/:id", kc.GetKey)
	r.PUT("/v1/key/:id", kc.UpdateKey)
	r.POST("/v1/key", kc.CreateKey)
	r.DELETE("/v1/key/:id", kc.RemoveKey)
}
