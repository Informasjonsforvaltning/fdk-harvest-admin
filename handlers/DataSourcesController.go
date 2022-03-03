package handlers

import (
	"net/http"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var GetAllHandler = func() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		logrus.Info("Getting all data sources")

		dataSources, err := service.GetAllDataSources()
		if err != nil {
			logrus.Error("Get all data sources failed ", err)
		}

		c.JSON(http.StatusOK, dataSources)
	}
}
