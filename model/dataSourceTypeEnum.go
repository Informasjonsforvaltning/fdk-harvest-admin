package model

import "fmt"

type DataSourceTypeEnum string

const (
	SkosApNo DataSourceTypeEnum = "SKOS-AP-NO"
	DcatApNo DataSourceTypeEnum = "DCAT-AP-NO"
	CpsvApNo DataSourceTypeEnum = "CPSV-AP-NO"
	Tbx      DataSourceTypeEnum = "TBX"
)

func (dataSourceType DataSourceTypeEnum) Validate() error {
	switch dataSourceType {
	case SkosApNo, DcatApNo, CpsvApNo, Tbx:
		return nil
	}
	return fmt.Errorf("%s is not a valid data source type", dataSourceType)
}
