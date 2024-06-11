package awsapicallviacloudtrail

type RequestParametersItem_1 struct {
    Commit string `json:"commit,omitempty"`
    Ref string `json:"ref,omitempty"`
}

func (r *RequestParametersItem_1) SetCommit(commit string) {
    r.Commit = commit
}

func (r *RequestParametersItem_1) SetRef(ref string) {
    r.Ref = ref
}
