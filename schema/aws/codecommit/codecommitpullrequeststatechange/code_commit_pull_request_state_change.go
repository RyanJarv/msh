package codecommitpullrequeststatechange

type CodeCommitPullRequestStateChange struct {
    ApprovalStatus string `json:"approvalStatus,omitempty"`
    Author string `json:"author"`
    CallerUserArn string `json:"callerUserArn"`
    CreationDate string `json:"creationDate"`
    Description string `json:"description"`
    DestinationCommit string `json:"destinationCommit"`
    DestinationReference string `json:"destinationReference"`
    Event string `json:"event"`
    IsMerged string `json:"isMerged"`
    LastModifiedDate string `json:"lastModifiedDate"`
    MergeOption string `json:"mergeOption,omitempty"`
    NotificationBody string `json:"notificationBody"`
    OverrideStatus string `json:"overrideStatus,omitempty"`
    PullRequestId string `json:"pullRequestId"`
    PullRequestStatus string `json:"pullRequestStatus"`
    RepositoryNames []string `json:"repositoryNames"`
    RevisionId string `json:"revisionId"`
    SourceCommit string `json:"sourceCommit"`
    SourceReference string `json:"sourceReference"`
    Title string `json:"title"`
}

func (c *CodeCommitPullRequestStateChange) SetApprovalStatus(approvalStatus string) {
    c.ApprovalStatus = approvalStatus
}

func (c *CodeCommitPullRequestStateChange) SetAuthor(author string) {
    c.Author = author
}

func (c *CodeCommitPullRequestStateChange) SetCallerUserArn(callerUserArn string) {
    c.CallerUserArn = callerUserArn
}

func (c *CodeCommitPullRequestStateChange) SetCreationDate(creationDate string) {
    c.CreationDate = creationDate
}

func (c *CodeCommitPullRequestStateChange) SetDescription(description string) {
    c.Description = description
}

func (c *CodeCommitPullRequestStateChange) SetDestinationCommit(destinationCommit string) {
    c.DestinationCommit = destinationCommit
}

func (c *CodeCommitPullRequestStateChange) SetDestinationReference(destinationReference string) {
    c.DestinationReference = destinationReference
}

func (c *CodeCommitPullRequestStateChange) SetEvent(event string) {
    c.Event = event
}

func (c *CodeCommitPullRequestStateChange) SetIsMerged(isMerged string) {
    c.IsMerged = isMerged
}

func (c *CodeCommitPullRequestStateChange) SetLastModifiedDate(lastModifiedDate string) {
    c.LastModifiedDate = lastModifiedDate
}

func (c *CodeCommitPullRequestStateChange) SetMergeOption(mergeOption string) {
    c.MergeOption = mergeOption
}

func (c *CodeCommitPullRequestStateChange) SetNotificationBody(notificationBody string) {
    c.NotificationBody = notificationBody
}

func (c *CodeCommitPullRequestStateChange) SetOverrideStatus(overrideStatus string) {
    c.OverrideStatus = overrideStatus
}

func (c *CodeCommitPullRequestStateChange) SetPullRequestId(pullRequestId string) {
    c.PullRequestId = pullRequestId
}

func (c *CodeCommitPullRequestStateChange) SetPullRequestStatus(pullRequestStatus string) {
    c.PullRequestStatus = pullRequestStatus
}

func (c *CodeCommitPullRequestStateChange) SetRepositoryNames(repositoryNames []string) {
    c.RepositoryNames = repositoryNames
}

func (c *CodeCommitPullRequestStateChange) SetRevisionId(revisionId string) {
    c.RevisionId = revisionId
}

func (c *CodeCommitPullRequestStateChange) SetSourceCommit(sourceCommit string) {
    c.SourceCommit = sourceCommit
}

func (c *CodeCommitPullRequestStateChange) SetSourceReference(sourceReference string) {
    c.SourceReference = sourceReference
}

func (c *CodeCommitPullRequestStateChange) SetTitle(title string) {
    c.Title = title
}
