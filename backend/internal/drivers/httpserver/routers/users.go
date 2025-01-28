package routers

import (
	"backend/internal/drivers/httpserver/handlers"

	"github.com/go-chi/chi/v5"
)

func SetUsersRoute(r chi.Router) {
	r.Post("/", handlers.HandleUserCreation)
}
