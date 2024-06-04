package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Nerzal/gocloak/v10"
	"net/http"
	"strings"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/rabbit"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataSourceService struct {
	DataSourceRepository repository.DataSourceRepository
	ReportsRepository    repository.ReportsRepository
	Publisher            rabbit.Publisher
}

func InitService() *DataSourceService {
	service := DataSourceService{
		DataSourceRepository: repository.InitDataSourceRepository(),
		ReportsRepository:    repository.InitReportsRepository(),
		Publisher:            &rabbit.PublisherImpl{},
	}
	return &service
}

func (service *DataSourceService) GetAllowedDataSources(ctx context.Context, token string, dataType string, dataSourceType string) (*[]model.DataSource, int) {
	client := gocloak.NewClient(env.KeycloakHost())
	_, claims, _ := client.DecodeAccessToken(ctx, token, "fdk")
	authorities := (*claims)["authorities"].(string)

	if HasSystemAdminRole(authorities) {
		return service.GetDataSources(ctx, nil, dataType, dataSourceType)
	} else {
		return service.GetDataSources(ctx, AllAuthorizedOrgs(authorities), dataType, dataSourceType)
	}
}

func (service *DataSourceService) GetDataSources(ctx context.Context, orgs []string, dataType string, dataSourceType string) (*[]model.DataSource, int) {
	query := bson.D{}
	if orgs != nil {
		query = append(query, bson.E{Key: "publisherId", Value: bson.D{{Key: "$in", Value: orgs}}})
	}
	if len(dataType) > 0 {
		query = append(query, bson.E{Key: "dataType", Value: dataType})
	}
	if len(dataSourceType) > 0 {
		query = append(query, bson.E{Key: "dataSourceType", Value: dataSourceType})
	}
	dataSources, err := service.DataSourceRepository.GetDataSources(ctx, query)
	if err != nil {
		logrus.Error("Get data sources failed ")
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	}

	if dataSources == nil {
		dataSources = []model.DataSource{}
	}
	return &dataSources, http.StatusOK
}

func (service *DataSourceService) GetDataSource(ctx context.Context, id string) (*model.DataSource, int) {
	dataSource, err := service.DataSourceRepository.GetDataSource(ctx, id)
	if err != nil {
		logrus.Errorf("Get data source with id %s failed, ", id)
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	} else if dataSource == nil {
		return nil, http.StatusNotFound
	} else {
		return dataSource, http.StatusOK
	}
}

func (service *DataSourceService) DeleteDataSource(ctx context.Context, id string) int {
	err := service.DataSourceRepository.DeleteDataSource(ctx, id)
	if err == nil {
		return http.StatusNoContent
	} else if err == mongo.ErrNoDocuments {
		return http.StatusNotFound
	} else {
		logrus.Error("Delete data source with id %s failed, ", id)
		logging.LogAndPrintError(err)
		return http.StatusInternalServerError
	}
}

func (service *DataSourceService) CreateDataSource(ctx context.Context, bytes []byte, org string) (*model.DataSource, *string, *string, int) {
	dataSource, err := unmarshalAndValidateDataSource(bytes)
	var msg string
	if err != nil {
		logging.LogAndPrintError(err)
		msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return nil, &msg, nil, http.StatusBadRequest
	}

	if org != dataSource.PublisherID {
		logging.LogAndPrintError(errors.New("Create failed, trying to create data source for other organization"))
		msg = "Bad Request - trying to create data source for other organization"
		return nil, &msg, nil, http.StatusBadRequest
	}

	dataSource.ID = uuid.New().String()
	err = service.DataSourceRepository.CreateDataSource(ctx, *dataSource)
	if err != nil {
		logrus.Error("Create failed")
		logging.LogAndPrintError(err)
		return nil, nil, nil, http.StatusInternalServerError
	} else {
		location := fmt.Sprintf("/%s/%s/%s/%s", env.PathValues.Organizations, org, env.PathValues.Datasources, dataSource.ID)
		return dataSource, nil, &location, http.StatusCreated
	}
}

func (service *DataSourceService) CreateDataSourceFromRabbitMessage(ctx context.Context, bytes []byte) error {
	dataSource, err := unmarshalAndValidateDataSource(bytes)
	if err != nil {
		return err
	} else {
		dataSource.ID = uuid.New().String()
		err = service.DataSourceRepository.CreateDataSource(ctx, *dataSource)
		if err != nil {
			logrus.Error("Create failed")
			return err
		} else {
			logrus.Info("Data source created from rabbit message")
			return nil
		}
	}
}

func (service *DataSourceService) UpdateDataSource(ctx context.Context, id string, bytes []byte, org string) (*model.DataSource, *string, int) {
	toUpdate, err := unmarshalAndValidateDataSource(bytes)
	var msg string
	if err != nil {
		logging.LogAndPrintError(err)
		msg = fmt.Sprintf("Bad Request - %s", err.Error())
		return nil, &msg, http.StatusBadRequest
	}

	var dbSource *model.DataSource
	dbSource, err = service.DataSourceRepository.GetDataSource(ctx, id)
	if err != nil {
		logrus.Errorf("Data source with id %s failed, ", id)
		logging.LogAndPrintError(err)
		return nil, nil, http.StatusInternalServerError
	} else if dbSource == nil {
		return nil, nil, http.StatusNotFound
	}

	if org != dbSource.PublisherID {
		logging.LogAndPrintError(errors.New("Update failed, trying to update data source for other organization"))
		msg = "Bad Request - trying to update data source for other organization"
		return nil, &msg, http.StatusBadRequest
	}

	toUpdate.ID = dbSource.ID
	err = service.DataSourceRepository.UpdateDataSource(ctx, *toUpdate)

	var updated *model.DataSource
	updated, err = service.DataSourceRepository.GetDataSource(ctx, id)
	if err != nil {
		logrus.Error("Update failed")
		logging.LogAndPrintError(err)
		return nil, nil, http.StatusInternalServerError
	} else {
		return updated, nil, http.StatusOK
	}
}

