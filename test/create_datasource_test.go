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

func TestCreateDataSource(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "987654321",
		Description:       "created source",
		AuthHeader:        &model.AuthHeader{
			Name: 		   "X-API-KEY",
			Value: 		   "MyAPIKey",
		},
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	post := httptest.NewRecorder()
	postReq, _ := http.NewRequest("GET", w.HeaderMap.Get("Location"), nil)
	router.ServeHTTP(post, postReq)

	assert.Equal(t, http.StatusOK, post.Code)

	var actualResponse model.DataSource
	json.Unmarshal(post.Body.Bytes(), &actualResponse)

	toBeCreated.Id = actualResponse.Id
	assert.Equal(t, toBeCreated, actualResponse)
}

func TestInvalidDataSourceType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "BAD-REQUEST",
		DataType:          "concept",
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "987654321",
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
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "987654321",
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
		Url:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherId:       "123456789",
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
