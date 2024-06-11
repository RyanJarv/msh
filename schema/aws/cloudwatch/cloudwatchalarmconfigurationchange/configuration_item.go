package cloudwatchalarmconfigurationchange

type ConfigurationItem struct {
    MetricStat MetricStat `json:"metricStat"`
    ReturnData bool `json:"returnData"`
    Id string `json:"id"`
}

func (c *ConfigurationItem) SetMetricStat(metricStat MetricStat) {
    c.MetricStat = metricStat
}

func (c *ConfigurationItem) SetReturnData(returnData bool) {
    c.ReturnData = returnData
}

func (c *ConfigurationItem) SetId(id string) {
    c.Id = id
}
