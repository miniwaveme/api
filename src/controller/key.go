package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	KeyController struct{}
)

func NewKeyController() *KeyController {
	return &KeyController{}
}

func (kc KeyController) GeKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (kc KeyController) CreateKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (kc KeyController) UpdateKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (kc KeyController) RemoveKey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
