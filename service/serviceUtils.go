package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

func unmarshalAndValidateDataSource(bytes []byte) (*model.DataSource, error) {
	var dataSource model.DataSource
	err := json.Unmarshal(bytes, &dataSource)
	if err != nil {
		return nil, err
	}

	err = dataSource.Validate()
	if err != nil {
		return nil, err
	}

	return &dataSource, nil
}

func dataTypeToMessageKey(dataType model.DataTypeEnum) (*string, error) {
	switch dataType {
	case model.Concept:
		msgKey := messageKey("concept")
		return &msgKey, nil
	case model.Dataset:
		msgKey := messageKey("dataset")
		return &msgKey, nil
	case model.InformationModel:
		msgKey := messageKey("informationmodel")
		return &msgKey, nil
	case model.DataService:
		msgKey := messageKey("dataservice")
		return &msgKey, nil
	case model.PublicService:
		msgKey := messageKey("publicservice")
		return &msgKey, nil
	}

	return nil, fmt.Errorf("%s is not a valid data type", dataType)
}

func messageKey(messageType string) string {
	return fmt.Sprintf("%s.%s.%s", messageType, env.ConstantValues.RabbitMsgKeyMiddle, env.ConstantValues.RabbitMsgKeyEnd)
}

func HarvestTypeFromRoutingKey(routingKeyPrefix string) (*string, error) {
	switch routingKeyPrefix {
	case "concepts":
		harvestType := model.ConceptHarvestType
		return &harvestType, nil
	case "datasets":
		harvestType := model.DatasetHarvestType
		return &harvestType, nil
	case "dataservices":
		harvestType := model.DataServiceHarvestType
		return &harvestType, nil
	case "informationmodels":
		harvestType := model.InformationModelHarvestType
		return &harvestType, nil
	case "public_services":
		harvestType := model.PublicServiceHarvestType
		return &harvestType, nil
	case "events":
		harvestType := model.EventHarvestType
		return &harvestType, nil
	}

	return nil, fmt.Errorf("%s is not a valid harvest type", routingKeyPrefix)
}

func ReasonedOrIngestedReport(routingKey string, startAndEndTime model.StartAndEndTime) (*model.HarvestReport, error) {
	splitKey := strings.Split(routingKey, ".")
	if len(splitKey) != 2 {
		return nil, fmt.Errorf("%s is not a valid routing key", routingKey)
	}

	harvestType, err := HarvestTypeFromRoutingKey(splitKey[0])
	if err != nil {
		return nil, err
	}

	report := model.HarvestReport{
		ID:               splitKey[1],
		URL:              nil,
		DataType:         model.HarvestTypeEnum(*harvestType),
		HarvestError:     false,
		StartTime:        startAndEndTime.StartTime,
		EndTime:          startAndEndTime.EndTime,
		ErrorMessage:     nil,
		ChangedCatalogs:  []model.FdkIdAndUri{},
		ChangedResources: []model.FdkIdAndUri{},
	}

	return &report, nil
}

func calculateHarvestStatusesFromReports(
	harvestReports model.HarvestReports,
	reasoningReports model.HarvestReports,
	ingestReports model.HarvestReports,
) (*model.HarvestStatuses, error) {
	var statusList []model.HarvestStatus
	for _, report := range harvestReports.Reports {
		reasoningReport, _ := reasoningReports.Reports[string(report.DataType)]
		ingestReport, _ := ingestReports.Reports[string(report.DataType)]
		status, err := harvestStatusFromRelevantReports(report, &reasoningReport, &ingestReport)
		if err != nil {
			return nil, err
		}
		statusList = append(statusList, *status)
	}

	harvestStatuses := model.HarvestStatuses{
		ID:       harvestReports.ID,
		Statuses: statusList,
	}

	return &harvestStatuses, nil
}

func harvestStatusFromRelevantReports(
	harvestReport model.HarvestReport,
	reasoningReport *model.HarvestReport,
	ingestReport *model.HarvestReport,
) (*model.HarvestStatus, error) {
	status := model.HarvestStatus{
		HarvestType: string(harvestReport.DataType),
		StartTime:   harvestReport.StartTime,
	}
	if harvestReport.HarvestError {
		status.Status = model.HarvestError
		status.ErrorMessage = harvestReport.ErrorMessage

		return &status, nil
	}

	harvestIsInProgress, err := IsInProgress(harvestReport, reasoningReport, ingestReport)
	if err != nil {
		return nil, err
	}

	if harvestIsInProgress {
		status.Status = model.HarvestInProgress
	} else {
		status.Status = model.HarvestDone
		status.EndTime = &ingestReport.EndTime
	}

	return &status, nil
}

func IsInProgress(
	harvestReport model.HarvestReport,
	reasoningReport *model.HarvestReport,
	ingestReport *model.HarvestReport,
) (bool, error) {
	if reasoningReport == nil {
		return true, nil
	} else if ingestReport == nil {
		return true, nil
	}

	harvestEnd, err := parseDateTime(harvestReport.EndTime)
	if err != nil {
		return false, err
	}
	reasoningStart, err := parseDateTime(reasoningReport.StartTime)
	if err != nil {
		return false, err
	}
	if harvestEnd.After(reasoningStart) {
		return true, nil
	}

	reasoningEnd, err := parseDateTime(reasoningReport.EndTime)
	if err != nil {
		return false, err
	}
	ingestStart, err := parseDateTime(ingestReport.StartTime)
	if err != nil {
		return false, err
	}
	if reasoningEnd.After(ingestStart) {
		return true, nil
	}

	return false, nil
}

func parseDateTime(dateString string) (time.Time, error) {
	layout := "2006-01-02 15:04:05 -0700"
	return time.Parse(layout, dateString)
}
