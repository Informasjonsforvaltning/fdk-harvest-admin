package model

import "errors"

type DataSourceTypeEnum string

const (
	SkosApNo = "SKOS-AP-NO"
	DcatApNo = "DCAT-AP-NO"
	CpsvApNo = "CPSV-AP-NO"
)

func (dataSourceType DataSourceTypeEnum) Validate() error {
	switch dataSourceType {
	case SkosApNo, DcatApNo, CpsvApNo:
		return nil
	}
	return errors.New(string(dataSourceType) + " is not a valid data source type")
}
