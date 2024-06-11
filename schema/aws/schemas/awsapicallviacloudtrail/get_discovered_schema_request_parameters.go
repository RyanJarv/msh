package awsapicallviacloudtrail

type GetDiscoveredSchemaRequestParameters struct {
    GetDiscoveredSchemaRequestParametersType string `json:"Type"`
    Events []GetDiscoveredSchemaVersionItemInput `json:"Events"`
}

func (g *GetDiscoveredSchemaRequestParameters) SetGetDiscoveredSchemaRequestParametersType(getDiscoveredSchemaRequestParametersType string) {
    g.GetDiscoveredSchemaRequestParametersType = getDiscoveredSchemaRequestParametersType
}

func (g *GetDiscoveredSchemaRequestParameters) SetEvents(events []GetDiscoveredSchemaVersionItemInput) {
    g.Events = events
}
