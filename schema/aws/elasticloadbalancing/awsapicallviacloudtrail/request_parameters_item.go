package awsapicallviacloudtrail

type RequestParametersItem struct {
    InstancePort float64 `json:"instancePort"`
    Protocol string `json:"protocol"`
    LoadBalancerPort float64 `json:"loadBalancerPort"`
}

func (r *RequestParametersItem) SetInstancePort(instancePort float64) {
    r.InstancePort = instancePort
}

func (r *RequestParametersItem) SetProtocol(protocol string) {
    r.Protocol = protocol
}

func (r *RequestParametersItem) SetLoadBalancerPort(loadBalancerPort float64) {
    r.LoadBalancerPort = loadBalancerPort
}
