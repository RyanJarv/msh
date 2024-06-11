package awsapicallviacloudtrail

type PredefinedMetricSpecification struct {
    PredefinedMetricType string `json:"predefinedMetricType,omitempty"`
}

func (p *PredefinedMetricSpecification) SetPredefinedMetricType(predefinedMetricType string) {
    p.PredefinedMetricType = predefinedMetricType
}
