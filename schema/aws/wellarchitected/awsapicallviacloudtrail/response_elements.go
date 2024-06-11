package awsapicallviacloudtrail

type ResponseElements struct {
    LensArn string `json:"LensArn,omitempty"`
    LensVersion string `json:"LensVersion,omitempty"`
}

func (r *ResponseElements) SetLensArn(lensArn string) {
    r.LensArn = lensArn
}

func (r *ResponseElements) SetLensVersion(lensVersion string) {
    r.LensVersion = lensVersion
}
