package awsapicallviacloudtrail

type Monitoring struct {
    Enabled bool `json:"enabled,omitempty"`
}

func (m *Monitoring) SetEnabled(enabled bool) {
    m.Enabled = enabled
}
