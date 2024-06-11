package devopsgurunewinsightopen

type MetricQuery struct {
    GroupBy GroupBy `json:"groupBy,omitempty"`
    Metric string `json:"metric,omitempty"`
}

func (m *MetricQuery) SetGroupBy(groupBy GroupBy) {
    m.GroupBy = groupBy
}

func (m *MetricQuery) SetMetric(metric string) {
    m.Metric = metric
}
