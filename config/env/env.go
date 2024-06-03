package env

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func MongoHost() string {
	return getEnv("MONGO_HOST", "localhost:27017")
}

func MongoPassword() string {
	return getEnv("MONGO_PASSWORD", "admin")
}

func MongoUsername() string {
	return getEnv("MONGO_USERNAME", "admin")
}

func KeycloakHost() string {
	return getEnv("KEYCLOAK_HOST", "https://sso.staging.fellesdatakatalog.digdir.no")
}

func RabbitPassword() string {
	return getEnv("RABBIT_PASSWORD", "guest")
}

func RabbitUsername() string {
	return getEnv("RABBIT_USERNAME", "guest")
}

func RabbitHost() string {
	return getEnv("RABBIT_HOST", "localhost")
}

func RabbitPort() string {
	return getEnv("RABBIT_PORT", "5672")
}

func ApiKey() string {
	return getEnv("API_KEY", "test-key")
}

type Constants struct {
	MongoAuthParams        string
	DataSourcesCollection  string
	ReportsCollection      string
	MongoDatabase          string
	RabbitExchange         string
	RabbitExchangeKind     string
	RabbitMsgKeyMiddle     string
	RabbitMsgKeyEnd        string
	RabbitListenQueue      string
	RabbitNewDataSourceKey string
	RabbitHarvestedKey     string
	RabbitIngestedKey      string
	RabbitReasonedKey      string
}

type Paths struct {
	Datasource             string
	InternalDatasource     string
	HarvestStatus          string
	Datasources            string
	InternalDatasources    string
	OrgDatasources         string
	InternalOrgDatasources string
	Organizations          string
	Ping                   string
	Ready                  string
	StartHarvest           string
	Internal               string
}

type Security struct {
	TokenAudience   string
	SysAdminAuth    string
	OrgType         string
	AdminPermission string
	WritePermission string
}

var ConstantValues = Constants{
	MongoAuthParams:        "authSource=admin&authMechanism=SCRAM-SHA-1",
	DataSourcesCollection:  "datasources",
	ReportsCollection:      "reports",
	MongoDatabase:          "fdk-harvest-admin",
	RabbitExchange:         "harvests",
	RabbitExchangeKind:     "topic",
	RabbitMsgKeyMiddle:     "publisher",
	RabbitMsgKeyEnd:        "HarvestTrigger",
	RabbitListenQueue:      "fdkHarvestAdmin",
	RabbitNewDataSourceKey: "*.publisher.NewDataSource",
	RabbitHarvestedKey:     "*.harvested",
	RabbitIngestedKey:      "*.ingested",
	RabbitReasonedKey:      "*.reasoned",
}

var PathValues = Paths{
	Datasource:             "organizations/:org/datasources/:id",
	InternalDatasource:     "internal/organizations/:org/datasources/:id",
	HarvestStatus:          "organizations/:org/datasources/:id/status",
	Datasources:            "datasources",
	InternalDatasources:    "internal/datasources",
	OrgDatasources:         "organizations/:org/datasources",
	InternalOrgDatasources: "internal/organizations/:org/datasources",
	Organizations:          "organizations",
	Ping:                   "ping",
	Ready:                  "ready",
	StartHarvest:           "organizations/:org/datasources/:id/start-harvesting",
}

var SecurityValues = Security{
	TokenAudience:   "fdk-harvest-admin",
	SysAdminAuth:    "system:root:admin",
	OrgType:         "organization",
	AdminPermission: "admin",
	WritePermission: "write",
}
