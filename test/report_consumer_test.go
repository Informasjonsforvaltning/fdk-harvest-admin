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
		ID:               "data-source-id",
		URL:              &url,
		DataType:         model.ConceptHarvestType,
		HarvestError:     false,
		StartTime:        "2022-04-06 15:00:07 +0200",
		EndTime:          "2022-04-06 15:00:17 +0200",
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FDKIDAndURI{},
		ChangedResources: []model.FDKIDAndURI{},
	}
	reports = append(reports, report)
	body, _ := json.Marshal(reports)

	errors := service.ConsumeReport(context.TODO(), "concepts.harvested", body)

	assert.Nil(t, errors)
}

func TestCreateHarvestedReport(t *testing.T) {
	service := service.InitService()
	url := "http://example.com"

	var reports []model.HarvestReport
	report := model.HarvestReport{
		ID:               "data-source-id-2",
		URL:              &url,
		DataType:         model.ConceptHarvestType,
		HarvestError:     false,
		StartTime:        "2022-04-06 15:00:07 +0200",
		EndTime:          "2022-04-06 15:00:17 +0200",
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FDKIDAndURI{},
		ChangedResources: []model.FDKIDAndURI{},
	}
	reports = append(reports, report)
	body, _ := json.Marshal(reports)

	errors := service.ConsumeReport(context.TODO(), "concepts.harvested", body)

	assert.Nil(t, errors)
}

func TestHarvestedReportError(t *testing.T) {
	service := service.InitService()
	url := "http://example.com"

	report := model.HarvestReport{
		ID:               "data-source-id",
		URL:              &url,
		DataType:         model.ConceptHarvestType,
		HarvestError:     false,
		StartTime:        "2022-04-06 15:00:07 +0200",
		EndTime:          "2022-04-06 15:00:17 +0200",
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FDKIDAndURI{},
		ChangedResources: []model.FDKIDAndURI{},
	}
	body, _ := json.Marshal(report)

	errors := service.ConsumeReport(context.TODO(), "concepts.harvested", body)

	assert.NotEmpty(t, errors)
}

func TestConsumeReasonedReport(t *testing.T) {
	service := service.InitService()
	url := "http://example.com"

	var reasoningReports []model.HarvestReport
	report := model.HarvestReport{
		ID:               "data-source-id-2",
		URL:              &url,
		DataType:         model.ConceptHarvestType,
		HarvestError:     false,
		StartTime:        "2022-04-06 15:00:07 +0200",
		EndTime:          "2022-04-06 15:00:17 +0200",
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FDKIDAndURI{},
		ChangedResources: []model.FDKIDAndURI{},
	}
	reasoningReports = append(reasoningReports, report)
	body, _ := json.Marshal(reasoningReports)

	errors := service.ConsumeReport(context.TODO(), "concepts.reasoned", body)

	assert.Nil(t, errors)
}

func TestReasonedReportError(t *testing.T) {
	service := service.InitService()

	reasoningReport := model.StartAndEndTime{
		StartTime: "2022-04-06 15:00:07 +0200",
		EndTime:   "2022-04-06 15:00:17 +0200",
	}
	body, _ := json.Marshal(reasoningReport)

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

	assert.Nil(t, errors)
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
