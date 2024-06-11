package batchjobstatechange

type Container struct {
    Image string `json:"image"`
    Memory float64 `json:"memory"`
    LogStreamName string `json:"logStreamName,omitempty"`
    TaskArn string `json:"taskArn,omitempty"`
    Volumes []Volumes `json:"volumes"`
    Vcpus float64 `json:"vcpus"`
    Command []string `json:"command"`
    ResourceRequirements []ResourceRequirement `json:"resourceRequirements"`
    Environment []ContainerItem `json:"environment"`
    Ulimits []ULimit `json:"ulimits"`
    NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`
    MountPoints []MountPoint `json:"mountPoints"`
    ExitCode float64 `json:"exitCode,omitempty"`
    ContainerInstanceArn string `json:"containerInstanceArn,omitempty"`
}

func (c *Container) SetImage(image string) {
    c.Image = image
}

func (c *Container) SetMemory(memory float64) {
    c.Memory = memory
}

func (c *Container) SetLogStreamName(logStreamName string) {
    c.LogStreamName = logStreamName
}

func (c *Container) SetTaskArn(taskArn string) {
    c.TaskArn = taskArn
}

func (c *Container) SetVolumes(volumes []Volumes) {
    c.Volumes = volumes
}

func (c *Container) SetVcpus(vcpus float64) {
    c.Vcpus = vcpus
}

func (c *Container) SetCommand(command []string) {
    c.Command = command
}

func (c *Container) SetResourceRequirements(resourceRequirements []ResourceRequirement) {
    c.ResourceRequirements = resourceRequirements
}

func (c *Container) SetEnvironment(environment []ContainerItem) {
    c.Environment = environment
}

func (c *Container) SetUlimits(ulimits []ULimit) {
    c.Ulimits = ulimits
}

func (c *Container) SetNetworkInterfaces(networkInterfaces []NetworkInterface) {
    c.NetworkInterfaces = networkInterfaces
}

func (c *Container) SetMountPoints(mountPoints []MountPoint) {
    c.MountPoints = mountPoints
}

func (c *Container) SetExitCode(exitCode float64) {
    c.ExitCode = exitCode
}

func (c *Container) SetContainerInstanceArn(containerInstanceArn string) {
    c.ContainerInstanceArn = containerInstanceArn
}
