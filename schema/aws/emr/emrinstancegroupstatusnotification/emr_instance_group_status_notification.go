package emrinstancegroupstatusnotification

type EMRInstanceGroupStatusNotification struct {
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

func (e *EMRInstanceGroupStatusNotification) SetMarket(market string) {
    e.Market = market
}

func (e *EMRInstanceGroupStatusNotification) SetSeverity(severity string) {
    e.Severity = severity
}

func (e *EMRInstanceGroupStatusNotification) SetRequestedInstanceCount(requestedInstanceCount int32) {
    e.RequestedInstanceCount = requestedInstanceCount
}

func (e *EMRInstanceGroupStatusNotification) SetInstanceType(instanceType string) {
    e.InstanceType = instanceType
}

func (e *EMRInstanceGroupStatusNotification) SetInstanceGroupType(instanceGroupType string) {
    e.InstanceGroupType = instanceGroupType
}

func (e *EMRInstanceGroupStatusNotification) SetInstanceGroupId(instanceGroupId string) {
    e.InstanceGroupId = instanceGroupId
}

func (e *EMRInstanceGroupStatusNotification) SetClusterId(clusterId string) {
    e.ClusterId = clusterId
}

func (e *EMRInstanceGroupStatusNotification) SetRunningInstanceCount(runningInstanceCount int32) {
    e.RunningInstanceCount = runningInstanceCount
}

func (e *EMRInstanceGroupStatusNotification) SetState(state string) {
    e.State = state
}

func (e *EMRInstanceGroupStatusNotification) SetMessage(message string) {
    e.Message = message
}

func (e *EMRInstanceGroupStatusNotification) SetBidPrice(bidPrice string) {
    e.BidPrice = bidPrice
}

func (e *EMRInstanceGroupStatusNotification) SetBidPriceAsPercentageOfOnDemandPrice(bidPriceAsPercentageOfOnDemandPrice float64) {
    e.BidPriceAsPercentageOfOnDemandPrice = bidPriceAsPercentageOfOnDemandPrice
}
