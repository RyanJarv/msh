package codepipelineactionexecutionstatechange

type OutputArtifacts struct {
    Name string `json:"name"`
    S3location S3Location `json:"s3location"`
}

func (o *OutputArtifacts) SetName(name string) {
    o.Name = name
}

func (o *OutputArtifacts) SetS3location(s3location S3Location) {
    o.S3location = s3location
}
