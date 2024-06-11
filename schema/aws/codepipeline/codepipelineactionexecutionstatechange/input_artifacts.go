package codepipelineactionexecutionstatechange

type InputArtifacts struct {
    Name string `json:"name"`
    S3location S3Location `json:"s3location"`
}

func (i *InputArtifacts) SetName(name string) {
    i.Name = name
}

func (i *InputArtifacts) SetS3location(s3location S3Location) {
    i.S3location = s3location
}
