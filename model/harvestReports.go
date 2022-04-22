package model

type HarvestReports struct {
	ID      string                   `bson:"id"`
	Reports map[string]HarvestReport `bson:"reports"`
}
