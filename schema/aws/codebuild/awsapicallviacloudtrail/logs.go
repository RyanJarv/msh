package awsapicallviacloudtrail

type Logs struct {
    DeepLink string `json:"deepLink,omitempty"`
    GroupName string `json:"groupName,omitempty"`
    StreamName string `json:"streamName,omitempty"`
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
