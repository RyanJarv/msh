package awsapicallviacloudtrail

type RequestParameters struct {
    HealthCheck HealthCheck_1 `json:"healthCheck,omitempty"`
    Listeners []RequestParametersItem `json:"listeners,omitempty"`
    Instances []ResponseElementsItem `json:"instances,omitempty"`
    LoadBalancerName string `json:"loadBalancerName,omitempty"`
    TargetGroupArn string `json:"targetGroupArn,omitempty"`
    Subnets []string `json:"subnets,omitempty"`
    Targets []RequestParametersItem_1 `json:"targets,omitempty"`
}

func (r *RequestParameters) SetHealthCheck(healthCheck HealthCheck_1) {
    r.HealthCheck = healthCheck
}

func (r *RequestParameters) SetListeners(listeners []RequestParametersItem) {
    r.Listeners = listeners
}

func (r *RequestParameters) SetInstances(instances []ResponseElementsItem) {
    r.Instances = instances
}

func (r *RequestParameters) SetLoadBalancerName(loadBalancerName string) {
    r.LoadBalancerName = loadBalancerName
}

func (r *RequestParameters) SetTargetGroupArn(targetGroupArn string) {
    r.TargetGroupArn = targetGroupArn
}

func (r *RequestParameters) SetSubnets(subnets []string) {
    r.Subnets = subnets
}

func (r *RequestParameters) SetTargets(targets []RequestParametersItem_1) {
    r.Targets = targets
}
