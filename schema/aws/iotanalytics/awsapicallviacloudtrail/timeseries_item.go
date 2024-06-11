package awsapicallviacloudtrail

type TimeseriesItem struct {
    Attribute string `json:"attribute"`
    Name string `json:"name"`
}

func (t *TimeseriesItem) SetAttribute(attribute string) {
    t.Attribute = attribute
}

func (t *TimeseriesItem) SetName(name string) {
    t.Name = name
}
