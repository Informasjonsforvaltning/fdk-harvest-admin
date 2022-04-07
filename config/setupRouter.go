package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/security"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/handlers"
)

func InitializeRoutes(e *gin.Engine) {
	e.SetTrustedProxies(nil)

	e.GET(env.PathValues.Ping, handlers.PingHandler())
	e.GET(env.PathValues.Ready, handlers.ReadyHandler())
	e.GET(env.PathValues.Datasource, handlers.GetDataSourceHandler())
	e.GET(env.PathValues.HarvestStatus, handlers.GetHarvestStatusHandler())
	e.PUT(env.PathValues.Datasource, security.RequireAdminAuth(), handlers.UpdateDataSourceHandler())
	e.DELETE(env.PathValues.Datasource, security.RequireAdminAuth(), handlers.DeleteDataSourceHandler())
	e.GET(env.PathValues.Datasources, handlers.GetAllHandler())
	e.GET(env.PathValues.OrgDatasources, handlers.GetOrgDataSourcesHandler())
	e.POST(env.PathValues.OrgDatasources, security.RequireAdminAuth(), handlers.CreateDataSourceHandler())
	e.POST(env.PathValues.StartHarvest, security.RequireAdminAuth(), handlers.StartHarvestingHandler())
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	InitializeRoutes(router)
	return router
}

func corsMiddleware() gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}
	corsConfig.AllowAllOrigins = true
	return cors.New(corsConfig)
}
