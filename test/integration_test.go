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

func TestMain(m *testing.M) {
	MongoContainerRunner(m)
}

func TestPingRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReadyRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ready", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDataSourcesRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/datasources", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse []model.DataSource
	expectedResponse = append(expectedResponse, model.DataSource{
		Id:                "test-id",
		DataSourceType:    "DCAT-AP-NO",
		DataType:          "dataset",
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "test source",
	})

	var actualResponse []model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetDataSourceRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/datasources/test-id", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := model.DataSource{
		Id:                "test-id",
		DataSourceType:    "DCAT-AP-NO",
		DataType:          "dataset",
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "test source",
	}

	var actualResponse model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetDataSourceNotFound(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/datasources/not-found", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
