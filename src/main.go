package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-eden/slf4go"
	"github.com/igor-sn/user-register/src/configs"
	"github.com/igor-sn/user-register/src/handlers"
	hcr "github.com/igor-sn/user-register/src/healthcheck/router"
	"net/http"
)

func main() {
	run()
}

func run() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/users", handlers.Create)
	r.Get("/users", handlers.List)
	r.Get("/users/{id}", handlers.Get)
	r.Put("/users/{id}", handlers.Update)
	r.Delete("/users/{id}", handlers.Delete)

	hcr.AddHealthCheckRoutes(r)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		panic(err)
	}

	slog.Infof("Server started | PORT: 8080")
	return
}
