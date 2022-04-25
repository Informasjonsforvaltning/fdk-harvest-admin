package model

import "fmt"

type DataTypeEnum string

const (
	Concept          DataTypeEnum = "concept"
	Dataset          DataTypeEnum = "dataset"
	InformationModel DataTypeEnum = "informationmodel"
	DataService      DataTypeEnum = "dataservice"
	PublicService    DataTypeEnum = "publicService"
)

func (dataType DataTypeEnum) Validate() error {
	switch dataType {
	case Concept, Dataset, InformationModel, DataService, PublicService:
		return nil
	}
	return fmt.Errorf("%s is not a valid data type", dataType)
}
