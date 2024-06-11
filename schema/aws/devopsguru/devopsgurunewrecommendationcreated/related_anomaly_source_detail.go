package devopsgurunewrecommendationcreated

type RelatedAnomalySourceDetail struct {
    CloudWatchMetrics []CloudWatchDimension `json:"cloudWatchMetrics"`
}

func (r *RelatedAnomalySourceDetail) SetCloudWatchMetrics(cloudWatchMetrics []CloudWatchDimension) {
    r.CloudWatchMetrics = cloudWatchMetrics
}
