package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

func TestGetDataSourcesRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/datasources", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualResponse []model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.True(t, len(actualResponse) > 2)
}

func TestGetDataSourcesByDataSourceType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/datasources?dataSourceType=DCAT-AP-NO", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse []model.DataSource
	expectedResponse = append(expectedResponse, model.DataSource{
		ID:                "test-id",
		DataSourceType:    "DCAT-AP-NO",
		DataType:          "dataset",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "123456789",
		Description:       "test source",
		AuthHeader: &model.AuthHeader{
			Name:  "X-API-KEY",
			Value: "MyApiKey",
		},
	})

	var actualResponse []model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetDataSourcesByDataType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/datasources?dataType=publicService", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse []model.DataSource
	expectedResponse = append(expectedResponse, model.DataSource{
		ID:                "test-id-2",
		DataSourceType:    "CPSV-AP-NO",
		DataType:          "publicService",
		URL:               "http://url2.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "111222333",
		Description:       "test source 2",
	})
	expectedResponse = append(expectedResponse, model.DataSource{
		ID:                "test-id-3",
		DataSourceType:    "CPSV-AP-NO",
		DataType:          "publicService",
		URL:               "http://url3.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "123456789",
		Description:       "test source 3",
	})

	var actualResponse []model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
