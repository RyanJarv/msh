package awsapicallviacloudtrail

type RepositoryMetadata struct {
    AccountId string `json:"accountId,omitempty"`
    Arn string `json:"arn,omitempty"`
    CloneUrlHttp string `json:"cloneUrlHttp,omitempty"`
    CloneUrlSsh string `json:"cloneUrlSsh,omitempty"`
    CreationDate string `json:"creationDate,omitempty"`
    LastModifiedDate string `json:"lastModifiedDate,omitempty"`
    RepositoryDescription string `json:"repositoryDescription,omitempty"`
    RepositoryId string `json:"repositoryId,omitempty"`
    RepositoryName string `json:"repositoryName,omitempty"`
}

func (r *RepositoryMetadata) SetAccountId(accountId string) {
    r.AccountId = accountId
}

func (r *RepositoryMetadata) SetArn(arn string) {
    r.Arn = arn
}

func (r *RepositoryMetadata) SetCloneUrlHttp(cloneUrlHttp string) {
    r.CloneUrlHttp = cloneUrlHttp
}

func (r *RepositoryMetadata) SetCloneUrlSsh(cloneUrlSsh string) {
    r.CloneUrlSsh = cloneUrlSsh
}

func (r *RepositoryMetadata) SetCreationDate(creationDate string) {
    r.CreationDate = creationDate
}

func (r *RepositoryMetadata) SetLastModifiedDate(lastModifiedDate string) {
    r.LastModifiedDate = lastModifiedDate
}

func (r *RepositoryMetadata) SetRepositoryDescription(repositoryDescription string) {
    r.RepositoryDescription = repositoryDescription
}

func (r *RepositoryMetadata) SetRepositoryId(repositoryId string) {
    r.RepositoryId = repositoryId
}

func (r *RepositoryMetadata) SetRepositoryName(repositoryName string) {
    r.RepositoryName = repositoryName
}
