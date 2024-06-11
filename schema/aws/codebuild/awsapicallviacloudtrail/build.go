package awsapicallviacloudtrail

type Build struct {
    Artifacts Artifacts `json:"artifacts,omitempty"`
    Cache Cache `json:"cache,omitempty"`
    Environment Environment `json:"environment,omitempty"`
    Logs Logs `json:"logs,omitempty"`
    Source Source `json:"source,omitempty"`
    VpcConfig VpcConfig `json:"vpcConfig,omitempty"`
    Arn string `json:"arn,omitempty"`
    BuildComplete bool `json:"buildComplete,omitempty"`
    BuildStatus string `json:"buildStatus,omitempty"`
    CurrentPhase string `json:"currentPhase,omitempty"`
    EncryptionKey string `json:"encryptionKey,omitempty"`
    EndTime string `json:"endTime,omitempty"`
    Id string `json:"id,omitempty"`
    Initiator string `json:"initiator,omitempty"`
    Phases []BuildPhase `json:"phases,omitempty"`
    ProjectName string `json:"projectName,omitempty"`
    QueuedTimeoutInMinutes float64 `json:"queuedTimeoutInMinutes,omitempty"`
    ResolvedSourceVersion string `json:"resolvedSourceVersion,omitempty"`
    ServiceRole string `json:"serviceRole,omitempty"`
    SourceVersion string `json:"sourceVersion,omitempty"`
    StartTime string `json:"startTime,omitempty"`
    TimeoutInMinutes float64 `json:"timeoutInMinutes,omitempty"`
}

func (b *Build) SetArtifacts(artifacts Artifacts) {
    b.Artifacts = artifacts
}

func (b *Build) SetCache(cache Cache) {
    b.Cache = cache
}

func (b *Build) SetEnvironment(environment Environment) {
    b.Environment = environment
}

func (b *Build) SetLogs(logs Logs) {
    b.Logs = logs
}

func (b *Build) SetSource(source Source) {
    b.Source = source
}

func (b *Build) SetVpcConfig(vpcConfig VpcConfig) {
    b.VpcConfig = vpcConfig
}

func (b *Build) SetArn(arn string) {
    b.Arn = arn
}

func (b *Build) SetBuildComplete(buildComplete bool) {
    b.BuildComplete = buildComplete
}

func (b *Build) SetBuildStatus(buildStatus string) {
    b.BuildStatus = buildStatus
}

func (b *Build) SetCurrentPhase(currentPhase string) {
    b.CurrentPhase = currentPhase
}

func (b *Build) SetEncryptionKey(encryptionKey string) {
    b.EncryptionKey = encryptionKey
}

func (b *Build) SetEndTime(endTime string) {
    b.EndTime = endTime
}

func (b *Build) SetId(id string) {
    b.Id = id
}

func (b *Build) SetInitiator(initiator string) {
    b.Initiator = initiator
}

func (b *Build) SetPhases(phases []BuildPhase) {
    b.Phases = phases
}

func (b *Build) SetProjectName(projectName string) {
    b.ProjectName = projectName
}

func (b *Build) SetQueuedTimeoutInMinutes(queuedTimeoutInMinutes float64) {
    b.QueuedTimeoutInMinutes = queuedTimeoutInMinutes
}

func (b *Build) SetResolvedSourceVersion(resolvedSourceVersion string) {
    b.ResolvedSourceVersion = resolvedSourceVersion
}

func (b *Build) SetServiceRole(serviceRole string) {
    b.ServiceRole = serviceRole
}

func (b *Build) SetSourceVersion(sourceVersion string) {
    b.SourceVersion = sourceVersion
}

func (b *Build) SetStartTime(startTime string) {
    b.StartTime = startTime
}

func (b *Build) SetTimeoutInMinutes(timeoutInMinutes float64) {
    b.TimeoutInMinutes = timeoutInMinutes
}
