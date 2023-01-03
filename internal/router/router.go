package router

import (
	"awesome-go-project/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Router() http.Handler {

	r := chi.NewRouter()

	r.Get("/test-http", handlers.TestHttp)
	r.Get("/upload", handlers.GetUpload)
	r.Post("/upload", handlers.PostUpload)

	return r
}
