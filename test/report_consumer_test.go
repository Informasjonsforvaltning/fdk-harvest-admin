package test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/service"
	"github.com/stretchr/testify/assert"
)

func TestUpdateHarvestedReport(t *testing.T) {
	service := service.InitService()
	url := "http://example.com"

	var reports []model.HarvestReport
	report := model.HarvestReport{
		Id:               "data-source-id",
		Url:              &url,
		DataType:         model.ConceptHarvestType,
		HarvestError:     false,
		StartTime:        "2022-04-06 15:00:07 +0200",
		EndTime:          "2022-04-06 15:00:17 +0200",
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FdkIdAndUri{},
		ChangedResources: []model.FdkIdAndUri{},
	}
	reports = append(reports, report)
	body, _ := json.Marshal(reports)

	errors := service.ConsumeReport(context.TODO(), "concepts.harvested", body)

	assert.Empty(t, errors)
}

func TestCreateHarvestedReport(t *testing.T) {
	service := service.InitService()
	url := "http://example.com"

	var reports []model.HarvestReport
	report := model.HarvestReport{
		Id:               "data-source-id-2",
		Url:              &url,
		DataType:         model.ConceptHarvestType,
		HarvestError:     false,
		StartTime:        "2022-04-06 15:00:07 +0200",
		EndTime:          "2022-04-06 15:00:17 +0200",
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FdkIdAndUri{},
		ChangedResources: []model.FdkIdAndUri{},
	}
	reports = append(reports, report)
	body, _ := json.Marshal(reports)

	errors := service.ConsumeReport(context.TODO(), "concepts.harvested", body)

	assert.Empty(t, errors)
}

func TestHarvestedReportError(t *testing.T) {
	service := service.InitService()
	url := "http://example.com"

	report := model.HarvestReport{
		Id:               "data-source-id",
		Url:              &url,
		DataType:         model.ConceptHarvestType,
		HarvestError:     false,
		StartTime:        "2022-04-06 15:00:07 +0200",
		EndTime:          "2022-04-06 15:00:17 +0200",
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FdkIdAndUri{},
		ChangedResources: []model.FdkIdAndUri{},
	}
	body, _ := json.Marshal(report)

	errors := service.ConsumeReport(context.TODO(), "concepts.harvested", body)

	assert.NotEmpty(t, errors)
}

func TestConsumeReasonedReport(t *testing.T) {
	service := service.InitService()

	reasoningReport := model.StartAndEndTime{
		StartTime: "2022-04-06 15:00:07 +0200",
		EndTime:   "2022-04-06 15:00:17 +0200",
	}
	body, _ := json.Marshal(reasoningReport)

	errors := service.ConsumeReport(context.TODO(), "concepts.reasoned", body)

	assert.Empty(t, errors)
}

func TestReasonedReportError(t *testing.T) {
	service := service.InitService()

	var reports []model.StartAndEndTime
	reasoningReport := model.StartAndEndTime{
		StartTime: "2022-04-06 15:00:07 +0200",
		EndTime:   "2022-04-06 15:00:17 +0200",
	}
	reports = append(reports, reasoningReport)
	body, _ := json.Marshal(reports)

	errors := service.ConsumeReport(context.TODO(), "concepts.reasoned", body)

	assert.NotEmpty(t, errors)
}

func TestConsumeIngestReport(t *testing.T) {
	service := service.InitService()

	ingestReport := model.StartAndEndTime{
		StartTime: "2022-04-06 15:00:07 +0200",
		EndTime:   "2022-04-06 15:00:17 +0200",
	}
	body, _ := json.Marshal(ingestReport)

	errors := service.ConsumeReport(context.TODO(), "concepts.ingested", body)

	assert.Empty(t, errors)
}

func TestIngestReportError(t *testing.T) {
	service := service.InitService()

	var reports []model.StartAndEndTime
	ingestReport := model.StartAndEndTime{
		StartTime: "2022-04-06 15:00:07 +0200",
		EndTime:   "2022-04-06 15:00:17 +0200",
	}
	reports = append(reports, ingestReport)
	body, _ := json.Marshal(reports)

	errors := service.ConsumeReport(context.TODO(), "concepts.ingested", body)

	assert.NotEmpty(t, errors)
}
