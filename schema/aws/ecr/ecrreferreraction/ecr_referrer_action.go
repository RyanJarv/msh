package ecrreferreraction

type ECRReferrerAction struct {
    ActionType string `json:"action-type,omitempty"`
    ImageDigest string `json:"image-digest,omitempty"`
    ImageTag string `json:"image-tag,omitempty"`
    RepositoryName string `json:"repository-name,omitempty"`
    Result string `json:"result,omitempty"`
    ManifestMediaType string `json:"manifest-media-type,omitempty"`
    ArtifactMediaType string `json:"artifact-media-type,omitempty"`
}

func (e *ECRReferrerAction) SetActionType(actionType string) {
    e.ActionType = actionType
}

func (e *ECRReferrerAction) SetImageDigest(imageDigest string) {
    e.ImageDigest = imageDigest
}

func (e *ECRReferrerAction) SetImageTag(imageTag string) {
    e.ImageTag = imageTag
}

func (e *ECRReferrerAction) SetRepositoryName(repositoryName string) {
    e.RepositoryName = repositoryName
}

func (e *ECRReferrerAction) SetResult(result string) {
    e.Result = result
}

func (e *ECRReferrerAction) SetManifestMediaType(manifestMediaType string) {
    e.ManifestMediaType = manifestMediaType
}

func (e *ECRReferrerAction) SetArtifactMediaType(artifactMediaType string) {
    e.ArtifactMediaType = artifactMediaType
}
