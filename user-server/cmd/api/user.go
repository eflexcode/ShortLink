package api

import (
	"net/http"
)

type CheckUserPayload struct {
	Username string
}

type CheckUserResponsePayload struct {
	Exist bool
}

type RegisterUserPayload struct {
	Display  string
	Username string
	Password string
}

func (api *ApiService) CheckUserExist(w http.ResponseWriter, r *http.Request) {

	var userPayload CheckUserPayload

	ReadJson(w, r, &userPayload)

	ctx := r.Context()

	exist := api.db.CheckuserExist(ctx, userPayload.Username)

	response := CheckUserResponsePayload{
		Exist: exist,
	}

	WriteJson(w, response, http.StatusOK)

}

func (api *ApiService) Register(w http.ResponseWriter, r *http.Request) {

	var registerUserPayload RegisterUserPayload

	ReadJson(w, r, &registerUserPayload)

	ctx := r.Context()

	err := api.db.Insert(ctx,registerUserPayload.Display,registerUserPayload.Username,registerUserPayload.Password)

	if err != nil {
		InsernalServalError(w,"failed to register user")
		return
	}

	standardResponse := StandardResponse{
		status: http.StatusOK,
		message: "user registered succesfully",
	}

	WriteJson(w,standardResponse,http.StatusOK)

}
