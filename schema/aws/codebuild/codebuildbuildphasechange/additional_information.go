package codebuildbuildphasechange

type Additional_information struct {
    Artifact Artifact `json:"artifact"`
    Cache Cache `json:"cache"`
    Environment Environment `json:"environment"`
    Logs Logs `json:"logs"`
    NetworkInterface Network_interface `json:"network-interface"`
    Source Source `json:"source"`
    VpcConfig Vpc_config `json:"vpc-config"`
    BuildComplete bool `json:"build-complete"`
    BuildStartTime string `json:"build-start-time"`
    Initiator string `json:"initiator"`
    Phases []Additional_informationItem `json:"phases"`
    QueuedTimeoutInMinutes float64 `json:"queued-timeout-in-minutes"`
    SourceVersion string `json:"source-version,omitempty"`
    TimeoutInMinutes float64 `json:"timeout-in-minutes"`
}

func (a *Additional_information) SetArtifact(artifact Artifact) {
    a.Artifact = artifact
}

func (a *Additional_information) SetCache(cache Cache) {
    a.Cache = cache
}

func (a *Additional_information) SetEnvironment(environment Environment) {
    a.Environment = environment
}

func (a *Additional_information) SetLogs(logs Logs) {
    a.Logs = logs
}

func (a *Additional_information) SetNetworkInterface(networkInterface Network_interface) {
    a.NetworkInterface = networkInterface
}

func (a *Additional_information) SetSource(source Source) {
    a.Source = source
}

func (a *Additional_information) SetVpcConfig(vpcConfig Vpc_config) {
    a.VpcConfig = vpcConfig
}

func (a *Additional_information) SetBuildComplete(buildComplete bool) {
    a.BuildComplete = buildComplete
}

func (a *Additional_information) SetBuildStartTime(buildStartTime string) {
    a.BuildStartTime = buildStartTime
}

func (a *Additional_information) SetInitiator(initiator string) {
    a.Initiator = initiator
}

func (a *Additional_information) SetPhases(phases []Additional_informationItem) {
    a.Phases = phases
}

func (a *Additional_information) SetQueuedTimeoutInMinutes(queuedTimeoutInMinutes float64) {
    a.QueuedTimeoutInMinutes = queuedTimeoutInMinutes
}

func (a *Additional_information) SetSourceVersion(sourceVersion string) {
    a.SourceVersion = sourceVersion
}

func (a *Additional_information) SetTimeoutInMinutes(timeoutInMinutes float64) {
    a.TimeoutInMinutes = timeoutInMinutes
}
