package handlers

import (
	"net/http"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetAllHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		logrus.Info("Getting all data sources")

		dataSources, status := service.GetDataSources(c.Request.Context(), nil, c.Query("dataType"), c.Query("dataSourceType"))
		if status == http.StatusOK {
			c.JSON(status, dataSources)
		} else {
			c.Status(status)
		}
	}
}

func GetOrgDataSourcesHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		org := c.Param("org")
		logrus.Infof("Getting data sources for %s", org)

		dataSources, status := service.GetDataSources(c.Request.Context(), &org, c.Query("dataType"), c.Query("dataSourceType"))
		if status == http.StatusOK {
			c.JSON(status, dataSources)
		} else {
			c.Status(status)
		}
	}
}

func GetDataSourceHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		id := c.Param("id")
		logrus.Infof("Getting data source with id %s", id)

		dataSource, status := service.GetDataSource(c.Request.Context(), id)
		if status == http.StatusOK {
			c.JSON(status, dataSource)
		} else {
			c.Status(status)
		}
	}
}

func GetHarvestStatusHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		id := c.Param("id")
		logrus.Infof("Getting harvest status for data source with id %s", id)

		harvestStatus, httpStatus := service.GetHarvestStatus(c.Request.Context(), id)
		if httpStatus == http.StatusOK {
			c.JSON(httpStatus, harvestStatus)
		} else {
			c.Status(httpStatus)
		}
	}
}

func DeleteDataSourceHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		id := c.Param("id")
		logrus.Infof("Deleting data source with id %s", id)

		status := service.DeleteDataSource(c.Request.Context(), id)
		c.Status(status)
	}
}

func CreateDataSourceHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		logrus.Infof("Creating data source")
		bytes, err := c.GetRawData()

		if err != nil {
			logrus.Errorf("Unable to get bytes from request.")
			logging.LogAndPrintError(err)
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			dataSource, msg, location, status := service.CreateDataSource(c.Request.Context(), bytes, c.Param("org"))
			if status == http.StatusBadRequest {
				c.JSON(status, msg)
			} else if status == http.StatusCreated {
				c.Writer.Header().Add("Location", *location)
				c.JSON(status, dataSource)
			} else {
				c.Status(status)
			}
		}
	}
}

func UpdateDataSourceHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		id := c.Param("id")
		logrus.Infof("Updating data source with id %s", id)
		bytes, err := c.GetRawData()

		if err != nil {
			logrus.Errorf("Unable to get bytes from request.")
			logging.LogAndPrintError(err)
			c.JSON(http.StatusBadRequest, err.Error())
		} else {
			dataSource, msg, status := service.UpdateDataSource(c.Request.Context(), id, bytes, c.Param("org"))
			if status == http.StatusBadRequest {
				c.JSON(status, msg)
			} else if status == http.StatusOK {
				c.JSON(status, dataSource)
			} else {
				c.Status(status)
			}
		}
	}
}

func StartHarvestingHandler() func(c *gin.Context) {
	service := service.InitService()
	return func(c *gin.Context) {
		status := service.StartHarvesting(c.Request.Context(), c.Param("id"), c.Param("org"))
		c.Status(status)
	}
}
