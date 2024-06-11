package devopsguruinsightclosed

type DataIdentifiers struct {
    MetricQuery MetricQuery `json:"metricQuery,omitempty"`
    ResourceId string `json:"ResourceId,omitempty"`
    ResourceType string `json:"ResourceType,omitempty"`
    Dimensions []CloudWatchDimension `json:"dimensions,omitempty"`
    AlarmCondition AlarmCondition `json:"alarmCondition,omitempty"`
    MetricDisplayName string `json:"metricDisplayName,omitempty"`
    Name string `json:"name,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    Period string `json:"period,omitempty"`
    Stat string `json:"stat,omitempty"`
    Unit string `json:"unit,omitempty"`
}

func (d *DataIdentifiers) SetMetricQuery(metricQuery MetricQuery) {
    d.MetricQuery = metricQuery
}

func (d *DataIdentifiers) SetResourceId(resourceId string) {
    d.ResourceId = resourceId
}

func (d *DataIdentifiers) SetResourceType(resourceType string) {
    d.ResourceType = resourceType
}

func (d *DataIdentifiers) SetDimensions(dimensions []CloudWatchDimension) {
    d.Dimensions = dimensions
}

func (d *DataIdentifiers) SetAlarmCondition(alarmCondition AlarmCondition) {
    d.AlarmCondition = alarmCondition
}

func (d *DataIdentifiers) SetMetricDisplayName(metricDisplayName string) {
    d.MetricDisplayName = metricDisplayName
}

func (d *DataIdentifiers) SetName(name string) {
    d.Name = name
}

func (d *DataIdentifiers) SetNamespace(namespace string) {
    d.Namespace = namespace
}

func (d *DataIdentifiers) SetPeriod(period string) {
    d.Period = period
}

func (d *DataIdentifiers) SetStat(stat string) {
    d.Stat = stat
}

func (d *DataIdentifiers) SetUnit(unit string) {
    d.Unit = unit
}
