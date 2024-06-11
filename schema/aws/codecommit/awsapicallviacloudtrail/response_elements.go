package awsapicallviacloudtrail

type ResponseElements struct {
    Comment Comment `json:"comment,omitempty"`
    DeletedBranch DeletedBranch `json:"deletedBranch,omitempty"`
    Location Location `json:"location,omitempty"`
    PullRequest PullRequest `json:"pullRequest,omitempty"`
    RepositoryMetadata RepositoryMetadata `json:"repositoryMetadata,omitempty"`
    AfterBlobId string `json:"afterBlobId,omitempty"`
    AfterCommitId string `json:"afterCommitId,omitempty"`
    BeforeBlobId string `json:"beforeBlobId,omitempty"`
    BeforeCommitId string `json:"beforeCommitId,omitempty"`
    BlobId string `json:"blobId,omitempty"`
    CommitId string `json:"commitId,omitempty"`
    FilePath string `json:"filePath,omitempty"`
    FilesAdded []ResponseElementsItem `json:"filesAdded,omitempty"`
    FilesDeleted []ResponseElementsItem `json:"filesDeleted,omitempty"`
    FilesUpdated []ResponseElementsItem `json:"filesUpdated,omitempty"`
    PullRequestId string `json:"pullRequestId,omitempty"`
    RepositoryId string `json:"repositoryId,omitempty"`
    RepositoryName string `json:"repositoryName,omitempty"`
    TreeId string `json:"treeId,omitempty"`
    UploadId string `json:"uploadId,omitempty"`
}

func (r *ResponseElements) SetComment(comment Comment) {
    r.Comment = comment
}

func (r *ResponseElements) SetDeletedBranch(deletedBranch DeletedBranch) {
    r.DeletedBranch = deletedBranch
}

func (r *ResponseElements) SetLocation(location Location) {
    r.Location = location
}

func (r *ResponseElements) SetPullRequest(pullRequest PullRequest) {
    r.PullRequest = pullRequest
}

func (r *ResponseElements) SetRepositoryMetadata(repositoryMetadata RepositoryMetadata) {
    r.RepositoryMetadata = repositoryMetadata
}

func (r *ResponseElements) SetAfterBlobId(afterBlobId string) {
    r.AfterBlobId = afterBlobId
}

func (r *ResponseElements) SetAfterCommitId(afterCommitId string) {
    r.AfterCommitId = afterCommitId
}

func (r *ResponseElements) SetBeforeBlobId(beforeBlobId string) {
    r.BeforeBlobId = beforeBlobId
}

func (r *ResponseElements) SetBeforeCommitId(beforeCommitId string) {
    r.BeforeCommitId = beforeCommitId
}

func (r *ResponseElements) SetBlobId(blobId string) {
    r.BlobId = blobId
}

func (r *ResponseElements) SetCommitId(commitId string) {
    r.CommitId = commitId
}

func (r *ResponseElements) SetFilePath(filePath string) {
    r.FilePath = filePath
}

func (r *ResponseElements) SetFilesAdded(filesAdded []ResponseElementsItem) {
    r.FilesAdded = filesAdded
}

func (r *ResponseElements) SetFilesDeleted(filesDeleted []ResponseElementsItem) {
    r.FilesDeleted = filesDeleted
}

func (r *ResponseElements) SetFilesUpdated(filesUpdated []ResponseElementsItem) {
    r.FilesUpdated = filesUpdated
}

func (r *ResponseElements) SetPullRequestId(pullRequestId string) {
    r.PullRequestId = pullRequestId
}

func (r *ResponseElements) SetRepositoryId(repositoryId string) {
    r.RepositoryId = repositoryId
}

func (r *ResponseElements) SetRepositoryName(repositoryName string) {
    r.RepositoryName = repositoryName
}

func (r *ResponseElements) SetTreeId(treeId string) {
    r.TreeId = treeId
}

func (r *ResponseElements) SetUploadId(uploadId string) {
    r.UploadId = uploadId
}
