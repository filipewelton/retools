package httpserver

import (
	"backend/config"
	"backend/internal/drivers/httpserver/middlewares"
	"backend/internal/drivers/httpserver/routers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/go-chi/jwtauth"
)

func Setup() chi.Router {
	r := chi.NewRouter()

	cors := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	tokenSettings := jwtauth.New("HS256", config.Env.JWT_SECRET, nil)

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(middleware.Throttle(60))
	r.Use(httprate.LimitByIP(30, 1*time.Minute))
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(cors)
	r.Use(middleware.StripSlashes)
	r.Use(middlewares.LimitRequestBodySize)
	r.Use(jwtauth.Verifier(tokenSettings))
	r.Use(jwtauth.Authenticator)

	// Routers
	r.Route("/users", routers.SetUsersRoute)

	return r
}

func Run() {
	server := &http.Server{
		Addr:         config.Env.HTTP_SERVER_ADDR,
		Handler:      Setup(),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
