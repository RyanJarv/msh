package awsapicallviacloudtrail

type MergeMetadata struct {
    IsMerged bool `json:"isMerged"`
    MergeCommitId string `json:"mergeCommitId,omitempty"`
    MergeOption string `json:"mergeOption,omitempty"`
    MergedBy string `json:"mergedBy,omitempty"`
}

func (m *MergeMetadata) SetIsMerged(isMerged bool) {
    m.IsMerged = isMerged
}

func (m *MergeMetadata) SetMergeCommitId(mergeCommitId string) {
    m.MergeCommitId = mergeCommitId
}

func (m *MergeMetadata) SetMergeOption(mergeOption string) {
    m.MergeOption = mergeOption
}

func (m *MergeMetadata) SetMergedBy(mergedBy string) {
    m.MergedBy = mergedBy
}
