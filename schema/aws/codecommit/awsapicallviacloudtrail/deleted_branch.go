package awsapicallviacloudtrail

type DeletedBranch struct {
    BranchName string `json:"branchName,omitempty"`
    CommitId string `json:"commitId,omitempty"`
}

func (d *DeletedBranch) SetBranchName(branchName string) {
    d.BranchName = branchName
}

func (d *DeletedBranch) SetCommitId(commitId string) {
    d.CommitId = commitId
}
