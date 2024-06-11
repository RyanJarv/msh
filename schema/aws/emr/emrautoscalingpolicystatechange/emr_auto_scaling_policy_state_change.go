package emrautoscalingpolicystatechange

type EMRAutoScalingPolicyStateChange struct {
    ResourceId string `json:"resourceId,omitempty"`
    ClusterId string `json:"clusterId,omitempty"`
    State string `json:"state,omitempty"`
    Message string `json:"message,omitempty"`
    ScalingResourceType string `json:"scalingResourceType,omitempty"`
    Severity string `json:"severity,omitempty"`
}

func (e *EMRAutoScalingPolicyStateChange) SetResourceId(resourceId string) {
    e.ResourceId = resourceId
}

func (e *EMRAutoScalingPolicyStateChange) SetClusterId(clusterId string) {
    e.ClusterId = clusterId
}

func (e *EMRAutoScalingPolicyStateChange) SetState(state string) {
    e.State = state
}

func (e *EMRAutoScalingPolicyStateChange) SetMessage(message string) {
    e.Message = message
}

func (e *EMRAutoScalingPolicyStateChange) SetScalingResourceType(scalingResourceType string) {
    e.ScalingResourceType = scalingResourceType
}

func (e *EMRAutoScalingPolicyStateChange) SetSeverity(severity string) {
    e.Severity = severity
}
