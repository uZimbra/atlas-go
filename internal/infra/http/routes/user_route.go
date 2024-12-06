package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/korg/atlas/internal/infra/http/handlers"
)

func UserRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/", handlers.CreateUserHandler)

	return r
}
