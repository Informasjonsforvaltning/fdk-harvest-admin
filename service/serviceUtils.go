package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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

	return nil, errors.New(string(dataType) + " is not a valid data type")
}

func messageKey(messageType string) string {
	return fmt.Sprintf("%s.%s.%s", messageType, env.ConstantValues.RabbitMsgKeyMiddle, env.ConstantValues.RabbitMsgKeyEnd)
}

func harvestTypeFromRoutingKey(routingKeyPrefix string) (*string, error) {
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

	return nil, errors.New(string(routingKeyPrefix) + " is not a valid harvest type")
}

func reasonedOrIngestedReport(routingKey string, startAndEndTime model.StartAndEndTime) (*model.HarvestReport, error) {
	splitKey := strings.Split(routingKey, ".")
	if len(splitKey) != 2 {
		return nil, errors.New(string(routingKey) + " is not a valid routing key")
	}

	harvestType, err := harvestTypeFromRoutingKey(splitKey[0])
	if err != nil {
		return nil, err
	}

	report := model.HarvestReport{
		Id:               splitKey[1],
		Url:              nil,
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
