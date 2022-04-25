package model

type HarvestReport struct {
	ID               string          `json:"id" bson:"id"`
	URL              *string         `json:"url" bson:"url"`
	DataType         HarvestTypeEnum `json:"dataType" bson:"dataType"`
	HarvestError     bool            `json:"harvestError" bson:"harvestError"`
	StartTime        string          `json:"startTime" bson:"startTime"`
	EndTime          string          `json:"endTime" bson:"endTime"`
	ErrorMessage     *string         `json:"errorMessage" bson:"errorMessage"`
	ChangedCatalogs  []FDKIDAndURI   `json:"changedCatalogs" bson:"changedCatalogs"`
	ChangedResources []FDKIDAndURI   `json:"changedResources" bson:"changedResources"`
}

type FDKIDAndURI struct {
	FDKID string `json:"fdkId"`
	URI   string `json:"uri"`
}

type HarvestTypeEnum string

const (
	ConceptHarvestType          HarvestTypeEnum = "concept"
	DatasetHarvestType          HarvestTypeEnum = "dataset"
	InformationModelHarvestType HarvestTypeEnum = "informationmodel"
	DataServiceHarvestType      HarvestTypeEnum = "dataservice"
	PublicServiceHarvestType    HarvestTypeEnum = "publicService"
	EventHarvestType            HarvestTypeEnum = "event"
)
