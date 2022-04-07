package test

import (
	"testing"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"
	"github.com/stretchr/testify/assert"
)

func TestHarvestTypeFromRoutingKey(t *testing.T) {
	result, err := service.HarvestTypeFromRoutingKey("concepts")
	expected := model.ConceptHarvestType
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)

	result, err = service.HarvestTypeFromRoutingKey("datasets")
	expected = model.DatasetHarvestType
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)

	result, err = service.HarvestTypeFromRoutingKey("dataservices")
	expected = model.DataServiceHarvestType
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)

	result, err = service.HarvestTypeFromRoutingKey("informationmodels")
	expected = model.InformationModelHarvestType
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)

	result, err = service.HarvestTypeFromRoutingKey("public_services")
	expected = model.PublicServiceHarvestType
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)

	result, err = service.HarvestTypeFromRoutingKey("events")
	expected = model.EventHarvestType
	assert.Nil(t, err)
	assert.Equal(t, &expected, result)

	result, err = service.HarvestTypeFromRoutingKey("invalid")
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestReasonedOrIngestedReportErrors(t *testing.T) {
	validStartAndEndTime := model.StartAndEndTime{
		StartTime: "2022-04-06 15:00:07 +0200",
		EndTime:   "2022-04-06 15:00:17 +0200",
	}

	result, err := service.ReasonedOrIngestedReport("missing_delimiter", validStartAndEndTime)
	assert.NotNil(t, err)
	assert.Nil(t, result)

	result, err = service.ReasonedOrIngestedReport("invalid.values", validStartAndEndTime)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestIsInProgress(t *testing.T) {
	validHarvestTimes := model.HarvestReport{
		StartTime: "2022-04-06 14:00:07 +0100",
		EndTime:   "2022-04-06 14:00:17 +0100",
	}
	result, err := service.IsInProgress(validHarvestTimes, nil, nil)
	assert.True(t, result)
	assert.Nil(t, err)

	validReasoningTimes := model.HarvestReport{
		StartTime: "2022-04-06 15:00:27 +0200",
		EndTime:   "2022-04-06 15:00:37 +0200",
	}
	result, err = service.IsInProgress(validHarvestTimes, &validReasoningTimes, nil)
	assert.True(t, result)
	assert.Nil(t, err)

	validIngestTimes := model.HarvestReport{
		StartTime: "2022-04-06 15:00:47 +0200",
		EndTime:   "2022-04-06 15:00:57 +0200",
	}
	result, err = service.IsInProgress(validHarvestTimes, &validReasoningTimes, &validIngestTimes)
	assert.False(t, result)
	assert.Nil(t, err)

	notUpdatedIngestTimes := model.HarvestReport{
		StartTime: "2022-04-06 14:00:47 +0200",
		EndTime:   "2022-04-06 14:00:57 +0200",
	}
	result, err = service.IsInProgress(validHarvestTimes, &validReasoningTimes, &notUpdatedIngestTimes)
	assert.True(t, result)
	assert.Nil(t, err)

	notUpdatedReasoningTimes := model.HarvestReport{
		StartTime: "2022-04-06 14:00:27 +0200",
		EndTime:   "2022-04-06 14:00:37 +0200",
	}
	result, err = service.IsInProgress(validHarvestTimes, &notUpdatedReasoningTimes, &notUpdatedIngestTimes)
	assert.True(t, result)
	assert.Nil(t, err)

	invalidHarvestEnd := model.HarvestReport{
		StartTime: "2022-04-06 14:00:07 +0100",
		EndTime:   "invalid",
	}
	result, err = service.IsInProgress(invalidHarvestEnd, &validReasoningTimes, &validIngestTimes)
	assert.NotNil(t, err)

	invalidReasoningStart := model.HarvestReport{
		StartTime: "2022-04-06 15-00-27 +0200",
		EndTime:   "2022-04-06 15:00:37 +0200",
	}
	result, err = service.IsInProgress(validHarvestTimes, &invalidReasoningStart, &validIngestTimes)
	assert.NotNil(t, err)

	invalidReasoningEnd := model.HarvestReport{
		StartTime: "2022-04-06 15:00:27 +0200",
		EndTime:   "2022:04:06 15:00:37 +0200",
	}
	result, err = service.IsInProgress(validHarvestTimes, &invalidReasoningEnd, &validIngestTimes)
	assert.NotNil(t, err)

	invalidIngestStart := model.HarvestReport{
		StartTime: "2022-99-99 99:99:99 +0200",
		EndTime:   "2022-04-06 15:00:57 +0200",
	}
	result, err = service.IsInProgress(validHarvestTimes, &validReasoningTimes, &invalidIngestStart)
	assert.NotNil(t, err)
}
