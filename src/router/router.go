package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request, Params)

type Router struct {
	router     *httprouter.Router
	middleware []Middleware
}

func New() *Router {
	r := &Router{}
	r.router = httprouter.New()
	r.router.NotFound = notFoundHandler()

	return r
}

func (r *Router) Run(address string) error {
	return http.ListenAndServe(address, r.chain())
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h := r.chain()
	h.ServeHTTP(w, req)
}

func notFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		WriteError(w, 404, "404 Not Found")
	})
}

func (r *Router) GetRouter() *httprouter.Router {
	return r.router
}

func (r *Router) AddRoute(m, p string, fn Handle) {
	wf := func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		fn(w, req, paramsFromHTTPRouter(p))
	}

	r.router.Handle(m, p, wf)
}
