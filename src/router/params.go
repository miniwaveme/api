package router

import "github.com/julienschmidt/httprouter"

type Params map[string]string

func paramsFromHTTPRouter(hrps httprouter.Params) Params {
	var ps = Params{}

	for _, p := range hrps {
		ps[p.Key] = p.Value
	}

	return ps
}
