package awsapicallviacloudtrail

type TargetCapacitySpecification struct {
    DefaultTargetCapacityType string `json:"DefaultTargetCapacityType,omitempty"`
    TotalTargetCapacity float64 `json:"TotalTargetCapacity,omitempty"`
    OnDemandTargetCapacity float64 `json:"OnDemandTargetCapacity,omitempty"`
    SpotTargetCapacity float64 `json:"SpotTargetCapacity,omitempty"`
}

func (t *TargetCapacitySpecification) SetDefaultTargetCapacityType(defaultTargetCapacityType string) {
    t.DefaultTargetCapacityType = defaultTargetCapacityType
}

func (t *TargetCapacitySpecification) SetTotalTargetCapacity(totalTargetCapacity float64) {
    t.TotalTargetCapacity = totalTargetCapacity
}

func (t *TargetCapacitySpecification) SetOnDemandTargetCapacity(onDemandTargetCapacity float64) {
    t.OnDemandTargetCapacity = onDemandTargetCapacity
}

func (t *TargetCapacitySpecification) SetSpotTargetCapacity(spotTargetCapacity float64) {
    t.SpotTargetCapacity = spotTargetCapacity
}
