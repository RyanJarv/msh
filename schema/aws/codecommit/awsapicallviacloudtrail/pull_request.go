package awsapicallviacloudtrail

type PullRequest struct {
    ApprovalRules []interface{} `json:"approvalRules,omitempty"`
    AuthorArn string `json:"authorArn,omitempty"`
    ClientRequestToken string `json:"clientRequestToken,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    Description string `json:"description,omitempty"`
    LastActivityDate string `json:"lastActivityDate,omitempty"`
    PullRequestId string `json:"pullRequestId,omitempty"`
    PullRequestStatus string `json:"pullRequestStatus,omitempty"`
    PullRequestTargets []PullRequestItem `json:"pullRequestTargets,omitempty"`
    RevisionId string `json:"revisionId,omitempty"`
    Title string `json:"title,omitempty"`
}

func (p *PullRequest) SetApprovalRules(approvalRules []interface{}) {
    p.ApprovalRules = approvalRules
}

func (p *PullRequest) SetAuthorArn(authorArn string) {
    p.AuthorArn = authorArn
}

func (p *PullRequest) SetClientRequestToken(clientRequestToken string) {
    p.ClientRequestToken = clientRequestToken
}

func (p *PullRequest) SetCreationDate(creationDate string) {
    p.CreationDate = creationDate
}

func (p *PullRequest) SetDescription(description string) {
    p.Description = description
}

func (p *PullRequest) SetLastActivityDate(lastActivityDate string) {
    p.LastActivityDate = lastActivityDate
}

func (p *PullRequest) SetPullRequestId(pullRequestId string) {
    p.PullRequestId = pullRequestId
}

func (p *PullRequest) SetPullRequestStatus(pullRequestStatus string) {
    p.PullRequestStatus = pullRequestStatus
}

func (p *PullRequest) SetPullRequestTargets(pullRequestTargets []PullRequestItem) {
    p.PullRequestTargets = pullRequestTargets
}

func (p *PullRequest) SetRevisionId(revisionId string) {
    p.RevisionId = revisionId
}

func (p *PullRequest) SetTitle(title string) {
    p.Title = title
}
