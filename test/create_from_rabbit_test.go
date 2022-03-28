package test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateFromRabbit(t *testing.T) {
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "987654321",
		Description:       "source created from rabbit",
	}
	mockRepository := MockDataSourceRepository{&toBeCreated, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	body, _ := json.Marshal(toBeCreated)

	err := mockService.CreateDataSourceFromRabbitMessage(context.Background(), body)
	assert.Nil(t, err)
}

func TestDoesNotCreateInvalidFromRabbit(t *testing.T) {
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "invalid",
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "987654321",
		Description:       "source created from rabbit",
	}
	mockRepository := MockDataSourceRepository{&toBeCreated, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	body, _ := json.Marshal(toBeCreated)

	err := mockService.CreateDataSourceFromRabbitMessage(context.Background(), body)
	assert.NotNil(t, err)
}

func TestHandlesRepositoryError(t *testing.T) {
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "987654321",
		Description:       "source created from rabbit",
	}
	mockRepository := MockDataSourceRepository{nil, errors.New("repo error")}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	body, _ := json.Marshal(toBeCreated)

	err := mockService.CreateDataSourceFromRabbitMessage(context.Background(), body)
	assert.NotNil(t, err)
}
