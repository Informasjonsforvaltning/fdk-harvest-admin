package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/connection"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

type DataSourceRepository struct {
	collection *mongo.Collection
}

var dataSourceRepository *DataSourceRepository

func InitRepository() *DataSourceRepository {
	if dataSourceRepository == nil {
		dataSourceRepository = &DataSourceRepository{collection: connection.MongoCollection()}
	}
	return dataSourceRepository
}

func (r *DataSourceRepository) GetAllDataSources(ctx context.Context) ([]model.DataSource, error) {
	current, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer current.Close(ctx)
	var dataSources []model.DataSource
	for current.Next(ctx) {
		var dataSource model.DataSource
		err := bson.Unmarshal(current.Current, &dataSource)
		if err != nil {
			return nil, err
		}
		dataSources = append(dataSources, dataSource)
	}
	if err := current.Err(); err != nil {
		return nil, err
	}
	return dataSources, nil
}

func (r *DataSourceRepository) GetDataSource(ctx context.Context, id string) (*model.DataSource, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	bytes, err := r.collection.FindOne(ctx, filter).DecodeBytes()

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var dataSource model.DataSource
	unmarshalError := bson.Unmarshal(bytes, &dataSource)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &dataSource, nil
}

func (r *DataSourceRepository) DeleteDataSource(ctx context.Context, id string) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := r.collection.FindOneAndDelete(ctx, filter).DecodeBytes()
	return err
}
