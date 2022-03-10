package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type DataSourceService struct {
	repository *repository.DataSourceRepository
}

func InitService() *DataSourceService {
	service := DataSourceService{repository.InitRepository()}
	return &service
}

func (service *DataSourceService) GetDataSources(ctx context.Context, org *string) (*[]model.DataSource, error) {
	query := bson.D{}
	if org != nil {
		query = append(query, bson.E{Key: "publisherId", Value: org})
	}
	dataSources, err := service.repository.GetDataSources(ctx, query)
	if err != nil {
		return nil, err
	}

	return &dataSources, nil
}

func (service *DataSourceService) GetDataSource(ctx context.Context, id string) (*model.DataSource, error) {
	dataSource, err := service.repository.GetDataSource(ctx, id)
	if err != nil {
		return nil, err
	}

	return dataSource, nil
}

func (service *DataSourceService) DeleteDataSource(ctx context.Context, id string) error {
	return service.repository.DeleteDataSource(ctx, id)
}

func (service *DataSourceService) CreateDataSource(ctx context.Context, bytes []byte, org string) (*string, error) {
	var dataSource model.DataSource
	err := json.Unmarshal(bytes, &dataSource)
	if err != nil {
		logrus.Error("create failed, ", err)
		return nil, errors.New("Bad Request - " + err.Error())
	}
	err = dataSource.Validate()
	if err != nil {
		logrus.Error("create failed, ", err)
		return nil, errors.New("Bad Request - " + err.Error())
	}
	if org != dataSource.PublisherId {
		logrus.Error("create failed, wrong org")
		return nil, errors.New("Bad Request - trying to create data source for other organization")
	}

	dataSource.Id = uuid.New().String()
	return service.repository.CreateDataSource(ctx, dataSource)
}
