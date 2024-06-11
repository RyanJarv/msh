package awsapicallviacloudtrail

type Project struct {
    Artifacts Artifacts_1 `json:"artifacts,omitempty"`
    Badge Badge `json:"badge,omitempty"`
    Cache Cache `json:"cache,omitempty"`
    Environment Environment_1 `json:"environment,omitempty"`
    Source Source_1 `json:"source,omitempty"`
    Arn string `json:"arn,omitempty"`
    Created string `json:"created,omitempty"`
    Description string `json:"description,omitempty"`
    EncryptionKey string `json:"encryptionKey,omitempty"`
    LastModified string `json:"lastModified,omitempty"`
    Name string `json:"name,omitempty"`
    QueuedTimeoutInMinutes float64 `json:"queuedTimeoutInMinutes,omitempty"`
    ServiceRole string `json:"serviceRole,omitempty"`
    SourceVersion string `json:"sourceVersion,omitempty"`
    Tags []ProjectItem `json:"tags,omitempty"`
    TimeoutInMinutes float64 `json:"timeoutInMinutes,omitempty"`
}

func (p *Project) SetArtifacts(artifacts Artifacts_1) {
    p.Artifacts = artifacts
}

func (p *Project) SetBadge(badge Badge) {
    p.Badge = badge
}

func (p *Project) SetCache(cache Cache) {
    p.Cache = cache
}

func (p *Project) SetEnvironment(environment Environment_1) {
    p.Environment = environment
}

func (p *Project) SetSource(source Source_1) {
    p.Source = source
}

func (p *Project) SetArn(arn string) {
    p.Arn = arn
}

func (p *Project) SetCreated(created string) {
    p.Created = created
}

func (p *Project) SetDescription(description string) {
    p.Description = description
}

func (p *Project) SetEncryptionKey(encryptionKey string) {
    p.EncryptionKey = encryptionKey
}

func (p *Project) SetLastModified(lastModified string) {
    p.LastModified = lastModified
}

func (p *Project) SetName(name string) {
    p.Name = name
}

func (p *Project) SetQueuedTimeoutInMinutes(queuedTimeoutInMinutes float64) {
    p.QueuedTimeoutInMinutes = queuedTimeoutInMinutes
}

func (p *Project) SetServiceRole(serviceRole string) {
    p.ServiceRole = serviceRole
}

func (p *Project) SetSourceVersion(sourceVersion string) {
    p.SourceVersion = sourceVersion
}

func (p *Project) SetTags(tags []ProjectItem) {
    p.Tags = tags
}

func (p *Project) SetTimeoutInMinutes(timeoutInMinutes float64) {
    p.TimeoutInMinutes = timeoutInMinutes
}
