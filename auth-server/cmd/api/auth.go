package api

import (
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
	username string
	password string
}

type RegisterUserPayload struct {
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

var userServerBaseUrl string = "http://localhost:8082"
var jwtKey = "I want to share how to set up JWT authentication in GO. It’s easy to implement, but before we dive into implementation, let’s first understand what JWT authentication is and why we need to implement it in our code."

func (api *apiService) Login(w http.ResponseWriter, r *http.Request) {

	var payload LoginPayload

	if err := ReadJson(r, w, &payload); err != nil {
		BadRequestHttpError(w)
		return
	}

	// resp,err := http.Get(userServerBaseUrl+"/v1/"+payload.username)

	user := api.database.GetUser(payload.username, r.Context())

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.password))

	if err != nil {
		UnauthorizedHttpError(w, "Something went wrong")
		return
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(jwtKey)

	if err != nil {
		UnauthorizedHttpError(w, "Error generating jwt")
		return
	}

	WriteJson(w, t, http.StatusOK)

}

func (api *apiService) Auth(w http.ResponseWriter, r *http.Request) {
	
}

func (api *apiService) ResetPassword(w http.ResponseWriter, r *http.Request) {

	
	
}

func (api *apiService) Register(w http.ResponseWriter, r *http.Request) {

	var registerUserPayload db.RegUser

	if err := ReadJson(r, w, &registerUserPayload); err != nil {
		BadRequestHttpError(w)
		return
	}

	err := api.database.Register(registerUserPayload, r.Context())

	if err != nil {
		InternalServerErrorHttpError(w)
		return
	}

	s := StandardResponse{
		status:  200,
		message: "registeed succefully",
	}

	WriteJson(w, s, http.StatusOK)

}
