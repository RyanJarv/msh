package devopsgurunewanomalyassociation

type AnomalySourceMetadata struct {
    Source string `json:"source,omitempty"`
    SourceResourceType string `json:"sourceResourceType,omitempty"`
    SourceResourceArn string `json:"sourceResourceArn,omitempty"`
}

func (a *AnomalySourceMetadata) SetSource(source string) {
    a.Source = source
}

func (a *AnomalySourceMetadata) SetSourceResourceType(sourceResourceType string) {
    a.SourceResourceType = sourceResourceType
}

func (a *AnomalySourceMetadata) SetSourceResourceArn(sourceResourceArn string) {
    a.SourceResourceArn = sourceResourceArn
}
