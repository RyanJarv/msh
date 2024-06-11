package awsapicallviacloudtrail

type EnclaveOptions struct {
    Enabled bool `json:"enabled,omitempty"`
}

func (e *EnclaveOptions) SetEnabled(enabled bool) {
    e.Enabled = enabled
}
