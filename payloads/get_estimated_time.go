package payloads

type TimeEstimationResult struct {
	ID                  string         `json:"JobID"`
	TotalTimeEstimation string         `json:"TotalTimeEstimation"`
	Resources           []ResourceData `json:"Resources"`
}
type ResourceData struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name"`
	Region             string `json:"region"`
	TimeEstimation     string `json:"timeEstimation"`
	ServiceType        string `json:"serviceType"`
	Action             string `json:"action"`
	StartTime          string `json:"startTime"`
	Day                string `json:"day"`
	AccuracyPercentage int64  `json:"accuracyPercentage"`
}
