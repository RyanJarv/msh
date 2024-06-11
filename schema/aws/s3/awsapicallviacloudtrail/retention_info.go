package awsapicallviacloudtrail

type RetentionInfo struct {
    LastModifiedTime int64 `json:"lastModifiedTime,omitempty"`
    RetainUntilTime int64 `json:"retainUntilTime,omitempty"`
    RetainUntilMode string `json:"retainUntilMode,omitempty"`
}

func (r *RetentionInfo) SetLastModifiedTime(lastModifiedTime int64) {
    r.LastModifiedTime = lastModifiedTime
}

func (r *RetentionInfo) SetRetainUntilTime(retainUntilTime int64) {
    r.RetainUntilTime = retainUntilTime
}

func (r *RetentionInfo) SetRetainUntilMode(retainUntilMode string) {
    r.RetainUntilMode = retainUntilMode
}
