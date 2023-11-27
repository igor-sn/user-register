package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/igor-sn/user-register/src/healthcheck/handler"
)

func AddHealthCheckRoutes(r *chi.Mux) {
	r.Get("/ping", handler.Pong)
}
