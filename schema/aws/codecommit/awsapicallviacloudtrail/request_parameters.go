package awsapicallviacloudtrail

type RequestParameters struct {
    Location Location `json:"location,omitempty"`
    Tags Tags `json:"tags,omitempty"`
    AfterCommitId string `json:"afterCommitId,omitempty"`
    ApprovalRuleTemplateContent string `json:"approvalRuleTemplateContent,omitempty"`
    ApprovalRuleTemplateDescription string `json:"approvalRuleTemplateDescription,omitempty"`
    ApprovalRuleTemplateName string `json:"approvalRuleTemplateName,omitempty"`
    ApprovalState string `json:"approvalState,omitempty"`
    ArchiveType string `json:"archiveType,omitempty"`
    BeforeCommitId string `json:"beforeCommitId,omitempty"`
    BranchName string `json:"branchName,omitempty"`
    ClientRequestToken string `json:"clientRequestToken,omitempty"`
    CommentId string `json:"commentId,omitempty"`
    CommitId string `json:"commitId,omitempty"`
    CommitIds []string `json:"commitIds,omitempty"`
    CommitMessage string `json:"commitMessage,omitempty"`
    ConflictDetailLevel string `json:"conflictDetailLevel,omitempty"`
    ConflictResolutionStrategy string `json:"conflictResolutionStrategy,omitempty"`
    Content string `json:"content,omitempty"`
    DefaultBranchName string `json:"defaultBranchName,omitempty"`
    DeleteFiles []RequestParametersItem `json:"deleteFiles,omitempty"`
    Description string `json:"description,omitempty"`
    DestinationCommitSpecifier string `json:"destinationCommitSpecifier,omitempty"`
    FileMode string `json:"fileMode,omitempty"`
    FilePath string `json:"filePath,omitempty"`
    FilePaths []string `json:"filePaths,omitempty"`
    InReplyTo string `json:"inReplyTo,omitempty"`
    KeepEmptyFolders bool `json:"keepEmptyFolders,omitempty"`
    MaxConflictFiles float64 `json:"maxConflictFiles,omitempty"`
    MaxMergeHunks float64 `json:"maxMergeHunks,omitempty"`
    MergeOption string `json:"mergeOption,omitempty"`
    Name string `json:"name,omitempty"`
    NewName string `json:"newName,omitempty"`
    OldName string `json:"oldName,omitempty"`
    ParentCommitId string `json:"parentCommitId,omitempty"`
    PullRequestId string `json:"pullRequestId,omitempty"`
    PullRequestIds []string `json:"pullRequestIds,omitempty"`
    PullRequestStatus string `json:"pullRequestStatus,omitempty"`
    PutFiles []RequestParametersItem `json:"putFiles,omitempty"`
    References []RequestParametersItem_1 `json:"references,omitempty"`
    RepositoryDescription string `json:"repositoryDescription,omitempty"`
    RepositoryName string `json:"repositoryName,omitempty"`
    ResourceArn string `json:"resourceArn,omitempty"`
    RevisionId string `json:"revisionId,omitempty"`
    S3Bucket string `json:"s3Bucket,omitempty"`
    S3Key string `json:"s3Key,omitempty"`
    SourceCommitSpecifier string `json:"sourceCommitSpecifier,omitempty"`
    TagKeys []string `json:"tagKeys,omitempty"`
    TargetBranch string `json:"targetBranch,omitempty"`
    Targets []RequestParametersItem_2 `json:"targets,omitempty"`
    ThrowIfClosed bool `json:"throwIfClosed,omitempty"`
    Title string `json:"title,omitempty"`
    UploadId string `json:"uploadId,omitempty"`
}

func (r *RequestParameters) SetLocation(location Location) {
    r.Location = location
}

func (r *RequestParameters) SetTags(tags Tags) {
    r.Tags = tags
}

func (r *RequestParameters) SetAfterCommitId(afterCommitId string) {
    r.AfterCommitId = afterCommitId
}

func (r *RequestParameters) SetApprovalRuleTemplateContent(approvalRuleTemplateContent string) {
    r.ApprovalRuleTemplateContent = approvalRuleTemplateContent
}

func (r *RequestParameters) SetApprovalRuleTemplateDescription(approvalRuleTemplateDescription string) {
    r.ApprovalRuleTemplateDescription = approvalRuleTemplateDescription
}

func (r *RequestParameters) SetApprovalRuleTemplateName(approvalRuleTemplateName string) {
    r.ApprovalRuleTemplateName = approvalRuleTemplateName
}

func (r *RequestParameters) SetApprovalState(approvalState string) {
    r.ApprovalState = approvalState
}

func (r *RequestParameters) SetArchiveType(archiveType string) {
    r.ArchiveType = archiveType
}

func (r *RequestParameters) SetBeforeCommitId(beforeCommitId string) {
    r.BeforeCommitId = beforeCommitId
}

func (r *RequestParameters) SetBranchName(branchName string) {
    r.BranchName = branchName
}

