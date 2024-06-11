package ecrimageaction

type ECRImageAction struct {
    ActionType string `json:"action-type,omitempty"`
    ImageDigest string `json:"image-digest,omitempty"`
    ImageTag string `json:"image-tag,omitempty"`
    RepositoryName string `json:"repository-name,omitempty"`
    Result string `json:"result,omitempty"`
    ManifestMediaType string `json:"manifest-media-type,omitempty"`
    ArtifactMediaType string `json:"artifact-media-type,omitempty"`
}

func (e *ECRImageAction) SetActionType(actionType string) {
    e.ActionType = actionType
}

func (e *ECRImageAction) SetImageDigest(imageDigest string) {
    e.ImageDigest = imageDigest
}

func (e *ECRImageAction) SetImageTag(imageTag string) {
    e.ImageTag = imageTag
}

func (e *ECRImageAction) SetRepositoryName(repositoryName string) {
    e.RepositoryName = repositoryName
}

func (e *ECRImageAction) SetResult(result string) {
    e.Result = result
}

func (e *ECRImageAction) SetManifestMediaType(manifestMediaType string) {
    e.ManifestMediaType = manifestMediaType
}

func (e *ECRImageAction) SetArtifactMediaType(artifactMediaType string) {
    e.ArtifactMediaType = artifactMediaType
}
