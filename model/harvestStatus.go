package model

type HarvestStatuses struct {
	ID       string          `json:"id"`
	Statuses []HarvestStatus `json:"statuses"`
}

type HarvestStatus struct {
	HarvestType  string     `json:"harvestType"`
	Status       StatusEnum `json:"status"`
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	StartTime    string     `json:"startTime"`
	EndTime      *string    `json:"endTime,omitempty"`
}

type StatusEnum string

const (
	HarvestDone       = "done"
	HarvestError      = "error"
	HarvestInProgress = "in-progress"
)
