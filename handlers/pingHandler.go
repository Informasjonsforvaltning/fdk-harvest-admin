package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
