package api

import "net/http"

func BadRequest(w http.ResponseWriter, message string) {

	response := StandardResponse{

		status:  http.StatusBadRequest,
		message: message,
	}

	WriteJson(w, &response, http.StatusBadRequest)

}

func InsernalServalError(w http.ResponseWriter, message string) {

	response := StandardResponse{

		status:  http.StatusInternalServerError,
		message: message,
	}

	WriteJson(w, &response, http.StatusInternalServerError)

}

func NotFound(w http.ResponseWriter,message string){

	response := StandardResponse{

		status:  http.StatusNotFound,
		message: message,
	}

	WriteJson(w, &response, http.StatusInternalServerError)

}