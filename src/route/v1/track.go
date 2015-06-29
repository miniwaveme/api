package route

import (
	"github.com/julienschmidt/httprouter"
	"api/src/controller"
)

func RegisterTrackRoutes(r *httprouter.Router) {

	tc := controller.NewTrackController()

	r.GET("/v1/track/list/:page", tc.GetTrackList)
	r.GET("/v1/track/get/:id", tc.GetTrack)
	r.PUT("/v1/track/:id", tc.UpdateTrack)
	r.POST("/v1/track", tc.CreateTrack)
	r.DELETE("/v1/track/:id", tc.RemoveTrack)
}
