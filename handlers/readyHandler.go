package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadyHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
