package codebuildbuildphasechange

type Logs struct {
    DeepLink string `json:"deep-link"`
    GroupName string `json:"group-name"`
    StreamName string `json:"stream-name"`
}

func (l *Logs) SetDeepLink(deepLink string) {
    l.DeepLink = deepLink
}

func (l *Logs) SetGroupName(groupName string) {
    l.GroupName = groupName
}

func (l *Logs) SetStreamName(streamName string) {
    l.StreamName = streamName
}