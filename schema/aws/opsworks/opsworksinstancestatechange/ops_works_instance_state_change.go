package opsworksinstancestatechange

type OpsWorksInstanceStateChange struct {
    InitiatedBy string `json:"initiated_by"`
    Hostname string `json:"hostname"`
    StackId string `json:"stack-id"`
    LayerIds []string `json:"layer-ids"`
    InstanceId string `json:"instance-id"`
    Ec2InstanceId string `json:"ec2-instance-id"`
    Status string `json:"status"`
}

func (o *OpsWorksInstanceStateChange) SetInitiatedBy(initiatedBy string) {
    o.InitiatedBy = initiatedBy
}

func (o *OpsWorksInstanceStateChange) SetHostname(hostname string) {
    o.Hostname = hostname
}

func (o *OpsWorksInstanceStateChange) SetStackId(stackId string) {
    o.StackId = stackId
}

func (o *OpsWorksInstanceStateChange) SetLayerIds(layerIds []string) {
    o.LayerIds = layerIds
}

func (o *OpsWorksInstanceStateChange) SetInstanceId(instanceId string) {
    o.InstanceId = instanceId
}

func (o *OpsWorksInstanceStateChange) SetEc2InstanceId(ec2InstanceId string) {
    o.Ec2InstanceId = ec2InstanceId
}

func (o *OpsWorksInstanceStateChange) SetStatus(status string) {
    o.Status = status
}
