package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/internal/env"
	"golang.org/x/crypto/bcrypt"
)

type CheckUserPayload struct {
	Username string `json:"username"`
}

type CheckUserResponsePayload struct {
	Exist bool `json:"exist"`
}

type RegisterUserPayload struct {
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type UpdateUserPayload struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResultPayload struct {
	Token string `json:"token"`
}

type RestPasswordPayload struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (api *ApiService) checkUserExist(w http.ResponseWriter, r *http.Request) {

	username := r.PathValue("username")

	ctx := r.Context()

	exist := api.db.CheckuserExist(ctx, username)

	response := CheckUserResponsePayload{
		Exist: exist,
	}

	WriteJson(w, response, http.StatusOK)

}

func (api *ApiService) Register(w http.ResponseWriter, r *http.Request) {

	var registerUserPayload RegisterUserPayload

	if err := ReadJson(w, r, &registerUserPayload); err != nil {
		BadRequest(w, "Error reading json payload")
		return
	}

	ctx := r.Context()

	err := api.db.Insert(ctx, registerUserPayload.DisplayName, registerUserPayload.Username, registerUserPayload.Password)

	if err != nil {
		log.Print(err.Error())
		InternalServalError(w, "failed to register user")
		return
	}

	standardResponse := StandardResponse{
		status:  http.StatusOK,
		message: "user registered successfully",
	}

	WriteJson(w, standardResponse, http.StatusOK)

}

func (api *ApiService) GetUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	ctx := r.Context()

	user, err := api.db.GetUser(ctx, id, "id")

	if err != nil {

		log.Print(err.Error())

		if err == sql.ErrNoRows {
			NotFound(w, "no user found with id: "+id)
			return
		}

		InternalServalError(w, "failed to register user")
		return
	}

	WriteJson(w, user, http.StatusOK)

}

func (api *ApiService) GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	username := r.PathValue("username")
	ctx := r.Context()

	user, err := api.db.GetUser(ctx, username, "username")

	if err != nil {
		log.Print(err.Error())

		if err == sql.ErrNoRows {
			NotFound(w, "no user found with username: "+username)
			return
		}

		InternalServalError(w, "something went wrong")
		return
	}

	WriteJson(w, user, http.StatusOK)

}

func (api *ApiService) Update(w http.ResponseWriter, r *http.Request) {

	var payload UpdateUserPayload

	if err := ReadJson(w, r, &payload); err != nil {
		BadRequest(w, "Error reading json payload")
		return
	}

	ctx := r.Context()

	err := api.db.Update(ctx, payload.DisplayName, payload.Id)

	if err != nil {
		InternalServalError(w, "Failed to update user details")
		return
	}

	s := StandardResponse{
		status:  http.StatusOK,
		message: "User details updated.",
	}

	WriteJson(w, s, http.StatusOK)

}

func (api *ApiService) Login(w http.ResponseWriter, r *http.Request) {

	var loginPayload LoginPayload

	if err := ReadJson(w, r, &loginPayload); err != nil {
		BadRequest(w, "Error reading json payload")
		return
	}

	ctx := r.Context()

	user, err := api.db.GetUser(ctx, loginPayload.Username, "username")

	if err != nil {
		log.Print(err.Error())
		UnAuthorized(w, "Login failed")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPayload.Password))

	if err != nil {
		log.Print(err.Error())
		UnAuthorized(w, "Login failed")
		return
	}

	claim := &jwt.MapClaims{
		"user": user.Username,
		"exp":  time.Now().AddDate(0, 2, 0).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	secret := env.GetString("TOKEN_SECRET", "And in the next step, we will create two columns, the first one is an id column, which is our Primary Key (if you don’t know what that is don’t worry, I will explain it in a later tutorial) and the other one is a text column which represents the cat’s name. We can also specify constraints such as “Not Null?” which guarantees that all cells in this column have a value.")

	jwt, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Print(err.Error())
		UnAuthorized(w, "Login failed")
		return
	}

	s := LoginResultPayload{
		Token: jwt,
	}

	WriteJson(w, s, http.StatusOK)

}

func (api *ApiService) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var resetPasswordPayload RestPasswordPayload

	if err := ReadJson(w, r, &resetPasswordPayload); err != nil {
		BadRequest(w, "Error reading json payload")
		return
	}

	ctx := r.Context()

	user, err := api.db.GetUser(ctx,  resetPasswordPayload.Username,"username")

	if err != nil {
		UnAuthorized(w, "Password reset failed")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(resetPasswordPayload.Password))

	if err != nil {
		UnAuthorized(w, "Password reset failed")
		return
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(resetPasswordPayload.NewPassword), bcrypt.DefaultCost)

	err = api.db.UpdatePassword(ctx, string(newHashedPassword), user.Id)

	if err != nil {
		UnAuthorized(w, "Password reset failed")
		return
	}

	s := StandardResponse{
		status:  http.StatusOK,
		message: "Password rest successful",
	}

	WriteJson(w, s, http.StatusOK)

}
