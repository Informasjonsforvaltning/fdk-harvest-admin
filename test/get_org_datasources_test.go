package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

func TestGetOrgDataSourcesRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources", nil)
	orgAdminAuth := OrgAdminAuth("123456789")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	req.Header.Set("Authorization", *jwt)
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

func TestGetOrgDataSourcesUnauthorized(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetOrgDataSourcesForbidden(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources", nil)
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestGetOrgDataSourcesInternalRoute(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/internal/organizations/123456789/datasources", nil)
	req.Header.Set("X-API-KEY", "test-key")
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

func TestGetOrgDataSourcesInternalForbidden(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/internal/organizations/123456789/datasources", nil)
	req.Header.Set("X-API-KEY", "wrong-key")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestGetOrgDataSourcesByDataSourceType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources?dataSourceType=CPSV-AP-NO", nil)
	orgAdminAuth := OrgAdminAuth("123456789")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse []model.DataSource
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

func TestGetOrgDataSourcesByDataType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources?dataType=publicService", nil)
	orgAdminAuth := OrgAdminAuth("123456789")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var expectedResponse []model.DataSource
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
