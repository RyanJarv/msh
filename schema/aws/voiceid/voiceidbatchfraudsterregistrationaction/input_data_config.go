package voiceidbatchfraudsterregistrationaction

type InputDataConfig struct {
    S3Uri string `json:"s3Uri,omitempty"`
}

func (i *InputDataConfig) SetS3Uri(s3Uri string) {
    i.S3Uri = s3Uri
}
