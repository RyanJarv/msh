package awsapicallviacloudtrail

type ListRegistriesRequestParameters struct {
    Scope string `json:"scope,omitempty"`
    Limit string `json:"limit,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    RegistryNamePrefix string `json:"registryNamePrefix,omitempty"`
}

func (l *ListRegistriesRequestParameters) SetScope(scope string) {
    l.Scope = scope
}

func (l *ListRegistriesRequestParameters) SetLimit(limit string) {
    l.Limit = limit
}

func (l *ListRegistriesRequestParameters) SetNextToken(nextToken string) {
    l.NextToken = nextToken
}

func (l *ListRegistriesRequestParameters) SetRegistryNamePrefix(registryNamePrefix string) {
    l.RegistryNamePrefix = registryNamePrefix
}
