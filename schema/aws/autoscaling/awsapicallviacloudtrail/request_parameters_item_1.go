package awsapicallviacloudtrail

type RequestParametersItem_1 struct {
    MetricIntervalLowerBound float64 `json:"metricIntervalLowerBound"`
    ScalingAdjustment float64 `json:"scalingAdjustment"`
}

func (r *RequestParametersItem_1) SetMetricIntervalLowerBound(metricIntervalLowerBound float64) {
    r.MetricIntervalLowerBound = metricIntervalLowerBound
}

func (r *RequestParametersItem_1) SetScalingAdjustment(scalingAdjustment float64) {
    r.ScalingAdjustment = scalingAdjustment
}
