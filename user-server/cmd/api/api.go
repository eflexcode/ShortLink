package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/internal/config"
	"github.com/internal/db"
	"github.com/internal/env"
)

type ApiService struct {
	db *db.DatabaseRepo
}

func Init() {

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

	log.Print("User-Server database conncetion established")

	databseRepo := db.NewDatabaseRepo(database)

	apiService := ApiService{
		db: databseRepo,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {

		s := StandardResponse{
			status:  http.StatusOK,
			message: "UP",
		}

		WriteJson(w, s, http.StatusOK)

	})

	r.Route("/v1", func(r chi.Router) {

		r.Get("/check-user-exist", apiService.CheckUserExist)
		r.Get("/{id}", apiService.GetUser)
		r.Get("/{username}", apiService.GetUserByUsername)
		r.Put("/update/{id}", apiService.Update)

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", apiService.Login)
			r.Post("/resetPassword", apiService.ResetPassword)
			r.Post("/register", apiService.Register)
		})

	})

	http.ListenAndServe(":8082", r)

}
