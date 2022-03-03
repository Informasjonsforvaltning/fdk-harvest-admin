package service

import (
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

func (service *DataSourceService) GetAllDataSources() (*[]model.DataSource, error) {
	dataSources, err := service.repository.GetAllDataSources()
	if err != nil {
		return nil, err
	}

	return &dataSources, nil
}
