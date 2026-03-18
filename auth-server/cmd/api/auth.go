package api

import (

	"log"
	"net/http"
	"time"

	"github.com/cmd/internal/db"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type apiService struct {
	database db.Database
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserPayload struct {
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type ResetPassword struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

var userServerBaseUrl string = "http://localhost:8082"
var jwtKey = "I want to share how to set up JWT authentication in GO. It’s easy to implement, but before we dive into implementation, let’s first understand what JWT authentication is and why we need to implement it in our code."

func (api *apiService) Login(w http.ResponseWriter, r *http.Request) {
	
	var payload LoginPayload

	if err := ReadJson(r, w, &payload); err != nil {
		print(err.Error())
		BadRequestHttpError(w)
		return
	}

	// resp,err := http.Get(userServerBaseUrl+"/v1/"+payload.username)

	user := api.database.GetUser(payload.Username, r.Context())
	
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		println(err.Error())
		UnauthorizedHttpError(w, "Something went wrong")
		return
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		log.Println(err.Error())
		UnauthorizedHttpError(w, "Error generating jwt")
		return
	}

	WriteJson(w, t, http.StatusOK)

}

func (api *apiService) ResetPassword(w http.ResponseWriter, r *http.Request) {

	var resetPassword ResetPassword

	if err := ReadJson(r, w, &resetPassword); err != nil {
		BadRequestHttpError(w)
		return
	}

	user := api.database.GetUser(resetPassword.Username, r.Context())

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(resetPassword.Username))

	if err != nil {
		UnauthorizedHttpError(w, "Unauthorized")
		return
	}

	newPasswordHashed, err := bcrypt.GenerateFromPassword([]byte(resetPassword.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		UnauthorizedHttpError(w, "Error hashing password")
		return
	}

	err = api.database.ResetPassword(user.Username, string(newPasswordHashed), r.Context())

	if err != nil {
		InternalServerErrorHttpError(w)
		return
	}

	s := StandardResponse{
		message: "password changed successfully",
		status:  http.StatusOK,
	}

	WriteJson(w, s, http.StatusOK)

}

func (api *apiService) Register(w http.ResponseWriter, r *http.Request) {

	var registerUserPayload db.RegUser

	if err := ReadJson(r, w, &registerUserPayload); err != nil {
		BadRequestHttpError(w)
		return
	}

	err := api.database.Register(registerUserPayload, r.Context())

	if err != nil {
		println(err.Error())
		InternalServerErrorHttpError(w)
		return
	}

	s := StandardResponse{
		status:  200,
		message: "registered successfully",
	}

	WriteJson(w, &s, http.StatusOK)

}
