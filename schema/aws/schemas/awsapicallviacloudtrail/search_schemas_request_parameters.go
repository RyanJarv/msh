package awsapicallviacloudtrail

type SearchSchemasRequestParameters struct {
    RegistryName string `json:"registryName"`
    Limit string `json:"limit,omitempty"`
    NextToken string `json:"nextToken,omitempty"`
    Keywords string `json:"keywords,omitempty"`
}

func (s *SearchSchemasRequestParameters) SetRegistryName(registryName string) {
    s.RegistryName = registryName
}

func (s *SearchSchemasRequestParameters) SetLimit(limit string) {
    s.Limit = limit
}

func (s *SearchSchemasRequestParameters) SetNextToken(nextToken string) {
    s.NextToken = nextToken
}

func (s *SearchSchemasRequestParameters) SetKeywords(keywords string) {
    s.Keywords = keywords
}
