package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"
	"github.com/stretchr/testify/assert"
)

func TestStartHarvestingNotFound(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	orgAdminAuth := OrgAdminAuth("123456789")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	req, _ := http.NewRequest("POST", "/organizations/123456789/datasources/invalid-id/start-harvesting", nil)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestStartHarvestingWrongOrg(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources/test-id/start-harvesting", nil)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestStartHarvestingConnctionError(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	orgAdminAuth := OrgAdminAuth("123456789")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	req, _ := http.NewRequest("POST", "/organizations/123456789/datasources/test-id/start-harvesting", nil)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestStartHarvestingConcepts(t *testing.T) {
	mockDataSource := model.DataSource{
		ID:                "concept-source",
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "concepts source",
	}
	mockRepository := MockDataSourceRepository{&mockDataSource, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	status := mockService.StartHarvesting(context.TODO(), "concept-source", "123456789")

	assert.Equal(t, http.StatusNoContent, status)
}

func TestStartHarvestingDataServices(t *testing.T) {
	mockDataSource := model.DataSource{
		ID:                "dataservice-source",
		DataSourceType:    "DCAT-AP-NO",
		DataType:          "dataservice",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "data services source",
	}
	mockRepository := MockDataSourceRepository{&mockDataSource, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	status := mockService.StartHarvesting(context.TODO(), "dataservice-source", "123456789")

	assert.Equal(t, http.StatusNoContent, status)
}

func TestStartHarvestingDatasets(t *testing.T) {
	mockDataSource := model.DataSource{
		ID:                "dataset-source",
		DataSourceType:    "DCAT-AP-NO",
		DataType:          "dataset",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "datasets source",
	}
	mockRepository := MockDataSourceRepository{&mockDataSource, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	status := mockService.StartHarvesting(context.TODO(), "dataset-source", "123456789")

	assert.Equal(t, http.StatusNoContent, status)
}

func TestStartHarvestingInformationModels(t *testing.T) {
	mockDataSource := model.DataSource{
		ID:                "infomodel-source",
		DataSourceType:    "DCAT-AP-NO",
		DataType:          "informationmodel",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "information model source",
	}
	mockRepository := MockDataSourceRepository{&mockDataSource, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	status := mockService.StartHarvesting(context.TODO(), "infomodel-source", "123456789")

	assert.Equal(t, http.StatusNoContent, status)
}

func TestStartHarvestingPublicServices(t *testing.T) {
	mockDataSource := model.DataSource{
		ID:                "service-source",
		DataSourceType:    "CPSV-AP-NO",
		DataType:          "publicService",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "public services source",
	}
	mockRepository := MockDataSourceRepository{&mockDataSource, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	status := mockService.StartHarvesting(context.TODO(), "service-source", "123456789")

	assert.Equal(t, http.StatusNoContent, status)
}

func TestStartHarvestingInvalidType(t *testing.T) {
	mockDataSource := model.DataSource{
		ID:                "invalid-source",
		DataSourceType:    "CPSV-AP-NO",
		DataType:          "invalid",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "invalid source",
	}
	mockRepository := MockDataSourceRepository{&mockDataSource, nil}
	mockPublisher := MockPublisher{nil}
	mockService := service.DataSourceService{DataSourceRepository: &mockRepository, Publisher: &mockPublisher}
	status := mockService.StartHarvesting(context.TODO(), "invalid-source", "123456789")

	assert.Equal(t, http.StatusInternalServerError, status)
}
