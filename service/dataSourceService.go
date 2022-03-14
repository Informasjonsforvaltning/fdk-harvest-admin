package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSourceService struct {
	repository *repository.DataSourceRepository
}

func InitService() *DataSourceService {
	service := DataSourceService{repository.InitRepository()}
	return &service
}

func (service *DataSourceService) GetDataSources(ctx context.Context, org *string, dataSourceType string) (*[]model.DataSource, int) {
	query := bson.D{}
	if org != nil {
		query = append(query, bson.E{Key: "publisherId", Value: org})
	}
	if len(dataSourceType) > 0 {
		query = append(query, bson.E{Key: "dataSourceType", Value: dataSourceType})
	}
	dataSources, err := service.repository.GetDataSources(ctx, query)
	if err != nil {
		logrus.Error("Get data sources failed ", err)
		return nil, http.StatusInternalServerError
	}

	return &dataSources, http.StatusOK
}

func (service *DataSourceService) GetDataSource(ctx context.Context, id string) (*model.DataSource, int) {
	dataSource, err := service.repository.GetDataSource(ctx, id)
	if err != nil {
		logrus.Errorf("Get data source with id %s failed, ", id, err)
		return nil, http.StatusInternalServerError
	} else if dataSource == nil {
		return nil, http.StatusNotFound
	} else {
		return dataSource, http.StatusOK
	}
}

func (service *DataSourceService) DeleteDataSource(ctx context.Context, id string) int {
	err := service.repository.DeleteDataSource(ctx, id)
	if err == nil {
		return http.StatusOK
	} else if err == mongo.ErrNoDocuments {
		return http.StatusNotFound
	} else {
		logrus.Error("Delete data source with id %s failed, ", id, err)
		return http.StatusInternalServerError
	}
}

func (service *DataSourceService) CreateDataSource(ctx context.Context, bytes []byte, org string) (*string, *string, int) {
	var dataSource model.DataSource
	var msg string
	err := json.Unmarshal(bytes, &dataSource)
	if err != nil {
		logrus.Error("Create failed, ", err)
		msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return &msg, nil, http.StatusBadRequest
	}

	err = dataSource.Validate()
	if err != nil {
		logrus.Error("Create failed, ", err)
		msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return &msg, nil, http.StatusBadRequest
	}
	if org != dataSource.PublisherId {
		logrus.Error("Create failed, wrong org")
		msg = "Bad Request - trying to create data source for other organization"
		return &msg, nil, http.StatusBadRequest
	}

	dataSource.Id = uuid.New().String()
	var createdId *string
	createdId, err = service.repository.CreateDataSource(ctx, dataSource)
	if err != nil {
		logrus.Error("Create failed, ", err)
		return nil, nil, http.StatusInternalServerError
	} else {
		location := fmt.Sprintf("/%s/%s/%s/%s", env.PathValues.Organizations, org, env.PathValues.Datasources, *createdId)
		return nil, &location, http.StatusCreated
	}
}
