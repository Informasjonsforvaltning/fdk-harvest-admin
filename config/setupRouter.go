package config

import (
	"github.com/gin-gonic/gin"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/handlers"
)

func InitializeRoutes(e *gin.Engine) {
	e.SetTrustedProxies(nil)

	e.GET("ping", handlers.PingHandler())
	e.GET("ready", handlers.ReadyHandler())
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	InitializeRoutes(router)
	return router
}
