package api

import "net/http"

func BadRequestHttpError (w http.ResponseWriter){
	
	s := StandardResponse{
		message: "Bad Request",
	 	status: http.StatusBadRequest,
	}
	
	WriteJson(w,s,http.StatusBadRequest)
	
}

func InternalServerErrorHttpError (w http.ResponseWriter){
	
	s := StandardResponse{
		message: "Internal Server Error",
	 	status: http.StatusInternalServerError,
	}
	
	WriteJson(w,s,http.StatusInternalServerError)
	
}

func UnauthorizedHttpError (w http.ResponseWriter,  message string){
	
	s := StandardResponse{
		message: "Unauthorized "+message,
	 	status: http.StatusUnauthorized,
	}
	
	WriteJson(w,s,http.StatusUnauthorized)
	
}