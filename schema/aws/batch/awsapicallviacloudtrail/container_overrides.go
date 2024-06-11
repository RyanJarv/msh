package awsapicallviacloudtrail

type ContainerOverrides struct {
    Environment []ContainerOverridesItem `json:"environment,omitempty"`
}

func (c *ContainerOverrides) SetEnvironment(environment []ContainerOverridesItem) {
    c.Environment = environment
}
