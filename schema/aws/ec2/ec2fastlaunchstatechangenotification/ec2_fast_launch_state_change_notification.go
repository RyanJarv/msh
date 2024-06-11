package ec2fastlaunchstatechangenotification

type EC2FastLaunchStateChangeNotification struct {
    ImageId string `json:"imageId"`
    ResourceType string `json:"resourceType"`
    State string `json:"state"`
    StateTransitionReason string `json:"stateTransitionReason"`
}

func (e *EC2FastLaunchStateChangeNotification) SetImageId(imageId string) {
    e.ImageId = imageId
}

func (e *EC2FastLaunchStateChangeNotification) SetResourceType(resourceType string) {
    e.ResourceType = resourceType
}

func (e *EC2FastLaunchStateChangeNotification) SetState(state string) {
    e.State = state
}

func (e *EC2FastLaunchStateChangeNotification) SetStateTransitionReason(stateTransitionReason string) {
    e.StateTransitionReason = stateTransitionReason
}
