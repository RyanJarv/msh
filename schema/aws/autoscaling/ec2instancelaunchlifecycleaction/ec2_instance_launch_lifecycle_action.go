package ec2instancelaunchlifecycleaction

type EC2InstanceLaunchLifecycleAction struct {
    LifecycleHookName string `json:"LifecycleHookName"`
    LifecycleTransition string `json:"LifecycleTransition"`
    AutoScalingGroupName string `json:"AutoScalingGroupName"`
    EC2InstanceId string `json:"EC2InstanceId"`
    LifecycleActionToken string `json:"LifecycleActionToken"`
    NotificationMetadata string `json:"NotificationMetadata"`
    Origin string `json:"Origin"`
    Destination string `json:"Destination"`
}

func (e *EC2InstanceLaunchLifecycleAction) SetLifecycleHookName(lifecycleHookName string) {
    e.LifecycleHookName = lifecycleHookName
}

func (e *EC2InstanceLaunchLifecycleAction) SetLifecycleTransition(lifecycleTransition string) {
    e.LifecycleTransition = lifecycleTransition
}

func (e *EC2InstanceLaunchLifecycleAction) SetAutoScalingGroupName(autoScalingGroupName string) {
    e.AutoScalingGroupName = autoScalingGroupName
}

func (e *EC2InstanceLaunchLifecycleAction) SetEC2InstanceId(eC2InstanceId string) {
    e.EC2InstanceId = eC2InstanceId
}

func (e *EC2InstanceLaunchLifecycleAction) SetLifecycleActionToken(lifecycleActionToken string) {
    e.LifecycleActionToken = lifecycleActionToken
}

func (e *EC2InstanceLaunchLifecycleAction) SetNotificationMetadata(notificationMetadata string) {
    e.NotificationMetadata = notificationMetadata
}

func (e *EC2InstanceLaunchLifecycleAction) SetOrigin(origin string) {
    e.Origin = origin
}

func (e *EC2InstanceLaunchLifecycleAction) SetDestination(destination string) {
    e.Destination = destination
}
