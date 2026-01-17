package api

import "net/http"

func BadRequest(w http.ResponseWriter, message string) {

	response := StandardResponse{

		status:  http.StatusBadRequest,
		message: message,
	}

	WriteJson(w, &response, http.StatusBadRequest)

}

func InternalServalError(w http.ResponseWriter, message string) {

	response := StandardResponse{

		status:  http.StatusInternalServerError,
		message: message,
	}

	WriteJson(w, &response, http.StatusInternalServerError)

}

func NotFound(w http.ResponseWriter, message string) {

	response := StandardResponse{

		status:  http.StatusNotFound,
		message: message,
	}

	WriteJson(w, &response, http.StatusNotFound)

}

func UnAuthorized(w http.ResponseWriter, message string) {

	response := StandardResponse{

		status:  http.StatusUnauthorized,
		message: message,
	}

	WriteJson(w, &response, http.StatusUnauthorized)

}
