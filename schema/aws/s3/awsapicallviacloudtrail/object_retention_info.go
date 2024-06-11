package awsapicallviacloudtrail

type ObjectRetentionInfo struct {
    LegalHoldInfo LegalHoldInfo `json:"legalHoldInfo,omitempty"`
    RetentionInfo RetentionInfo `json:"retentionInfo,omitempty"`
}

func (o *ObjectRetentionInfo) SetLegalHoldInfo(legalHoldInfo LegalHoldInfo) {
    o.LegalHoldInfo = legalHoldInfo
}

func (o *ObjectRetentionInfo) SetRetentionInfo(retentionInfo RetentionInfo) {
    o.RetentionInfo = retentionInfo
}
