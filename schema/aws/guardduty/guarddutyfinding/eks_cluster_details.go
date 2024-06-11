package guarddutyfinding

type EksClusterDetails struct {
    Arn string `json:"arn,omitempty"`
    CreatedAt float64 `json:"createdAt,omitempty"`
    Name string `json:"name,omitempty"`
    Status string `json:"status,omitempty"`
    Tags []EcsClusterDetailsItem `json:"tags,omitempty"`
    VpcId string `json:"vpcId,omitempty"`
}

func (e *EksClusterDetails) SetArn(arn string) {
    e.Arn = arn
}

func (e *EksClusterDetails) SetCreatedAt(createdAt float64) {
    e.CreatedAt = createdAt
}

func (e *EksClusterDetails) SetName(name string) {
    e.Name = name
}

func (e *EksClusterDetails) SetStatus(status string) {
    e.Status = status
}

func (e *EksClusterDetails) SetTags(tags []EcsClusterDetailsItem) {
    e.Tags = tags
}

func (e *EksClusterDetails) SetVpcId(vpcId string) {
    e.VpcId = vpcId
}
