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
		URL:               "http://concept-url0.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "source created from rabbit",
	}
	mockRepository := MockDataSourceRepository{nil, nil}
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
		URL:               "http://concept-url1.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "source created from rabbit",
	}
	mockRepository := MockDataSourceRepository{&toBeCreated, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	body, _ := json.Marshal(toBeCreated)

	err := mockService.CreateDataSourceFromRabbitMessage(context.Background(), body)
	assert.NotNil(t, err)
}

func TestCreateFromRabbitAlreadyExists(t *testing.T) {
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://concept-url0.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "source created from rabbit",
	}
	mockRepository := MockDataSourceRepository{&model.DataSource{
		ID:                "id-of-existing-source",
		DataSourceType:    toBeCreated.DataSourceType,
		DataType:          toBeCreated.DataType,
		URL:               toBeCreated.URL,
		AcceptHeaderValue: toBeCreated.AcceptHeaderValue,
		PublisherID:       toBeCreated.PublisherID,
		Description:       toBeCreated.Description,
	}, nil}
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
		URL:               "http://concept-url2.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "source created from rabbit",
	}
	mockRepository := MockDataSourceRepository{nil, errors.New("repo error")}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	body, _ := json.Marshal(toBeCreated)

	err := mockService.CreateDataSourceFromRabbitMessage(context.Background(), body)
	assert.NotNil(t, err)
}
