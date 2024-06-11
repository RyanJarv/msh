package devopsgurunewrecommendationcreated

type CloudWatchDimension struct {
    MetricName string `json:"metricName"`
    Namespace string `json:"namespace"`
}

func (c *CloudWatchDimension) SetMetricName(metricName string) {
    c.MetricName = metricName
}

func (c *CloudWatchDimension) SetNamespace(namespace string) {
    c.Namespace = namespace
}
