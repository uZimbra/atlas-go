package process

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/korg/atlas/internal/infra/http/routes"
)

func SeverBootstrap() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Mount("/user", routes.UserRouter())

	fmt.Println("Server is starting...")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Server fail to start")
	}
}
