package codecommitrepositorystatechange

type CodeCommitRepositoryStateChange struct {
    BaseCommitId string `json:"baseCommitId,omitempty"`
    CallerUserArn string `json:"callerUserArn"`
    CommitId string `json:"commitId,omitempty"`
    ConflictDetailLevel string `json:"conflictDetailLevel,omitempty"`
    ConflictResolutionStrategy string `json:"conflictResolutionStrategy,omitempty"`
    DestinationCommitId string `json:"destinationCommitId,omitempty"`
    Event string `json:"event"`
    MergeOption string `json:"mergeOption,omitempty"`
    OldCommitId string `json:"oldCommitId,omitempty"`
    ReferenceFullName string `json:"referenceFullName,omitempty"`
    ReferenceName string `json:"referenceName,omitempty"`
    ReferenceType string `json:"referenceType,omitempty"`
    RepositoryId string `json:"repositoryId"`
    RepositoryName string `json:"repositoryName"`
    SourceCommitId string `json:"sourceCommitId,omitempty"`
    ConflictDetailsLevel string `json:"conflictDetailsLevel,omitempty"`
}

func (c *CodeCommitRepositoryStateChange) SetBaseCommitId(baseCommitId string) {
    c.BaseCommitId = baseCommitId
}

func (c *CodeCommitRepositoryStateChange) SetCallerUserArn(callerUserArn string) {
    c.CallerUserArn = callerUserArn
}

func (c *CodeCommitRepositoryStateChange) SetCommitId(commitId string) {
    c.CommitId = commitId
}

func (c *CodeCommitRepositoryStateChange) SetConflictDetailLevel(conflictDetailLevel string) {
    c.ConflictDetailLevel = conflictDetailLevel
}

func (c *CodeCommitRepositoryStateChange) SetConflictResolutionStrategy(conflictResolutionStrategy string) {
    c.ConflictResolutionStrategy = conflictResolutionStrategy
}

func (c *CodeCommitRepositoryStateChange) SetDestinationCommitId(destinationCommitId string) {
    c.DestinationCommitId = destinationCommitId
}

func (c *CodeCommitRepositoryStateChange) SetEvent(event string) {
    c.Event = event
}

func (c *CodeCommitRepositoryStateChange) SetMergeOption(mergeOption string) {
    c.MergeOption = mergeOption
}

func (c *CodeCommitRepositoryStateChange) SetOldCommitId(oldCommitId string) {
    c.OldCommitId = oldCommitId
}

func (c *CodeCommitRepositoryStateChange) SetReferenceFullName(referenceFullName string) {
    c.ReferenceFullName = referenceFullName
}

func (c *CodeCommitRepositoryStateChange) SetReferenceName(referenceName string) {
    c.ReferenceName = referenceName
}

func (c *CodeCommitRepositoryStateChange) SetReferenceType(referenceType string) {
    c.ReferenceType = referenceType
}

func (c *CodeCommitRepositoryStateChange) SetRepositoryId(repositoryId string) {
    c.RepositoryId = repositoryId
}

func (c *CodeCommitRepositoryStateChange) SetRepositoryName(repositoryName string) {
    c.RepositoryName = repositoryName
}

func (c *CodeCommitRepositoryStateChange) SetSourceCommitId(sourceCommitId string) {
    c.SourceCommitId = sourceCommitId
}

func (c *CodeCommitRepositoryStateChange) SetConflictDetailsLevel(conflictDetailsLevel string) {
    c.ConflictDetailsLevel = conflictDetailsLevel
}
