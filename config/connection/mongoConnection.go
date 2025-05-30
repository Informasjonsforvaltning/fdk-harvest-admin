package connection

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/sirupsen/logrus"
)

func mongoConnectionString() string {
	var connectionString strings.Builder
	connectionString.WriteString("mongodb://")
	connectionString.WriteString(env.MongoUsername())
	connectionString.WriteString(":")
	connectionString.WriteString(env.MongoPassword())
	connectionString.WriteString("@")
	connectionString.WriteString(env.MongoHost())
	connectionString.WriteString("/")
	connectionString.WriteString(env.ConstantValues.MongoDatabase)
	connectionString.WriteString("?authSource=")
	connectionString.WriteString(env.MongoAuthSource())
	connectionString.WriteString("&replicaSet=")
	connectionString.WriteString(env.MongoReplicaSet())

	return connectionString.String()
}

func MongoClient() *mongo.Client {
	mongoOptions := options.Client().ApplyURI(mongoConnectionString())
	client, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		logrus.Error("mongo client failed")
		logging.LogAndPrintError(err)
	}

	return client
}

func DataSourcesCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(env.ConstantValues.MongoDatabase).Collection(env.ConstantValues.DataSourcesCollection)
}

func ReportsCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(env.ConstantValues.MongoDatabase).Collection(env.ConstantValues.ReportsCollection)
}
