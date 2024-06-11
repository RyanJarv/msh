package awsapicallviacloudtrail

type OnDemandOptions struct {
    AllocationStrategy string `json:"AllocationStrategy,omitempty"`
    MaxTargetCapacity float64 `json:"MaxTargetCapacity,omitempty"`
    MaxInstanceCount float64 `json:"MaxInstanceCount,omitempty"`
    InstancePoolConstraintFilterDisabled bool `json:"InstancePoolConstraintFilterDisabled,omitempty"`
}

func (o *OnDemandOptions) SetAllocationStrategy(allocationStrategy string) {
    o.AllocationStrategy = allocationStrategy
}

func (o *OnDemandOptions) SetMaxTargetCapacity(maxTargetCapacity float64) {
    o.MaxTargetCapacity = maxTargetCapacity
}

func (o *OnDemandOptions) SetMaxInstanceCount(maxInstanceCount float64) {
    o.MaxInstanceCount = maxInstanceCount
}

func (o *OnDemandOptions) SetInstancePoolConstraintFilterDisabled(instancePoolConstraintFilterDisabled bool) {
    o.InstancePoolConstraintFilterDisabled = instancePoolConstraintFilterDisabled
}
