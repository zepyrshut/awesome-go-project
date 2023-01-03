package router

import (
	"awesome-go-project/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Router a simple router that uses chi.
func Router() http.Handler {

	r := chi.NewRouter()

	r.Get("/upload", handlers.GetUpload)
	r.Post("/upload", handlers.PostUpload)

	return r
}
