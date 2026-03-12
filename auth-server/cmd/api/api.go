package api

import (
	"database/sql"
	"net/http"

	"github.com/cmd/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitApi(s *sql.DB) {
	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	db := db.Database{
		Postgres: s,
	}
	
	apiService :=  apiService{
		database: db,
	}

	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", apiService.Login)
		r.Post("/resetPassword", apiService.ResetPassword)
		r.Post("/register", apiService.Register)
	})

	http.ListenAndServe(":8084", r)
}
