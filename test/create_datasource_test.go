package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
	}
	body, _ := json.Marshal(toBeCreated)
	req, _ := http.NewRequest("POST", "/organizations/987654321/datasources", bytes.NewReader(body))
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
