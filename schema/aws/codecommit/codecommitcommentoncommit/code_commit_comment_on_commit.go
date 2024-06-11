package codecommitcommentoncommit

type CodeCommitCommentOnCommit struct {
    BeforeCommitId string `json:"beforeCommitId"`
    RepositoryId string `json:"repositoryId"`
    InReplyTo string `json:"inReplyTo"`
    NotificationBody string `json:"notificationBody"`
    CommentId string `json:"commentId"`
    AfterCommitId string `json:"afterCommitId"`
    Event string `json:"event"`
    RepositoryName string `json:"repositoryName"`
    CallerUserArn string `json:"callerUserArn"`
}

func (c *CodeCommitCommentOnCommit) SetBeforeCommitId(beforeCommitId string) {
    c.BeforeCommitId = beforeCommitId
}

func (c *CodeCommitCommentOnCommit) SetRepositoryId(repositoryId string) {
    c.RepositoryId = repositoryId
}

func (c *CodeCommitCommentOnCommit) SetInReplyTo(inReplyTo string) {
    c.InReplyTo = inReplyTo
}

func (c *CodeCommitCommentOnCommit) SetNotificationBody(notificationBody string) {
    c.NotificationBody = notificationBody
}

func (c *CodeCommitCommentOnCommit) SetCommentId(commentId string) {
    c.CommentId = commentId
}

func (c *CodeCommitCommentOnCommit) SetAfterCommitId(afterCommitId string) {
    c.AfterCommitId = afterCommitId
}

func (c *CodeCommitCommentOnCommit) SetEvent(event string) {
    c.Event = event
}

func (c *CodeCommitCommentOnCommit) SetRepositoryName(repositoryName string) {
    c.RepositoryName = repositoryName
}

func (c *CodeCommitCommentOnCommit) SetCallerUserArn(callerUserArn string) {
    c.CallerUserArn = callerUserArn
}
