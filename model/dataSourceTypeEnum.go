package model

import "fmt"

type DataSourceTypeEnum string

const (
	SkosApNo = "SKOS-AP-NO"
	DcatApNo = "DCAT-AP-NO"
	CpsvApNo = "CPSV-AP-NO"
	Tbx      = "TBX"
)

func (dataSourceType DataSourceTypeEnum) Validate() error {
	switch dataSourceType {
	case SkosApNo, DcatApNo, CpsvApNo, Tbx:
		return nil
	}
	return fmt.Errorf("%s is not a valid data source type", dataSourceType)
}
