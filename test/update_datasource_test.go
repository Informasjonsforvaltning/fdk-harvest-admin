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

func TestUpdateDataSource(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeUpdated := model.DataSource{
		ID:                "to-be-updated",
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://updated.com",
		AcceptHeaderValue: "application/rdf+json",
		PublisherID:       "987654321",
		Description:       "updated source",
		AuthHeader: &model.AuthHeader{
			Name:  "X-API-KEY",
			Value: "UpdatedAPIKey",
		},
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeUpdated)
	req, _ := http.NewRequest("PUT", "/organizations/987654321/datasources/to-be-updated", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualResponse model.DataSource
	json.Unmarshal(w.Body.Bytes(), &actualResponse)

	assert.Equal(t, toBeUpdated, actualResponse)
}

func TestUpdateBadRequestWhenCopyingExistingSource(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeUpdated := model.DataSource{
		ID:                "to-be-updated",
		DataSourceType:    "DCAT-AP-NO",
		DataType:          "dataset",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "test source",
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeUpdated)
	req, _ := http.NewRequest("PUT", "/organizations/987654321/datasources/to-be-updated", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateInvalidDataSourceType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "BAD-REQUEST",
		DataType:          "concept",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("PUT", "/organizations/987654321/datasources/to-be-updated", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateInvalidDataType(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "invalid",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
	}
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("PUT", "/organizations/987654321/datasources/to-be-updated", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateInvalidJSON(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	orgAdminAuth := OrgAdminAuth("987654321")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body := []byte("{''}")
	req, _ := http.NewRequest("PUT", "/organizations/987654321/datasources/to-be-updated", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateWrongOrganization(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	toBeCreated := model.DataSource{
		DataSourceType:    "SKOS-AP-NO",
		DataType:          "concept",
		URL:               "http://url.com",
		AcceptHeaderValue: "text/turtle",
		PublisherID:       "987654321",
		Description:       "created source",
	}
	orgAdminAuth := OrgAdminAuth("123456789")
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &orgAdminAuth, &TestValues.Audience)
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("PUT", "/organizations/123456789/datasources/to-be-updated", bytes.NewReader(body))
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
