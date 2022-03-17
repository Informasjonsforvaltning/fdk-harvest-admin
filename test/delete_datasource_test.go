package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
)

func TestDeleteDataSourceNotFound(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/organizations/987654321/datasources/not-found", nil)
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &TestValues.SysAdminAuth, &TestValues.Audience)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteDataSource(t *testing.T) {
	router := config.SetupRouter()

	pre := httptest.NewRecorder()
	preReq, _ := http.NewRequest("GET", "/organizations/987654321/datasources/to-be-deleted", nil)
	router.ServeHTTP(pre, preReq)

	assert.Equal(t, http.StatusOK, pre.Code)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/organizations/987654321/datasources/to-be-deleted", nil)
	jwt := CreateMockJwt(time.Now().Add(time.Hour).Unix(), &TestValues.SysAdminAuth, &TestValues.Audience)
	req.Header.Set("Authorization", *jwt)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	post := httptest.NewRecorder()
	postReq, _ := http.NewRequest("GET", "/organizations/987654321/datasources/to-be-deleted", nil)
	router.ServeHTTP(post, postReq)

	assert.Equal(t, http.StatusNotFound, post.Code)
}
