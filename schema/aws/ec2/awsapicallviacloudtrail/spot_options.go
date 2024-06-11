package awsapicallviacloudtrail

type SpotOptions struct {
    AllocationStrategy string `json:"AllocationStrategy,omitempty"`
    MaxTargetCapacity float64 `json:"MaxTargetCapacity,omitempty"`
    MaxInstanceCount float64 `json:"MaxInstanceCount,omitempty"`
    InstancePoolsToUseCount float64 `json:"InstancePoolsToUseCount,omitempty"`
    InstancePoolConstraintFilterDisabled bool `json:"InstancePoolConstraintFilterDisabled,omitempty"`
}

func (s *SpotOptions) SetAllocationStrategy(allocationStrategy string) {
    s.AllocationStrategy = allocationStrategy
}

func (s *SpotOptions) SetMaxTargetCapacity(maxTargetCapacity float64) {
    s.MaxTargetCapacity = maxTargetCapacity
}

func (s *SpotOptions) SetMaxInstanceCount(maxInstanceCount float64) {
    s.MaxInstanceCount = maxInstanceCount
}

func (s *SpotOptions) SetInstancePoolsToUseCount(instancePoolsToUseCount float64) {
    s.InstancePoolsToUseCount = instancePoolsToUseCount
}

func (s *SpotOptions) SetInstancePoolConstraintFilterDisabled(instancePoolConstraintFilterDisabled bool) {
    s.InstancePoolConstraintFilterDisabled = instancePoolConstraintFilterDisabled
}
