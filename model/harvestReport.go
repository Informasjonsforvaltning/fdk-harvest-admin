package model

type HarvestReport struct {
	Id               string          `json:"id" bson:"id"`
	Url              *string         `json:"url" bson:"url"`
	DataType         HarvestTypeEnum `json:"dataType" bson:"dataType"`
	HarvestError     bool            `json:"harvestError" bson:"harvestError"`
	StartTime        string          `json:"startTime" bson:"startTime"`
	EndTime          string          `json:"endTime" bson:"endTime"`
	ErrorMessage     *string         `json:"errorMessage" bson:"errorMessage"`
	ChangedCatalogs  []FdkIdAndUri   `json:"changedCatalogs" bson:"changedCatalogs"`
	ChangedResources []FdkIdAndUri   `json:"changedResources" bson:"changedResources"`
}

type FdkIdAndUri struct {
	FdkId string `json:"fdkId"`
	Uri   string `json:"uri"`
}

type HarvestTypeEnum string

const (
	ConceptHarvestType          = "concept"
	DatasetHarvestType          = "dataset"
	InformationModelHarvestType = "informationmodel"
	DataServiceHarvestType      = "dataservice"
	PublicServiceHarvestType    = "publicService"
	EventHarvestType            = "event"
)
