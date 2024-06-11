package awsapicallviacloudtrail

type CustomizedMetricSpecification struct {
    Unit string `json:"unit,omitempty"`
    Statistic string `json:"statistic,omitempty"`
    MetricName string `json:"metricName,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    Dimensions []CustomizedMetricSpecificationItem `json:"dimensions,omitempty"`
}

func (c *CustomizedMetricSpecification) SetUnit(unit string) {
    c.Unit = unit
}

func (c *CustomizedMetricSpecification) SetStatistic(statistic string) {
    c.Statistic = statistic
}

func (c *CustomizedMetricSpecification) SetMetricName(metricName string) {
    c.MetricName = metricName
}

func (c *CustomizedMetricSpecification) SetNamespace(namespace string) {
    c.Namespace = namespace
}

func (c *CustomizedMetricSpecification) SetDimensions(dimensions []CustomizedMetricSpecificationItem) {
    c.Dimensions = dimensions
}
