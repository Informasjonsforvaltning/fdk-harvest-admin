package config

import (
	"github.com/gin-gonic/gin"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/handlers"
)

func InitializeRoutes(e *gin.Engine) {
	e.SetTrustedProxies(nil)

	e.GET(env.PathValues.Ping, handlers.PingHandler())
	e.GET(env.PathValues.Ready, handlers.ReadyHandler())
	e.GET(env.PathValues.Datasource, handlers.GetDataSourceHandler())
	e.DELETE(env.PathValues.Datasource, handlers.DeleteDataSourceHandler())
	e.GET(env.PathValues.Datasources, handlers.GetAllHandler())
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	InitializeRoutes(router)
	return router
}
