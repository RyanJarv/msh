package emrinstancefleetstatechange

type EMRInstanceFleetStateChange struct {
    Severity string `json:"severity,omitempty"`
    ProvisionedSpotCapacity float64 `json:"provisionedSpotCapacity,omitempty"`
    TargetSpotCapacity float64 `json:"targetSpotCapacity,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
    TargetOnDemandCapacity float64 `json:"targetOnDemandCapacity,omitempty"`
    State string `json:"state,omitempty"`
    InstanceFleetId string `json:"instanceFleetId,omitempty"`
    ProvisionedOnDemandCapacity float64 `json:"provisionedOnDemandCapacity,omitempty"`
    Message string `json:"message,omitempty"`
    InstanceFleetType string `json:"instanceFleetType,omitempty"`
}

func (e *EMRInstanceFleetStateChange) SetSeverity(severity string) {
    e.Severity = severity
}

func (e *EMRInstanceFleetStateChange) SetProvisionedSpotCapacity(provisionedSpotCapacity float64) {
    e.ProvisionedSpotCapacity = provisionedSpotCapacity
}

func (e *EMRInstanceFleetStateChange) SetTargetSpotCapacity(targetSpotCapacity float64) {
    e.TargetSpotCapacity = targetSpotCapacity
}

func (e *EMRInstanceFleetStateChange) SetClusterId(clusterId string) {
    e.ClusterId = clusterId
}

func (e *EMRInstanceFleetStateChange) SetTargetOnDemandCapacity(targetOnDemandCapacity float64) {
    e.TargetOnDemandCapacity = targetOnDemandCapacity
}

func (e *EMRInstanceFleetStateChange) SetState(state string) {
    e.State = state
}

func (e *EMRInstanceFleetStateChange) SetInstanceFleetId(instanceFleetId string) {
    e.InstanceFleetId = instanceFleetId
}

func (e *EMRInstanceFleetStateChange) SetProvisionedOnDemandCapacity(provisionedOnDemandCapacity float64) {
    e.ProvisionedOnDemandCapacity = provisionedOnDemandCapacity
}

func (e *EMRInstanceFleetStateChange) SetMessage(message string) {
    e.Message = message
}

func (e *EMRInstanceFleetStateChange) SetInstanceFleetType(instanceFleetType string) {
    e.InstanceFleetType = instanceFleetType
}
