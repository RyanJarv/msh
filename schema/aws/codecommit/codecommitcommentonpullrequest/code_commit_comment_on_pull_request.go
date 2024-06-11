package codecommitcommentonpullrequest

type CodeCommitCommentOnPullRequest struct {
    BeforeCommitId string `json:"beforeCommitId"`
    RepositoryId string `json:"repositoryId"`
    InReplyTo string `json:"inReplyTo"`
    NotificationBody string `json:"notificationBody"`
    CommentId string `json:"commentId"`
    AfterCommitId string `json:"afterCommitId"`
    Event string `json:"event"`
    RepositoryName string `json:"repositoryName"`
    CallerUserArn string `json:"callerUserArn"`
    PullRequestId string `json:"pullRequestId"`
}

func (c *CodeCommitCommentOnPullRequest) SetBeforeCommitId(beforeCommitId string) {
    c.BeforeCommitId = beforeCommitId
}

func (c *CodeCommitCommentOnPullRequest) SetRepositoryId(repositoryId string) {
    c.RepositoryId = repositoryId
}

func (c *CodeCommitCommentOnPullRequest) SetInReplyTo(inReplyTo string) {
    c.InReplyTo = inReplyTo
}

func (c *CodeCommitCommentOnPullRequest) SetNotificationBody(notificationBody string) {
    c.NotificationBody = notificationBody
}

func (c *CodeCommitCommentOnPullRequest) SetCommentId(commentId string) {
    c.CommentId = commentId
}

func (c *CodeCommitCommentOnPullRequest) SetAfterCommitId(afterCommitId string) {
    c.AfterCommitId = afterCommitId
}

func (c *CodeCommitCommentOnPullRequest) SetEvent(event string) {
    c.Event = event
}

func (c *CodeCommitCommentOnPullRequest) SetRepositoryName(repositoryName string) {
    c.RepositoryName = repositoryName
}

func (c *CodeCommitCommentOnPullRequest) SetCallerUserArn(callerUserArn string) {
    c.CallerUserArn = callerUserArn
}

func (c *CodeCommitCommentOnPullRequest) SetPullRequestId(pullRequestId string) {
    c.PullRequestId = pullRequestId
}
