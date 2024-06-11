package awsapicallviacloudtrail

type ListDiscoverersRequestParameters struct {
    SourceArnPrefix string `json:"sourceArnPrefix,omitempty"`
    Limit string `json:"limit,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    DiscovererIdPrefix string `json:"discovererIdPrefix,omitempty"`
}

func (l *ListDiscoverersRequestParameters) SetSourceArnPrefix(sourceArnPrefix string) {
    l.SourceArnPrefix = sourceArnPrefix
}

func (l *ListDiscoverersRequestParameters) SetLimit(limit string) {
    l.Limit = limit
}

func (l *ListDiscoverersRequestParameters) SetNextToken(nextToken string) {
    l.NextToken = nextToken
}

func (l *ListDiscoverersRequestParameters) SetDiscovererIdPrefix(discovererIdPrefix string) {
    l.DiscovererIdPrefix = discovererIdPrefix
}
