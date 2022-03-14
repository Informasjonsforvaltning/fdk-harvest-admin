package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/rabbit"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSourceService struct {
	DataSourceRepository repository.DataSourceRepository
	Publisher            rabbit.Publisher
}

func InitService() *DataSourceService {
	service := DataSourceService{
		DataSourceRepository: repository.InitRepository(),
		Publisher:            &rabbit.PublisherImpl{},
	}
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
	dataSources, err := service.DataSourceRepository.GetDataSources(ctx, query)
	if err != nil {
		logrus.Error("Get data sources failed ")
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	}

	return &dataSources, http.StatusOK
}

func (service *DataSourceService) GetDataSource(ctx context.Context, id string) (*model.DataSource, int) {
	dataSource, err := service.DataSourceRepository.GetDataSource(ctx, id)
	if err != nil {
		logrus.Errorf("Get data source with id %s failed, ", id)
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	} else if dataSource == nil {
		return nil, http.StatusNotFound
	} else {
		return dataSource, http.StatusOK
	}
}

func (service *DataSourceService) DeleteDataSource(ctx context.Context, id string) int {
	err := service.DataSourceRepository.DeleteDataSource(ctx, id)
	if err == nil {
		return http.StatusOK
	} else if err == mongo.ErrNoDocuments {
		return http.StatusNotFound
	} else {
		logrus.Error("Delete data source with id %s failed, ", id)
		logging.LogAndPrintError(err)
		return http.StatusInternalServerError
	}
}

func (service *DataSourceService) CreateDataSource(ctx context.Context, bytes []byte, org string) (*string, *string, int) {
	var dataSource model.DataSource
	var msg string
	err := json.Unmarshal(bytes, &dataSource)
	if err != nil {
		logging.LogAndPrintError(err)
		msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return &msg, nil, http.StatusBadRequest
	}

	err = dataSource.Validate()
	if err != nil {
		logging.LogAndPrintError(err)
		msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return &msg, nil, http.StatusBadRequest
	}
	if org != dataSource.PublisherId {
		logging.LogAndPrintError(errors.New("Create failed, trying to create data source for other organization"))
		msg = "Bad Request - trying to create data source for other organization"
		return &msg, nil, http.StatusBadRequest
	}

	dataSource.Id = uuid.New().String()
	var createdId *string
	createdId, err = service.DataSourceRepository.CreateDataSource(ctx, dataSource)
	if err != nil {
		logrus.Error("Create failed")
		logging.LogAndPrintError(err)
		return nil, nil, http.StatusInternalServerError
	} else {
		location := fmt.Sprintf("/%s/%s/%s/%s", env.PathValues.Organizations, org, env.PathValues.Datasources, *createdId)
		return nil, &location, http.StatusCreated
	}
}

func (service *DataSourceService) StartHarvesting(ctx context.Context, id string, org string) int {
	dataSource, err := service.DataSourceRepository.GetDataSource(ctx, id)
	if err != nil {
		logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
		logging.LogAndPrintError(err)
		return http.StatusInternalServerError
	} else if dataSource == nil {
		return http.StatusNotFound
	} else if dataSource.PublisherId != org {
		return http.StatusBadRequest
	} else {
		var msgKey *string
		msgKey, err = dataTypeToMessageKey(dataSource.DataType)
		if err != nil {
			logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
			logging.LogAndPrintError(err)
			return http.StatusInternalServerError
		}

		harvestParams := make(map[string]string)
		harvestParams["org"] = dataSource.PublisherId

		var msgBody []byte
		msgBody, err = json.Marshal(harvestParams)
		if err != nil {
			logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
			logging.LogAndPrintError(err)
			return http.StatusInternalServerError
		}

		err = service.Publisher.Publish(*msgKey, msgBody)
		if err != nil {
			logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
			logging.LogAndPrintError(err)
			return http.StatusInternalServerError
		} else {
			logrus.Infof("Harvest triggered for %s in org %s with type %s", id, org, dataSource.DataType)
			return http.StatusOK
		}
	}
}

func dataTypeToMessageKey(dataType model.DataTypeEnum) (*string, error) {
	switch dataType {
	case model.Concept:
		msgKey := messageKey("concept")
		return &msgKey, nil
	case model.Dataset:
		msgKey := messageKey("dataset")
		return &msgKey, nil
	case model.InformationModel:
		msgKey := messageKey("informationmodel")
		return &msgKey, nil
	case model.DataService:
		msgKey := messageKey("dataservice")
		return &msgKey, nil
	case model.PublicService:
		msgKey := messageKey("publicservice")
		return &msgKey, nil
	}

	return nil, errors.New(string(dataType) + " is not a valid data type")
}

func messageKey(messageType string) string {
	return fmt.Sprintf("%s.%s.%s", messageType, env.ConstantValues.RabbitMsgKeyMiddle, env.ConstantValues.RabbitMsgKeyEnd)
}
