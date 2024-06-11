package awsapicallviacloudtrail

type Overrides_1Item struct {
    ResourceRequirements []interface{} `json:"resourceRequirements,omitempty"`
    Environment []interface{} `json:"environment,omitempty"`
    Memory float64 `json:"memory,omitempty"`
    Name string `json:"name"`
    Cpu float64 `json:"cpu,omitempty"`
    Command []string `json:"command,omitempty"`
}

func (o *Overrides_1Item) SetResourceRequirements(resourceRequirements []interface{}) {
    o.ResourceRequirements = resourceRequirements
}

func (o *Overrides_1Item) SetEnvironment(environment []interface{}) {
    o.Environment = environment
}

func (o *Overrides_1Item) SetMemory(memory float64) {
    o.Memory = memory
}

func (o *Overrides_1Item) SetName(name string) {
    o.Name = name
}

func (o *Overrides_1Item) SetCpu(cpu float64) {
    o.Cpu = cpu
}

func (o *Overrides_1Item) SetCommand(command []string) {
    o.Command = command
}
