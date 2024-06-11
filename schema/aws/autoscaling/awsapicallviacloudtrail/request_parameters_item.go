package awsapicallviacloudtrail

type RequestParametersItem struct {
    ResourceId string `json:"resourceId,omitempty"`
    PropagateAtLaunch bool `json:"propagateAtLaunch,omitempty"`
    Value string `json:"value,omitempty"`
    Key string `json:"key,omitempty"`
    ResourceType string `json:"resourceType,omitempty"`
}

func (r *RequestParametersItem) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

func (r *RequestParametersItem) SetPropagateAtLaunch(propagateAtLaunch bool) {
    r.PropagateAtLaunch = propagateAtLaunch
}

func (r *RequestParametersItem) SetValue(value string) {
    r.Value = value
}

func (r *RequestParametersItem) SetKey(key string) {
    r.Key = key
}

func (r *RequestParametersItem) SetResourceType(resourceType string) {
    r.ResourceType = resourceType
}
