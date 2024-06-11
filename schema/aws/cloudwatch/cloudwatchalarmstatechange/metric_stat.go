package cloudwatchalarmstatechange

type MetricStat struct {
    Metric Metric `json:"metric"`
    Period float64 `json:"period"`
    Stat string `json:"stat"`
}

func (m *MetricStat) SetMetric(metric Metric) {
    m.Metric = metric
}

func (m *MetricStat) SetPeriod(period float64) {
    m.Period = period
}

func (m *MetricStat) SetStat(stat string) {
    m.Stat = stat
}
