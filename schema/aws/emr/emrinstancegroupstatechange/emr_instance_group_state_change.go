package emrinstancegroupstatechange

type EMRInstanceGroupStateChange struct {
    Market string `json:"market,omitempty"`
    Severity string `json:"severity,omitempty"`
    RequestedInstanceCount int32 `json:"requestedInstanceCount,omitempty"`
    InstanceType string `json:"instanceType,omitempty"`
    InstanceGroupType string `json:"instanceGroupType,omitempty"`
    InstanceGroupId string `json:"instanceGroupId,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
    RunningInstanceCount int32 `json:"runningInstanceCount,omitempty"`
    State string `json:"state,omitempty"`
    Message string `json:"message,omitempty"`
    BidPrice string `json:"bidPrice,omitempty"`
    BidPriceAsPercentageOfOnDemandPrice float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty"`
}

func (e *EMRInstanceGroupStateChange) SetMarket(market string) {
    e.Market = market
}

func (e *EMRInstanceGroupStateChange) SetSeverity(severity string) {
    e.Severity = severity
}

func (e *EMRInstanceGroupStateChange) SetRequestedInstanceCount(requestedInstanceCount int32) {
    e.RequestedInstanceCount = requestedInstanceCount
}

func (e *EMRInstanceGroupStateChange) SetInstanceType(instanceType string) {
    e.InstanceType = instanceType
}

func (e *EMRInstanceGroupStateChange) SetInstanceGroupType(instanceGroupType string) {
    e.InstanceGroupType = instanceGroupType
}

func (e *EMRInstanceGroupStateChange) SetInstanceGroupId(instanceGroupId string) {
    e.InstanceGroupId = instanceGroupId
}

func (e *EMRInstanceGroupStateChange) SetClusterId(clusterId string) {
    e.ClusterId = clusterId
}

func (e *EMRInstanceGroupStateChange) SetRunningInstanceCount(runningInstanceCount int32) {
    e.RunningInstanceCount = runningInstanceCount
}

func (e *EMRInstanceGroupStateChange) SetState(state string) {
    e.State = state
}

func (e *EMRInstanceGroupStateChange) SetMessage(message string) {
    e.Message = message
}

func (e *EMRInstanceGroupStateChange) SetBidPrice(bidPrice string) {
    e.BidPrice = bidPrice
}

func (e *EMRInstanceGroupStateChange) SetBidPriceAsPercentageOfOnDemandPrice(bidPriceAsPercentageOfOnDemandPrice float64) {
    e.BidPriceAsPercentageOfOnDemandPrice = bidPriceAsPercentageOfOnDemandPrice
}
