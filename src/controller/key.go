package controller

import (
	"github.com/miniwaveme/api/src/router"
	"net/http"
)

type (
	KeyController struct{}
)

func NewKeyController() *KeyController {
	return &KeyController{}
}

func (kc KeyController) GetKeyList(w http.ResponseWriter, r *http.Request, p router.Params) {

}

func (kc KeyController) GetKey(w http.ResponseWriter, r *http.Request, p router.Params) {
}

func (kc KeyController) CreateKey(w http.ResponseWriter, r *http.Request, p router.Params) {
}

func (kc KeyController) UpdateKey(w http.ResponseWriter, r *http.Request, p router.Params) {
}

func (kc KeyController) RemoveKey(w http.ResponseWriter, r *http.Request, p router.Params) {
}
