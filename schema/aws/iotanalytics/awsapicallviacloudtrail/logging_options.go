package awsapicallviacloudtrail

type LoggingOptions struct {
    Enabled bool `json:"enabled,omitempty"`
}

func (l *LoggingOptions) SetEnabled(enabled bool) {
    l.Enabled = enabled
}
