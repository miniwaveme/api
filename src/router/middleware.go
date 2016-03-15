package router

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func (r *Router) chain() http.Handler {
	var final http.Handler

	final = r.router
	mw := r.allMiddleware()
	for i := len(mw) - 1; i >= 0; i-- {
		final = mw[i](final)
	}

	return final
}

func (r *Router) allMiddleware() []Middleware {
	return r.middleware
}

func (r *Router) Use(middleware ...Middleware) {

	r.middleware = append(r.middleware, middleware...)
}
