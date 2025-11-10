package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/connection"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

type DataSourceRepository interface {
	GetDataSources(ctx context.Context, query bson.D) ([]model.DataSource, error)
	GetDataSource(ctx context.Context, id string) (*model.DataSource, error)
	DeleteDataSource(ctx context.Context, id string) error
	CreateDataSource(ctx context.Context, dataSource model.DataSource) error
	UpdateDataSource(ctx context.Context, toUpdate model.DataSource) error
}

type DataSourceRepositoryImpl struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var dataSourceRepository *DataSourceRepositoryImpl

func InitDataSourceRepository() *DataSourceRepositoryImpl {
	if dataSourceRepository == nil {
		client := connection.MongoClient()
		dataSourceRepository = &DataSourceRepositoryImpl{client: client, collection: connection.DataSourcesCollection(client)}
	}
	return dataSourceRepository
}

func (r *DataSourceRepositoryImpl) GetDataSources(ctx context.Context, query bson.D) ([]model.DataSource, error) {
	// Validate query is not nil
	if query == nil {
		return nil, fmt.Errorf("query cannot be nil")
	}

	current, err := r.collection.Find(ctx, query)
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

func (r *DataSourceRepositoryImpl) GetDataSource(ctx context.Context, id string) (*model.DataSource, error) {
	if !isValidID(id) {
		return nil, fmt.Errorf("invalid id format: %s", id)
	}

	filter := bson.D{{Key: "id", Value: id}}
	singleResult := r.collection.FindOne(ctx, filter)
	if err := singleResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	bytes, err := singleResult.Raw()
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

func (r *DataSourceRepositoryImpl) DeleteDataSource(ctx context.Context, id string) error {
	if !isValidID(id) {
		return fmt.Errorf("invalid id format: %s", id)
	}

	return r.client.UseSession(ctx, func(sctx mongo.SessionContext) error {
		err := sctx.StartTransaction(options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.Majority()),
		)

		if err != nil {
			return err
		}

		filter := bson.D{{Key: "id", Value: id}}
		singleResult := r.collection.FindOne(ctx, filter)

		if err = singleResult.Err(); err != nil {
			sctx.AbortTransaction(sctx)
			return err
		}

		_, err = r.collection.DeleteOne(ctx, filter)
		if err != nil {
			sctx.AbortTransaction(sctx)
			return err
		}
		return nil
	})
}

func (r *DataSourceRepositoryImpl) CreateDataSource(ctx context.Context, dataSource model.DataSource) error {
	// Validate data source
	if err := validateDataSource(&dataSource); err != nil {
		return fmt.Errorf("invalid data source: %w", err)
	}

	// Validate PublisherID format
	if !isValidPublisherID(dataSource.PublisherID) {
		return fmt.Errorf("invalid publisherId format: %s", dataSource.PublisherID)
	}

	return r.client.UseSession(ctx, func(sctx mongo.SessionContext) error {
		err := sctx.StartTransaction(options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.Majority()),
		)

		if err != nil {
			return err
		}

		_, err = r.collection.InsertOne(ctx, dataSource, nil)

		if err != nil {
			sctx.AbortTransaction(sctx)
			return err
		}
		return nil
	})
}

func (r *DataSourceRepositoryImpl) UpdateDataSource(ctx context.Context, toUpdate model.DataSource) error {
	// Validate ID format
	if !isValidID(toUpdate.ID) {
		return fmt.Errorf("invalid id format: %s", toUpdate.ID)
	}

	// Validate data source
	if err := validateDataSource(&toUpdate); err != nil {
		return fmt.Errorf("invalid data source: %w", err)
	}

	// Validate PublisherID format
	if !isValidPublisherID(toUpdate.PublisherID) {
		return fmt.Errorf("invalid publisherId format: %s", toUpdate.PublisherID)
	}

	return r.client.UseSession(ctx, func(sctx mongo.SessionContext) error {
		err := sctx.StartTransaction(options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.Majority()),
		)

		if err != nil {
			return err
		}

		filter := bson.D{{Key: "id", Value: toUpdate.ID}}
		result := r.collection.FindOneAndReplace(ctx, filter, toUpdate, nil)
		err = result.Err()

		if err != nil {
			sctx.AbortTransaction(sctx)
			return err
		} else {
			return nil
		}
	})
}
