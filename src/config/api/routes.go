package api

import (
	"github.com/gin-gonic/gin"
	"github.com/igor-sn/user-register/src/config/healthcheck"
)

type Routes struct {
}

func (p *Routes) GetRoutes(router *gin.Engine) {
	router.GET("/ping", healthcheck.Pong)
}
