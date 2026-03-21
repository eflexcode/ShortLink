package api

import (
	"encoding/json"
	"net/http"
)

var maxJsonSize = 1024 * 1024

type StandardResponse struct {
	status  int
	message string
}

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxJsonSize))
	deco := json.NewDecoder(r.Body)
	deco.DisallowUnknownFields()

	err := deco.Decode(data)

	return err
}

func WriteJson(w http.ResponseWriter, data any, status int) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}


