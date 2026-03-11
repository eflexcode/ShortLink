package api

import (
	"encoding/json"
	"net/http"

)

var maxJsonSize = 1024 * 1024

type StandardResponse struct{
	message string
	status int
}

func ReadJson(r *http.Request, w http.ResponseWriter, data any) error {

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxJsonSize))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)

}

func WriteJson(w http.ResponseWriter, data any,status int ) error{ 
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
	
}