func (service *DataSourceService) StartHarvesting(ctx context.Context, id string, org string) int {
	dataSource, err := service.DataSourceRepository.GetDataSource(ctx, id)
	if err != nil {
		logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
		logging.LogAndPrintError(err)
		return http.StatusInternalServerError
	} else if dataSource == nil {
		return http.StatusNotFound
	} else if dataSource.PublisherID != org {
		return http.StatusBadRequest
	} else {
		var msgKey *string
		msgKey, err = dataTypeToMessageKey(dataSource.DataType)
		if err != nil {
			logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
			logging.LogAndPrintError(err)
			return http.StatusInternalServerError
		}

		harvestParams := make(map[string]string)
		harvestParams["dataSourceId"] = dataSource.ID
		harvestParams["publisherId"] = dataSource.PublisherID
		harvestParams["forceUpdate"] = "true"

		var msgBody []byte
		msgBody, err = json.Marshal(harvestParams)
		if err != nil {
			logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
			logging.LogAndPrintError(err)
			return http.StatusInternalServerError
		}

		err = service.Publisher.Publish(*msgKey, msgBody)
		if err != nil {
			logrus.Errorf("Unable to trigger harvest of data source with id %s", id)
			logging.LogAndPrintError(err)
			return http.StatusInternalServerError
		} else {
			logrus.Infof("Harvest triggered for %s in org %s with type %s", id, org, dataSource.DataType)
			service.setHarvestStartTimeAsNow(ctx, id)
			return http.StatusNoContent
		}
	}
}

func (service *DataSourceService) setHarvestStartTimeAsNow(ctx context.Context, id string) {
	savedReports, err := service.ReportsRepository.GetReports(ctx, id)
	if err == nil && savedReports != nil {
		now := nowDateString()

		for _, report := range savedReports.Reports {
			service.ReportsRepository.UpsertReports(ctx, model.HarvestReport{
				ID:               report.ID,
				URL:              report.URL,
				DataType:         report.DataType,
				HarvestError:     false,
				StartTime:        now,
				EndTime:          now,
				ErrorMessage:     nil,
				ChangedCatalogs:  report.ChangedCatalogs,
				ChangedResources: report.ChangedResources,
			})
		}
	}
}

func (service *DataSourceService) GetHarvestStatus(ctx context.Context, id string) (*model.HarvestStatuses, int) {
	harvestReports, err := service.ReportsRepository.GetReports(ctx, id)
	if err != nil {
		logrus.Errorf("Get harvest reports for id %s failed", id)
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	} else if harvestReports == nil {
		return nil, http.StatusNotFound
	}

	reasoningReports, err := service.ReportsRepository.GetReports(ctx, ReasoningReportID(id))
	if err != nil {
		logrus.Error("Get reasoning reports failed")
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	} else if reasoningReports == nil {
		reasoningReports = &model.HarvestReports{}
	}

	ingestReports, err := service.ReportsRepository.GetReports(ctx, "ingested")
	if err != nil {
		logrus.Error("Get ingest reports failed")
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	}

	statuses, err := calculateHarvestStatusesFromReports(*harvestReports, *reasoningReports, *ingestReports)
	if err != nil {
		logrus.Error("Harvest status calculation failed")
		logging.LogAndPrintError(err)
		return nil, http.StatusInternalServerError
	}

	return statuses, http.StatusOK
}

func (service *DataSourceService) ConsumeReport(ctx context.Context, routingKey string, body []byte) []error {
	switch {
	case strings.Contains(routingKey, "harvested"):
		return service.consumeHarvestedReport(ctx, body)
	case strings.Contains(routingKey, "reasoned"):
		return service.consumeReasonedReport(ctx, body)
	case strings.Contains(routingKey, "ingested"):
		return service.consumeIngestedReport(ctx, routingKey, body)
	default:
		return nil
	}
}

func (service *DataSourceService) consumeHarvestedReport(ctx context.Context, body []byte) []error {
	var errors []error
	var reports []model.HarvestReport
	err := json.Unmarshal(body, &reports)
	if err != nil {
		return []error{err}
	}

	for _, report := range reports {
		err := service.ReportsRepository.UpsertReports(ctx, report)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func (service *DataSourceService) consumeReasonedReport(ctx context.Context, body []byte) []error {
	var reports []model.HarvestReport
	err := json.Unmarshal(body, &reports)
	if err != nil {
		return []error{err}
	}

	var errors []error
	for _, report := range reports {
		report.ID = ReasoningReportID(report.ID)
		err := service.ReportsRepository.UpsertReports(ctx, report)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func (service *DataSourceService) consumeIngestedReport(ctx context.Context, routingKey string, body []byte) []error {
	var startAndEnd model.StartAndEndTime
	err := json.Unmarshal(body, &startAndEnd)
	if err != nil {
		return []error{err}
	}

	report, err := IngestedReport(routingKey, startAndEnd)
	if err != nil {
		return []error{err}
	}

	err = service.ReportsRepository.UpsertReports(ctx, *report)
	if err != nil {
		return []error{err}
	}
	return nil
}
