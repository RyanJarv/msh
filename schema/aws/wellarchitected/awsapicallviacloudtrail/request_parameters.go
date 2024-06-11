package awsapicallviacloudtrail

type RequestParameters struct {
    IsMajorVersion bool `json:"IsMajorVersion,omitempty"`
    LensVersion string `json:"LensVersion,omitempty"`
    ClientRequestToken string `json:"ClientRequestToken,omitempty"`
    LensAlias string `json:"LensAlias,omitempty"`
}

func (r *RequestParameters) SetIsMajorVersion(isMajorVersion bool) {
    r.IsMajorVersion = isMajorVersion
}

func (r *RequestParameters) SetLensVersion(lensVersion string) {
    r.LensVersion = lensVersion
}

func (r *RequestParameters) SetClientRequestToken(clientRequestToken string) {
    r.ClientRequestToken = clientRequestToken
}

func (r *RequestParameters) SetLensAlias(lensAlias string) {
    r.LensAlias = lensAlias
}
