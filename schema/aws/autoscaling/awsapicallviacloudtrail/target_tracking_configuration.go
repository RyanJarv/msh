package awsapicallviacloudtrail

type TargetTrackingConfiguration struct {
    PredefinedMetricSpecification PredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty"`
    CustomizedMetricSpecification CustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty"`
    TargetValue float64 `json:"targetValue,omitempty"`
}

func (t *TargetTrackingConfiguration) SetPredefinedMetricSpecification(predefinedMetricSpecification PredefinedMetricSpecification) {
    t.PredefinedMetricSpecification = predefinedMetricSpecification
}

func (t *TargetTrackingConfiguration) SetCustomizedMetricSpecification(customizedMetricSpecification CustomizedMetricSpecification) {
    t.CustomizedMetricSpecification = customizedMetricSpecification
}

func (t *TargetTrackingConfiguration) SetTargetValue(targetValue float64) {
    t.TargetValue = targetValue
}
