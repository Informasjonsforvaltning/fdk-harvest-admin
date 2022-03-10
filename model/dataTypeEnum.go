package model

import "errors"

type DataTypeEnum string

const (
	Concept          = "concept"
	Dataset          = "dataset"
	InformationModel = "informationmodel"
	DataService      = "dataservice"
	PublicService    = "publicService"
)

func (dataType DataTypeEnum) Validate() error {
	switch dataType {
	case Concept, Dataset, InformationModel, DataService, PublicService:
		return nil
	}
	return errors.New(string(dataType) + " is not a valid data type")
}
