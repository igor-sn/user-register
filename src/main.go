package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-eden/slf4go"
	"github.com/igor-sn/user-register/src/configs"
	hcr "github.com/igor-sn/user-register/src/healthcheck/router"
	"net/http"
)

func main() {
	run()
}

func run() () {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	hcr.AddHealthCheckRoutes(r)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		panic(err)
	}

	slog.Infof("Server started | PORT: 8080")
	return
}
