package sagemakertransformjobstatechange

type TransformOutput struct {
    AssembleWith string `json:"AssembleWith"`
    S3OutputPath string `json:"S3OutputPath"`
}

func (t *TransformOutput) SetAssembleWith(assembleWith string) {
    t.AssembleWith = assembleWith
}

func (t *TransformOutput) SetS3OutputPath(s3OutputPath string) {
    t.S3OutputPath = s3OutputPath
}
