package awsapicallviacloudtrail

type Comment struct {
    AuthorArn string `json:"authorArn,omitempty"`
    ClientRequestToken string `json:"clientRequestToken,omitempty"`
    CommentId string `json:"commentId,omitempty"`
    Content string `json:"content,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    Deleted bool `json:"deleted,omitempty"`
    InReplyTo string `json:"inReplyTo,omitempty"`
    LastModifiedDate string `json:"lastModifiedDate,omitempty"`
}

func (c *Comment) SetAuthorArn(authorArn string) {
    c.AuthorArn = authorArn
}

func (c *Comment) SetClientRequestToken(clientRequestToken string) {
    c.ClientRequestToken = clientRequestToken
}

func (c *Comment) SetCommentId(commentId string) {
    c.CommentId = commentId
}

func (c *Comment) SetContent(content string) {
    c.Content = content
}

func (c *Comment) SetCreationDate(creationDate string) {
    c.CreationDate = creationDate
}

func (c *Comment) SetDeleted(deleted bool) {
    c.Deleted = deleted
}

func (c *Comment) SetInReplyTo(inReplyTo string) {
    c.InReplyTo = inReplyTo
}

func (c *Comment) SetLastModifiedDate(lastModifiedDate string) {
    c.LastModifiedDate = lastModifiedDate
}
