package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var GetAllHandler = func() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		logrus.Info("Getting all data sources")

		dataSources, err := service.GetDataSources(c.Request.Context(), nil, c.Query("dataSourceType"))
		if err != nil {
			logrus.Error("Get all data sources failed ", err)
		}

		c.JSON(http.StatusOK, dataSources)
	}
}

var GetOrgDataSourcesHandler = func() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		org := c.Param("org")
		logrus.Infof("Getting data sources for %s", org)

		dataSources, err := service.GetDataSources(c.Request.Context(), &org, c.Query("dataSourceType"))
		if err != nil {
			logrus.Errorf("Get data sources for %s failed ", org, err)
		}

		c.JSON(http.StatusOK, dataSources)
	}
}

var GetDataSourceHandler = func() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		id := c.Param("id")
		logrus.Infof("Getting data source with id %s", id)

		dataSource, err := service.GetDataSource(c.Request.Context(), id)
		if err != nil {
			logrus.Errorf("Get data sources with id %s failed ", id, err)
		}

		if dataSource == nil {
			c.Status(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, dataSource)
		}
	}
}

var DeleteDataSourceHandler = func() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		id := c.Param("id")
		logrus.Infof("Deleting data source with id %s", id)

		err := service.DeleteDataSource(c.Request.Context(), id)

		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusNotFound)
		} else if err != nil {
			logrus.Errorf("Delete data source with id %s failed. ", id, err)
			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
	}
}

var CreateDataSourceHandler = func() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		logrus.Infof("Creating data source")
		bytes, err := c.GetRawData()

		if err != nil {
			logrus.Errorf("Unable to get bytes from request. ", err)
			c.Status(http.StatusBadRequest)
		} else {
			id, err := service.CreateDataSource(c.Request.Context(), bytes, c.Param("org"))
			if err != nil && strings.Contains(err.Error(), "Bad Request") {
				c.JSON(http.StatusBadRequest, err.Error())
			} else if err != nil || id == nil {
				logrus.Errorf("Data source creation failed. ", err)
				c.Status(http.StatusInternalServerError)
			} else {
				location := fmt.Sprintf("/%s/%s/%s/%s", env.PathValues.Organizations, c.Param("org"), env.PathValues.Datasources, *id)
				c.Writer.Header().Add("Location", location)
				c.Status(http.StatusCreated)
			}
		}
	}
}
