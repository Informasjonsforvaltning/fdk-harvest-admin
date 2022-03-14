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

func TestGetOrgDataSourcesRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources", nil)
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
		AuthHeader:        &model.AuthHeader{
            Name: "X-API-KEY",
            Value: "MyApiKey",
        },
	})
	expectedResponse = append(expectedResponse, model.DataSource{
		Id:                "test-id-3",
		DataSourceType:    "CPSV-AP-NO",
		DataType:          "publicService",
		Url:               "http://url3.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "test source 3",
	})

	var actualResponse []model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetOrgDataSourcesByDataSourceType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources?dataSourceType=CPSV-AP-NO", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse []model.DataSource
	expectedResponse = append(expectedResponse, model.DataSource{
		Id:                "test-id-3",
		DataSourceType:    "CPSV-AP-NO",
		DataType:          "publicService",
		Url:               "http://url3.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
		Description:       "test source 3",
	})

	var actualResponse []model.DataSource
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
