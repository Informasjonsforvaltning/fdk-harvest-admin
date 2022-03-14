package model

type AuthHeader struct {
	Name                string             	`json:"name" bson:"name"`
	Value               string              `json:"value" bson:"value"`
}
