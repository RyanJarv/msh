package awsapicallviacloudtrail

type LegalHoldInfo struct {
    LastModifiedTime int64 `json:"lastModifiedTime,omitempty"`
    IsUnderLegalHold bool `json:"isUnderLegalHold,omitempty"`
}

func (l *LegalHoldInfo) SetLastModifiedTime(lastModifiedTime int64) {
    l.LastModifiedTime = lastModifiedTime
}

func (l *LegalHoldInfo) SetIsUnderLegalHold(isUnderLegalHold bool) {
    l.IsUnderLegalHold = isUnderLegalHold
}