func (r *RequestParameters) SetClientRequestToken(clientRequestToken string) {
    r.ClientRequestToken = clientRequestToken
}

func (r *RequestParameters) SetCommentId(commentId string) {
    r.CommentId = commentId
}

func (r *RequestParameters) SetCommitId(commitId string) {
    r.CommitId = commitId
}

func (r *RequestParameters) SetCommitIds(commitIds []string) {
    r.CommitIds = commitIds
}

func (r *RequestParameters) SetCommitMessage(commitMessage string) {
    r.CommitMessage = commitMessage
}

func (r *RequestParameters) SetConflictDetailLevel(conflictDetailLevel string) {
    r.ConflictDetailLevel = conflictDetailLevel
}

func (r *RequestParameters) SetConflictResolutionStrategy(conflictResolutionStrategy string) {
    r.ConflictResolutionStrategy = conflictResolutionStrategy
}

func (r *RequestParameters) SetContent(content string) {
    r.Content = content
}

func (r *RequestParameters) SetDefaultBranchName(defaultBranchName string) {
    r.DefaultBranchName = defaultBranchName
}

func (r *RequestParameters) SetDeleteFiles(deleteFiles []RequestParametersItem) {
    r.DeleteFiles = deleteFiles
}

func (r *RequestParameters) SetDescription(description string) {
    r.Description = description
}

func (r *RequestParameters) SetDestinationCommitSpecifier(destinationCommitSpecifier string) {
    r.DestinationCommitSpecifier = destinationCommitSpecifier
}

func (r *RequestParameters) SetFileMode(fileMode string) {
    r.FileMode = fileMode
}

func (r *RequestParameters) SetFilePath(filePath string) {
    r.FilePath = filePath
}

func (r *RequestParameters) SetFilePaths(filePaths []string) {
    r.FilePaths = filePaths
}

func (r *RequestParameters) SetInReplyTo(inReplyTo string) {
    r.InReplyTo = inReplyTo
}

func (r *RequestParameters) SetKeepEmptyFolders(keepEmptyFolders bool) {
    r.KeepEmptyFolders = keepEmptyFolders
}

func (r *RequestParameters) SetMaxConflictFiles(maxConflictFiles float64) {
    r.MaxConflictFiles = maxConflictFiles
}

func (r *RequestParameters) SetMaxMergeHunks(maxMergeHunks float64) {
    r.MaxMergeHunks = maxMergeHunks
}

func (r *RequestParameters) SetMergeOption(mergeOption string) {
    r.MergeOption = mergeOption
}

func (r *RequestParameters) SetName(name string) {
    r.Name = name
}

func (r *RequestParameters) SetNewName(newName string) {
    r.NewName = newName
}

func (r *RequestParameters) SetOldName(oldName string) {
    r.OldName = oldName
}

func (r *RequestParameters) SetParentCommitId(parentCommitId string) {
    r.ParentCommitId = parentCommitId
}

func (r *RequestParameters) SetPullRequestId(pullRequestId string) {
    r.PullRequestId = pullRequestId
}

func (r *RequestParameters) SetPullRequestIds(pullRequestIds []string) {
    r.PullRequestIds = pullRequestIds
}

func (r *RequestParameters) SetPullRequestStatus(pullRequestStatus string) {
    r.PullRequestStatus = pullRequestStatus
}

func (r *RequestParameters) SetPutFiles(putFiles []RequestParametersItem) {
    r.PutFiles = putFiles
}

func (r *RequestParameters) SetReferences(references []RequestParametersItem_1) {
    r.References = references
}

func (r *RequestParameters) SetRepositoryDescription(repositoryDescription string) {
    r.RepositoryDescription = repositoryDescription
}

func (r *RequestParameters) SetRepositoryName(repositoryName string) {
    r.RepositoryName = repositoryName
}

func (r *RequestParameters) SetResourceArn(resourceArn string) {
    r.ResourceArn = resourceArn
}

func (r *RequestParameters) SetRevisionId(revisionId string) {
    r.RevisionId = revisionId
}

func (r *RequestParameters) SetS3Bucket(s3Bucket string) {
    r.S3Bucket = s3Bucket
}

func (r *RequestParameters) SetS3Key(s3Key string) {
    r.S3Key = s3Key
}

func (r *RequestParameters) SetSourceCommitSpecifier(sourceCommitSpecifier string) {
    r.SourceCommitSpecifier = sourceCommitSpecifier
}

func (r *RequestParameters) SetTagKeys(tagKeys []string) {
    r.TagKeys = tagKeys
}

func (r *RequestParameters) SetTargetBranch(targetBranch string) {
    r.TargetBranch = targetBranch
}

func (r *RequestParameters) SetTargets(targets []RequestParametersItem_2) {
    r.Targets = targets
}

func (r *RequestParameters) SetThrowIfClosed(throwIfClosed bool) {
    r.ThrowIfClosed = throwIfClosed
}

func (r *RequestParameters) SetTitle(title string) {
    r.Title = title
}

func (r *RequestParameters) SetUploadId(uploadId string) {
    r.UploadId = uploadId
}
