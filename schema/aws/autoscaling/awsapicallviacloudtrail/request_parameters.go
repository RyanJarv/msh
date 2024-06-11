package awsapicallviacloudtrail

type RequestParameters struct {
    TargetTrackingConfiguration TargetTrackingConfiguration `json:"targetTrackingConfiguration,omitempty"`
    MixedInstancesPolicy MixedInstancesPolicy `json:"mixedInstancesPolicy,omitempty"`
    LaunchTemplate LaunchTemplate `json:"launchTemplate,omitempty"`
    HealthCheckType string `json:"healthCheckType,omitempty"`
    UserData string `json:"userData,omitempty"`
    SpotPrice string `json:"spotPrice,omitempty"`
    BreachThreshold float64 `json:"breachThreshold,omitempty"`
    AdjustmentType string `json:"adjustmentType,omitempty"`
    DefaultCooldown float64 `json:"defaultCooldown,omitempty"`
    MetricValue float64 `json:"metricValue,omitempty"`
    AvailabilityZones []string `json:"availabilityZones,omitempty"`
    MaxSize float64 `json:"maxSize,omitempty"`
    TargetGroupARNs []string `json:"targetGroupARNs,omitempty"`
    HonorCooldown bool `json:"honorCooldown,omitempty"`
    LaunchConfigurationName string `json:"launchConfigurationName,omitempty"`
    AutoScalingGroupName string `json:"autoScalingGroupName,omitempty"`
    ForceDelete bool `json:"forceDelete,omitempty"`
    ScheduledUpdateGroupActions []interface{} `json:"scheduledUpdateGroupActions,omitempty"`
    ScheduledActionName string `json:"scheduledActionName,omitempty"`
    HealthCheckGracePeriod float64 `json:"healthCheckGracePeriod,omitempty"`
    NewInstancesProtectedFromScaleIn bool `json:"newInstancesProtectedFromScaleIn,omitempty"`
    InstanceIds []string `json:"instanceIds,omitempty"`
    MinSize float64 `json:"minSize,omitempty"`
    StartTime string `json:"startTime,omitempty"`
    TopicARN string `json:"topicARN,omitempty"`
    LifecycleHookSpecificationList []interface{} `json:"lifecycleHookSpecificationList,omitempty"`
    ImageId string `json:"imageId,omitempty"`
    LoadBalancerNames []string `json:"loadBalancerNames,omitempty"`
    PolicyName string `json:"policyName,omitempty"`
    ScheduledActionNames []string `json:"scheduledActionNames,omitempty"`
    InstanceType string `json:"instanceType,omitempty"`
    Tags []RequestParametersItem `json:"tags,omitempty"`
    ServiceLinkedRoleARN string `json:"serviceLinkedRoleARN,omitempty"`
    StepAdjustments []RequestParametersItem_1 `json:"stepAdjustments,omitempty"`
    Granularity string `json:"granularity,omitempty"`
    PolicyType string `json:"policyType,omitempty"`
    ScalingAdjustment float64 `json:"scalingAdjustment,omitempty"`
    SecurityGroups []string `json:"securityGroups,omitempty"`
    NotificationTypes []string `json:"notificationTypes,omitempty"`
    Time string `json:"time,omitempty"`
    Metrics []string `json:"metrics,omitempty"`
    ProtectedFromScaleIn bool `json:"protectedFromScaleIn,omitempty"`
    DesiredCapacity float64 `json:"desiredCapacity,omitempty"`
    VPCZoneIdentifier string `json:"vPCZoneIdentifier,omitempty"`
}

func (r *RequestParameters) SetTargetTrackingConfiguration(targetTrackingConfiguration TargetTrackingConfiguration) {
    r.TargetTrackingConfiguration = targetTrackingConfiguration
}

func (r *RequestParameters) SetMixedInstancesPolicy(mixedInstancesPolicy MixedInstancesPolicy) {
    r.MixedInstancesPolicy = mixedInstancesPolicy
}

func (r *RequestParameters) SetLaunchTemplate(launchTemplate LaunchTemplate) {
    r.LaunchTemplate = launchTemplate
}

func (r *RequestParameters) SetHealthCheckType(healthCheckType string) {
    r.HealthCheckType = healthCheckType
}

func (r *RequestParameters) SetUserData(userData string) {
    r.UserData = userData
}

func (r *RequestParameters) SetSpotPrice(spotPrice string) {
    r.SpotPrice = spotPrice
}

func (r *RequestParameters) SetBreachThreshold(breachThreshold float64) {
    r.BreachThreshold = breachThreshold
}

func (r *RequestParameters) SetAdjustmentType(adjustmentType string) {
    r.AdjustmentType = adjustmentType
}

func (r *RequestParameters) SetDefaultCooldown(defaultCooldown float64) {
    r.DefaultCooldown = defaultCooldown
}

