package api

import (
	"net/http"
	"time"
)

type apiService struct{
	
}

type LoginPayload struct{
	username string
	password string
}


var userServerBaseUrl string = "http://localhost:8082"


func (api *apiService) Login(w http.ResponseWriter, r *http.Request){
	
	var payload LoginPayload
 	if err := ReadJson(r,w,&payload); err !=  nil{
		BadRequestHttpError(w)
		return
  	}
    
  // resp,err := http.Get(userServerBaseUrl+"/v1/"+payload.username)
  
  
  
  if err != nil{
 	UnauthorizedHttpError(w,"Something went wrong")
  	return
  }
  
    
	
}

func (api *apiService) ResetPassword(w http.ResponseWriter, r *http.Request){
	
	
	
}

func (api *apiService) Register(w http.ResponseWriter, r *http.Request){
	
	
	
}