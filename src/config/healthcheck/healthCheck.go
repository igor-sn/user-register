package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
