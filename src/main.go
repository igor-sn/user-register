package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-eden/slf4go"
	"github.com/igor-sn/user-register/src/configs/api"
)

func main() {
	srv := run()
	if err := srv.Run(); err != nil {
		slog.Errorf("error running server", err)
		panic(fmt.Errorf(err.Error()))
	}
}

func run() (svr *gin.Engine) {
	svr = gin.Default()
	routes := api.Routes{}
	routes.GetRoutes(svr)
	slog.Infof("Server started | PORT: 8080")
	return svr
}
