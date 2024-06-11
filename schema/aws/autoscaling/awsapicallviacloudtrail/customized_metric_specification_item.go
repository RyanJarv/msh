package awsapicallviacloudtrail

type CustomizedMetricSpecificationItem struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

func (c *CustomizedMetricSpecificationItem) SetName(name string) {
    c.Name = name
}

func (c *CustomizedMetricSpecificationItem) SetValue(value string) {
    c.Value = value
}
