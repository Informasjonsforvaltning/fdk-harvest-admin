package model

import "fmt"

type DataSourceTypeEnum string

const (
	SkosApNo DataSourceTypeEnum = "SKOS-AP-NO"
	DcatApNo DataSourceTypeEnum = "DCAT-AP-NO"
	CpsvApNo DataSourceTypeEnum = "CPSV-AP-NO"
	Tbx      DataSourceTypeEnum = "TBX"
	ModellDcatApNo DataSourceTypeEnum = "ModellDCAT-AP-NO"
)

func (dataSourceType DataSourceTypeEnum) Validate() error {
	switch dataSourceType {
	case SkosApNo, DcatApNo, CpsvApNo, Tbx, ModellDcatApNo:
		return nil
	}
	return fmt.Errorf("%s is not a valid data source type", dataSourceType)
}
