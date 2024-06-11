package awsapicallviacloudtrail

type Overrides struct {
    ContainerOverrides []OverridesItem `json:"containerOverrides,omitempty"`
}

func (o *Overrides) SetContainerOverrides(containerOverrides []OverridesItem) {
    o.ContainerOverrides = containerOverrides
}
