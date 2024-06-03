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

func TestGetDataSourceRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources/test-id", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := model.DataSource{
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
	}

	var actualResponse model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetDataSourceInternalRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/internal/organizations/123456789/datasources/test-id", nil)
	req.Header.Set("X-API-KEY", "test-key")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := model.DataSource{
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
	}

	var actualResponse model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetDataSourceInternalForbidden(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/internal/organizations/123456789/datasources/test-id", nil)
	req.Header.Set("X-API-KEY", "wrong-key")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestGetDataSourceNotFound(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources/not-found", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
