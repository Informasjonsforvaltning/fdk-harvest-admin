package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/stretchr/testify/assert"
)

func TestGetDatasetStatus(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources/test-id/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualResponse model.HarvestStatuses
	_ = json.Unmarshal(w.Body.Bytes(), &actualResponse)

	end := "2022-04-06 14:00:47 +0200"

	var expectedStatuses []model.HarvestStatus
	expectedStatuses = append(expectedStatuses, model.HarvestStatus{
		HarvestType:  "dataset",
		Status:       "done",
		ErrorMessage: nil,
		StartTime:    "2022-04-06 14:00:07 +0200",
		EndTime:      &end,
	})

	expected := model.HarvestStatuses{
		Id:       "test-id",
		Statuses: expectedStatuses,
	}

	assert.Equal(t, expected, actualResponse)
}

func TestGetHarvestStatuses(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources/test-id-2/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualResponse model.HarvestStatuses
	_ = json.Unmarshal(w.Body.Bytes(), &actualResponse)

	eventEnd := "2022-04-06 14:00:37 +0200"
	publicServiceEnd := "2022-04-06 14:01:17 +0200"

	expectedServiceStatus := model.HarvestStatus{
		HarvestType:  "publicService",
		Status:       "done",
		ErrorMessage: nil,
		StartTime:    "2022-04-06 14:00:07 +0200",
		EndTime:      &publicServiceEnd,
	}
	expectedEventStatus := model.HarvestStatus{
		HarvestType:  "event",
		Status:       "done",
		ErrorMessage: nil,
		StartTime:    "2022-04-06 14:00:07 +0200",
		EndTime:      &eventEnd,
	}

	assert.Equal(t, "test-id-2", actualResponse.Id)
	assert.Equal(t, 2, len(actualResponse.Statuses))
	assert.Contains(t, actualResponse.Statuses, expectedServiceStatus)
	assert.Contains(t, actualResponse.Statuses, expectedEventStatus)
}

func TestGetFailedHarvestStatuses(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources/test-id-3/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actualResponse model.HarvestStatuses
	_ = json.Unmarshal(w.Body.Bytes(), &actualResponse)

	errorMessage := "error message"

	expectedServiceStatus := model.HarvestStatus{
		HarvestType:  "publicService",
		Status:       "error",
		ErrorMessage: &errorMessage,
		StartTime:    "2022-04-06 14:00:07 +0200",
	}
	expectedEventStatus := model.HarvestStatus{
		HarvestType:  "event",
		Status:       "in-progress",
		ErrorMessage: nil,
		StartTime:    "2022-04-06 15:00:07 +0200",
		EndTime:      nil,
	}

	assert.Equal(t, "test-id-3", actualResponse.Id)
	assert.Equal(t, 2, len(actualResponse.Statuses))
	assert.Contains(t, actualResponse.Statuses, expectedServiceStatus)
	assert.Contains(t, actualResponse.Statuses, expectedEventStatus)
}

func TestHarvestStatusErrorMissingReasoningReport(t *testing.T) {
	router := config.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/organizations/123456789/datasources/data-source-id/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
