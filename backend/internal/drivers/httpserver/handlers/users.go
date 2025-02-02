package handlers

import (
	"backend/internal/application/usecases/createsession"
	"backend/internal/application/usecases/createuser"
	"backend/internal/drivers/httpserver/handlers/internal"
	"backend/internal/persistence/repositories"
	"net/http"

	"github.com/go-chi/render"
)

var userRepository = repositories.PostgresUserRepository{}

func HandleUserCreation(w http.ResponseWriter, r *http.Request) {
	body, ok := internal.ReadRequestBody[createuser.Payload](w, r)

	if !ok {
		return
	}

	result := createuser.Run(createuser.Params{
		UserRepository: userRepository,
		Payload:        body,
	})

	if result.IsLeft {
		render.Status(r, result.LeftResponse.StatusCode)
		render.JSON(w, r, result.LeftResponse)
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, result.RightResponse)
}

func HandleSessionCreation(w http.ResponseWriter, r *http.Request) {
	body, ok := internal.ReadRequestBody[createsession.Payload](w, r)

	if !ok {
		return
	}

	result := createsession.Run(createsession.Params{
		UserRepository: userRepository,
		Payload:        body,
	})

	if result.IsLeft {
		render.Status(r, result.LeftResponse.StatusCode)
		render.JSON(w, r, result.LeftResponse)
		return
	}

	render.Status(r, 200)
	render.JSON(w, r, result.RightResponse)
}
