package awsapicallviacloudtrail

type OverridesItem struct {
    ResourceRequirements []interface{} `json:"resourceRequirements,omitempty"`
    Environment []interface{} `json:"environment"`
    Memory float64 `json:"memory,omitempty"`
    Name string `json:"name"`
    Cpu float64 `json:"cpu,omitempty"`
    Command []string `json:"command,omitempty"`
}

func (o *OverridesItem) SetResourceRequirements(resourceRequirements []interface{}) {
    o.ResourceRequirements = resourceRequirements
}

func (o *OverridesItem) SetEnvironment(environment []interface{}) {
    o.Environment = environment
}

func (o *OverridesItem) SetMemory(memory float64) {
    o.Memory = memory
}

func (o *OverridesItem) SetName(name string) {
    o.Name = name
}

func (o *OverridesItem) SetCpu(cpu float64) {
    o.Cpu = cpu
}

func (o *OverridesItem) SetCommand(command []string) {
    o.Command = command
}
