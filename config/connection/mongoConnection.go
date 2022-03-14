package connection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/config/env"
	"github.com/Informasjonsforvaltning/fdk-harvest-admin/logging"
	"github.com/sirupsen/logrus"
)

func ConnectionString() string {
	authParams := env.ConstantValues.MongoAuthParams
	dbName := env.ConstantValues.MongoDatabase
	host := env.MongoHost()
	password := env.MongoPassword()
	user := env.MongoUsername()

	connectionString := "mongodb://" + user + ":" + password + "@" + host + "/" + dbName + "?" + authParams

	return connectionString
}

func MongoCollection() *mongo.Collection {
	mongoOptions := options.Client().ApplyURI(ConnectionString())
	client, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		logrus.Error("mongo client failed")
		logging.LogAndPrintError(err)
	}
	collection := client.Database(env.ConstantValues.MongoDatabase).Collection(env.ConstantValues.MongoCollection)

	return collection
}
