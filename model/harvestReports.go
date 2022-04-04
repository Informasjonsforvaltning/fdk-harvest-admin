package model

type HarvestReports struct {
	Id      string                   `bson:"id"`
	Reports map[string]HarvestReport `bson:"reports"`
}