func (r *RequestParameters) SetMetricValue(metricValue float64) {
    r.MetricValue = metricValue
}

func (r *RequestParameters) SetAvailabilityZones(availabilityZones []string) {
    r.AvailabilityZones = availabilityZones
}

func (r *RequestParameters) SetMaxSize(maxSize float64) {
    r.MaxSize = maxSize
}

func (r *RequestParameters) SetTargetGroupARNs(targetGroupARNs []string) {
    r.TargetGroupARNs = targetGroupARNs
}

func (r *RequestParameters) SetHonorCooldown(honorCooldown bool) {
    r.HonorCooldown = honorCooldown
}

func (r *RequestParameters) SetLaunchConfigurationName(launchConfigurationName string) {
    r.LaunchConfigurationName = launchConfigurationName
}

func (r *RequestParameters) SetAutoScalingGroupName(autoScalingGroupName string) {
    r.AutoScalingGroupName = autoScalingGroupName
}

func (r *RequestParameters) SetForceDelete(forceDelete bool) {
    r.ForceDelete = forceDelete
}

func (r *RequestParameters) SetScheduledUpdateGroupActions(scheduledUpdateGroupActions []interface{}) {
    r.ScheduledUpdateGroupActions = scheduledUpdateGroupActions
}

func (r *RequestParameters) SetScheduledActionName(scheduledActionName string) {
    r.ScheduledActionName = scheduledActionName
}

func (r *RequestParameters) SetHealthCheckGracePeriod(healthCheckGracePeriod float64) {
    r.HealthCheckGracePeriod = healthCheckGracePeriod
}

func (r *RequestParameters) SetNewInstancesProtectedFromScaleIn(newInstancesProtectedFromScaleIn bool) {
    r.NewInstancesProtectedFromScaleIn = newInstancesProtectedFromScaleIn
}

func (r *RequestParameters) SetInstanceIds(instanceIds []string) {
    r.InstanceIds = instanceIds
}

func (r *RequestParameters) SetMinSize(minSize float64) {
    r.MinSize = minSize
}

func (r *RequestParameters) SetStartTime(startTime string) {
    r.StartTime = startTime
}

func (r *RequestParameters) SetTopicARN(topicARN string) {
    r.TopicARN = topicARN
}

func (r *RequestParameters) SetLifecycleHookSpecificationList(lifecycleHookSpecificationList []interface{}) {
    r.LifecycleHookSpecificationList = lifecycleHookSpecificationList
}

func (r *RequestParameters) SetImageId(imageId string) {
    r.ImageId = imageId
}

func (r *RequestParameters) SetLoadBalancerNames(loadBalancerNames []string) {
    r.LoadBalancerNames = loadBalancerNames
}

func (r *RequestParameters) SetPolicyName(policyName string) {
    r.PolicyName = policyName
}

func (r *RequestParameters) SetScheduledActionNames(scheduledActionNames []string) {
    r.ScheduledActionNames = scheduledActionNames
}

func (r *RequestParameters) SetInstanceType(instanceType string) {
    r.InstanceType = instanceType
}

func (r *RequestParameters) SetTags(tags []RequestParametersItem) {
    r.Tags = tags
}

func (r *RequestParameters) SetServiceLinkedRoleARN(serviceLinkedRoleARN string) {
    r.ServiceLinkedRoleARN = serviceLinkedRoleARN
}

func (r *RequestParameters) SetStepAdjustments(stepAdjustments []RequestParametersItem_1) {
    r.StepAdjustments = stepAdjustments
}

func (r *RequestParameters) SetGranularity(granularity string) {
    r.Granularity = granularity
}

func (r *RequestParameters) SetPolicyType(policyType string) {
    r.PolicyType = policyType
}

func (r *RequestParameters) SetScalingAdjustment(scalingAdjustment float64) {
    r.ScalingAdjustment = scalingAdjustment
}

func (r *RequestParameters) SetSecurityGroups(securityGroups []string) {
    r.SecurityGroups = securityGroups
}

func (r *RequestParameters) SetNotificationTypes(notificationTypes []string) {
    r.NotificationTypes = notificationTypes
}

func (r *RequestParameters) SetTime(time string) {
    r.Time = time
}

func (r *RequestParameters) SetMetrics(metrics []string) {
    r.Metrics = metrics
}

func (r *RequestParameters) SetProtectedFromScaleIn(protectedFromScaleIn bool) {
    r.ProtectedFromScaleIn = protectedFromScaleIn
}

func (r *RequestParameters) SetDesiredCapacity(desiredCapacity float64) {
    r.DesiredCapacity = desiredCapacity
}

func (r *RequestParameters) SetVPCZoneIdentifier(vPCZoneIdentifier string) {
    r.VPCZoneIdentifier = vPCZoneIdentifier
}
