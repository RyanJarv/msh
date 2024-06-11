package guarddutyfinding

type Action struct {
    AwsApiCallAction AwsApiCallAction_1 `json:"awsApiCallAction,omitempty"`
    DnsRequestAction DnsRequestAction `json:"dnsRequestAction,omitempty"`
    KubernetesApiCallAction KubernetesApiCallAction `json:"kubernetesApiCallAction,omitempty"`
    NetworkConnectionAction NetworkConnectionAction `json:"networkConnectionAction,omitempty"`
    PortProbeAction PortProbeAction `json:"portProbeAction,omitempty"`
    ActionType string `json:"actionType,omitempty"`
}

func (a *Action) SetAwsApiCallAction(awsApiCallAction AwsApiCallAction_1) {
    a.AwsApiCallAction = awsApiCallAction
}

func (a *Action) SetDnsRequestAction(dnsRequestAction DnsRequestAction) {
    a.DnsRequestAction = dnsRequestAction
}

func (a *Action) SetKubernetesApiCallAction(kubernetesApiCallAction KubernetesApiCallAction) {
    a.KubernetesApiCallAction = kubernetesApiCallAction
}

func (a *Action) SetNetworkConnectionAction(networkConnectionAction NetworkConnectionAction) {
    a.NetworkConnectionAction = networkConnectionAction
}

func (a *Action) SetPortProbeAction(portProbeAction PortProbeAction) {
    a.PortProbeAction = portProbeAction
}

func (a *Action) SetActionType(actionType string) {
    a.ActionType = actionType
}
