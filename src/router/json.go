package router

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, code int, err string) {
	w.WriteHeader(code)

	WriteJSON(w, map[string]string{
		"error": err,
	})
}

func WriteJSON(w http.ResponseWriter, v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
	return nil
}
