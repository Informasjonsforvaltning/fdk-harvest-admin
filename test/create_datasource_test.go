package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

func TestCreateDataSourceWithAdminRole(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://url0.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
		AuthHeader: &model.AuthHeader{
			Name:  "X-API-KEY",
			Value: "MyAPIKey",
		},
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var actualResponse model.DataSource
	json.Unmarshal(w.Body.Bytes(), &actualResponse)

	toBeCreated.ID = actualResponse.ID
	assert.Equal(t, toBeCreated, actualResponse)
}

func TestCreateDataSourceWithWriteRole(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://url1.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
		AuthHeader: &model.AuthHeader{
			Name:  "X-API-KEY",
			Value: "MyAPIKey",
		},
	}
	orgAdminAuth := OrgWriteAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var actualResponse model.DataSource
	json.Unmarshal(w.Body.Bytes(), &actualResponse)

	toBeCreated.ID = actualResponse.ID
	assert.Equal(t, toBeCreated, actualResponse)
}

func TestCreateDataSourceWithSysAdminRole(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://url2.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
		AuthHeader: &model.AuthHeader{
			Name:  "X-API-KEY",
			Value: "MyAPIKey",
		},
	}
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &TestValues.SysAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var actualResponse model.DataSource
	json.Unmarshal(w.Body.Bytes(), &actualResponse)

	toBeCreated.ID = actualResponse.ID
	assert.Equal(t, toBeCreated, actualResponse)
}

func TestCreateAlreadyExist(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "dataset",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
		AuthHeader: &model.AuthHeader{
			Name:  "X-API-KEY",
			Value: "MyAPIKey",
		},
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInvalidDataSourceType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "BAD-REQUEST",
		DataType:          "concept",
		URL:               "http://url3.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInvalidDataType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "invalid",
		URL:               "http://url4.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInvalidJSON(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body := []byte("{''}")
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWrongOrganization(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://url5.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "123456789",
		Description:       "created source",
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
