package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/internal/config"
	"github.com/internal/db"
	"github.com/internal/env"
)

func initTestDb() ApiService {

	//you cab use different config based on test
	dbConfig := config.DatabaseConfig{
		DbType:       env.GetString("DB_TYPE", "postgres"),
		Addr:         env.GetString("DB_ADDR", "postgres://postgres:12345@localhost/shortlinkuser?sslmode=disable"),
		MaxOpenConn:  env.GetInt("MAX_OPEN_CONN", 20),
		MaxIdealConn: env.GetInt("MAX_IDEA_CONN", 20),
		MaxIdealTime: env.GetString("MAX_IDEAL_TIME", "15m"),
	}

	database, err := db.ConnectDb(dbConfig)

	if err != nil {
		panic(err)
	}

	// log.Print("User-Server  test database conncetion established")

	databseRepo := db.NewDatabaseRepo(database)

	return ApiService{
		db: databseRepo,
	}
}

func TestCheckUserExist(t *testing.T) {

	apiServiceTest := initTestDb()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/check-user-exist/{username}", apiServiceTest.checkUserExist)

	server := httptest.NewServer(mux)
	defer server.Close()

	baseUrl := server.URL + "/v1/check-user-exist/ifeanyi25"

	resp, err := http.Get(baseUrl)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 but got %d ", resp.StatusCode)
	}

}

func TestRegister(t *testing.T) {

	apiServiceTest := initTestDb()

	server := httptest.NewServer(http.HandlerFunc(apiServiceTest.Register))
	defer server.Close()

	body := RegisterUserPayload{
		DisplayName: "ifeanyi4",
		Username:    "ifeanyi25",
		Password:    "12345",
	}

	byteBody, err := json.Marshal(body)

	if err != nil {
		t.Error(err)
	}

	bodyReader := bytes.NewBuffer(byteBody)

	resp, err := http.Post(server.URL, "application/json", bodyReader)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 but got %d ", resp.StatusCode)
	}

}

func TestGetUserUsername(t *testing.T) {

	apiServiceTest := initTestDb()

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/{username}", apiServiceTest.GetUserByUsername)

	server := httptest.NewServer(mux)
	defer server.Close()

	urlm := server.URL + "/v1/ifeanyi25"

	resp, err := http.Get(urlm)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 but got %d ", resp.StatusCode)
	}

}

func TestGetUserId(t *testing.T) {

	apiServiceTest := initTestDb()
	var id = "/v1/get-with-id/50392ebb-8ac1-4cce-9ce2-e4f9d84f9f2c"

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/get-with-id/{id}", apiServiceTest.GetUser)

	server := httptest.NewServer(mux)
	defer server.Close()

	var urlm = server.URL + id

	resp, err := http.Get(urlm)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 but got %d ", resp.StatusCode)
	}

}

func TestLogin(t *testing.T) {

	api := initTestDb()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/auth/login", api.Login)

	server := httptest.NewServer(mux)
	defer server.Close()

	urlm := server.URL + "/v1/auth/login"

	loginPayload := LoginPayload{
		Username: "ifeanyi25",
		Password: "12345",
	}

	bytesPayload, err := json.Marshal(loginPayload)

	if err != nil {
		t.Error(err)
	}

	resp, err := http.Post(urlm, "application/json", bytes.NewBuffer(bytesPayload))

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 but got %d ", resp.StatusCode)
	}

}

func TestResetPassword(t *testing.T) {

	api := initTestDb()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/auth/resetPassword", api.ResetPassword)

	server := httptest.NewServer(mux)
	defer server.Close()

	resetPayload := RestPasswordPayload{
		Username:    "ifeanyi25",
		Password:    "12345",
		NewPassword: "54321",
	}

	resetPayloadBytes, err := json.Marshal(resetPayload)

	if err != nil {
		t.Error(err)
	}

	resp, err := http.Post(server.URL+"/v1/auth/resetPassword", "application/json", bytes.NewBuffer(resetPayloadBytes))

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 but got %d ", resp.StatusCode)
	}

}
