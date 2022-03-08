package service

import (
	"context"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/repository"
)

type DataSourceService struct {
	repository *repository.DataSourceRepository
}

func InitService() *DataSourceService {
	service := DataSourceService{repository.InitRepository()}
	return &service
}

func (service *DataSourceService) GetAllDataSources(ctx context.Context) (*[]model.DataSource, error) {
	dataSources, err := service.repository.GetAllDataSources(ctx)
	if err != nil {
		return nil, err
	}

	return &dataSources, nil
}
