package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Post("/login", apiService.Login)
		r.Post("/resetPassword", apiService.ResetPassword)
		r.Post("/register", apiService.Register)
	})

	http.ListenAndServe(":8083", r)
}
