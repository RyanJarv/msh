package awsapicallviacloudtrail

type ResponseElementsItemItem_1 struct {
    Image string `json:"image"`
    NetworkInterfaces []interface{} `json:"networkInterfaces"`
    Memory string `json:"memory,omitempty"`
    TaskArn string `json:"taskArn"`
    Name string `json:"name"`
    Cpu string `json:"cpu"`
    ContainerArn string `json:"containerArn"`
    LastStatus string `json:"lastStatus"`
}

func (r *ResponseElementsItemItem_1) SetImage(image string) {
    r.Image = image
}

func (r *ResponseElementsItemItem_1) SetNetworkInterfaces(networkInterfaces []interface{}) {
    r.NetworkInterfaces = networkInterfaces
}

func (r *ResponseElementsItemItem_1) SetMemory(memory string) {
    r.Memory = memory
}

func (r *ResponseElementsItemItem_1) SetTaskArn(taskArn string) {
    r.TaskArn = taskArn
}

func (r *ResponseElementsItemItem_1) SetName(name string) {
    r.Name = name
}

func (r *ResponseElementsItemItem_1) SetCpu(cpu string) {
    r.Cpu = cpu
}

func (r *ResponseElementsItemItem_1) SetContainerArn(containerArn string) {
    r.ContainerArn = containerArn
}

func (r *ResponseElementsItemItem_1) SetLastStatus(lastStatus string) {
    r.LastStatus = lastStatus
}
