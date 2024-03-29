package connection

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/sirupsen/logrus"
)

func mongoConnectionString() string {
	authParams := env.ConstantValues.MongoAuthParams
	dbName := env.ConstantValues.MongoDatabase
	host := env.MongoHost()
	password := env.MongoPassword()
	user := env.MongoUsername()

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s/%s?%s", user, password, host, dbName, authParams)

	return connectionString
}

func DataSourcesCollection() *mongo.Collection {
	mongoOptions := options.Client().ApplyURI(mongoConnectionString())
	client, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		logrus.Error("mongo client failed")
		logging.LogAndPrintError(err)
	}
	collection := client.Database(env.ConstantValues.MongoDatabase).Collection(env.ConstantValues.DataSourcesCollection)

	return collection
}

func ReportsCollection() *mongo.Collection {
	mongoOptions := options.Client().ApplyURI(mongoConnectionString())
	client, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		logrus.Error("mongo client failed")
		logging.LogAndPrintError(err)
	}
	collection := client.Database(env.ConstantValues.MongoDatabase).Collection(env.ConstantValues.ReportsCollection)

	return collection
}
