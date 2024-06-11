package awsapicallviacloudtrail

type InstancesDistribution struct {
    OnDemandBaseCapacity float64 `json:"onDemandBaseCapacity,omitempty"`
    SpotInstancePools float64 `json:"spotInstancePools,omitempty"`
    SpotAllocationStrategy string `json:"spotAllocationStrategy,omitempty"`
    OnDemandAllocationStrategy string `json:"onDemandAllocationStrategy,omitempty"`
    OnDemandPercentageAboveBaseCapacity float64 `json:"onDemandPercentageAboveBaseCapacity,omitempty"`
}

func (i *InstancesDistribution) SetOnDemandBaseCapacity(onDemandBaseCapacity float64) {
    i.OnDemandBaseCapacity = onDemandBaseCapacity
}

func (i *InstancesDistribution) SetSpotInstancePools(spotInstancePools float64) {
    i.SpotInstancePools = spotInstancePools
}

func (i *InstancesDistribution) SetSpotAllocationStrategy(spotAllocationStrategy string) {
    i.SpotAllocationStrategy = spotAllocationStrategy
}

func (i *InstancesDistribution) SetOnDemandAllocationStrategy(onDemandAllocationStrategy string) {
    i.OnDemandAllocationStrategy = onDemandAllocationStrategy
}

func (i *InstancesDistribution) SetOnDemandPercentageAboveBaseCapacity(onDemandPercentageAboveBaseCapacity float64) {
    i.OnDemandPercentageAboveBaseCapacity = onDemandPercentageAboveBaseCapacity
}
