package awsapicallviacloudtrail

type ExistingInstances struct {
    OperatingSystem string `json:"OperatingSystem,omitempty"`
    AvailabilityZone string `json:"AvailabilityZone,omitempty"`
    Tag float64 `json:"tag,omitempty"`
    Count float64 `json:"Count,omitempty"`
    InstanceType string `json:"InstanceType,omitempty"`
    MarketOption string `json:"MarketOption,omitempty"`
}

func (e *ExistingInstances) SetOperatingSystem(operatingSystem string) {
    e.OperatingSystem = operatingSystem
}

func (e *ExistingInstances) SetAvailabilityZone(availabilityZone string) {
    e.AvailabilityZone = availabilityZone
}

func (e *ExistingInstances) SetTag(tag float64) {
    e.Tag = tag
}

func (e *ExistingInstances) SetCount(count float64) {
    e.Count = count
}

func (e *ExistingInstances) SetInstanceType(instanceType string) {
    e.InstanceType = instanceType
}

func (e *ExistingInstances) SetMarketOption(marketOption string) {
    e.MarketOption = marketOption
}
