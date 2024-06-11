package awsapicallviacloudtrail

type PullRequestItem struct {
    MergeMetadata MergeMetadata `json:"mergeMetadata"`
    DestinationCommit string `json:"destinationCommit"`
    DestinationReference string `json:"destinationReference"`
    MergeBase string `json:"mergeBase,omitempty"`
    RepositoryName string `json:"repositoryName"`
    SourceCommit string `json:"sourceCommit"`
    SourceReference string `json:"sourceReference"`
}

func (p *PullRequestItem) SetMergeMetadata(mergeMetadata MergeMetadata) {
    p.MergeMetadata = mergeMetadata
}

func (p *PullRequestItem) SetDestinationCommit(destinationCommit string) {
    p.DestinationCommit = destinationCommit
}

func (p *PullRequestItem) SetDestinationReference(destinationReference string) {
    p.DestinationReference = destinationReference
}

func (p *PullRequestItem) SetMergeBase(mergeBase string) {
    p.MergeBase = mergeBase
}

func (p *PullRequestItem) SetRepositoryName(repositoryName string) {
    p.RepositoryName = repositoryName
}

func (p *PullRequestItem) SetSourceCommit(sourceCommit string) {
    p.SourceCommit = sourceCommit
}

func (p *PullRequestItem) SetSourceReference(sourceReference string) {
    p.SourceReference = sourceReference
}
